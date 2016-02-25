package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
)

type Game struct {
	Title      string
	Publisher  string
	Developer  string
	Released   string
	Metacritic int
	Platform   string
}

type byTitle []*Game

// byTitle implements the sort.Interface by defining three methods Len(), Less() and Swap() on the type

// sort.Interface
// type Interface interface {
// Len() int             // Len is the number of elements in a collection
// Less(i, j int) bool   // Less reports true if i is less than j
// Swap(i, j int)        // Swaps the elements with indexes i and j
// }

func (x byTitle) Len() int {
	return len(x)
}

func (x byTitle) Less(i, j int) bool {
	return x[i].Title < x[j].Title
}

func (x byTitle) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

var games = []*Game{
	{"Little Big Planet", "Sony Computer Entertainment", "Media Molecule", "2008-Oct-21", 95, "PS3"},
	{"The Last of Us", "Sony Computer Entertainment", "Naughty Dog", "2013-Jun-14", 95, "PS3"},
	{"Ni no Kuni: Wrath of the White Witch", "Bandai Namco Entertainment", "Level-5", "2010-Dec-09", 85, "PS3"},
}

func main() {

	values := []int{5, 7, 2, 3, 7, 8, 0, 4}
	fmt.Printf("Values are: %v \n", values)
	fmt.Printf("sort.IntsAreSorted(values) %v \n", sort.IntsAreSorted(values))

	sort.Ints(values)

	fmt.Printf("After sort.Ints(values) %v \n", values)
	fmt.Printf("sort.IntsAreSorted(values) %v \n", sort.IntsAreSorted(values))

	sort.Sort(sort.Reverse(sort.IntSlice(values)))

	fmt.Printf("After sort.Sort(sort.Reverse(sort.IntSlice(values))) %v \n", values)
	fmt.Println()

	fmt.Println("a list of some games:")
	printGames(games)
	fmt.Println()

	fmt.Println("games after sort.Sort(byTitle(games))")
	sort.Sort(byTitle(games))
	printGames(games)
	fmt.Println()

	fmt.Println("games after applying reverse function -	sort.Sort(sort.Reverse(byTitle(games))) ")
	sort.Sort(sort.Reverse(byTitle(games)))
	printGames(games)
}

func printGames(games []*Game) {
	const format = "%v\t%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 4, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Publisher", "Developer", "Released", "Metacritic", "Platform")
	fmt.Fprintf(tw, format, "-----", "---------", "---------", "--------", "----------", "--------")
	for _, t := range games {
		fmt.Fprintf(tw, format, t.Title, t.Publisher, t.Developer, t.Released, t.Metacritic, t.Platform)
	}
	tw.Flush() // calculate column widths and print table
}
