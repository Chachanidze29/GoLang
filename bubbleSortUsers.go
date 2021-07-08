package main

import (
	"fmt"
)

type User struct {
	firstName string
	lastName  string
}

func main() {
	var users []User = []User{
		User{
			firstName: "Jon",
			lastName:  "Calhoun",
		},
		User{
			firstName: "Bob",
			lastName:  "Smith",
		},
		User{
			firstName: "Jhon",
			lastName:  "Smith",
		},
		User{
			firstName: "Joseph",
			lastName:  "Page",
		},
		User{
			firstName: "George",
			lastName:  "Costanza",
		},
		User{
			firstName: "Jerry",
			lastName:  "Seinfeld",
		},
		User{
			firstName: "Shane",
			lastName:  "Calhoun",
		},
	}
	fmt.Println("Our list of Users", users)
	for i := 0; i < len(users); i++ {
		if !sweep(users, i) {
			break
		}
	}
	fmt.Println("Our list of Users (sorted)", users)
}

func sweep(users []User, prevPasses int) bool {
	var N int = len(users)
	var firstIndex int = 0
	var secondIndex int = 1
	var didSwap bool = false
	for secondIndex < N-prevPasses {
		var firstUser User = users[firstIndex]
		var secondUser User = users[secondIndex]
		if firstUser.lastName > secondUser.lastName {
			users[firstIndex] = secondUser
			users[secondIndex] = firstUser
			didSwap = true
		}
		if firstUser.lastName == secondUser.lastName {
			if firstUser.firstName > secondUser.firstName {
				users[firstIndex] = secondUser
				users[secondIndex] = firstUser
				didSwap = true
			}
		}
		firstIndex++
		secondIndex++
	}
	return didSwap
}
