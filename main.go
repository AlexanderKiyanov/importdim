package main

import (
	"fmt"
	"importdim/helpers"
	_ "importdim/helpers"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	path, err := helpers.GetOptions(os.Args[1:], currentDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fileList, err := helpers.FindFilesByPath(path, "*.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	params, err := helpers.CheckFilesFormat(fileList)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	utility := "C:\\Middleware\\user_projects\\epmsystem1\\Planning\\planning1\\OutlineLoad.cmd"
	var dim, app, importFile string

	logFile := fmt.Sprintf("%s\\epmpi_%s.log", currentDir, time.Now().Format("20060102_150405"))
	fmt.Printf("\n\nLog file is located: %s\n\n", logFile)
	logFile = fmt.Sprintf("%q", logFile)

	passFile := "C:\\epmpi\\enc_pass"
	passFile = fmt.Sprintf("%q", passFile)

	// Elevate Privileges to Administrator level
	if !helpers.AmIAdmin() {
		helpers.ElevateAsAdmin()
	}

	fmt.Println("\nThe following files were found to be imported:")

	for i := range params {
		fmt.Printf("%d: %s\n\n", i, params[i])
		dim = strings.TrimSpace(params[i][0])
		app = strings.TrimSpace(params[i][1])
		importFile = fmt.Sprintf("%q", params[i][2])

		helpers.StartImport(utility, dim, app, passFile, importFile, logFile)
	}
	//time.Sleep(1*time.Second)

}
