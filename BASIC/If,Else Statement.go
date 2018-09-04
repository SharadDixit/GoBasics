package main

import "fmt"

func main()  {
	Yourage:= 18

	if Yourage>16{
		fmt.Println("you can drive")
		if Yourage>25{
			fmt.Println("Can Vote")
		}else{
			fmt.Println("Cannot Vote")
		}
		}else{
		fmt.Println("can't drive")
	}

}
//elseif{}    syntax
