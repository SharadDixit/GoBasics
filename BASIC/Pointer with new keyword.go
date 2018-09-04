package main
//new first.
// It's a built-in function that allocates memory, but unlike its namesakes in some other languages it does not initialize the memory, it only zeros it.
// That is, new(T) allocates zeroed storage for a new item of type T and returns its address, a value of type *T.
// In Go terminology, it returns a pointer to a newly allocated zero value of type T.
import "fmt"
func main(){
yptr :=new(int)           //new keyword used for allocating memory to the pointer
changeval(yptr)
fmt.Println(yptr)         //prints address
fmt.Println(&yptr)        //prints address
fmt.Println(*yptr)        //prints pointer value
}
func changeval(yptr *int){
*yptr=100
}

