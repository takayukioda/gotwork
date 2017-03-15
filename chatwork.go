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
	RoomId         int    `json:room_id`
	Name           string `json:name`
	Type           string `json:type`
	Role           string `json:role`
	Sticky         bool   `json:sticky`
	UnreadNum      int    `json:unread_num`
	MentionNum     int    `json:mention_num`
	MytaskNum      int    `json:mytask_num`
	MessageNum     int    `json:message_num`
	FileNum        int    `json:file_num`
	TaskNum        int    `json:task_num`
	IconPath       string `json:icon_path`
	LastUpdateTime string `json:last_update_time`
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

func (c *Chatwork) GetRooms() []Room {
	const (
		method = "GET"
		path   = "/rooms"
	)
	req, e := http.NewRequest(method, c.endpoint(path), nil)
	if e != nil {
		println("Failed on creating new request")
		log.Fatal(e)
	}
	req.Header.Add("User-Agent", "chatwork-client by Go")
	req.Header.Add("X-ChatWorkToken", c.token)

	resp, er := c.client.Do(req)
	if er != nil {
		println("Failed on api call")
		log.Fatal(er)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	var rooms []Room
	err := json.Unmarshal(buf.Bytes(), &rooms)
	if err != nil {
		println("Failed on unmarshal the result")
		log.Fatal(er)
	}
	return rooms
}
