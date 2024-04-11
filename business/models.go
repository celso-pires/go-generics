package business

import "fmt"

// Solar handles all the different energy offers powered by solar.
type Solar struct {
	Name  string
	Netto float64
}

// Wind handles all the different energy offers powered by wind.
type Wind struct {
	Name  string
	Netto float64
}

type Energy interface {
	Solar | Wind
	// Cost() float64
}

type Coster interface {
	Cost() float64
}

// Print prints the information for a solar product.
// The string is enriched with the required kineteco legal information.
func (s Solar) Print() string {
	return fmt.Sprintf("%s - %v\n", kinetecoPrint, s)
}

// Print prints the information for a wind product.
// The string is enriched with the required kineteco legal information.
func (w Wind) Print() string {
	return fmt.Sprintf("%s - %v\n", kinetecoPrint, w)
}

// PrintGeneric prints the information for a product of type T.
// T must implement the Energy interface.
// The string is enriched with the required kineteco legal information.
func PrintGeneric[T Energy](t T) string {
	// T must implement the Energy interface
	// to ensure that it has a Cost method.

	return fmt.Sprintf("%s - %v\n", kinetecoPrint, t)
}

// PrintSlice prints the information for a slice of energy offers.
// The information for each offer is printed on a new line,
// with the index of the offer and the information as returned by
// PrintGeneric.
func PrintSlice[T Energy](tt []T) {
	for i, t := range tt {
		// Ensure that T implements the Energy interface,
		// which is required to have a Cost method.
		fmt.Printf("%d: %s\n", i, PrintGeneric[T](t))
	}
}

var kinetecoPrint string = "Kineteco Deal:"
