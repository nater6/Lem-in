package lemin

import (
	"bufio"
	"strings"
)

type Rooms struct {
	Roomname string
	x        string
	y        string
}

type Link struct {
	startroom string
	endroom   string
	distnace  float64
}

type Path struct {
	start bool
	final bool
}

func SortFiles() {
	scanner := bufio.NewScanner(os.Args[1])
	scanner.Split(bufio.ScanLines)

	for scanner.Scan(){
		space := strings.Split(scanner.Text(), " ")

		if len(lines) > 1 {
			//Roomname
		} 

		lines := strings.Split(scanner.Text(), "-")

		if len(lines) > 1 {
			//Link
		} 
	}

)