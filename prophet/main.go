package main

import (
	"context"
	"time"

	"github.com/warden-protocol/wardenprotocol/go-client"
)

func main() {
	inbox := Inbox{
		c: make(chan Request, 100),
	}
	outbox := Outbox{
		c: make(chan Response, 100),
	}

	client, err := client.NewQueryClient("localhost:9090", true)
	if err != nil {
		panic(err)
	}

	go fetchRequests(client, &inbox)
	go processRequests(&inbox, &outbox)

	for {
		time.Sleep(time.Second)
	}
}

func fetchRequests(c *client.QueryClient, in *Inbox) {
	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		reqs, err := c.PendingInferenceRequests(ctx, &client.PageRequest{
			Limit: 100,
		})
	}
}

func processRequests(*Inbox, *Outbox) {

}

type Request struct {
	Id    uint64
	Input []byte
}

type Inbox struct {
	c chan Request
}

func (i *Inbox) Push(r Request) {
	i.c <- r
}

func (i *Inbox) Take() Request {
	return <-i.c
}

type Response struct {
	Id      uint64
	Output  []byte
	Receipt []byte
}

type Outbox struct {
	c chan Response
}
