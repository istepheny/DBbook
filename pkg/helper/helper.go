package helper

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func AppPath() string {
	appPath, _ := filepath.Abs(".")

	return appPath
}

func ConfigFilePath() string {
	return strings.Join([]string{AppPath(), "database.json"}, string(os.PathSeparator))
}

func TemplatePath() string {
	return strings.Join([]string{AppPath(), "web", "template", ""}, string(os.PathSeparator))
}

func BookPath() string {
	return strings.Join([]string{AppPath(), "web", "dbbook", ""}, string(os.PathSeparator))
}

func Mkdir(dir string) {
	if _, e := os.Stat(dir); os.IsNotExist(e) {
		e := os.MkdirAll(dir, 0755)
		if e != nil {
			log.Fatal(e)
		}
	}
}
