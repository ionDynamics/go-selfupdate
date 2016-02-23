package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Specify directoy and version")
		return
	}

	dir, err := filepath.Abs(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if path != dir && path != dir+string(filepath.Separator)+info.Name() && info.IsDir() && info.Name() != os.Args[2] {
			return os.RemoveAll(path)
		}
		//uncomment this if you don't want to provide diffs
		/*if path != dir && path != dir+string(filepath.Separator)+info.Name() && !info.IsDir() && filepath.Base(filepath.Dir(path)) != os.Args[2] {
			return os.Remove(path)
		}*/
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
