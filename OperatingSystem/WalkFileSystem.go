package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"runtime"
)

func GetPlatform() string {
	os := runtime.GOOS
	return os
}

func main() {
	os := GetPlatform()

	if os == "windows" {
		root := "C:\\Users\\robert.gaines\\"

		err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				fmt.Println("[*] Directory: ", path)
			} else {
				fmt.Println("[*] File: ", path)
			}
			return nil
		})
		if err != nil {
			fmt.Println("[!] Error walking the path: ", err)
			return
		}
	}
}
