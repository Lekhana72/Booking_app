package main

import (
	"fmt"
	"sync"
	"time"
)

// Best Practice: define variable as "local" as possible
var conferenceName string = "Go Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// removed for loop so the time , goroutine is never excuted

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()

	isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		bookTicket(userTickets, firstName, lastName, email)
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("These are all first names: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out. Come back next year.")

		}

	} else {
		if !isValidName {
			fmt.Println("first name or last name you entered is too short")

		}
		if !isValidEmail {
			fmt.Println("email you entered is not valid")
		}
		if !isValidTicketNumber {
			fmt.Println("number of tickets you entered is not valid")
		}

	}

	city := "London"

	switch city {
	case "New York":
		fmt.Println("The city is New York")
	case "Singapore":
		fmt.Println("The city is Singapore")
	case "London":
		fmt.Println("The city is London")
	default:
		fmt.Println("The city is not listed")
	}
	wg.Wait()

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")

}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {

		firstNames = append(firstNames, booking.firstName)
	}

	// fmt.Printf("These are all first names: %v\n", firstNames)
	return firstNames
}

// func validateUserInput(firstname string, lastName string, email string, userTickets uint) (bool, bool, bool) {
// 	isValidName := len(firstname) >= 2 && len(lastName) >= 2
// 	isValidEmail := strings.Contains(email, "@")
// 	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
// 	return isValidName, isValidEmail, isValidTicketNumber
// }

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your firstname:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email:")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create a map for a user
	// var myslice []string
	// var mymap map[string]string

	// var userData = make(map[string]string)
	//
	// STRUCT : objects are below
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	// NO NEED TO CONVERT TO STRING BECAUSE WE ARE USING STRUCT
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of Bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets . You will reeive a confirmation email at  %v\n", firstName, lastName, userTickets, email)

	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

// Encapsulated the logic , which belongs together
/* greetUsers()
getUserInput()
validateUserInput()
bookTicket()
getFirstNames() */

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#####################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#####################")
	wg.Done()
}

// "time" - functionality for time
// The sleep function stops or blocks the current "thread" (goroutine) execution for the defined duration

// Waitgroup
// - Waits for the launched goroutine to finish
// - Package "sync" provides basic synchronization functionality
// - Add : Sets the number of goroutines to wait for (increases the counter by the provided number)
// - Wait: Blocks until the WaitGroup counter is 0

// Why? what exactly is different?
/* Goroutine
- Go is using, whats called a "Green thread"
-Abstraction of an actual thread
- Managed by the go runtime, we are only interacting with these high level goroutines
- Cheaper & lightweight
- You can run hundreds of thousands or millions goroutines without affecting the performace
/*

/*OS Thread
- Managed by kernel
- Are hardware dependent
- Cost of these threads are higher
- Higher start up time
*/
