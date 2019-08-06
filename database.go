package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

/*
 * Struct Restaurant
 * ID int
 * Name string
 */
type Restaurant struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

/*
 * Struct Category
 * ID int
 * Name string
 */
type Category struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

/*
 * Struct RestaurantCategory
 * ID int
 * RestaurantID int
 * CategoryID int
 */
type RestaurantCategory struct {
	ID   int    `db:"id"`
	RestaurantID int `db:"restaurantid"`
	CategoryID int `db:"categoryid"`
}

/*
 * Function connectToDb
 * Param: none
 * Return: *sqlx.DB, error
 * Description:
 *		Connects the Sqlite3 database using the string path to the .db file
 *		On success, returns a reference to the db connection
 *		On failure, returns the error
 */
func connectToDb() (*sqlx.DB, error) {
	return sqlx.Connect("sqlite3", "./data/randomrestaurantpicker.db")
}

/*
 * Function allRestaurant
 * Param: restaurants *[]Restaurant
 * Return: error
 * Description:
 * 		Retrieves all restaurants from the database
 *		Accepts a reference to an array of Restaurant
 *		Upon failure, it will return the error
 */
func allRestaurants(restaurants *[]Restaurant) error {
	// Create the databse connection
	// Log and return the error if there is an error connecting
	db, err := connectToDb()
	if err != nil {
		log.Println(err)
		return err
	}

	// Find all restaurants in the database and order them by name alphabetically from A - Z
	// If there is an error retrieving results, log and return the error
	err = db.Select(restaurants, "SELECT * FROM restaurant ORDER BY name ASC")
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

/*
 * Function addRestaurant
 * Param: restaurant string
 * Return: error
 * Description:
 * 		Adds a new restaurant to the database by the restaurant name
 *		Returns an error if the record could not be added
 */
func addRestaurant(restaurant string) error {
	// Create the database connection
	// Log and raturn the error if there was an error connecting
	db, err := connectToDb()
	if err != nil {
		log.Println(err)
		return err
	}

	// Create a new Resturant struct r with the default ID value 0, and Name restaurant
	r := Restaurant{0, restaurant}

	// Insert the new restaurant struct into the database
	// If there was an error inserting the record, log and return the error
	_, err = db.NamedExec(`INSERT INTO restaurant (name) VALUES (:name)`, r)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

/*
 * Function addCategory
 * Param: category string
 * Return: error
 * Description:
 * 		Adds a new category to the database by the category name
 *		Returns an error if the record could not be added
 */
func addCategory(category string) error {
	// Create the database connection
	// Log and raturn the error if there was an error connecting
	db, err := connectToDb()
	if err != nil {
		log.Println(err)
		return err
	}

	// Create a new Category struct c with the default ID value 0, and Name category
	c := Category{0, category}

	// Insert the new restaurant struct into the database
	// If there was an error inserting the record, log and return the error
	_, err = db.NamedExec(`INSERT INTO category (name) VALUES (:name)`, c)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
