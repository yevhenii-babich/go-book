package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func readDir() {
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fn := file.Name()
		if len(fn) < 8 {
			fn = fn + "\t"
		}
		var sInfo string
		info, err := file.Info()
		if err == nil {
			sInfo = fmt.Sprintf("%v\t %v %v", info.Size(), info.Mode(), info.ModTime())
		}
		fmt.Println(fn, "\t", sInfo)
	}
}

func main() {
	readDir() // read current directory
	path := "./newdir"
	time.Sleep(1 * time.Second)
	println("\n\nCreate directory:", path)
	if err := os.Mkdir(path, 0755); err != nil {
		log.Fatal(err)
	}
	readDir()
	time.Sleep(1 * time.Second)
	println("\n\nRemove directory:", path)
	if err := os.RemoveAll(path); err != nil {
		log.Fatal(err)
	}
	readDir()
}
