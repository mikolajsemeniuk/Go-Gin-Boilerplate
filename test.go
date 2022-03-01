package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func AddFileError(name string) string {
	return ""
}

func ListAllFiles(path string) (int, int, []string) {
	filesAmount, foldersAmount := 0, 0
	filesErrors := []string{}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			files, folders, errors := ListAllFiles(path + "/" + file.Name())
			filesAmount += files
			foldersAmount += folders + 1
			filesErrors = append(filesErrors, errors...)
		}

		if !file.IsDir() {
			filesAmount += 1
			// FIXME:
			func() {
				if file.Name() == "hi.go" {
					filesErrors = append(filesErrors, path+"/"+file.Name())
				}
			}()
		}
	}

	return filesAmount, foldersAmount, filesErrors
}

// func main() {
// 	filesAmount, foldersAmount, errors := ListAllFiles(".")
// 	fmt.Println("files scanned: ", filesAmount, ", folders scanned: ", foldersAmount, "\nok files: ", filesAmount-len(errors), ", files with errors: ", len(errors))
// 	for _, name := range errors {
// 		fmt.Println(name)
// 	}
// }
