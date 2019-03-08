package main

import (
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	flag.Parse()
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	if len(flag.Args()) != 0 && flag.Args()[0] == "true" {
		// $ photopic true
		///////////////////////////////////
		// TODO :: 날짜별 분류
		///////////////////////////////////
	} else {
		// $ photopic
		makeDir()
		files, _ := ioutil.ReadDir(dir)
		if len(files) == 0 {
			return
		}
		for _, file := range files {
			fn := strings.Split(file.Name(), ".")
			fileType := fn[len(fn)-1]
			if fileType == "JPG" {
				os.Rename(file.Name(), "./jpg/"+file.Name())
			} else if fileType == "RAF" || fileType == "CR2" {
				os.Rename(file.Name(), "./raw/"+file.Name())
			}
		}
	}
}

func makeDir() {
	jpgDir, _ := exists("./jpg")
	if !jpgDir {
		os.MkdirAll("./jpg", os.ModePerm)
	}
	rawDir, _ := exists("./raw")
	if !rawDir {
		os.MkdirAll("./raw", os.ModePerm)
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
