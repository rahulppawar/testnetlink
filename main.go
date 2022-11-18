package main

import (
	"fmt"
	"log"

	"github.com/mdlayher/netlink"
)

func main() {
	const (
		// Speak to route netlink using netlink
		familyRoute = 0

		// Listen for events triggered by addition or deletion of
		// network interfaces
		rtmGroupLink = 0x1
	)

	c, err := netlink.Dial(familyRoute, &netlink.Config{
		// Groups is a bitmask; more than one group can be specified
		// by OR'ing multiple group values together
		Groups: rtmGroupLink,
	})
	if err != nil {
		log.Fatalf("failed to dial netlink: %v", err)
	}
	defer c.Close()

	fmt.Println("entering receive loop")
	for {
		// Listen for netlink messages triggered by multicast groups
		msgs, err := c.Receive()
		if err != nil {
			log.Fatalf("failed to receive messages: %v", err)
		}

		for _, m := range msgs {
			fmt.Println(string(m.Data))
		}
	}
}
