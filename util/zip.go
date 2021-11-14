package util

import (
	"archive/zip"
	"github.com/axgle/mahonia"
	"io"
	"os"
	"path/filepath"
)

func UnZip(ZipFile,UnZipPath string) error {
	var dec mahonia.Decoder
	//var enc mahonia.Encoder
	zipReader, err := zip.OpenReader(ZipFile)
	if err != nil {
		return err
	}
	defer zipReader.Close()
	for _, f := range zipReader.File {
		dec = mahonia.NewDecoder("gbk")
		fpath := filepath.Join(UnZipPath, f.Name)
		fpath = dec.ConvertString(fpath)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
		} else {
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return err
			}

			inFile, err := f.Open()
			if err != nil {
				return err
			}
			defer inFile.Close()

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer outFile.Close()
			_, err = io.Copy(outFile, inFile)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
