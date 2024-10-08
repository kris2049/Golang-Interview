package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	inputFiles := []string{"test1.txt", "test2.txt", "test3.txt", "test4.txt"}
	outputFile := "o.txt"

	//创建输出文件
	outFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating output file: ", err)
		return
	}

	defer outFile.Close()

	t1 := time.Now()

	//创建一个写入器
	writer := bufio.NewWriter(outFile)

	//遍历每一个写入文件
	for _, inputFile := range inputFiles {
		// 打开输入文件
		inFile, err := os.Open(inputFile)
		if err != nil {
			fmt.Println("Error opening input file: ", err)
			continue
		}

		// 使用bufio.NewScanner逐行读取输入文件
		scanner := bufio.NewScanner(inFile)
		for scanner.Scan() {
			_, err := writer.WriteString(scanner.Text() + "\n")
			if err != nil {
				fmt.Println("Error writing to output file: ", err)
				break
			}
		}

		// 检查读取过程中是否有错误
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading from input file: ", inputFile, err)
		}

		//关闭输入文件
		inFile.Close()
	}

	if err := writer.Flush(); err != nil {
		fmt.Println("Error flushing writer: ", err)
	}

	t2 := time.Now()

	fmt.Println("Data has been successfully written to ", outputFile)

	fmt.Println("Total cost of time is: ", t2.Sub(t1))

}
