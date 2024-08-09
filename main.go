package main

import (
	"fmt"
	"strings"
)
func main(){
	var conferenceName ="go conference"
	const conferenceTicket =50
    var remainingTickets uint =50

	fmt.Printf("***welcome to our %v book application***\n",conferenceName)
	fmt.Printf("**we have total of %v tickets and %v are still available**\n",conferenceTicket,remainingTickets)
	fmt.Println("*get your tickets here to attend*")
	


    var bookings[] string
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	//ask user for their name
for{
		fmt.Println("enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("enter your email:")
	fmt.Scan(&email)

	fmt.Println("how much tickets do you need:")
	fmt.Scan(&userTickets)

	 isValidName := len(firstName)>=2 && len(lastName)>=2
	 isValidEmail := strings.Contains(email,"@")
     isValidTicketNumber := userTickets>0 && userTickets <= remainingTickets


	if isValidName&&isValidEmail&&isValidTicketNumber{
		remainingTickets = remainingTickets - userTickets
bookings=append(bookings,firstName+" "+lastName)

fmt.Printf("Thank you %v %v for booking %v tickets you will receive a confirmation email at %v\n",firstName,lastName, userTickets , email)
fmt.Printf("%v tickets remaining for %v\n",remainingTickets,conferenceName)

firstNames := []string{}
for _, booking :=range bookings {
	var names=strings .Fields(booking)
	 firstNames = append(firstNames,names[0])
}
fmt.Printf("The first names of bookings are :%v\n",firstNames)
// var noTicketsRemaining bool = remainingTickets==0
if remainingTickets ==0{
	//end the program
	fmt.Println("our conference is booked out. come back next year")
	break
}

}else{
	if !isValidName{
		fmt.Println("first name or last name you entered is too short")
	}
	if !isValidEmail{
		fmt.Println("email address you entered doesn't contain @ sign")
	}
	if !isValidTicketNumber{
		fmt.Println("number of tickets you entered is invalid")
	}
	
	}

	city:= "london"
	switch city{
	case "newyork":
		//something
	case "singapore":
		//something 
	case "berlin":
		//something
	case "london":
		//something
	default:
		fmt.Print("no valid city is selected")
	}
}
}