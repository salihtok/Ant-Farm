package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Room struct {
	ID          int
	Coordinates []int
	Connections []int
}

type Ant struct {
	ID          int
	StartRoom   int
	CurrentRoom int
}

func parseRooms(filename string) ([]Room, error) {
	var rooms []Room

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "##start") || strings.HasPrefix(line, "##end") {
			continue
		}

		parts := strings.Split(line, " ")
		roomID, _ := strconv.Atoi(parts[0])
		coordinates := []int{0, 0} // Default coordinates
		for i := 1; i <= 2; i++ {
			coordinates[i-1], _ = strconv.Atoi(parts[i])
		}

		connections := make([]int, 0)
		if len(parts) > 3 {
			for _, connStr := range parts[3:] {
				connID, _ := strconv.Atoi(connStr)
				connections = append(connections, connID)
			}
		}

		room := Room{
			ID:          roomID,
			Coordinates: coordinates,
			Connections: connections,
		}
		rooms = append(rooms, room)
	}

	return rooms, nil
}

func parseAnts(filename string) ([]Ant, error) {
	var ants []Ant

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "##start") || strings.HasPrefix(line, "##end") {
			continue
		}

		parts := strings.Split(line, " ")
		antID, _ := strconv.Atoi(parts[0])
		startRoom, _ := strconv.Atoi(parts[1])

		ant := Ant{
			ID:          antID,
			StartRoom:   startRoom,
			CurrentRoom: startRoom,
		}
		ants = append(ants, ant)
	}

	return ants, nil
}

func main() {
	filename := "input.txt" // Change this to the name of your input file

	rooms, err := parseRooms(filename)
	if err != nil {
		fmt.Println("Error parsing rooms:", err)
		return
	}

	ants, err := parseAnts(filename)
	if err != nil {
		fmt.Println("Error parsing ants:", err)
		return
	}

	// Print rooms
	fmt.Println("Rooms:")
	for _, room := range rooms {
		fmt.Printf("ID: %d, Coordinates: %v, Connections: %v\n", room.ID, room.Coordinates, room.Connections)
	}

	// Print ants
	fmt.Println("Ants:")
	for _, ant := range ants {
		fmt.Printf("ID: %d, StartRoom: %d\n", ant.ID, ant.StartRoom)
	}
}
