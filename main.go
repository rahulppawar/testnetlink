package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/vishvananda/netlink"
	"github.com/vishvananda/netlink/nl"
)

func main() {
	Listner()
}

func Listner(ifs ...string) {

	fmt.Println("netlink listner...")

	ch := make(chan netlink.NeighUpdate)
	done := make(chan struct{})
	defer close(done)
	if err := netlink.NeighSubscribe(ch, done); err != nil {
		log.Fatal(err)
	}

	for data := range ch {
		ip := data.Neigh.IP.String()
		// ignore empty IP || IPv4 || link local address
		if ip == "::" || (nl.GetIPFamily(data.Neigh.IP) == netlink.FAMILY_V4) || strings.HasPrefix(ip, "fe80") {
			continue
		}
		fmt.Printf("%s,%+v\n", time.Now().Format(time.RFC3339), data)
	}
}
