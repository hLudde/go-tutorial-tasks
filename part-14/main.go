package main

import "fmt"

func main() {
	carPrices := make(map[string]int)
	carPrices["Tesla Model 3"] = 367700  //nok
	carPrices["Tesla Model S"] = 713000  //nok
	carPrices["Tesla Model X"] = 768000  //nok
	carPrices["Tesla Model Z"] = 1213000 //nok

	for k, v := range carPrices {
		fmt.Printf("The %s costs %d\n", k, v)
	}

	//oppps the Tesla Model Z doesn't exist yet :o

	delete(carPrices, "Tesla Model Z")
	fmt.Println("===============================")
	for k, v := range carPrices {
		fmt.Printf("The %s costs %d\n", k, v)
	}

}
