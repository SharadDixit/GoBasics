package main

import (
	"fmt"
	"os"
	"bufio"
)
func main(){
	srcfile,err:=os.Open("Hello.txt")
	if err!=nil{
		fmt.Println(err)
	}
	defer srcfile.Close()
	scanner:= bufio.NewScanner(srcfile)
	//It is of buffio package
	//func NewScanner(r io.Reader)*scanner
	//receives io.reader file and returns .scanner file
	for scanner.Scan(){     //scanning the .scanner file until the end of input
		line:=scanner.Text()        //func (s *scanner) text() string - function name is text, receives *scanner file and then returns string
	fmt.Println(line)
		}
}

