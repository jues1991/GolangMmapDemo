// GolangMmapDemoproject main.go
package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	var err error
	var fn string
	var fd int
	var fi os.FileInfo
	var data []byte
	//打开文件
	fn = "1.data"
	fd, err = syscall.Open(fn, syscall.O_RDWR, 0)
	if nil != err {
		fmt.Println("open file fail!!!")
	}
	//获取文件大小
	fi, err = os.Stat(fn)
	if nil != err {
		fmt.Println("get file size fail!!!")
	}
	//映射到内存
	data, err = syscall.Mmap(fd, 0, int(fi.Size()), syscall.PROT_READ|syscall.PROT_WRITE|syscall.PROT_EXEC, syscall.MAP_SHARED)

	//
	if nil != err {
		fmt.Println("mmap fail!!!")
		return
	}
	//
	addr := &data[0]
	fmt.Println("mmap success,addr=", addr, "size=", len(data))

	//取消映射
	syscall.Munmap(data)
}
