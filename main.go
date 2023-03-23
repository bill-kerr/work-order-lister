package main

import (
	"errors"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/bill-kerr/work-order-lister/excel"
	"github.com/bill-kerr/work-order-lister/indexer"
)

func main() {
	workOrders := indexer.Index()
	excel.CreateFile(workOrders)
}

func getPrefixesFromFiles(files []fs.FileInfo) ([]string, error) {
	for _, f := range files {
		if f.Name() == "prefix.txt" {
			path, _ := filepath.Abs(filepath.Join(filepath.Dir(os.Args[0]), f.Name()))
			contents := readTextFile(path)
			return strings.Split(contents, "\n"), nil
		}
	}
	return nil, errors.New("No prefix.txt present in root directory. Create a prefix.txt file and add the prefixes you would like indexed.")
}

func hasPrefix(name string, prefixes []string) bool {
	for _, p := range prefixes {
		if strings.HasPrefix(name, p) {
			return true
		}
	}
	return false
}

func getDescription(directory string) string {
	path, _ := filepath.Abs(filepath.Join(filepath.Dir(os.Args[0]), directory))
	files, err := ioutil.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".txt") {
			path, _ = filepath.Abs(filepath.Join(filepath.Dir(os.Args[0]), directory, f.Name()))
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

func writeToFile(workOrders map[string]string) {
	fileContents := ""

	for key, elem := range workOrders {
		fileContents += key + " - " + elem + "\n"
	}

	err := ioutil.WriteFile("work_orders.txt", []byte(fileContents), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
