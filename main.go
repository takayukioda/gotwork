package main

import (
	"fmt"
	"os"
)

func main() {
	var chatwork *Chatwork = NewChatwork(os.Getenv("CHATWORK_TOKEN"))
	var rooms []Room = chatwork.GetRooms()
	for _, room := range rooms {
		fmt.Printf("%#v\n", room)
	}
}
