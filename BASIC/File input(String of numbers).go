//func (f *file) Name() String , therefore in the above function the receiving type is file(a pointer to the file )
// and Name is the function name
// and then the return type is a string
package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"os"
	"io/ioutil"
)

// ReadInts reads whitespace-separated ints from r. If there's an error, it
// returns the ints successfully read so far as well as the error value.
func ReadInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	//func NewScanner(r io.Reader) *Scanner
	scanner.Split(bufio.ScanWords)
	//func (s *Scanner) Split(split SplitFunc)
	//Receives *Scanner file
	//func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error)
	//bufio.ScanWords is a split function
	//ScanWords is a "split function" for a Scanner that returns each space-separated word of text, with surrounding spaces deleted.
	//It will never return an empty string. The definition of space is set by unicode.IsSpace.
	var result []int
	for scanner.Scan() {   //scanning the *scanner file until the end of input
//Scan advances the Scanner to the next token, which will then be available through the Bytes or Text method.
//It returns false when the scan stops, either by reaching the end of the input or an error.
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

func main() {
	f,err:=os.Open("numbers.txt")
	if err !=nil {
		fmt.Println(err)
	}
	fmt.Println(f)
	defer f.Close()
	bs,err:=ioutil.ReadAll(f)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(bs)
	s:= string(bs)

	ints, err := ReadInts(strings.NewReader(s))
	//ReadInts is a created function by myself which is parsing (strings.NewReader) which is a generated io.reader file
	//This is only information for strings.NewReader
	//func NewReader(s string) *Reader
	//It is of the strings package
	//It receives a string which is need to be converted to io.reader file and sent
	// returns io.reader file
	fmt.Println(ints, err)
}