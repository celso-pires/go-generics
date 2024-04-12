package main

import "fmt"

// Generic CleanEnergy
type CleanEnergy interface {
	*Solar | *Wind
	Cost() float64
	GetCommonValues() EnergyCommonAttributes
	GetCleanValues() (name string, netto float64)
}

type EnergyCommonAttributes struct {
	Name  string
	Netto float64
}

// Generic Energy
type Energy interface {
	Cost() float64
	GetCommonValues() EnergyCommonAttributes
}

// Petrol

type Petrol struct {
	Name  string
	Netto float64
}

func (p *Petrol) Cost() float64 {
	return p.Netto * 0.8
}

func (p *Petrol) GetCommonValues() EnergyCommonAttributes {
	return EnergyCommonAttributes{p.Name, p.Netto}
}

// Solar

type Solar struct {
	Name  string
	Netto float64
}

func (s *Solar) Cost() float64 {
	return s.Netto * 0.4
}

func (s *Solar) GetCommonValues() EnergyCommonAttributes {
	return EnergyCommonAttributes{s.Name, s.Netto}
}

func (s *Solar) GetCleanValues() (Name string, Netto float64) {
	Name = s.Name
	Netto = s.Netto
	return
}

// Wind
type Wind struct {
	Name  string
	Netto float64
}

func (w *Wind) Cost() float64 {
	return w.Netto * 0.3
}

func (w *Wind) GetCommonValues() EnergyCommonAttributes {
	return EnergyCommonAttributes{w.Name, w.Netto}
}

func (w *Wind) GetCleanValues() (Name string, Netto float64) {
	Name = w.Name
	Netto = w.Netto
	return
}

// Main
func main() {
	w := &Wind{Name: "Wind Energy", Netto: 5000}
	s := &Solar{Name: "Solar Energy", Netto: 4000}
	PrintCleanEnergy(w)
	PrintCleanEnergy(s)

	// Because of type constraints petrol is not CleanEnergy
	p := &Petrol{Name: "Petrol Energy", Netto: 9000}
	PrintEnergy(p)
	PrintEnergy(w)
	PrintEnergy(s)
	PrintSpecific(w)
	PrintSpecific(s)
}

func PrintCleanEnergy[T CleanEnergy](e T) {
	fmt.Printf("The total cost for %s is: %.2f\n", e.GetCommonValues().Name, e.Cost())
	fmt.Printf("The fields of %s is: %v\n", e.GetCommonValues().Name, e.GetCommonValues())
}

func PrintSpecific[T CleanEnergy](e T) {
	name, netto := e.GetCleanValues()
	fmt.Printf("Name: %s and Netto: %.2f\n", name, netto)
}

func PrintEnergy[T Energy](e T) {
	fmt.Printf("The total cost for %s is: %.2f\n", e.GetCommonValues().Name, e.Cost())
	fmt.Printf("The fields of %s is: %v\n", e.GetCommonValues().Name, e.GetCommonValues())
}
