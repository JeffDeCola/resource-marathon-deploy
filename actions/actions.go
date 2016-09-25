// resource-marathon-deploy actions.go

package actions

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/ckaznocha/marathon-resource/marathonclient"
	gomarathon "github.com/gambol99/go-marathon"
)

type (
	version struct {
		Ref string `json:"ref"`
	}
	// InputJSON ...
	InputJSON struct {
		Params  map[string]string `json:"params"`
		Source  map[string]string `json:"source"`
		Version version           `json:"version"`
	}
	metadata struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}
	checkOutputJSON []version
	inOutputJSON    struct {
		Version  version    `json:"version"`
		Metadata []metadata `json:"metadata"`
	}
	outOutputJSON inOutputJSON
)

func getversions() []string {

	return []string{
		"123",
		"3de",
		"456",
		"336",
	}

}

// Check will return the versions available.
func Check(input InputJSON, logger *log.Logger) (checkOutputJSON, error) {

	// PARSE THE JSON FILE /tmp/input.json
	source1, ok := input.Source["source1"]
	if !ok {
		return checkOutputJSON{}, errors.New("Source1 not set")
	}
	source2, ok := input.Source["source2"]
	if !ok {
		return checkOutputJSON{}, errors.New("Source2 not set")
	}
	var ref = input.Version.Ref
	logger.Print("source are")
	logger.Print(source1, source2)
	logger.Print("ref is")
	logger.Print(ref)

	// CHECK (THE RESOURCE VERSION(s)) AND OUTPUT *****************************************************
	// Mimic a fetch versions(s) and output the following versions for IN.

	var output = checkOutputJSON{}
	for _, ver := range getversions() {
		output = append(output, version{Ref: ver})
	}

	return output, nil

}

// IN will fetch something and place in the working directory.
func In(input InputJSON, logger *log.Logger) (inOutputJSON, error) {

	// PARSE THE JSON FILE /tmp/input.json
	source1, ok := input.Source["source1"]
	if !ok {
		return inOutputJSON{}, errors.New("source1 not set")
	}
	source2, ok := input.Source["source2"]
	if !ok {
		return inOutputJSON{}, errors.New("source2 not set")
	}
	param1, ok := input.Params["param1"]
	//if !ok {
	//	return inOutputJSON{}, errors.New("param1 not set")
	//}
	param2, ok := input.Params["param2"]
	//if !ok {
	//	return inOutputJSON{}, errors.New("param2 not set")
	//}
	var ref = input.Version.Ref
	logger.Print("source are")
	logger.Print(source1, source2)
	logger.Print("params are")
	logger.Print(param1, param2)
	logger.Print("ref is")
	logger.Print(ref)

	// SOME METATDATA YOU CAN USE
	logger.Print("BUILD_ID = ", os.Getenv("BUILD_ID"))
	logger.Print("BUILD_NAME = ", os.Getenv("BUILD_NAME"))
	logger.Print("BUILD_JOB_NAME = ", os.Getenv("BUILD_JOB_NAME"))
	logger.Print("BUILD_PIPELINE_NAME = ", os.Getenv("BUILD_PIPELINE_NAME"))
	logger.Print("ATC_EXTERNAL_URL = ", os.Getenv("ATC_EXTERNAL_URL"))

	// IN (FETCH THE RESOURCE) *************************************************************************
	// Mimic a fetch and place a fetched.json file in the working directory that contains the following.

	jsonfile := "Hi everone, This is a file I made"

	// Create a fake fetched file
	filewrite, err := os.Create("fetch.json")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer filewrite.Close()
	fmt.Fprintf(filewrite, jsonfile)

	//ls -lat $WORKING_DIR
	logger.Print("List whats in the directory:")
	files, _ := ioutil.ReadDir("./")
	for _, f := range files {
		logger.Print(f.Name())
	}

	// Cat the file
	logger.Print("Cat fetch.json")
	file, err := os.Open("fetch.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	bb, err := ioutil.ReadAll(file)
	logger.Print(string(bb))

	var monkeyname = "Larry"

	// OUTPUT **************************************************************************************
	output := inOutputJSON{
		Version: version{Ref: ref},
		Metadata: []metadata{
			{Name: "nameofmonkey", Value: monkeyname},
			{Name: "author", Value: "Jeff DeCola"},
		},
	}

	return output, nil

}

// Out ...
func Out(input InputJSON, logger *log.Logger) (outOutputJSON, error) {

	// PARSE THE JSON FILE /tmp/input.json
	source1, ok := input.Source["source1"]
	if !ok {
		return outOutputJSON{}, errors.New("source1 not set")
	}
	source2, ok := input.Source["source2"]
	if !ok {
		return outOutputJSON{}, errors.New("source2 not set")
	}
	param1, ok := input.Params["param1"]
	if !ok {
		return outOutputJSON{}, errors.New("param1 not set")
	}
	param2, ok := input.Params["param2"]
	if !ok {
		return outOutputJSON{}, errors.New("param2 not set")
	}
	var ref = input.Version.Ref
	logger.Print("source are")
	logger.Print(source1, source2)
	logger.Print("params are")
	logger.Print(param1, param2)
	logger.Print("ref is")
	logger.Print(ref)

	// SOME METATDATA YOU CAN USE
	logger.Print("BUILD_ID = ", os.Getenv("BUILD_ID"))
	logger.Print("BUILD_NAME = ", os.Getenv("BUILD_NAME"))
	logger.Print("BUILD_JOB_NAME = ", os.Getenv("BUILD_JOB_NAME"))
	logger.Print("BUILD_PIPELINE_NAME = ", os.Getenv("BUILD_PIPELINE_NAME"))
	logger.Print("ATC_EXTERNAL_URL = ", os.Getenv("ATC_EXTERNAL_URL"))

	// OUT (UPDATE THE RESOURCE) *************************************************************************

	uri, err := url.Parse("http://10.141.141.10:8080")
	logger.Print("uri is: ", uri)
	if err != nil {
		log.Fatal("Malformed URI", err)
	}

	apiclient := marathonclient.NewMarathoner(http.DefaultClient, uri)

	mem := 16.0
	forcepull := true

	did, err := apiclient.UpdateApp(gomarathon.Application{
		ID:   "jeffisawesome5",
		CPUs: 0.1,
		Mem:  &mem,
		Container: &gomarathon.Container{
			Type: "DOCKER",
			Docker: &gomarathon.Docker{
				ForcePullImage: &forcepull,
				Image:          "jeffdecola/hello-go",
			},
		},
	})
	if err != nil {
		log.Fatal("Failed UpdateApp", err)
	}

	logger.Print("did is: ", did)
	fmt.Fprintln(os.Stderr, did)

	var monkeyname = "Henry"
	// ???????????????????????????????????????????????????????????????????????
	// ??????????????????????????????????????????????????????????????????????
	ref = "123"

	// OUTPUT **************************************************************************************
	output := outOutputJSON{
		Version: version{Ref: ref},
		Metadata: []metadata{
			{Name: "nameofmonkey", Value: monkeyname},
			{Name: "author", Value: "Jeff DeCola"},
		},
	}

	return output, nil

}
