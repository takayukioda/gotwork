package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Chatwork struct {
	client       http.Client
	rootEndpoint string
	token        string
}

type Room struct {
	RoomId         int    `json:"room_id"`
	Name           string `json:"name"`
	Type           string `json:"type"`
	Role           string `json:"role"`
	Sticky         bool   `json:"sticky"`
	UnreadNum      int    `json:"unread_num"`
	MentionNum     int    `json:"mention_num"`
	MytaskNum      int    `json:"mytask_num"`
	MessageNum     int    `json:"message_num"`
	FileNum        int    `json:"file_num"`
	TaskNum        int    `json:"task_num"`
	IconPath       string `json:"icon_path"`
	LastUpdateTime int    `json:"last_update_time"`
}

func NewChatwork(token string) *Chatwork {
	var client http.Client = http.Client{
		Timeout: 30 * time.Second,
	}
	return &Chatwork{
		client:       client,
		rootEndpoint: "https://api.chatwork.com/v2",
		token:        token,
	}
}

func (c *Chatwork) endpoint(path string) string {
	var rooturl *bytes.Buffer = bytes.NewBufferString(c.rootEndpoint)
	if path[0:1] != "/" {
		rooturl.WriteString("/")
	}

	return fmt.Sprintf("%s%s", rooturl.String(), path)
}

func (c *Chatwork) request(method string, path string) (*http.Response, error) {
	req, err := http.NewRequest(method, c.endpoint(path), nil)
	if err != nil {
		println("Failed on creating new request")
		log.Fatal(err)
	}
	req.Header.Add("User-Agent", "chatwork-client by Go")
	req.Header.Add("X-ChatWorkToken", c.token)

	return c.client.Do(req)
}

func (c *Chatwork) GetRooms() []Room {
	const (
		method = "GET"
		path   = "/rooms"
	)
	resp, err := c.request(method, path)
	if err != nil {
		println("Failed on api call")
		log.Fatal(err)
	}

	var rooms []Room
	err = json.NewDecoder(resp.Body).Decode(&rooms)
	if err != nil {
		println("Failed on decode the result")
		log.Fatal(err)
	}
	return rooms
}
