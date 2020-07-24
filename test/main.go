package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func main() {
	dir, _ := filepath.Abs("build/")

	x, err := SearchDirectoryForAbi(dir)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(x)
	x0 := strings.Replace(x[0],"\\","/", -1)
	fmt.Println(x0)
	fmt.Println(path.Base(x0))
	fmt.Println(path.Base(x[0]))
	fmt.Println(path.Ext(x[0]))

}

func SearchDirectory(dir string) ([]string, error) {
	files := []string{}
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})

	if err != nil {
		return nil, err
	}
	return files, nil
}


func SearchDirectoryForAbi(dir string) ([]string, error) {
	files, err := SearchDirectory(dir)
	if err != nil {
		return []string{}, err
	}
	contracts := []string{}

	//fmt.Println(files)
	for _, file := range files {
		if(path.Ext(file) == ".abi") {
			contracts = append(contracts, file)
		}
	}

	return contracts, nil
}