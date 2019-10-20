package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	dir := ""
	flag.Parse()
	// fmt.Println("Your flags :: ", flag.Args())

	if len(flag.Args()) != 0 {
		dir = flag.Args()[0]
	} else {
		dir, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	}

	makeDir(dir)
	files, err := ioutil.ReadDir(dir)

	if len(files) == 0 || err != nil {
		fmt.Println("ERROR")
		fmt.Println(err.Error())
		return
	}

	rawType := []string{"RAF", "CR2", "CR3", "ARW"}

	for _, file := range files {
		fn := strings.Split(file.Name(), ".")
		fileType := fn[len(fn)-1]

		if fileType == "JPG" {
			os.Rename(dir+"/"+file.Name(), dir+"/jpg/"+file.Name())
		} else if Contains(rawType, fileType) {
			os.Rename(dir+"/"+file.Name(), dir+"/raw/"+file.Name())
		}
	}
}

func makeDir(path string) {
	jpgDir, _ := exists(path + "/jpg")
	if !jpgDir {
		os.MkdirAll(path+"/jpg", os.ModePerm)
	}
	rawDir, _ := exists(path + "/raw")
	if !rawDir {
		os.MkdirAll(path+"/raw", os.ModePerm)
	}
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
