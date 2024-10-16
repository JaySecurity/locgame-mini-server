// Copyright © 2022 FuryLion Group LLC. All Rights Reserved.
//
// Unauthorized copying of this file, via any medium is strictly prohibited
// without the express permission of FuryLion Group LLC.
// Proprietary and confidential.
//
// Created by Vadim Vlasov on 04.10.22

package config

import (
	"io/ioutil"
	"locgame-mini-server/pkg/log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"gopkg.in/yaml.v3"
)

type Repository interface {
	CreateRecord(path string, value interface{}) error
	UpdateRecord(path string, value interface{}) error
	GetRecords(path string) ([]string, error)
	GetRecord(path string, dst interface{}) error
	DeleteRecord(path string) error
	ListBranches() []string
	RefreshWorkspace() bool
}

type authGetter func() *http.BasicAuth

type GitRepository struct {
	repo      *git.Repository
	parentDir string
	getAuth   authGetter

	name string
}

func New(path, name string) *GitRepository {
	auth := func() *http.BasicAuth {
		return &http.BasicAuth{
			Username: "ConfigUser",
			Password: "gldt-2Anxm82H2yLMHztR-Vay",
		}
	}
	return NewWithAuth(auth, "https://gitlab.com/rbllabs1/legends-of-crypto/locgame-configs.git", path, name)
}

func NewWithAuth(getAuth func() *http.BasicAuth, repoUrl, path, name string) *GitRepository {
	r := new(GitRepository)
	r.name = name
	r.parentDir = path
	r.getAuth = getAuth

	if _, err := os.Stat(r.parentDir); os.IsNotExist(err) {
		repo, err := git.PlainClone(r.parentDir, false, &git.CloneOptions{
			Auth:     r.getAuth(),
			URL:      repoUrl,
			Progress: os.Stdout,
		})
		if err != nil {
			log.Fatal(err)
		}
		r.repo = repo
	}

	if err := r.Checkout(name); err != nil {
		log.Fatal(err)
	}
	return r
}

func (r *GitRepository) Checkout(name string) error {
	if r.repo == nil {
		repo, err := git.PlainOpen(r.parentDir)
		if err != nil {
			log.Fatal(err)
		}
		r.repo = repo
		log.Debug("Fetching git configs...")
		err = r.repo.Fetch(&git.FetchOptions{Auth: r.getAuth(), Progress: os.Stdout, Force: true})
		if err != nil && err != git.NoErrAlreadyUpToDate {
			return err
		}
	}

	workTree, err := r.repo.Worktree()
	if err != nil {
		log.Fatal("Can't get a working tree:", err)
	}

	branch, hash := r.determinateBranchAndHash(name)
	pullBranch := branch
	if !hash.IsZero() {
		log.Info("Force checkout to branch:", branch.Short())

		localBranchName := strings.Replace(branch.Short(), "origin/", "", 1)
		localBranchRef := plumbing.NewBranchReferenceName(localBranchName)
		pullBranch = localBranchRef

		if !r.HasBranch(localBranchName) {
			err = r.repo.CreateBranch(&config.Branch{Name: localBranchName, Remote: "origin", Merge: localBranchRef})
		}

		if err == nil {
			newReference := plumbing.NewSymbolicReference(localBranchRef, branch)
			err = r.repo.Storer.SetReference(newReference)
			if err == nil {
				err = r.repo.Storer.SetReference(plumbing.NewHashReference(localBranchRef, hash))
				if err == nil {
					err = workTree.Checkout(&git.CheckoutOptions{
						Branch: plumbing.ReferenceName(localBranchRef.String()),
						Create: !r.HasBranch(localBranchName),
						Force:  true,
					})
				}
			}
		}

		if err != nil {
			return err
		}
	}

	log.Info("Pull changes from config branch... " + pullBranch)
	err = workTree.Pull(&git.PullOptions{Auth: r.getAuth(), ReferenceName: pullBranch})
	if err == git.NoErrAlreadyUpToDate {
		err = nil
	}
	if err != nil {
		log.Info("Error pulling configs: " + err.Error())
	}
	ref, _ := r.repo.Head()
	log.Info("Config branch:", strings.Replace(ref.Name().String(), "refs/heads/", "", 1), "Commit:", ref.Hash())

	return err
}

