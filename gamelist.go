package main

import (
	"fmt"
	"sort"
  "os"
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

var games = []*Game{
	{"Little Big Planet", "Sony Computer Entertainment", "Media Molecule", "2008-Oct-21", 95, "PS3"},
	{"The Last of Us", "Sony Computer Entertainment", "Naughty Dog", "2013-Jun-14", 95, "PS3"},
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
