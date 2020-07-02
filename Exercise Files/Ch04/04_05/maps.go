package main

import (
	"fmt"
	"sort"
)

// map is an ordered collection of key value pairs; or hashtable using strings for key and int for val usually
func main() {
	states := make(map[string]string)
	fmt.Println(states)
	states["WA"] = "Washington"
	states["CA"] = "Californina"
	states["TX"] = "Texas"
	fmt.Println(states)

	california := states["CA"]
	fmt.Println(california)
	
	delete(states, "TX")
	fmt.Println(states) 

	states["NY"] = "New York"

	for k, v := range states{
		fmt.Printf("%v: %v\n", k, v)
	}

	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println("\nSorted")
	for i:= range keys {
		fmt.Println(states[keys[i]])
	}

}