package resources

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v3"
	"locgame-mini-server/pkg/dto/resources"
	"locgame-mini-server/pkg/log"
)

var Resources []*resources.ResourceData
var Categories []*resources.ResourceCategory

var resourceByID map[int32]*resources.ResourceData
var resourceByKey map[string]*resources.ResourceData

var resourcesByCategoryID map[int32][]*resources.ResourceData
var resourcesByCategoryKey map[string][]*resources.ResourceData

var categoryByID map[int32]*resources.ResourceCategory
var categoryByKey map[string]*resources.ResourceCategory

func Init(configPath string) {
	readResources(configPath)
	readCategories(configPath)
}

func readResources(configPath string) {
	bytes, err := ioutil.ReadFile(filepath.Join(configPath, "resources.yaml"))
	if err != nil {
		log.Fatal("Unable read resources config:", err)
	}
	err = yaml.Unmarshal(bytes, &Resources)
	if err != nil {
		log.Fatal("Unable parse resources config:", err)
	}

	resourceByID = make(map[int32]*resources.ResourceData)
	resourceByKey = make(map[string]*resources.ResourceData)

	for _, resource := range Resources {
		resourceByID[resource.ID] = resource
		resourceByKey[resource.Key] = resource
	}

	resources.SetResources(Resources)
}

func readCategories(configPath string) {
	bytes, err := ioutil.ReadFile(filepath.Join(configPath, "resource_categories.yaml"))
	if err != nil {
		log.Fatal("Unable read resource categories config:", err)
	}
	err = yaml.Unmarshal(bytes, &Categories)
	if err != nil {
		log.Fatal("Unable parse resource categories config:", err)
	}

	resourcesByCategoryID = make(map[int32][]*resources.ResourceData)
	resourcesByCategoryKey = make(map[string][]*resources.ResourceData)

	categoryByID = make(map[int32]*resources.ResourceCategory)
	categoryByKey = make(map[string]*resources.ResourceCategory)

	for _, category := range Categories {
		categoryByID[category.ID] = category
		categoryByKey[category.Key] = category
	}

	for _, resource := range Resources {
		resourcesByCategoryID[resource.CategoryID] = append(resourcesByCategoryID[resource.CategoryID], resource)
		resourcesByCategoryKey[categoryByID[resource.CategoryID].Key] = append(resourcesByCategoryKey[categoryByID[resource.CategoryID].Key], resource)
	}
}

func GetByID(id int32) *resources.ResourceData {
	return resourceByID[id]
}

func GetCategoryByID(id int32) *resources.ResourceCategory {
	return categoryByID[id]
}

func GetByKey(key string) *resources.ResourceData {
	return resourceByKey[key]
}

func GetCategoryByKey(key string) *resources.ResourceCategory {
	return categoryByKey[key]
}

func GetByCategoryID(id int32) []*resources.ResourceData {
	return resourcesByCategoryID[id]
}

func GetByCategoryKey(key string) []*resources.ResourceData {
	return resourcesByCategoryKey[key]
}
