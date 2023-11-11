package main

import "fmt"

// Description: Adapter is a structural design pattern that allows objects with
// incompatible interfaces to collaborate.

// The Adapter pattern lets you create a middle-layer class that serves as a
// translator between your code and a legacy class, a 3rd-party class or any
// other class with a weird interface.

type Client struct {
}

func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
  fmt.Println("Client inserts Lightning connector into computer.")
  
}
