	package main

import (
	"flag"
	"fmt"
	"os""
	"path/filepath"
	"strings"
	"log"

	"github.com/docker/go-plugins-helpers/volume"
)

var (
	defaultDir  =filepath.Join(volume.DefaultDockerRootDirectory)
	serversList =flag.String("servers","","List of servers")
	restAddress =flag.String("rest","","URL to docker-volume-driver rest api ")
//	gfsBase =flag.String("root",defaultDir,"
)
func main() {
// create the server that will serve our Unix socket for the Docker daemon. 
	driver := NewExampleDriver()
	handler := volume.NewHandler(driver)
	if err := handler.ServeUnix("dvd", 0); err != nil {
		log.Fatalf("Error %v", err)
	}
// The empty for loop is here so that the main function becomes blocking since the server will go in a separate goroutine.
	for {

	}
}
