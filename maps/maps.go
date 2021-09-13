package maps

import "fmt"

func PrintMaps() string {
	m := make(map[string]int)

	m["k1"] = 3
	m["k2"] = 6

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1", v1)
	fmt.Println("length of map m:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)
	return "vin"
}
