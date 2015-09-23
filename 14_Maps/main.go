package main

import (
	"fmt"
)




func main() {
// This map  is a map of words and their definitions
words := map[string]string{
"Ethnocentric: ": "evaluating other peoples and cultures according to the standards of one's own culture",
"Emic: ": "relating to or denoting an approach to the study or description of a particular language or culture in terms of its internal elements and their functioning rather than in terms of any existing external scheme",
"Etic: ": "relating to or denoting an approach to the study or description of a particular language or culture that is general, nonstructural, and objective in its perspective",
"Anthropology: ": "the study of humankind, in particular",
}


// This will loop through the map and print out the keys and values of that map
for k, v := range words {
fmt.Println(k, v)
}

// This is part of an in class exercise using maps
map1 := map[string]string{} // Creates and empty map
	map1["name"] = "MikeyRobby" // This line and the 2 following it are adding keys and setting their values
	map1["age"] = "21" // key = age. value = 21
	map1["major"] = "Computer Engineering"
	fmt.Println(map1) // Prints out the map
	map1["name"] = "Todd" // This changes the value in name from MikeyRobby to Todd
	fmt.Println(map1)
	delete(map1, "name") // This deletes the key "name" and the value that was in it
	fmt.Println(map1)
	// Print out the map using range
	for k,v := range map1  {
		fmt.Println(k,v)
	}
	fmt.Println(len(map1)) // Length of map1
	// Next is using the "Comma OK Idiom
	if val, exists := map1["major"]; exists {
		fmt.Println("This entry exists")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	}else {
		fmt.Println("This entry does not exist")
	}





}
