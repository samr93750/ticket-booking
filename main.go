package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
	
)
const conferenceTicket = 50

var conferenceName = "go conference"
var remainingTickets uint = 50
var bookings = make([]map[string]string,0)

func main() {
	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)
			firstNames := pickFNames()
			fmt.Printf("The first names of bookings are :%v\n", firstNames)
			// var noTicketsRemaining bool = remainingTickets==0
			if remainingTickets == 0 {
				fmt.Println("our conference is booked out. come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
		}
	}
}
func greetUsers() {
	fmt.Printf("Welcome to %v  booking application\n", conferenceName)
	fmt.Printf("**we have total of %v tickets and %v are still available**\n", conferenceTicket, remainingTickets)
	fmt.Println("*get your tickets here to attend*")
}
func pickFNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	fmt.Println("enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("enter your email:")
	fmt.Scan(&email)

	fmt.Println("how much tickets do you need:")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}
func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//create a map for a user
	var userData= make(map[string]string)
	userData["firstName"]=firstName
	userData["hostname"]=lastName
	userData["email"]=email
	userData["numberOfTicket"]=strconv.FormatUint(uint64(userTickets),10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n",bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets you will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}
