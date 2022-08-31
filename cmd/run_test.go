package cmd

import (
	"fmt"
	"github.com/jonghyeons/photopic/utils"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// The image file must be located in the path(/sample).
func TestAnalysisExifAndMoveFile(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Errorf(err.Error())
	}

	dir = filepath.Join(dir, "/../sample")
	err = utils.MakeDir(dir)
	if err != nil {
		t.Errorf(err.Error())
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		t.Errorf(err.Error())
	}
	if len(files) == 0 {
		fmt.Println("empty directory")
		t.Error()
	}

	type args struct {
		dir   string
		files []fs.FileInfo
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "exif analysis test",
			args: args{
				dir:   dir,
				files: files,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AnalysisExifAndMoveFile(tt.args.dir, tt.args.files)
		})
	}

	if wideAngleLensCnt > 0 || standardLensCnt > 0 || telephotoLensCnt > 0 || errCnt > 0 {
		fmt.Println("## It's done!")
		fmt.Println("## Photopic Report")
		fmt.Println("1. Lens Angle")
		fmt.Println("# This data is analyzed only in the raw files.")
		fmt.Println("wide-angle	", wideAngleLensCnt, "shot")
		fmt.Println("standard	", standardLensCnt, "shot")
		fmt.Println("telephoto	", telephotoLensCnt, "shot")
		fmt.Println("analysis failure: ", errCnt)
	} else {
		t.Error()
	}
}
