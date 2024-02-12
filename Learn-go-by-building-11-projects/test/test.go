package main

import (
	"fmt"
)


type Person struct {
	name string
	age int16
	school string
}

type welcome interface {
	greet() string
	giveDirection() string 
	canDrink() bool 
}

func (p Person) welcome(){
	
}