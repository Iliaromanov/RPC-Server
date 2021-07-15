package main

import "fmt"

type Item struct {
	title string
	body string
}

type API int

var database []Item

// Retrieves an item and its index from database given its title
//  returns empty Item if not found
func GetByTitle(title string) (int, Item) {
	for i, item := range database {
		if item.title == title {
			return i, item
		}
	}
	return -1, Item{}
}

// Adds a new Item to the database and returns it
func AddItem(item Item) Item {
	database = append(database, item)
	return item
}

// Updates an Items body and returns the new item
//  if not found, returns empty item
func UpdateItem(title string, edit Item) Item {
	idx, _ := GetByTitle(title)

	if idx != -1 {
		database[idx] = edit
		return edit
	}

	return Item{}
}

// Removes an item from the database
// **Requires that item exists in database
func DeleteItem(item Item) Item {
	idx, _ := GetByTitle(item.title)

	database = append(database[:idx], database[idx+1:] ...)

	return item
}

func main() {
	// Testing
	fmt.Println("initial database: ", database)
	a := Item{"one", "this is the first item"}
	b := Item{"two", "second item here"}
	c := Item{"three", "third item"}

	AddItem(a)
	AddItem(b)
	AddItem(c)
	fmt.Println("second database: ", database)

	DeleteItem(b)
	fmt.Println("third database: ", database)

	UpdateItem("three", Item{"second", "changed third to second"})
	fmt.Println("fourth database: ", database)

	_, x := GetByTitle("second")
	_, y := GetByTitle("one")
	fmt.Println(x, y)
}
