package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"
)

// GLOBAL FLAG VARIABLES //
var help bool
var list bool
var pick bool
var add bool
var del bool
var cat string
var rest string

// Init function on run //
func init() {
	flag.BoolVar(&add, "a", false, "Use add mode")
	flag.BoolVar(&del, "d", false, "Use delete mode")
	flag.BoolVar(&help, "h", false, "Display flag options and usage")
	flag.BoolVar(&list, "l", false, "List all possible restaurants")
	flag.BoolVar(&pick, "p", false, "Pick a random restaurant")
	flag.StringVar(&cat, "c", "", "Specify the `name` of the category to add, delete, pick from, or list")
	flag.StringVar(&rest, "r", "", "Specify the `name` of the restaurant to add, or delete")
	flag.Parse()
}

// Main function
func main() {
	// Check different modes
	// help = Help mode
	// pick = Pick mode
	// list = List mode
	// add = Add mode
	// del = Delete mode
	if help {
		flag.PrintDefaults()
	} else if pick {
		if cat != "" {
			// TODO: Pick restaurant based on category
		} else {
			pickRandomRestaurant()
		}
	} else if list {
		if cat != "" {
			// TODO: List restaurants in a category
		} else {
			// List all restaurants
			listAllRestaurants()
		}
	} else if add {
		if rest != "" && cat != "" {
			// TODO: Add restaurant with specified category
		}else if rest != "" {
			// Add restaurant with default category "Other"
			addRestaurant(rest)
		} else if cat != "" {
			// TODO: Add category
			addCategory(cat)
		}
	} else if del {
		if rest != "" {
			// TODO: Delete restaurant
		} else if cat != "" {
			// TODO: Delete category
		}
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
		return
	}

	// Loop through the list of restaurants and display the name of each one
	for _, r := range restaurants {
		fmt.Printf("%s\n", r.Name)
	}
}

/*
 * Function pickRandomRestaurant
 * Param: none
 * Returns: none
 * Description:
 *		Handles the actions when the program runs with the -p flag
 */
func pickRandomRestaurant() {
	// We need to retrive all of the restaurants
	restaurants := []Restaurant{}
	err := allRestaurants(&restaurants)
	if err != nil {
		log.Println(err)
		return
	}

	// Create a new source for random number using time
	// Use that soure to generate a random Int from 0 to restaurants length - 1
	// That integer will be the index of the restaurant selected
	source := rand.NewSource(time.Now().Unix())
	random := rand.New(source)
	number := random.Intn(len(restaurants))
	restaurant := restaurants[number]
	fmt.Println(restaurant.Name)

}
