// Everything in Go is organized as packages
// main package is the main module of the application
package main

// importing packages - builtin go packages
// A single package can be imported like this:
// import "fmt"
import (
	"fmt"
	"time"
	"sync"
	// "strconv"
	// "strings"
	// importing the helper package
	"booking-app/helper"
)

// Package level Scope
// ###############################################

// declaring and initializing a constant value : Package Scope
const conferenceTickets int = 50
var conferenceName string = "Go Conference"
// alternate way for defining var variables
// Example above var conferenceName string = "Go Conference" can be written as below
// conferenceName := "Go Conference" - Note: this way of defining variables will not work on the package level variables.

var remainingTickets uint = 50 //uint type can only contain positive numbers
// declaring dynamic array - known as slice in golang.
// var bookings = make([]map[string]string, 0) // this is the way to declare dynamic array of maps in golang

// this is a form of declaring an array
// var bookingUsers [50]string // array of booking users
// declaring and defining an array example is mentioned below
// var bookingUsers = [50]string{"user1", "user2", "user3} 

// Creating a custom structure to hold the bookings
type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var bookings = make([]UserData, 0) // this is the way to declare dynamic array of struct in golang

// Wait group is a way to tell the main function to wait for the goroutines to finish their execution.
var wg = sync.WaitGroup{} // this is a way to create a wait group in golang.


// main function is the entry point of the application from where the go starts running the application
func main() {

	// calling greetUser function
	greetUser()

	// for loop - this is a way to run an infinite loop in golang.
	for {

 		// calling getUserInput function
		firstName, lastName, email, userTickets := getUserInput()
		
		// calling the validateUserInput function
		isValidName, isValidEmail, isValidTicketNumber := 
					helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)


		if isValidName && isValidEmail && isValidTicketNumber {
			
			bookTickets(userTickets, firstName, lastName, email)

			// this is a way to add a goroutine to the wait group. 
			// 1 means we're adding a single thread to the wait group.
			// 2 means we're adding two threads to the wait group. 
			wg.Add(1) 
			
			// go keyword spins the goroutine thread in the background.
			go sendTicket(userTickets, firstName, lastName, email)

			// calling getFirstNames function		
			firstNames := getFirstNames()
			fmt.Printf("First names of the users who booked the tickets are: %v\n", firstNames)

			// conditions statements in golang.
			// breaking the infinite loop when all the tickets are sold out.
			if remainingTickets == 0 {
				// end program
				fmt.Println("Sorry, all tickets are sold out. Please come back next year.")
				break
			}

		} else {

			if !isValidName {
		        fmt.Println("Please enter a valid first name and last name.")
			}
			if !isValidEmail {
				fmt.Println("Please enter a valid email address.")
			}
            if !isValidTicketNumber {
				fmt.Println("Please enter a valid number of tickets.")
			} 
		}

	}

	wg.Wait() // this is a way to tell the main function to wait for the goroutines to finish their execution.
	
}

// greet user function.
func greetUser()  {
	// template string printing
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still remaining\n", conferenceTickets, remainingTickets)

	fmt.Println("Get your tickets here to attend")
}

// taking the names of the users and returning the first names of the users.
func getFirstNames() []string { // []string is the return type of the return value. In golang if the function return the value, then the return type
													// should be mentioned in the function declaration.	
	
	// declaring and defining an empty dynamic array - known as slice in golang.
	firstNames := []string{}
	
	// for-each loop in golang.
	// underscore (_) is the index of the array - underscore in golang refers to the variables which are to be ignored or
	// not in use
	// if you want to use the index of the array then you can use the index variable instead of underscore.
	for _, booking := range bookings {
		
		// strings.Fields splits the string on the blank space and returns an array of strings.
		// var names = strings.Fields(booking)
		// get the firstName from the split array.
		// var firstName = names[0]
		// appending the firstName to the firstNames array
		// firstNames = append(firstNames, booking["firstName"]) // Way to access the field from the map.
		firstNames = append(firstNames, booking.firstName) // Way to access the field from the structure.

	}

	return firstNames
}




func getUserInput() (string, string, string, uint) { // (string, string, string, uint) is the return type of the return value.
	// user basic details
	// declaring variables
	var firstName string
	var lastName string
	var email string
	var userTickets uint //uint type can only contain positive numbers

	// ask user for their first name.
	fmt.Println("Enter your first name: ")
	// fmt.Scan(&firstName) - this is a way to take the input from the user.
	// &firstName - this is a way to get the address of the variable firstName to store the value in the variable.
	// In golang - if you want to save the data in the variable then you need to pass the address of the variable.
	fmt.Scan(&firstName)

	// ask user for their last name.
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	// ask user for their email address.
	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	// ask user for no. of tickets required.
	fmt.Println("Enter your number of tickets: ")
	fmt.Scan(&userTickets)

	// returning multiple values in the return statement
	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	// remaining tickets after booking.
	remainingTickets = remainingTickets - userTickets
	
	// storing the value in the static sized array 
	// bookingUsers[0] = firstName + " " + lastName


	// creating a map to hold the data in key, value pairs
	// var userData = make(map[string]string)

	// // adding the user data to the map
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.Itoa(int(userTickets)) // converting uint to string
	
	// creating a struct object of the userData
	var userData = UserData{
		firstName: firstName,
        lastName: lastName,
        email: email,
        numberOfTickets: userTickets,
    }

	// storing the value in dynamic sized array
	bookings =  append(bookings, userData)

	fmt.Printf("List of bookings: %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)
}


func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	// sleep for 10 seconds
	time.Sleep(10 * time.Second)
	// fmt.Sprintf is a function that help in making statements but does not print anything and instead saves it in a variable.
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("##################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("##################")

	wg.Done() // this is a way to remove a goroutine from the wait group.
}