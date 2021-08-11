package main

import "fmt"

func main() {
	emails := map[string]string{"Sayani Mallick": "sayanimallick0218@gmail.com", "Sudipta Mallick": "mallicksudipta95@gmail.com", "Amit Kumar Mallick": "amitmallick01@gmail.com"}
	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}
}
