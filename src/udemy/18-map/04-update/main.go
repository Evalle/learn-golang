package main

import "fmt"

func main() {
	m := map[string]string{
		"Tod":    "teacher",
		"Evgeny": "student",
	}
	fmt.Println(m)
	m["Tod"] = "professor"
	fmt.Println(m)
}
