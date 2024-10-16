package main

import (
	"bufio"
	"fmt"
	"go/format"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"locgame-mini-server/pkg/log"
)

func main() {
	protoFiles, _ := filepath.Glob("proto/*.proto")

	// Generating Go proto files
	if err := runCommand("protoc -I=/usr/local/include -I proto --go_out=.. " + strings.Join(protoFiles, " ")); err != nil {
		log.Fatal(err)
	}

	_ = filepath.Walk("pkg/dto", func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		data, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}

		res := addYamlAndBsonTag(string(data))
		res = renameBsonID(res)
		res = removeOmitEmptyFromBooleans(res)
		// Remove omitempty and add omitnil
		fieldsToRemoveOmitEmpty := []string{"VirtualCards"}

		for _, fieldName := range fieldsToRemoveOmitEmpty {
			res = removeOmitEmptyFromField(res, fieldName)
			res = addOmitNilToField(res, fieldName)
		}

		return ioutil.WriteFile(path, []byte(res), 0644)
	})

	unityProjectPath := os.Getenv("LOCGAME_PROTO_UNITY_PATH")
	if unityProjectPath == "" {
		unityProjectPath = "../locgame-client-v3/Assets/Backend/Dto"
	}

	// Generating network code
	var networkGenCommand string
	if _, err := os.Stat(unityProjectPath); !os.IsNotExist(err) {
		networkGenCommand = "network-gen -csharp_client_out=" + unityProjectPath + " -go_server_out=pkg/dto proto/handlers.network"

		// Generating C# proto files
		if err := runCommand("protoc -I=proto --csharp_out " + unityProjectPath + " " + strings.Join(protoFiles, " ")); err != nil {
			log.Fatal(err)
		}
	} else {
		networkGenCommand = "network-gen -go_server_out=pkg/dto proto/handlers.network"
	}

	if err := runCommand(networkGenCommand); err != nil {
		log.Fatal(err)
	}

	processAdminPanel()
	processTestingTool()
}

func processTestingTool() {
	if _, err := os.Stat("../locgame-testing-tool"); !os.IsNotExist(err) {
		_ = os.RemoveAll("../locgame-testing-tool/pkg/dto")

		if err := runCommand("cp -r pkg/dto ../locgame-testing-tool/pkg/dto"); err != nil {
			log.Fatal(err)
		}

		_ = os.Remove("../locgame-testing-tool/pkg/dto/handlers.go")

		if err := runCommand("network-gen -go_client_out=../locgame-testing-tool/pkg/client -file_name=router -override_go_module_name=locgame-testing-tool -override_go_package=client -proto_package_path=pkg/dto proto/handlers.network"); err != nil {
			log.Fatal(err)
		}

		_ = filepath.Walk("../locgame-testing-tool/pkg/dto", func(path string, info fs.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}

			data, err := ioutil.ReadFile(path)
			if err != nil {
				panic(err)
			}

			res := strings.ReplaceAll(string(data), "locgame-mini-server/pkg/dto", "locgame-testing-tool/pkg/dto")

			return ioutil.WriteFile(path, []byte(res), 0644)
		})
	}
}

func processAdminPanel() {
	if _, err := os.Stat("../locgame-admin-panel"); !os.IsNotExist(err) {
		if err := runCommand("cp -r pkg/dto/. ../locgame-admin-panel/server/dto"); err != nil {
			log.Fatal(err)
		}

		_ = os.Remove("../locgame-admin-panel/server/dto/handlers.go")

		_ = filepath.Walk("../locgame-admin-panel/server/dto", func(path string, info fs.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}

			data, err := ioutil.ReadFile(path)
			if err != nil {
				panic(err)
			}

			res := strings.ReplaceAll(string(data), "locgame-mini-server/pkg/dto", "locgame-admin-panel/dto")

			return ioutil.WriteFile(path, []byte(res), 0644)
		})

		processAdminPanelFrontend()
	}
}

func processAdminPanelFrontend() {
	if err := runCommand("cp -r ../locgame-admin-panel/server/dto ../locgame-admin-panel/app"); err != nil {
		log.Fatal(err)
	}

	_ = os.RemoveAll("../locgame-admin-panel/app/dto/errors")

	_ = filepath.Walk("../locgame-admin-panel/app/dto", func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if path == "../locgame-admin-panel/app/dto/resources/extensions.go" {
			return nil
		}

		if strings.Contains(path, "extensions.go") {
			_ = os.Remove(path)
			return nil
		}

		data, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}

		data = removeProtobufDependency(data)

		res := strings.ReplaceAll(string(data), "locgame-admin-panel/dto", "locgame-admin-panel-frontend/dto")

		return ioutil.WriteFile(path, []byte(res), 0644)
	})
}

