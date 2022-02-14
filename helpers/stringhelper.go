package helpers

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func SubStrFromNum(word string, start int, fromEnd int) string {
	runes := []rune(word)
	lenn := len(runes)

	return string(runes[start : lenn-fromEnd])
}

func MakeFullPath(path string, currDir string) (string, error) {
	runes := []rune(path)
	firstSymbol := string(runes[0:1])

	if IsDirectory(path) || IsFile(path) {
		return path, nil
	} else if firstSymbol != "\\" {
		path = fmt.Sprint(currDir, "\\", path)
	} else {
		path = fmt.Sprint(currDir, path)
	}

	if IsDirectory(path) || IsFile(path) {
		return path, nil
	}
	return path, errors.New("error: unrecognized arguments")
}

func IsDirectory(path string) bool {

	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	if fileInfo.IsDir() {
		return true
	} else {
		return false
	}
}

func IsFile(path string) bool {

	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	if fileInfo.Mode().IsRegular() {
		return true
	} else {
		return false
	}
}

func DelParenthesis(word string) string {
	runes := []rune(word)
	l := len(runes)

	return string(runes[10 : l-1])
}

func ConvertCubeName(cube string) (string, error) {
	switch cube {
	case "Console":
		return "Console", nil
	case "Feed":
		return "Feed", nil
	case "MeatPl", "MeatProc":
		return "Meat", nil
	case "MEZ":
		return "MEZ", nil
	case "Plant":
		return "Plant", nil
	case "Pork":
		return "Pork", nil
	case "Plan1", "Plan2", "Plan3":
		return "Poultry", nil
	case "TradeCo":
		return "TradeCo", nil
	case "TrkPL", "TrkSales":
		return "Turkey", nil
	}

	return "", errors.New("error: invalid cube name: " + cube)
}

func ReadCsvFile(filePath string) ([][]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		if err := f.Close(); err != nil {
			return nil, err
		}
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	if err := f.Close(); err != nil {
		return nil, err
	}

	return records, nil
}

func FindFilesByPath(path, pattern string) ([]string, error) {

	var matches []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func CheckFilesFormat(fileList []string) ([][]string, error) {

	var fileListWithParams [][]string

	for i := range fileList {
		word, err := ReadCsvFile(fileList[i])
		if err != nil {
			log.Fatalf("error: while reading csv file: %s\n%s", err, fileList[i])
		}

		dimName := SubStrFromNum(word[0][0], 1, 0)

		j := len(word[0]) - 1
		cubeName, err := ConvertCubeName(DelParenthesis(word[0][j]))
		if err != nil {
			log.Fatal("error: invalid cube name: " + cubeName)
		}

		oneSet := []string{dimName, cubeName, fileList[i]}
		fileListWithParams = append(fileListWithParams, oneSet)

	}

	return fileListWithParams, nil
}
