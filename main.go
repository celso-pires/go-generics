package main

import (
	"fmt"
	"generics/business"
)

func main() {

	// main is our simple "playground" for the course.
	// Note, that in production code, it is a good practice to keep the main function short.
	solar2k := business.Solar{Name: "Solar 2000", Netto: 4.500}
	solar3k := business.Solar{Name: "Solar 3000", Netto: 4.000}
	windWest := business.Wind{Name: "Wind West", Netto: 3.950}

	// Print details for each energy offer with kineteco branding
	fmt.Println(solar3k.Print())
	fmt.Println(solar2k.Print())
	fmt.Println(windWest.Print())
}
