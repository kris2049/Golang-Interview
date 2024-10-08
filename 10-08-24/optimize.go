package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"sync"
// 	"time"
// )

// func main() {
// 	inputFiles := []string{"test1.txt", "test2.txt", "test3.txt", "test4.txt"}
// 	outputFile := "o.txt"

// 	// 创建输出文件
// 	outFile, err := os.Create(outputFile)
// 	if err != nil {
// 		fmt.Println("Error creating output file: ", outFile)
// 		return
// 	}

// 	defer outFile.Close()

// 	t1 := time.Now()

// 	writer := bufio.NewWriter(outFile)

// 	// 使用waitGroup 等待所有协程完成
// 	var wg sync.WaitGroup

// 	// 用于同步对writer的访问
// 	mutex := &sync.Mutex{}

// 	// 遍历每一个输入文件
// 	for _, inputFile := range inputFiles {
// 		wg.Add(1)

// 		go func(file string) {
// 			defer wg.Done() //协程完成时减少计数

// 			// 打开输入文件
// 			inFile, err := os.Open(file)
// 			if err != nil {
// 				fmt.Println("Error open inputfile: ", err)
// 				return
// 			}

// 			defer inFile.Close()

// 			// 使用scanner来逐行读取数据
// 			scanner := bufio.NewScanner(inFile)
// 			for scanner.Scan() {
// 				// 使用锁来保护写入操作
// 				mutex.Lock()
// 				_, err := writer.WriteString(scanner.Text() + "\n")
// 				mutex.Unlock()
// 				if err != nil {
// 					fmt.Println("Error write to output file: ", err)
// 					break
// 				}
// 			}

// 			// 检查读取过程中是否有错误
// 			if err := scanner.Err(); err != nil {
// 				fmt.Println("Error reading from input file: ", file, err)
// 			}
// 		}(inputFile)
// 	}

// 	// 等待所有协程完成
// 	wg.Wait()

// 	// 刷新写入器，确保所有的数据写入输出文件
// 	if err := writer.Flush(); err != nil {
// 		fmt.Println("Error flushing writer:", err)
// 	}

// 	t2 := time.Now()

// 	fmt.Println("Data has been successfully written to ", outputFile)
// 	fmt.Println("Total cost of time is: ", t2.Sub(t1))

// }
