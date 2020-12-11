package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const defaultV string = "//default//"

func main() {
	// flag vars
	var (
		dirLocation string
		oldExt      string
		newExt      string
	)

	// flags
	flag.StringVar(&dirLocation, "dirLocation", defaultV, "Location of the dir")
	flag.StringVar(&oldExt, "oldExt", defaultV, "old extension without .")
	flag.StringVar(&newExt, "newExt", defaultV, "new extension without .")
	flag.Parse()

	if val := isDefault(dirLocation, oldExt, newExt); val == true {
		log.Fatalf("flags are not defined")
	}
	if oldExt == newExt {
		log.Println("Noting to change")
		os.Exit(0)
	}

	oldExt = "." + oldExt
	newExt = "." + newExt

	files, err := getFiles(dirLocation)
	errorHandler(err)

	var changeCounter = 0

	for i := 0; i < len(files); i++ {
		file := files[i]
		ext := filepath.Ext(file.Name())

		if file.IsDir() != true && ext == oldExt {
			loc := dirLocation + "/" + file.Name()
			err := os.Rename(loc, strings.Replace(loc, oldExt, newExt, -1))
			errorHandler(err)
			changeCounter++
			log.Printf("Changed %s", file.Name())
		}
	}
	if changeCounter != 0 {
		log.Printf("Changed %d file(s)", changeCounter)
	} else {
		log.Printf("No file with %s extension found", oldExt)
	}

}

func getFiles(location string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(location)
	return files, err

}
func errorHandler(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
func isDefault(flags ...string) bool {
	for i := 0; i < len(flags); i++ {
		if flags[i] == defaultV {
			return true
		}
	}
	return false
}
