package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	workOrders := map[string]string{}

	path, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if strings.HasPrefix(f.Name(), "GR4_") && f.IsDir() {
			description := getDescription(f.Name())
			workOrders[f.Name()] = description
		}
	}

	writeToFile(workOrders)
}

func getDescription(directory string) string {
	path, _ := filepath.Abs(filepath.Dir(os.Args[0]) + "\\" + directory + "\\")
	files, err := ioutil.ReadDir(path)
	log.Print(files)

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".txt") {
			path, _ = filepath.Abs(filepath.Dir(os.Args[0]) + "\\" + directory + "\\" + f.Name())
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
	filebytes, err := ioutil.ReadAll(file)

	if err != nil {
		log.Fatal(err)
	}

	return parseBytes(filebytes)
}

func parseBytes(filebytes []byte) string {
	description := string(filebytes)
	description = strings.TrimSuffix(description, "\n")
	description = strings.Replace(description, "\r\n", ", ", -1)
	return description
}

func writeToFile(workOrders map[string]string) {
	fileContents := ""

	for key, elem := range workOrders {
		fileContents += key + " - " + elem
	}

	err := ioutil.WriteFile("work_orders.txt", []byte(fileContents), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
