package main

import (
	"fmt"
	"os"
	"bufio"

	"io/ioutil"
)
func main(){
	srcfile,err:=os.Open("Hello.txt")
	if err!=nil{
		fmt.Println(err)
	}
	defer srcfile.Close()
	bufferreader:= bufio.NewReader(srcfile)
	//func NewReader(rd io.Reader)*Reader
	fd,err:=ioutil.ReadAll(bufferreader)
	fd1:=string(fd)
	fmt.Println(fd1)
}

