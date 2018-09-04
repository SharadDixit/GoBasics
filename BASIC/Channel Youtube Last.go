//pizza problem: pizza number, pizza sauce, pizza topping
package main

import (
	"strconv"
	"fmt"
	"time"
)

var pizzaNo=0      //Global Variables can be accessed in any function
var pizzaName=""

func main(){
stringChan:=make(chan string)
	for i:=1;i<=3 ;i++  {
go makeDough(stringChan)
go addSauce(stringChan)
go addToppings(stringChan)
time.Sleep(time.Millisecond*5000)    //Give time so that it has enough time to complete one whole pizza made(makedough,addsauce,addtoppings all three go routines get run)
	}

}
func makeDough(stringChan chan<- string){
	pizzaNo++    //increment pizza number
	pizzaName="Pizza #"+strconv.Itoa(pizzaNo)   //put pizza name
	fmt.Println("Dough Done And Send For Sauce")
	stringChan <- pizzaName             //send to channel the pizza name, channel is receiving
	time.Sleep(time.Millisecond*10)
	}
func addSauce(stringChan chan string){    //NO NEED OF DEFINING HERE WHETHER RECEIVING OR SENDING
	pizza:=<- stringChan              //received string in the channel is sent through the channel into a variable
	fmt.Println("Sauce done And Send for Topping",pizza)
	stringChan<-pizzaName    //send the string back to channel that is channel is receiving
	time.Sleep(time.Millisecond*10)
	}
func addToppings(stringChan chan string){
	pizza:=<-stringChan       //received string in the channel is sent to a variable
	fmt.Println("Add Toppings to ",pizza,"and Ship")
	time.Sleep(time.Millisecond*10)
	//now no need of sending the channel anything, that is channel won't receive anything as all go routines for one pizza are complete and no need of other sequence
}
