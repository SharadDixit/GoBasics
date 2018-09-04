package main
import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"io/ioutil"
	"strconv"
	"time"
)
func ReadIntssss(r io.Reader) ([]int64) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int64
	for scanner.Scan() {
		aString := scanner.Text()
		aa := len(aString)
		a := make([]int, aa)
		for i := 0; i < len(aString); i++ {

			xa := int(aString[i])
			a[i] = xa
		}
		var IDs []string
		for i := 0; i < len(a); i++ {
			IDs = append(IDs, strconv.Itoa(a[i]))
		}
		x := strings.Join(IDs, "")
		newInt, _ := strconv.ParseInt(x, 0, 64)
		result=append(result,newInt)
		}
	return result
}
func elemtentstoeachworker(r []int64,elementspassigs int,aa int)[]int64{
	newarray:=[]int64{}
	newarray=r[aa:aa+elementspassigs]
	return(newarray)
	}

func main() {
	fmt.Println("File for input:ASCII FILE.txt")
	time.Sleep(time.Second*1)
	bs,err:=ioutil.ReadFile("ASCII FILE.txt")
	if err!=nil{
		fmt.Println(err)
	}
	s:= string(bs)
	ObtainedASCIIString := ReadIntssss(strings.NewReader(s))
	fmt.Println("String After Conversion:",ObtainedASCIIString)
	elementsinobtainedstring:=len(ObtainedASCIIString)
	time.Sleep(time.Second*1)
	var work int
	fmt.Println("Enter Number Of Workers:")
	fmt.Scan(&work)
	fmt.Println("Workers Entered:",work)
	WorkersWork:=elementsinobtainedstring/work
	fmt.Println("Each worker has the following number of elements for partial addition:",WorkersWork)
	a:=0
	Obtainedarrays:=[]int64{}
	for i:=1;i<=work ;i++  {
		Obtainedarrays=elemtentstoeachworker(ObtainedASCIIString,WorkersWork,a)
		a=a+WorkersWork
		fmt.Println(Obtainedarrays)
		go working(Obtainedarrays, )
		}

}
func working(r []int64){

}


