package main

import "fmt"

var totalseats int = 50
var remainingseats uint = 50

type tickets struct {
	f_name    string
	l_name    string
	tickit_no uint
	t_date    string
	email     string
}

func main() {
	for {
		greet()

		f_name, l_name, tickit_no, t_date, email := input()
		bookTicket(f_name, l_name, tickit_no, t_date, email)
	}
}

func greet() {
	fmt.Println("Welcome to the Movie Booking system")
	fmt.Printf("Total seats available are %v and remaining are %v\n",
		totalseats, remainingseats)
}

func input() (string, string, uint, string, string) {
	var user tickets

	fmt.Println("Enter first name:")
	fmt.Scan(&user.f_name)

	fmt.Println("Enter last name:")
	fmt.Scan(&user.l_name)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&user.tickit_no)

	fmt.Println("Enter booking date:")
	fmt.Scan(&user.t_date)

	fmt.Println("Enter email:")
	fmt.Scan(&user.email)

	return user.f_name, user.l_name, user.tickit_no, user.t_date, user.email
}

func bookTicket(f_name string, l_name string, tickit_no uint, t_date string, email string) {
	if tickit_no > remainingseats {
		fmt.Println("Not enough seats available")
		return
	}

	remainingseats -= tickit_no

	fmt.Println("Ticket Booked Successfully!")

	fmt.Println("Name:", f_name, l_name)
	fmt.Println("Tickets:", tickit_no)
	fmt.Println("Date:", t_date)
	fmt.Println("Email:", email)
	fmt.Printf("Remaining Seats: %v\n", remainingseats)
}
