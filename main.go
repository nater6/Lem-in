package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//graph structre
// graph is adjacency list
type Graph struct {
	Rooms     []*Room
	startRoom *Room
	endRoom   *Room
}

// vertex represents graph vertex
type Room struct {
	Roomname string
	adjacent map[string]*Room
	visited	bool
}

// add vertext
func (g *Graph) AddRoom(name string) {
	g.Rooms = append(g.Rooms, &Room{Roomname: name, adjacent: make(map[string]*Room), visited: false})
}

//add edge

func (g *Graph) AddLinks(from, to string) {
	// get vertex
	fromRoom := g.getRoom(from)
	toRoom := g.getRoom(to)

	// check
	if fromRoom == nil || toRoom == nil {
		//If either of the rooms dont exsist
		err := fmt.Errorf("Room doesn't exsist (%v --- %v)", from, to)
		fmt.Println(err.Error())
	} else if contains(fromRoom.adjacent, to) || contains(toRoom.adjacent, from) {
		// if the link already exsists
		err := fmt.Errorf(" Existing Link (%v --- %v)", from, to)
		fmt.Println(err.Error())
	} else if fromRoom == g.endRoom{
		//Checking for the endroom
		toRoom.adjacent[fromRoom.Roomname] = fromRoom
	} else if toRoom == g.endRoom {
		//Checking for the endroom
		fromRoom.adjacent[toRoom.Roomname] = toRoom
	} else if toRoom == g.startRoom {
		//Checking for the startroom
		toRoom.adjacent[fromRoom.Roomname] = fromRoom
	} else if fromRoom == g.startRoom {
		//Checking for the startroom
		fromRoom.adjacent[toRoom.Roomname] = toRoom
	} else if fromRoom != g.endRoom && toRoom != g.endRoom {
		
		fromRoom.adjacent[toRoom.Roomname] = toRoom
		toRoom.adjacent[fromRoom.Roomname] = fromRoom
	} 

}

// get vertex
func (g *Graph) getRoom(name string) *Room {
	for i, v := range g.Rooms {
		if v.Roomname == name {
			return g.Rooms[i]
		}
	}
	return nil
}

// contains
func contains(s map[string]*Room, name string) bool {
	for _, v := range s {
		if name == v.Roomname {
			return true
		}
	}
	return false
}

func main() {

	list1 := []*Room{}

	roomList := &Graph{Rooms: list1}

	SortFiles(roomList)

	for _, r := range roomList.Rooms{
		fmt.Println(r.Roomname)
		fmt.Println(r.adjacent)
		
	}
	
	for _, v:= range roomList.Rooms {
		x := *v
		fmt.Print(x.Roomname + "|")
	}

}



func PathFinder(g *Graph) [][]*Room {

// 	//Begin from the start room and check which rooms it is linked to
	start := g.startRoom
	final := [][]*Room{}
	path := []*Room{}

	for start.adjacent != nil {
		for i, v := range start.adjacent {
			if !v.visited {
				path = append(path, start.adjacent[i])
				// 	//As each room is checked mark it as visited so it isn't checked again
				start.visited = true
				start = start.adjacent[i]
			}
		// if v == g.endRoom {
		// 	final = append(final, path)
		// } 
	}
	}
	


// 	//Once a path is complete, run again
// 	// if the same path is found again,
	return final
}

func SortFiles(g *Graph) {
	file, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(file)
	start := false
	end := false
	i := 0

	scanner.Split(bufio.ScanLines)

	

	for scanner.Scan() {
		x := scanner.Text()
		fmt.Println(x)

		space := strings.Split(scanner.Text(), " ")

		if len(space) > 1 {
			g.AddRoom(space[0])
			i++
		}

		if start {
			g.startRoom = g.Rooms[i-1]
			start = false
		} else if end {
			g.endRoom = g.Rooms[i-1]
			end = false
		}

		hyphen := strings.Split(scanner.Text(), "-")
		if len(hyphen) > 1 {
			g.AddLinks(hyphen[0], hyphen[1])
			
		}
	

	if x == "##start" {
		start = true
		fmt.Println("Start")
		
	}

	if x == "##end" {
		end = true
	}

}

}

