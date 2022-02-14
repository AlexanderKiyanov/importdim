package helpers

import (
	"fmt"
	"golang.org/x/sys/windows"
	"log"
	"os"
	"strings"
	"syscall"
)

var (
	mod             = windows.NewLazyDLL("user32.dll")
	procCloseWindow = mod.NewProc("CloseWindow")
)

func AmIAdmin() bool {
	// if not elevated, relaunch by shell execute with runas verb set
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	if err != nil {
		fmt.Println("Am I admin? no\n\n")
		return false
	}
	fmt.Print("Am I admin? yes\n\n")

	//if err := root(os.Args[1:]); err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	return true
}

// Switch to administrator
func ElevateAsAdmin() {
	verb := "runas"
	exe, _ := os.Executable()
	cwd, _ := os.Getwd()
	args := strings.Join(os.Args[1:], " ")

	fmt.Println("Arguments: ", os.Args[1:])

	verbPtr, _ := syscall.UTF16PtrFromString(verb)
	exePtr, _ := syscall.UTF16PtrFromString(exe)
	cwdPtr, _ := syscall.UTF16PtrFromString(cwd)
	argPtr, _ := syscall.UTF16PtrFromString(args)

	var showCmd int32 = 1 //SW_NORMAL

	hwnd := getWindow("GetForegroundWindow")

	err := windows.ShellExecute(0, verbPtr, exePtr, argPtr, cwdPtr, showCmd)
	if err != nil {
		fmt.Println(err)
	} else {
		if hwnd != 0 {
			_, _, err := procCloseWindow.Call(uintptr(hwnd))
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func getWindow(funcName string) uintptr {
	mod := windows.NewLazyDLL("user32.dll")
	proc := mod.NewProc(funcName)
	hwnd, _, _ := proc.Call()
	return hwnd
}
