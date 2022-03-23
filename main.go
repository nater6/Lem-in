package main

import (
	"bufio"
	"fmt"
	"lemin/lemin"
	"os"
	"strconv"
	"strings"
)

type Room struct {
	Roomname string
	x        float64
	y        float64
}

type Link struct {
	startroom string
	endroom   string
	distance  float64
}

type Path struct {
	start bool
	final bool
}

func main() {
	Rooms, Links := SortFiles()
	Links = FindDistances(Rooms, Links)
	
	fmt.Println(Links)
}

func SortFiles() ([]Room, []Link){
	file, _:= os.Open(os.Args[1])
	scanner := bufio.NewScanner(file)
	
	scanner.Split(bufio.ScanLines)
	
	var Rooms []Room
 var Links []Link
	for scanner.Scan(){
		space := strings.Split(scanner.Text(), " ")
		if len(space) > 1 {
			//Roomname
			newx, _:= strconv.ParseFloat(space[1], 64)
			newy, _:= strconv.ParseFloat(space[2], 64)
			NewRoom := Room {Roomname: space[0], x: newx, y: newy}
			Rooms = append(Rooms, NewRoom)
		} 
		hyphen := strings.Split(scanner.Text(), "-")

		if len(hyphen) > 1 {
			//Link
			NewLink := Link{startroom:hyphen[0], endroom:hyphen[1]}
			Links = append(Links, NewLink)
		} 
	}

	return Rooms, Links
}

func FindDistances(Rooms []Room, Links []Link) []Link{
	var x1 float64
	var y1 float64
	var x2 float64
	var y2 float64

for i,l := range Links{
		for _,r := range Rooms {
			if r.Roomname == l.startroom {
				//x1, x2
				x1 = r.x
				y1 = r.y
			}
			if r.Roomname == l.endroom {
				//x2, y2
				x2 = r.x
				y2 = r.y
			}
		}
		distance:= lemin.DistanceCalc(x1,y1,x2,y2)
		p:= &Links[i]
		p.distance = distance
	}
		
	return Links
}
