package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var allAssets []string

func copyFolder(source string, dest string) (err error) {

	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, _ := os.Open(source)

	objects, err := directory.Readdir(-1)

	for _, obj := range objects {

		sourcefilepointer := source + "/" + obj.Name()

		destinationfilepointer := dest + "/" + obj.Name()

		if obj.IsDir() {
			panic("assets should not contains sub direcotry")
			err = copyFolder(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			err = copyFile(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
	return
}

func copyFile(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		sourceinfo, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, sourceinfo.Mode())
		}

	}

	return
}

func fileOrDirExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func processDir(path string, fileInfo os.FileInfo, errX error) error {

	if errX != nil {
		fmt.Printf("error 「%v」 at a path 「%q」\n", errX, path)
		return errX
	}

	if fileInfo.IsDir() {
		if strings.HasSuffix(fileInfo.Name(), ".assets") && fileOrDirExist(strings.TrimSuffix(path, ".assets")) {
			allAssets = append(allAssets, path)
		}
	}
	return nil
}

func moveAssets(assets string) {

}

func main() {

	allAssets = []string{}
	err := filepath.Walk("./docs/posts", processDir)

	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", ".", err)
	}

	for _, assets := range allAssets {
		markdownParentDir := assets[0:strings.LastIndex(assets, "/")]
		markdownFileNameDotAssets := assets[strings.LastIndex(assets, "/"):]
		markdownFileName := strings.TrimSuffix(markdownFileNameDotAssets, ".assets")
		dst := markdownParentDir + "/" + markdownFileName + "/" + markdownFileNameDotAssets

		fmt.Println(assets, "=>", dst)

		if err := copyFolder(assets, dst); err != nil {
			fmt.Println("Fail to copy dir:", assets)
			continue
		}
		if err := os.RemoveAll(assets); err != nil {
			fmt.Println("Fail to remove dir:", assets)
		}
	}
}