func (r *GitRepository) determinateBranchAndHash(name string) (plumbing.ReferenceName, plumbing.Hash) {
	var (
		hash       plumbing.Hash
		branchName plumbing.ReferenceName
	)

	if r.isVersion(name) {
		name = "origin/production"
	} else {
		name = "origin/" + name
	}

	iter, err := r.repo.References()
	if err != nil {
		log.Fatal("Unable get references:", err)
	}

	_ = iter.ForEach(func(reference *plumbing.Reference) error {
		if reference.Name().Short() == name {
			hash = reference.Hash()
			branchName = reference.Name()
		}
		return nil
	})

	return branchName, hash
}

func (r *GitRepository) isVersion(name string) bool {
	res, _ := regexp.MatchString("\\d+.\\d+.\\d+", name)
	return res
}

func (r *GitRepository) GetRecords(path string) ([]string, error) {
	files, err := ioutil.ReadDir(filepath.Join(r.parentDir, path))
	if err != nil {
		return []string{}, err
	}

	var actualFiles []string
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".yaml" {
			actualFiles = append(actualFiles, strings.ReplaceAll(file.Name(), yamlExt, ""))
		}
	}
	return actualFiles, nil
}

func (r *GitRepository) GetRecord(path string, dst interface{}) error {
	path += yamlExt

	data, err := os.ReadFile(filepath.Join(r.parentDir, path))
	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, dst)
}

func (r *GitRepository) CreateRecord(path string, value interface{}) error {
	path += yamlExt

	data, err := yaml.Marshal(value)
	if err != nil {
		log.Error(err)
	}
	err = os.WriteFile(filepath.Join(r.parentDir, path), data, 0644)
	if err != nil {
		return err
	}

	return r.save(path)
}

func (r *GitRepository) DeleteRecord(path string) error {
	path += yamlExt

	err := os.Remove(filepath.Join(r.parentDir, path))
	if err != nil {
		return err
	}

	return r.save(path)
}

func (r *GitRepository) save(path string) error {
	wt, _ := r.repo.Worktree()
	_, err := wt.Add(path)
	if err != nil {
		log.Error(err)
		return err
	}

	err = r.repo.Push(&git.PushOptions{Auth: r.getAuth()})
	if err != nil {
		log.Error(err)
	}
	return err
}

func (r *GitRepository) UpdateRecord(path string, value interface{}) error {
	return r.CreateRecord(path, value)
}

func (r *GitRepository) HasBranch(name string) bool {
	var exists bool
	iter, _ := r.repo.Branches()
	_ = iter.ForEach(func(reference *plumbing.Reference) error {
		if exists {
			return nil
		}

		if reference.Name().Short() == name {
			exists = true
		}
		return nil
	})
	return exists
}

func (r *GitRepository) ListBranches() []string {
	var branches []string
	iter, _ := r.repo.Branches()
	_ = iter.ForEach(func(reference *plumbing.Reference) error {
		branches = append(branches, reference.Name().String())
		return nil
	})

	return branches
}

func (r *GitRepository) ListTags() []string {
	var tags []string
	iter, _ := r.repo.Tags()
	_ = iter.ForEach(func(reference *plumbing.Reference) error {
		tags = append(tags, reference.Name().String())
		return nil
	})

	return tags
}

func (r *GitRepository) RefreshWorkspace() bool {
	previousHash, _ := r.repo.Head()
	_ = r.Checkout(r.name)
	currentHash, _ := r.repo.Head()

	return previousHash != currentHash
}
