package main

import (
	"fmt"
	"sort"
)

func main() {
	// Map
	states := make(map[string]string)
	states["DE"] = "Delaware"
	states["PA"] = "Pennsylvania"
	states["NJ"] = "New Jersey"
	states["CA"] = "California"
	fmt.Println(states)

	california := states["CA"]
	fmt.Println(california)

	delete(states, "PA")
	fmt.Println(states)

	states["NY"] = "New York"
	fmt.Println(states)

	for k, v := range states {
		fmt.Printf("%v = %v\n", k, v)
	}

	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println("\nSorted")

	for _, k := range keys {
		fmt.Printf("%v = %v\n", k, states[k])
	}
}
