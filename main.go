package main

import (
	"go-packge-demo/api"
	"go-packge-demo/controller"

	"fmt"
)

func main(){
	fmt.Println("it's main func")
	s1 := controller.HandleSearchByBusiness()
	fmt.Println(s1)

	s2 := controller.GetAddr()
	fmt.Println(s2)

	api.Http()
}