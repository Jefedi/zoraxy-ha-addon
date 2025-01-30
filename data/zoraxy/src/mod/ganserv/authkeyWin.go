//go:build windows
// +build windows

package ganserv

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"golang.org/x/sys/windows"
	"imuslab.com/zoraxy/mod/utils"
)

// Use admin permission to read auth token on Windows
func readAuthTokenAsAdmin() (string, error) {
	//Check if the previous startup already extracted the authkey
	if utils.FileExists("./conf/authtoken.secret") {
		authKey, err := os.ReadFile("./conf/authtoken.secret")
		if err == nil {
			return strings.TrimSpace(string(authKey)), nil
		}
	}

	verb := "runas"
	exe := "cmd.exe"
	cwd, _ := os.Getwd()

	output, _ := filepath.Abs(filepath.Join("./conf/", "authtoken.secret"))
	os.WriteFile(output, []byte(""), 0775)
	args := fmt.Sprintf("/C type \"C:\\ProgramData\\ZeroTier\\One\\authtoken.secret\" > \"" + output + "\"")

	verbPtr, _ := syscall.UTF16PtrFromString(verb)
	exePtr, _ := syscall.UTF16PtrFromString(exe)
	cwdPtr, _ := syscall.UTF16PtrFromString(cwd)
	argPtr, _ := syscall.UTF16PtrFromString(args)

	var showCmd int32 = 1 //SW_NORMAL

	err := windows.ShellExecute(0, verbPtr, exePtr, argPtr, cwdPtr, showCmd)
	if err != nil {
		return "", err
	}

	log.Println("Please click agree to allow access to ZeroTier authtoken from ProgramData")
	retry := 0
	time.Sleep(3 * time.Second)
	for !utils.FileExists("./conf/authtoken.secret") && retry < 10 {
		time.Sleep(3 * time.Second)
		log.Println("Waiting for ZeroTier authtoken extraction...")
		retry++
	}

	authKey, err := os.ReadFile("./conf/authtoken.secret")
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(authKey)), nil
}

// Check if admin on Windows
func isAdmin() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	if err != nil {
		return false
	}
	return true
}
