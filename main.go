package main

import "fmt"

func main() {
	var conferenceName = "Go conference"
	const conferenceTickets = 50
	var remainingTickets = 15
	greetUsers(conferenceName, conferenceTickets, remainingTickets)
	
	var firstName string
	var lastName string
	var email string
	var userTickets int 
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	fmt.Println("Enter the number of tickets: ")
	fmt.Scan(&userTickets)
	remainingTickets -= userTickets
	fmt.Printf("Thank you %v %v for booking tickets. You will receive a confirmation email at %v\n", firstName, lastName, email)
}
func greetUsers(confName string, confTickets int, remainingTickets int) {
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", confTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}