package tools

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func Combine(v1 uint32, v2 uint32) uint32 {
	return (v2 & 0xffff) | ((v1 & 0xffff) << 16)
}

func Separate(v uint32) (uint32, uint32) {
	v1 := v >> 16
	v2 := v & 0xffff

	return v1, v2
}

func GetStopScriptName() string {
	pathname := os.Args[1]
	paths, name := filepath.Split(pathname)

	return fmt.Sprintf("%sstop_%s.sh", paths, name)
}

func IsFileExist(filename string) bool {
	_, err := os.Stat(filename)

	return !os.IsNotExist(err)
}

func AllowExec() bool {
	scriptname := GetStopScriptName()

	return !IsFileExist(scriptname)
}

func BuildStopScript() {
	if runtime.GOOS == "windows" {
		return
	}

	scriptname := GetStopScriptName()
	f, err := os.Create(scriptname)
	if err != nil {
		fmt.Println("File is exist !")
	} else {
		pid := os.Getpid()
		script := fmt.Sprintf("kill %d", pid)
		f.WriteString(script)
		f.Sync()
		f.Close()
	}
}

func RemoveStopScript() {
	if runtime.GOOS == "windows" {
		return
	}

	scriptname := GetStopScriptName()

	os.Remove(scriptname)
}