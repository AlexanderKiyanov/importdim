package helpers

import (
	"errors"
	"fmt"
	"log"
)

func GetOptions(arguments []string, currentDir string) (string, error) {

	if len(arguments) < 1 {
		return "", errors.New("error: you must pass at least one argument")

	} else if arguments[0] == "all" && len(arguments) == 1 {
		fmt.Println("\n\neStart xecute: importdim all")
		return currentDir, nil

	} else if len(arguments) > 1 {

		path := arguments[1]
		for i := 2; i <= len(arguments[1:]); i++ {
			path = fmt.Sprint(path, arguments[i])
		}

		// check it's not a full path and make it
		path, err := MakeFullPath(path, currentDir)
		if err != nil {
			log.Fatal(err)
		}

		if arguments[0] == "dir" && IsDirectory(path) {
			fmt.Printf("\n\nStart execute: importdim dir %s\n", path)
			return path, nil

		} else if arguments[0] == "file" && IsFile(path) {
			fmt.Printf("\n\nStart execute: importdim %s\n", path)
			return path, nil

		} else {
			log.Fatalf("error: second argument must be a directory or a file %s", path)
		}
	}

	return "", errors.New("error: unrecognized arguments")
}
