package main 

import (
	"fmt"

);

// import route2 "application/route"
	

import route2 "github.com/gui-drumond/imersaofullcycle-simulator/application/route";

func main(){

	fmt.Println( route2.Route)

	route := route2.Route{
		ID: "1", 
		ClientID: "1",
	}
	
	route.LoadPositions()
	// stringjson, _ := route.ExportJsonPosition()
	// fmt.Println(stringjson[1])
}