package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 3

var conferenceName string = "Go Conference"
var remainingTickets int = conferenceTickets

type Booking struct { //struct is a collection of fields, each field has a name and a type
	firstName string
	lastName  string
	email     string
	tickets   int
}

var bookings = make([]Booking, 0) //create empty list of map

var wg = sync.WaitGroup{}

func main() {

	fmt.Printf("Welcome to %v booking app\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		// for remainingTickets > 0 && len(bookings ) < conferenceTickets {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTickets := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)
		// check if it has enough tickets
		if isValidTickets && isValidName && isValidEmail {
			booking(userTickets, firstName, lastName, email)

			wg.Add(1)
			go sendTicket(firstName, lastName, email)
			firstNames := getFirstNames()
			fmt.Printf("The first names of the bookings: %v\n", firstNames)

			// noTicketRemaining := remainingTickets == 0
			if remainingTickets == 0 {
				fmt.Println("Sorry, we are sold out")
				break // break out of loop
			}
		} else {
			if !isValidName {
				fmt.Println("Invalid name")
			}
			if !isValidEmail {
				fmt.Println("Invalid email")
			}
			if !isValidTickets {
				fmt.Println("Invalid number of tickets")
			}
		}
	}
	wg.Wait()
}

func greeting(name string) {
	fmt.Println("Hello", name)
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		// names := strings.Fields(booking)
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	greeting(firstName + " " + lastName)
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets you want: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func booking(userTickets int, firstName string, lastName string, email string) {
	// update remaining tickets
	remainingTickets = remainingTickets - userTickets

	//create a map for user
	user := Booking{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		tickets:   userTickets,
	}
	// user := make(map[string]string) //cannot mix types in one map
	// user["firstName"] = firstName
	// user["lastName"] = lastName
	// user["email"] = email
	// // convert int to string
	// user["tickets"] = strconv.FormatInt(int64(userTickets), 10)
	// user["tickets"] = fmt.Sprintf("%v", userTickets)

	// create booking
	bookings = append(bookings, user)
	fmt.Println("Booking list: ", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets\n", firstName, lastName, userTickets)
	fmt.Printf("You will receive the confirmation via email %v\n", email)
	fmt.Printf("We have %v tickets remaining for conference %v\n", remainingTickets, conferenceName)
}

func sendTicket(firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("ticket to %v for %v %v", email, firstName, lastName)
	fmt.Println("##############################")
	fmt.Printf("Sending %v\n", ticket)
	fmt.Println("##############################")
	wg.Done()
}
