package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	usage = `Usage: 
./jsonFormatter $filepath_to_json_file`
)

var (
	filePath       []string               // path of json file
	sourceFileStat os.FileInfo            // stat info about file
	fileContent    []byte                 // the file content
	dat            map[string]interface{} // thing to place json unmarshal'd stuff in
	newPath        string                 // destination path to backup original file
	jsonPretty     []byte                 // marshal'd json
)

func printUsage() {
	fmt.Println(usage)
	os.Exit(1)
}

func main() {

	filePath = os.Args[1:]
	if len(filePath) == 0 {
		printUsage()
	}

	for _, jsonFile := range filePath {

		// Validate the file.
		sourceFileStat, err := os.Stat(jsonFile)
		if err != nil {
			panic(err)
		}

		if !sourceFileStat.Mode().IsRegular() {
			fmt.Errorf("%s is not a regular file.", jsonFile)
		}

		fileContent, err = ioutil.ReadFile(jsonFile)
		if err != nil {
			panic(err)
		}

		if len(fileContent) == 0 {
			panic(fmt.Errorf("%s is empty.", jsonFile))
		}

		// Unmarshal file
		if err := json.Unmarshal(fileContent, &dat); err != nil {
			fmt.Println("Hey dawg, is that even json??")
			panic(err)
		}

		// move original file to /tmp
		newPath = filepath.Join("/tmp", jsonFile)
		os.Rename(jsonFile, newPath)

		// Write the file (json formatted)
		jsonPretty, err = json.MarshalIndent(dat, "", "  ")
		if err != nil {
			panic(err)
		}

		// TODO Change perms to match original file perms.
		if err := ioutil.WriteFile(jsonFile, jsonPretty, 0644); err != nil {
			panic(err)
		}
	}

}
