package main

import (
	"flag"
	"fmt"
	//"github.com/pivotal-golang/bytefmt"
	"os"
	"path/filepath"
)

func walkpath(path string, f os.FileInfo, err error) error {
	//fmt.Printf("%s with %d bytes\n", path,f.Size())
	//fmt.Printf("%s with %s", path, bytefmt.ByteSize(uint64(f.Size())))
	if filepath.Ext(path) == ".mp4" {

		fmt.Printf("%s", path)
	}
	fmt.Printf(" File Extension is %s \n", filepath.Ext(path))
	return nil
}

func main() {
	flag.Parse()
	root := flag.Arg(0) // 1st argument is the directory location
	filepath.Walk(root, walkpath)
}
