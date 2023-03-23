package indexer

import (
	"errors"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/bill-kerr/work-order-lister/lister"
)

var rootDir string = filepath.Dir(os.Args[0])

var ignoredFiles = []string{
	".git",
	"lister",
	"indexer",
	".vscode",
	"excel",
}

func isIgnoredFile(name string) bool {
	for _, ignored := range ignoredFiles {
		if name == ignored {
			return true
		}
	}

	return false
}

func Index() map[string]map[string]string {
	path, _ := filepath.Abs(rootDir)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	projectNames := []string{}
	projectWorkOrders := map[string]map[string]string{}

	for _, f := range files {
		if f.IsDir() && !isIgnoredFile(f.Name()) {
			projectNames = append(projectNames, f.Name())
			projectFolderPath, _ := filepath.Abs(filepath.Join(rootDir, f.Name()))

			workOrders := checkAndIndexProject(projectFolderPath)
			if err != nil {
				log.Fatal(err)
			}

			if len(workOrders) > 0 {
				projectWorkOrders[f.Name()] = workOrders
			}
		}
	}

	return projectWorkOrders
}

func checkAndIndexProject(path string) map[string]string {
	files, err := ioutil.ReadDir(path)

	workOrdersDirName, err := findWorkOrdersDir(files)
	if err != nil {
		log.Fatal(err)
	}

	workOrdersDir := filepath.Join(path, workOrdersDirName)
	return lister.List(workOrdersDir)
}

func findWorkOrdersDir(fileInfo []fs.FileInfo) (string, error) {
	for _, f := range fileInfo {
		if strings.ToLower(f.Name()) == "work orders" {
			return f.Name(), nil
		}
	}

	return "", errors.New("Could not find work order directory")
}
