package main
import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"io/ioutil"
)
func ReadIntss(r io.Reader) ([]int) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {

		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result
		}
		result = append(result, x)
	}
	return result
}
func main() {
	bs,err:=ioutil.ReadFile("numbers.txt")
	if err!=nil{
		fmt.Println(err)
	}
	s:= string(bs)
	fmt.Println(s)
	ints := ReadIntss(strings.NewReader(s))
	fmt.Println(ints)
}
