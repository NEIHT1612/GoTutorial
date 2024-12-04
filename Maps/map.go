package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google": "https://google.com",
		"Amazon": "https://amazon.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Google"])
	websites["LinkedIn"] = "https://linkedin/com"
	fmt.Println(websites)
	delete(websites, "Google")
	fmt.Println(websites)
	for index, val := range websites{
		fmt.Println("Index:" + index + ", Value:" + val)
	}
}