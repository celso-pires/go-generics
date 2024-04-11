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

	// With explicit typing
	fmt.Println(business.PrintGeneric[business.Solar](solar2k))
	fmt.Println(business.PrintGeneric[business.Wind](windWest))

	// With implicit typing
	fmt.Println(business.PrintGeneric(solar2k))
	fmt.Println(business.PrintGeneric(windWest))

	// Print a slice with a variable
	ss := []business.Solar{solar2k, solar3k}
	business.PrintSlice(ss)
	// Print a slice without a variable
	business.PrintSlice([]business.Wind{windWest, windWest})

	fmt.Println(costCalculation(windWest))
}

func costCalculation(p business.Coster) float64 {
	return p.Cost()
}
