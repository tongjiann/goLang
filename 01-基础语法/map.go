package main

import "fmt"

func main() {
	m := map[string]string{"key": "value", "k2": "v2"}
	fmt.Println(m)
	m2 := make(map[string]string)
	fmt.Println(m2)
	m2["k1"] = "v1"
	fmt.Println(m2)
	s, b := m2["k1"]
	fmt.Println(s, b)
	fmt.Println(m2["k1"])
	delete(m2, "k2")
	delete(m2, "k1")
	fmt.Println(m2)
}
