package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Users struct which contains
// an array of users
type Users []User

// User struct which contains a name and email
type User struct {
	Name  string
	Email string
}

func getUsersFromFile() []byte {
	// Open jsonFile
	jsonFile, err := os.Open("users.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	fmt.Printf(string(byteValue))

	var users Users
	json.Unmarshal(byteValue, &users)

	// Loop over structs and display them.
	for l := range users {
		fmt.Printf("Name= %v, Email = %v", users[l].Name, users[l].Email)
		fmt.Println()
	}
	return byteValue
}
