package web

import (
	"dbbook/pkg/flags"
	"dbbook/pkg/helper"
	"log"
	"net/http"
	"path/filepath"
)

func Serve() {

	appPath, _ := filepath.Abs(".")

	fileServer := http.FileServer(http.Dir(appPath + "/web"))

	helper.RunningLog()

	e := http.ListenAndServe(":"+flags.Port, fileServer)

	if e != nil {
		log.Fatalf("Failed to start service: %s", e)
	}
}