var matchStringMethod = regexp.MustCompile(`(?m)func .*$\n(.+\n)+}`)
var matchDeprecatedString = regexp.MustCompile(`// Deprecated: .*`)
var matchDescriptor = regexp.MustCompile(`(?m)var File_.*_proto .*$\n(.*\n)*`)
var matchProtoMessage = regexp.MustCompile(`func .* ProtoMessage\(\) {}`)
var matchProtoTag = regexp.MustCompile(`(?mU)protobuf:(.*)\s`)
var matchBsonTag = regexp.MustCompile(`(?mU)bson:\"(.*)\"`)
var matchProtoExt = regexp.MustCompile(`(?m)var file_.*_proto_extTypes(.*\n)*`)

func removeProtobufDependency(data []byte) []byte {
	data = matchStringMethod.ReplaceAll(data, nil)
	data = matchDescriptor.ReplaceAll(data, nil)
	data = matchProtoMessage.ReplaceAll(data, nil)
	data = matchDeprecatedString.ReplaceAll(data, nil)
	data = matchProtoTag.ReplaceAll(data, nil)
	data = matchBsonTag.ReplaceAll(data, nil)
	data = matchProtoExt.ReplaceAll(data, nil)
	result := strings.Replace(string(data), `const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)`, "", 1)
	result = strings.ReplaceAll(result, "\tstate         protoimpl.MessageState", "")
	result = strings.ReplaceAll(result, "\tsizeCache     protoimpl.SizeCache", "")
	result = strings.ReplaceAll(result, "\tunknownFields protoimpl.UnknownFields", "")

	result = strings.Replace(result, "protoreflect \"google.golang.org/protobuf/reflect/protoreflect\"", "", 1)
	result = strings.Replace(result, "protoimpl \"google.golang.org/protobuf/runtime/protoimpl\"", "", 1)
	result = strings.Replace(result, "_ \"google.golang.org/protobuf/types/descriptorpb\"", "", 1)
	result = strings.Replace(result, "descriptorpb \"google.golang.org/protobuf/types/descriptorpb\"", "", 1)
	result = strings.Replace(result, "reflect \"reflect\"", "", 1)
	result = strings.Replace(result, "sync \"sync\"", "", 1)
	data = []byte(result)
	data, _ = format.Source(data)
	return data
}

func runCommand(command string) error {
	args := strings.Split(command, " ")
	cmd := exec.Command(args[0], args[1:]...)

	outPipe, _ := cmd.StdoutPipe()
	errPipe, _ := cmd.StderrPipe()

	cmdReader := io.MultiReader(outPipe, errPipe)
	scanner := bufio.NewScanner(cmdReader)
	done := make(chan bool)
	go func() {
		for scanner.Scan() {
			fmt.Printf(scanner.Text() + "\n")
		}
		done <- true
	}()

	_ = cmd.Start()
	<-done
	return cmd.Wait()
}

// renameBsonID renames the tag for ObjectID to _id
func renameBsonID(data string) string {
	var re = regexp.MustCompile(`(?sU)(ID.*\*base\.ObjectID.*bson:)"(id)(,omitempty")`)
	return re.ReplaceAllString(data, "$1\"_id$3")
}

// removeOmitEmptyFromBooleans remove `omitempty` from bson tag, if this is boolean value
func removeOmitEmptyFromBooleans(data string) string {
	var re = regexp.MustCompile(`(?sU)(\sbool\s+.protobuf.*bson:.*)(,omitempty)"`)
	return re.ReplaceAllString(data, "$1\"")
}

func removeOmitEmptyFromField(data, fieldName string) string {
	// This regex pattern targets a Go struct field declaration by name with `omitempty` in its bson tag.
	// Adjust the pattern to match the specific syntax of your generated Go code.
	pattern := fmt.Sprintf(`(?sU)(\s%s\s+.*protobuf.*bson:.*)(,omitempty)"`, regexp.QuoteMeta(fieldName))
	var re = regexp.MustCompile(pattern)
	return re.ReplaceAllString(data, "$1\"")
}

// add OmitNil to field
func addOmitNilToField(data, fieldName string) string {
	pattern := fmt.Sprintf(`(?sU)(\s%s\s+.*protobuf.*bson:")(.*)"`, regexp.QuoteMeta(fieldName))
	var re = regexp.MustCompile(pattern)
	return re.ReplaceAllString(data, "$1$2,omitnil\"")
}

// addYamlAndBsonTag Adding an attribute for YAML and Bson
func addYamlAndBsonTag(data string) string {
	var re = regexp.MustCompile(`(?sU)json:"(.*)"`)
	return re.ReplaceAllStringFunc(data, func(s string) string {
		re := regexp.MustCompile("\".*\"")
		value := strings.ReplaceAll(re.FindString(s), "\"", "")
		return fmt.Sprintf("json:\"%s\" yaml:\"%s\" bson:\"%s\"", value, value, toSnakeCase(value))
	})
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z\\d])([A-Z])")

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
