package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

/*
由于go-zero的goctl生成gprc服务只能是一个proto，服务多的时候 比较乱，所以用这个工具来合成一个
Author: LiuFeiHua
Date: 2024/6/11 10:52
*/
func main() {
	folderPath := "rpc/pms/proto"         // 替换为你的文件夹路径
	outputFilePath := "rpc/pms/pms.proto" // 新文件的路径

	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	var fileContents []byte
	var startContents []byte
	for _, file := range files {
		if file.IsDir() || file.Name() == "main.go" {
			continue
		}

		fileName := filepath.Join(folderPath, file.Name())
		fileData, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Printf("Error reading file %s: %s\n", fileName, err)
			continue
		}
		startContents = fileData[:69]
		fileData = fileData[70:]

		fileContents = append(fileContents, fileData...)
	}
	startContents = append(startContents, fileContents...)

	start := strings.Replace(string(startContents), "package main", "package pmsclient", 1)
	start = strings.Replace(start, "option go_package = \"./proto\"", "option go_package = \"./pmsclient\"", 1)
	err = ioutil.WriteFile(outputFilePath, []byte(start), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Merged files to", outputFilePath)
}
