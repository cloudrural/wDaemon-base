package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var docroot = flag.String("root", "./html", "Document Root")

func main() {
	//docroot := flag.String("docroot", "./html", "Document Root")
	createDocRoot()
	fileServer := http.FileServer(http.Dir(*docroot))
	http.Handle("/", fileServer)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

//func createDocRoot() {
//	err := os.Mkdir("/var/www/wDaemon-static", 0755)
//	if err != nil {
//		log.Fatal(err)
//	}
//}

func createDocRoot() {
	if err := ensureDir(*docroot); err != nil {
		fmt.Println("Directory creation failed with error: " + err.Error())
		os.Exit(1)
	}
	// Proceed forward
}

func ensureDir(dirName string) error {
	err := os.Mkdir(dirName, 0755)
	if err == nil {
		return nil
	}
	if os.IsExist(err) {
		// check that the existing path is a directory
		info, err := os.Stat(dirName)
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return errors.New("path exists but is not a directory")
		}
		return nil
	}
	return err
}
