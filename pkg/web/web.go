package web

import (
	"dbbook/pkg/flags"
	"log"
	"net/http"
	"path/filepath"
)

func Serve() {

	appPath, _ := filepath.Abs(".")

	fileServer := http.FileServer(http.Dir(appPath + "/web"))

	log.Printf("Running at http://0.0.0.0:%s", flags.Port)

	e := http.ListenAndServe(":"+flags.Port, fileServer)

	if e != nil {
		log.Fatalf("Failed to start service: %s", e)
	}

}
