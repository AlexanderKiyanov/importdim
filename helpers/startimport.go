package helpers

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

func StartImport(utility string, dim string, app string, pass string, importFile string, logFile string) {
	cmdAllArg := fmt.Sprintf("%s -f:%s /U:admin /A:%s /D:%s /I:%s >> %s", utility, pass, app, dim, importFile, logFile)
	cmd := exec.Command("powershell.exe", "/C", cmdAllArg)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("error: cmd execution failed with %s and output is:\n %s\n", err, output)
		time.Sleep(14 * time.Second)
		log.Fatalf("error: cmd execution failed with %s and output is:\n %s\n", err, output)
	}

	fmt.Println("Import was successfully executed.")

}
