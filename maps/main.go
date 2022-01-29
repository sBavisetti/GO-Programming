package main

import "fmt"

func main() {

	// simple syntax to declare and intializing map
	colors := map[string]string{
		"red":   "#ff0000",
		"blue":  "#24ff11",
		"green": "#14abff",
	}

	fmt.Println(colors)
	printMap(colors)

	// iterating all the values in the map

	// var syntax for declaration and initialization of maps
	var food map[string]string
	fmt.Println(food)

	// using make to creat maps
	computers := make(map[string]string)
	computers["laptop"] = "rog"
	computers["desktop"] = "alienware"
	computers["mobile"] = "oneplus"

	fmt.Println(computers)

	// deleting keys and values from the map
	delete(computers, "laptop")
	fmt.Println(computers)

}

func printMap(c map[string]string) {

	for key, value := range c {

		fmt.Println("key : ", key, "value : ", value)
	}
}
