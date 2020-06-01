package main

import "fmt"

func maps() {
	// maps
	myMap := make(map[string]int)
	myMap["key"] = 15
	myMap["nameCode"] = 2165
	name, ok := myMap["key1"]
	delete(myMap, "key")

	fmt.Println(myMap, len(myMap), name, ok)

}
