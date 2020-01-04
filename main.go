package main

import (
	"flag"
	"fmt"
	"github.com/rwcarlsen/goexif/exif"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	wideAngleLensCnt = 0
	standardLensCnt  = 0
	telephotoLensCnt = 0
	errCnt           = 0
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

	MakeDir(dir)
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
			done := SetLensAngleCnt(dir + "/" + file.Name())
			if !done {
				errCnt++
			}

			os.Rename(dir+"/"+file.Name(), dir+"/raw/"+file.Name())
		}
	}

	fmt.Println("## Photopic Report")
	fmt.Println("1. Lens Angle")
	fmt.Println("# This data is analyzed only in the raw files. ")
	fmt.Println("wide-angle: ", wideAngleLensCnt, "shot")
	fmt.Println("standard: ", standardLensCnt, "shot")
	fmt.Println("telephoto: ", telephotoLensCnt, "shot")
	fmt.Println("analysis failure: ", errCnt)

}

func MakeDir(path string) {
	jpgDir, _ := Exists(path + "/jpg")
	if !jpgDir {
		os.MkdirAll(path+"/jpg", os.ModePerm)
	}
	rawDir, _ := Exists(path + "/raw")
	if !rawDir {
		os.MkdirAll(path+"/raw", os.ModePerm)
	}
}

func Exists(path string) (bool, error) {
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

func SetLensAngleCnt(fnameWithPath string) bool {
	f, _ := os.Open(fnameWithPath)

	exif.RegisterParsers()
	x, err := exif.Decode(f)
	if err != nil {
		return false
	}

	focal, _ := x.Get(exif.FocalLength)
	numer, denom, _ := focal.Rat2(0)

	if numer/denom < 35 {
		wideAngleLensCnt++
	} else if numer/denom >= 35 || numer/denom < 85 {
		standardLensCnt++
	} else {
		telephotoLensCnt++
	}

	return true
}
