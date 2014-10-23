package main

import "fmt"
import "os"
import "io"

func main() {
	CopyFile("world.txt", "hello.txt")
}
func CopyFile(dst, src string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println("File open error")
		return
	}
	defer srcFile.Close()
	dstFile, err := os.Create(dst)
	if err != nil {
		fmt.Println("File Create error")
		return
	}
	defer dstFile.Close()
	return io.Copy(dstFile, srcFile)
}
