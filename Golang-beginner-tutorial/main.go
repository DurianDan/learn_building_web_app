package main

// making a simple CLI booking apps
// go mod init booking-app

import (
	"booking-app/helper"
	"fmt"
	"strconv"
	"strings"
	"maps"
)

func main(){
	const CONFERENCE_TICKETS int16 = 50
	var conferenceName string = "Go Conference"
	var remainingTickets int= 50
	
	// fmt.Println() will automatically add a new line
	fmt.Println("Welcome to out",conferenceName," booking application")
	
	// fmt.Printf() format the string
	// %v is the default, printing plain strings
	// %T get the type of the var
	fmt.Printf("We have total of %v tickets\n", CONFERENCE_TICKETS)
	fmt.Printf("%v tickets left",remainingTickets)
	fmt.Println("Get your ticket, please!")
	
	firstNames := []string{}
	// arrays of size 50, and data type of string
	// array's max size is fixed
	// bookings can also be initialized by
	// var bookings = [50]string{"Dan","Huy"} 
	var bookings [50]string
	// bookings[0] = firstName + " " + lastName // asign data to the first element

	// map is like dictionary 
	// userData has keys of type string and values of type string
	// make() creates empty variable, for map, slice and array
	var bookings_map = make([]map[string]string, 0)
	for len(bookings) < 50 && remainingTickets > 0 {
		// interact with the user
		var firstName string
		var lastName string
		var userTickets int
		var email string
		var city string

		// var userTickets int16
		
		// ask user for their name
		// fmt.Scan() interpret spaces as seperators, seperate arguments 
		// cant use fmt.Scan() with spaces
		// use a pointer
		fmt.Println("Enter your first name")
		fmt.Scan(&firstName)
		
		fmt.Println("Enter your last name")
		fmt.Scan(&lastName)
		
		fmt.Println("Enter your email")
		fmt.Scan(&email)
		
		fmt.Println("Which city are you occupying ?:")
		fmt.Scan(&city)

		fmt.Println("Enter number of tickets")
		fmt.Scan(&userTickets)
		
		fmt.Printf("Thank you, %v %v for booking %v tickets\n",
		firstName, lastName, userTickets)

		if !helper.ValidateUserInput(firstName,lastName,email, userTickets, remainingTickets){
			fmt.Println("Please check your input, they are invalid")
			continue
		}
		
		// append() append new element to a slice/map/array 
		// and return a new slice/map/array
		var userData = map[string]string{
			"firstName":firstName,
			"lastName":lastName,
			"email":email,
			"numberOfTickets": strconv.FormatUint(uint64(userTickets), 10),
		}
		bookings_map = append(bookings_map, userData) 
		fmt.Printf("%T\n", bookings_map)

		// Application Login
		// math operation with the same type
		remainingTickets = remainingTickets - userTickets
		
		if remainingTickets <= 0 {
			fmt.Println("Out of Tickets !!!")
			break
		}else{
			fmt.Printf("The first name of customer is %v\n",firstNames)
			fmt.Printf("%v tickets left\n", remainingTickets)
		}

		switch city {
			case "Hanoi","Saigon": fmt.Println("you Live in Vietnam")
			case "Anger": fmt.Println("you Live in France")
			case "Tokyo": fmt.Println("you Live in Japan")
			default:
				// other posibilities
				fmt.Println("choosen country not in App")

		}
	}
}



