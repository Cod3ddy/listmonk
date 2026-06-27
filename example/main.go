package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/cod3ddy/listmonk"
)

func main() {
	c := listmonk.New(
		"http://127.0.0.1:9002", // mine was running at this port at the time...and please omit the /api suffix...
		"<api-user-name>",
		"<api-user-token>",
	)

	err := c.Health()
	if err != nil {
		log.Fatal("shit is not health at all: ", err.Error())
	}

	subscribers, err := c.GetSubscribers(listmonk.GetSubscribersParams{})
	if err != nil {
		log.Fatal(err)
	}

	fmtSubs, _ := json.MarshalIndent(subscribers, "", " ")

	fmt.Printf("%+v\n", string(fmtSubs))
}
