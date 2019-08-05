package main

import (
	"flag"
	"fmt"
	"log"
)

// GLOBAL FLAG VARIABLES //
var help bool
var listAll bool
var pick bool
var add string
var del string

// Init function on run //
func init() {
	flag.StringVar(&add, "a", "", "Add a restaurant by `name`")
	flag.StringVar(&del, "d", "", "Delete the restaurant by `name`")
	flag.BoolVar(&help, "h", false, "Display flag options and usage")
	flag.BoolVar(&listAll, "l", false, "List all possible restaurants")
	flag.BoolVar(&pick, "p", false, "Pick a random restaurant")
	flag.Parse()
}

// Main function
func main() {
	if help {
		flag.PrintDefaults()
	}

	if add != "" {
		addRestaurant(add)
	}

	if listAll {
		listAllRestaurants()
	}
}

/*
 * Function listAllRestaurants
 * Param: none
 * Returns: none
 * Description: 
 *		Handles the actions when the program is run with -l flag
 */
func listAllRestaurants() {
	// Create new array of Restaurant
	restaurants := []Restaurant{}

	// Retrieve all restaurants
	// If there was an error, log the error
	err := allRestaurants(&restaurants)
	if err != nil {
		log.Println(err)
	}

	// Loop through the list of restaurants and display the name of each one
	for _, r := range restaurants {
		fmt.Printf("%s\n", r.Name)
	}
}
