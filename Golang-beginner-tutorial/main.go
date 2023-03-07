package main

// making a simple CLI booking apps
// go mod init booking-app

import "fmt"

func main(){
	const CONFERENCE_TICKETS int16 = 50
	var conferenceName string = "Go Conference"
	var remainingTickets int16= 50
	
	// fmt.Println() will automatically add a new line
	fmt.Println("Welcome to out",conferenceName," booking application")
	
	// fmt.Printf() format the string
	// %v is the default, printing plain strings
	// %T get the type of the var
	fmt.Printf("We have total of %v tickets\n", CONFERENCE_TICKETS)
	fmt.Printf("%v tickets left",remainingTickets)
	fmt.Println("Get your ticket, please!")
	
	for {
		// interact with the user
		var firstName string
		var lastName string
		var userTickets int32
		// var userTickets int16
		
		// ask user for their name
		// fmt.Scan() interpret spaces as seperators, seperate arguments 
		// cant use fmt.Scan() with spaces
		// use a pointer
		fmt.Println("Enter your first name")
		fmt.Scan(&firstName)
		
		fmt.Println("Enter your last name")
		fmt.Scan(&lastName)
		
		fmt.Println("Enter number of tickets")
		fmt.Scan(&userTickets)
		
		fmt.Printf("Thank you, %v %v for booking %v tickets\n",
		firstName, lastName, userTickets)

		// arrays of size 50, and data type of string
		// array's max size is fixed
		// bookings can also be initialized by
		// var bookings = [50]string{"Dan","Huy"} 
		var bookings [50]string
		bookings[0] = firstName + " " + lastName // asign data to the first element
		
		// slices are like lists in python
		// dynamic size, 
		// and can add element without calling the actual index
		var bookings_list[]string
		// append() append new element to a slice
		// and return a new slice
		bookings_list = append(bookings_list, firstName+ " "+lastName) 
		fmt.Printf("%T\n",booking_list)

		// Application Login
		// math operation with the same type
		remainingTickets = remainingTickets - int16(userTickets)
		
		firstName := []string{}
		for 
		
		fmt.Printf("%v tickets left\n", remainingTickets)
	}
}