package main

import (
	"log"
	"net/rpc"
)

/*
RPC rules:
	- All functions must be methods
	- Must have two arguments
		* The first is what the client provides to the API
	     	and the second is what the API returns to the client
	- One of the return types must be error
*/

type Item struct {
	title string
	body string
}

type API int

var database []Item

// Retrieves an item and its index from database given its title
//  returns empty Item if not found
func (api *API) GetByTitle(title string, reply *Item) (int, error) {
	for i, item := range database {
		if item.title == title {
			*reply = item
			return i, nil
		}
	}
	return -1, nil
}

// Adds a new Item to the database and returns it
func (api *API) AddItem(item Item, reply *Item) error {
	database = append(database, item)
	*reply = item
	return nil
}

// Updates an Items body and returns the new item
//  if not found, returns empty item
func (api *API) UpdateItem(edit Item, reply *Item) error {
	var old_item Item
	idx, _ := api.GetByTitle(edit.title, &old_item)

	if idx != -1 {
		database[idx] = edit
		*reply = edit
	}

	return nil
}

// Removes an item from the database
// **Requires that item exists in database
func (api *API) DeleteItem(item Item, reply *Item) error {
	var old_item Item
	idx, _ := api.GetByTitle(item.title, &old_item)

	database = append(database[:idx], database[idx+1:] ...)
	*reply = old_item

	return nil
}

func main() {
	var api = new(API)  // new allocates memory for an empty struct
	err := rpc.Register((api))

	if err != nil {
		log.Fatal("error registering API", err)
	}

	// // Testing
	// fmt.Println("initial database: ", database)
	// a := Item{"one", "this is the first item"}
	// b := Item{"two", "second item here"}
	// c := Item{"three", "third item"}

	// var reply Item
	// api.AddItem(a, &reply)
	// api.AddItem(b, &reply)
	// api.AddItem(c, &reply)
	// fmt.Println("second database: ", database)

	// api.DeleteItem(b, &reply)
	// fmt.Println("third database: ", database)

	// api.UpdateItem(Item{"three", "changed third to second"}, &reply)
	// fmt.Println("fourth database: ", database)

	// var x, y Item
	// api.GetByTitle("three", &x)
	// api.GetByTitle("one", &y)
	// fmt.Println(x, y)
}