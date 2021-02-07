package common

import (
	"os"
	"strings"
)

func GetCurrentDir() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}

	return strings.Replace(dir, "\\", "/", -1), nil
}

func GetVersion() string {
	dir, err := GetCurrentDir()
	if err != nil {
		return "no version"
	}

}
