package main

import (
	"fmt"
	"os"
)

func main() {
	var chatwork *Chatwork = NewChatwork(os.Getenv("CHATWORK_TOKEN"))
	var rooms []Room = chatwork.GetRooms()
	var members []Room
	for _, room := range rooms {
		switch room.Role {
		case "member":
			members = append(members, room)
		default:
			println("Other than member", room.Role, room.Name)
		}
	}
	fmt.Printf("Members %v\n", members)
}
