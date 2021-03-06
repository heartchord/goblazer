package goblazer

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var appRootPath string

func init() {
	SetAppRootPath("")
}

// IsFileExisted : a function wrapper to check file existed or not
func IsFileExisted(szFilePath string) bool {
	_, err := os.Stat(szFilePath)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// IsFileDirectory :
func IsFileDirectory(szFilePath string) bool {
	fi, err := os.Stat(szFilePath)
	if os.IsNotExist(err) {
		fmt.Printf("File not existed - %s\n", szFilePath)
		return false
	}
	return fi.IsDir()
}

// FormatFilePath : a function wrapper to format path to "a/b/c/d"
func FormatFilePath(path *string) {
	*path = strings.Replace(*path, "\\", "/", -1)
	*path = strings.Trim(*path, "/ ")
}

// GetPathName : a function wrapper to get name from a path. eg: if path is "a/b/c/d/", it will return "d" string
func GetPathName(path string) string {
	FormatFilePath(&path)
	idx := strings.LastIndex(path, "/")
	if idx < 0 {
		idx = 0
	}
	return path[idx:len(path)]
}

// SetAppRootPath : a function wrapper to set current application base path, if parameter 'path' is nil or "",
// 					it will fetch base path by current executable program.
func SetAppRootPath(path string) bool {
	var err error

	if path != "" {
		if FormatFilePath(&path); IsFileExisted(path) {
			appRootPath = path
			return true
		}
	}

	if path, err = filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
		appRootPath = ""
		return false
	}

	if FormatFilePath(&path); IsFileExisted(path) {
		appRootPath = path
		return true
	}

	return false
}

// GetAppRootPath : a function wrapper to get current application path
func GetAppRootPath() string {
	return appRootPath
}
