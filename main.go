package main

import (
	"fmt"
	"strings"
)

// Item ist die gemeinsame Schnittstelle für Blätter und Komposita.
type Item interface {
	Price() float64
	Display() string
}

// Glass ist ein Blatt im Composite Pattern.
type Glass struct {
	Name       string
	PriceValue float64
}

func (g Glass) Price() float64 {
	return g.PriceValue
}

func (g Glass) Display() string {
	return fmt.Sprintf("Glass: %s (%.2f EUR)", g.Name, g.Price())
}

// Frame ist ein Blatt im Composite Pattern.
type Frame struct {
	Name       string
	PriceValue float64
}

func (f Frame) Price() float64 {
	return f.PriceValue
}

func (f Frame) Display() string {
	return fmt.Sprintf("Frame: %s (%.2f EUR)", f.Name, f.Price())
}

// Window ist das Kompositum. Es kann beliebige Items enthalten, auch weitere
// Windows.
type Window struct {
	Name  string
	Items []Item
}

func (w *Window) Add(item Item) {
	w.Items = append(w.Items, item)
}

func (w Window) Price() float64 {
	var total float64
	for _, item := range w.Items {
		total += item.Price()
	}
	return total
}

func (w Window) Display() string {
	var lines []string
	lines = append(lines, fmt.Sprintf("Window: %s (%.2f EUR)", w.Name, w.Price()))
	for _, item := range w.Items {
		for _, line := range strings.Split(item.Display(), "\n") {
			lines = append(lines, "  "+line)
		}
	}
	return strings.Join(lines, "\n")
}

func main() {
	singleWindow := Window{Name: "Kleines Badezimmerfenster"}
	singleWindow.Add(Frame{Name: "Holzrahmen", PriceValue: 120.00})
	singleWindow.Add(Glass{Name: "Doppelverglasung", PriceValue: 180.00})

	largeWindow := Window{Name: "Großes Wohnzimmerfenster"}
	largeWindow.Add(Frame{Name: "Aluminiumrahmen", PriceValue: 250.00})
	largeWindow.Add(Glass{Name: "Dreifachverglasung", PriceValue: 420.00})
	largeWindow.Add(singleWindow) // Verschachteltes Kompositum

	fmt.Println(largeWindow.Display())
	fmt.Printf("\nGesamtpreis: %.2f EUR\n", largeWindow.Price())
}
