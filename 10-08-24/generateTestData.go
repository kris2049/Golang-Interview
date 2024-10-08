package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"os"
// 	"time"
// )

// const (
// 	numLines   = 100 // 要生成的行数
// 	lineLength = 50  // 每行的字符长度
// )

// func randomString(r *rand.Rand, length int) string {
// 	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
// 	s := make([]rune, length)
// 	for i := range s {
// 		s[i] = letters[r.Intn(len(letters))]
// 	}
// 	return string(s)
// }

// func main() {
// 	// 创建一个新的随机生成器
// 	r := rand.New(rand.NewSource(time.Now().UnixNano()))

// 	// 创建输出文件
// 	file, err := os.Create("test4.txt")
// 	if err != nil {
// 		fmt.Println("Error creating file:", err)
// 		return
// 	}
// 	defer file.Close()

// 	// 写入随机内容
// 	for i := 0; i < numLines; i++ {
// 		_, err := file.WriteString(randomString(r, lineLength) + "\n")
// 		if err != nil {
// 			fmt.Println("Error writing to file:", err)
// 			return
// 		}
// 	}

// 	fmt.Println("Test file 'test.txt' has been created with", numLines, "lines.")
// }
