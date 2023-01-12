package cmd

import (
	"fmt"
	"github.com/jonghyeons/photopic/models/constant"
	"github.com/jonghyeons/photopic/utils"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/spf13/cobra"
	"io/fs"
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

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Classify JPEG and RAW files.",
	Long: `Classify JPEG and RAW files.
Usage:
photopic run [filepath]`,
	Run: func(cmd *cobra.Command, args []string) {
		dir := ""
		if len(args) != 0 {
			dir = args[0]
		} else {
			var err error
			if dir, err = os.Getwd(); err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
		}

		if err := utils.MakeDir(dir); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		files, err := ioutil.ReadDir(dir)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		if len(files) == 0 {
			fmt.Println("empty directory")
			os.Exit(1)
		}

		AnalysisExifAndMoveFile(dir, files)

		fmt.Println("## It's done!")
		fmt.Println("## Photopic Report")
		fmt.Println("1. Lens Angle")
		fmt.Println("# This data is analyzed only in the raw files.")
		fmt.Println("wide-angle	", wideAngleLensCnt, "shot")
		fmt.Println("standard	", standardLensCnt, "shot")
		fmt.Println("telephoto	", telephotoLensCnt, "shot")
		fmt.Println("analysis failure: ", errCnt)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	rootCmd.Flags().BoolP("directory", "d", false, "Specifies the directory. (default: current file path)")
}

func AnalysisExifAndMoveFile(dir string, files []fs.FileInfo) {
	for _, file := range files {
		if strings.Contains(file.Name(), "/raw") || strings.Contains(file.Name(), "/jpg") {
			continue
		}
		fn := strings.Split(file.Name(), ".")
		fileType := fn[len(fn)-1]
		if utils.Contains(constant.RawTypes, fileType) {
			if err := AnalysisAngle(dir + "/" + file.Name()); err != nil {
				fmt.Println(file.Name(), err.Error())
			}
		}

		newPath := ""
		if strings.ToUpper(fileType) == "JPG" {
			newPath = "/jpg/"
		} else if utils.Contains(constant.RawTypes, fileType) {
			newPath = "/raw/"
		}

		if err := os.Rename(filepath.Join(dir, "/", file.Name()), filepath.Join(dir, newPath, file.Name())); err != nil {
			fmt.Println(file.Name(), err.Error())
		}
	}
}

func AnalysisAngle(fnameWithPath string) error {
	f, err := os.Open(fnameWithPath)
	if err != nil {
		return err
	}

	exif.RegisterParsers()
	x, err := exif.Decode(f)
	if err != nil {
		return err
	}

	focal, err := x.Get(exif.FocalLength)
	if err != nil {
		errCnt++
		return err
	}

	numer, denom, err := focal.Rat2(0)
	if err != nil {
		errCnt++
		return err
	}

	if numer/denom < 35 {
		wideAngleLensCnt++
	} else if numer/denom >= 35 || numer/denom < 85 {
		standardLensCnt++
	} else {
		telephotoLensCnt++
	}

	return nil
}
