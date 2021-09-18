package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	path, e := exec.LookPath(os.Args[0])
	if e != nil {
		log.Fatalln(e)
	}
	path, e = filepath.Abs(path)
	if e != nil {
		log.Fatalln(e)
	}
	dir := filepath.Dir(path)
	NewGeneration(dir).Serve()
}
