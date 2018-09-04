//In programming languages and type theory, polymorphism is the provision of a single interface to entities of different types.
// A polymorphic type is one whose operations can also be applied to values of some other type, or types.
//An interface is a programming structure/syntax that allows the computer to enforce certain properties on an object (class).
// For example, say we have a car class and a scooter class and a truck class. 
// Each of these three classes should have a start_engine() action. 
// How the "engine is started" for each vehicle is left to each particular class, but the fact that they must have a start_engine action is the domain of the interface.
package main

import "fmt"
import "math"               //for mathematical uses of pow and pi
func main(){
	recta:= rectangles{20,30}
	circla:= circle{2}
	fmt.Println("rectangle area=",getarea(recta))
	fmt.Println("Circle area=",getarea(circla))
}

type shape interface {           //interface created to bind the program, that is we do not have to specifically run different area for different object
	areas() float64              //for each object no need of seeing area again and again just we need to know what type of object and then parse the value
}
type rectangles struct {        //rectangle structure to have as many as rectangles you want
	heights float64
	widths float64
}
type circle struct {             //circle structure to have as many as circles you want
	radius float64
	}
func(r rectangles)areas()float64{           //function which is referenced with rectangles to calculate area of rectangle with returning value as float
	return r.heights*r.widths
}
func(c circle)areas()float64{            //function which is referenced with circle to calculate area of circle with returning value as float
	return math.Pi*math.Pow(c.radius,2)
}
func getarea(s shape)float64{             //function returns the area of each different shape obtained
	return s.areas()
}
//automatically calls the area function whichever is needed according to data type all because of interface


