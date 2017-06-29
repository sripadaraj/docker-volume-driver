package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/docker/go-plugins-helpers/volume"
)

const volumedriverID = "_volumedriver"

var (
	defaultDir  = filepath.Join(volume.DefaultDockerRootDirectory, volumedriverID)
	serversList = flag.String("servers", "", "List of glusterfs servers")
	restAddress = flag.String("rest", "", "URL to volumedriverrest api")
	gfsBase     = flag.String("gfs-base", "/mnt/gfs", "Base directory where volumes are created in the cluster")
	root        = flag.String("root", defaultDir, "volumes root directory")
)

func main() {
	var Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()
	if len(*serversList) == 0 {
		Usage()
		os.Exit(1)
	}

	servers := strings.Split(*serversList, ":")

	d := newGlusterfsDriver(*root, *restAddress, *gfsBase, servers)
	h := volume.NewHandler(d)
	fmt.Println(h.ServeUnix("root", "volumedriver"))
}
