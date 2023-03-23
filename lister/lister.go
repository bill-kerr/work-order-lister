package lister

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func List(path string) map[string]string {
	workOrders := map[string]string{}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.IsDir() {
			description := getDescription(path, f.Name())
			workOrders[f.Name()] = description
		}
	}

	writeToTextFile(path, workOrders)
	return workOrders
}

func getDescription(rootPath string, directory string) string {
	path, _ := filepath.Abs(filepath.Join(rootPath, directory))
	files, err := ioutil.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".txt") {
			path, _ = filepath.Abs(filepath.Join(rootPath, directory, f.Name()))
			return readTextFile(path)
		}
	}

	return "\n"
}

func readTextFile(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	return getFileContents(file)
}

func getFileContents(file *os.File) string {
	fileBytes, err := ioutil.ReadAll(file)

	if err != nil {
		log.Fatal(err)
	}

	return parseBytes(fileBytes)
}

func parseBytes(fileBytes []byte) string {
	description := string(fileBytes)
	description = strings.TrimSuffix(description, "\n")
	description = strings.Replace(description, "\r\n", ", ", -1)
	return description
}

func writeToTextFile(rootPath string, workOrders map[string]string) {
	fileContents := ""

	for key, elem := range workOrders {
		fileContents += key + " - " + elem + "\n"
	}

	err := ioutil.WriteFile(filepath.Join(rootPath, "work_orders.txt"), []byte(fileContents), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
