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
	startRoom string
	endRoom   string
}

// vertex represents graph vertex
type Room struct {
	Roomname string
	adjacent []string
	visited  bool
}

// add vertext
func (g *Graph) AddRoom(name string) {
	g.Rooms = append(g.Rooms, &Room{Roomname: name, adjacent: []string{}, visited: false})
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
	} else if fromRoom.Roomname == g.endRoom {
		//Checking for the endroom
		toRoom.adjacent = append(toRoom.adjacent, fromRoom.Roomname)
	} else if toRoom.Roomname == g.endRoom {
		//Checking for the endroom
		fromRoom.adjacent = append(fromRoom.adjacent, toRoom.Roomname)
	} else if toRoom.Roomname == g.startRoom {
		//Checking for the startroom
		toRoom.adjacent = append(toRoom.adjacent, fromRoom.Roomname)
	} else if fromRoom.Roomname == g.startRoom {
		//Checking for the startroom
		fromRoom.adjacent = append(fromRoom.adjacent, toRoom.Roomname)
	} else if fromRoom.Roomname != g.endRoom && toRoom.Roomname != g.endRoom {

		fromRoom.adjacent = append(fromRoom.adjacent, toRoom.Roomname)
		toRoom.adjacent = append(toRoom.adjacent, fromRoom.Roomname)

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
func contains(s []string, name string) bool {
	for _, v := range s {
		if name == v {
			return true
		}
	}
	return false
}

func main() {

	list1 := []*Room{}

	roomList := &Graph{Rooms: list1}

	SortFiles(roomList)

	for _, r := range roomList.Rooms {
		fmt.Println(r.Roomname)
		fmt.Println(r.adjacent)

	}

	for _, v := range roomList.Rooms {
		x := *v
		fmt.Print(x.Roomname + "|")

	}
	fmt.Println()

	allPaths := [][]string{}
	path := []string{}
	fmt.Println("startroom: " + roomList.startRoom)
	fmt.Println("endroom: " + roomList.endRoom)

	FindPath(roomList.startRoom, roomList.endRoom, roomList, path, allPaths)

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
			g.startRoom = g.Rooms[i-1].Roomname
			start = false
		} else if end {
			g.endRoom = g.Rooms[i-1].Roomname
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

func FindPath(current, end string, g *Graph, path []string, pathList [][]string) {

	fmt.Println("Current Room: " + current)

	//Check if the current room is the end room
	if current == end {
		path = append(path, end)
		fmt.Printf("Path: %v \n", path)
		return
	}

	//Make new Path variable to append the current room to
	path1 := path
	curr := g.getRoom(current)
	path1 = append(path1, current)

	// Mark the room as visited
	curr.visited = true
	anyAdj := false

	//Loop through adjacent rooms and see if the end room is present or if there are any unvisited rooms
	for i := 0; i < len(curr.adjacent); i++ {

		y := g.getRoom(curr.adjacent[i])

		if !y.visited {
			anyAdj = true
		}

		// If the end room is present in the adjacent room move it to the start of the slice
		if curr.adjacent[i] == g.endRoom {
			curr.adjacent[0], curr.adjacent[i] = curr.adjacent[i], curr.adjacent[0]

		}
	}

	if !anyAdj {
		curr.visited = false
		return
	}

	// recurssively call the func to the end
	for i := 0; i < len(curr.adjacent); i++ {
		//When back at the startroom's adjacent rooms reset all rooms to unvisited
		if curr.Roomname == g.startRoom {
			for _, v := range g.Rooms {
				v.visited = false
			}
		}

		//Get information for the current room
		x := g.getRoom(curr.adjacent[i])

		if x.visited {
			//x.visited = false
			fmt.Println("Previously visited: " + x.Roomname)
			continue
		} else if !x.visited {
			// fmt.Println("Next Room")

			FindPath(x.Roomname, end, g, path1, pathList)
		}

	}
}

