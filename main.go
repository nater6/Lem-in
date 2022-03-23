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
}

// add vertext
func (g *Graph) AddRoom(name string) {
	g.Rooms = append(g.Rooms, &Room{Roomname: name, adjacent: make(map[string]*Room)})
}

//add edge

func (g *Graph) AddLinks(from, to string) {
	// get vertex
	fromRoom := g.getRoom(from)
	toRoom := g.getRoom(to)

	// check
	if fromRoom == nil || toRoom == nil {
		err := fmt.Errorf("Room doesn't exsist (%v --- %v)", from, to)
		fmt.Println(err.Error())
	} else if contains(fromRoom.adjacent, to) || contains(toRoom.adjacent, from) {
		err := fmt.Errorf(" Exsisting Link (%v --- %v)", from, to)
		fmt.Println(err.Error())
	} else {
		//add edRoom.adjacent = apRoom.adjacent, toRoom)
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

}

func SortFiles(g *Graph) {
	file, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(file)
	start := false
	end := false
	i := 0

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {

		space := strings.Split(scanner.Text(), " ")

		if len(space) > 1 {
			g.AddRoom(space[0])
			i++
		}

		if start {
			g.startRoom = g.Rooms[i]
			start = false
		} else if end {
			g.endRoom = g.Rooms[i]
			end = false
		}

		hyphen := strings.Split(scanner.Text(), "-")
		if len(hyphen) > 1 {
			g.AddLinks(hyphen[0], hyphen[1])
			
		}
	}

	if scanner.Text() == "##start" {
		start = true
	}

	if scanner.Text() == "##end" {
		end = true
	}

}
