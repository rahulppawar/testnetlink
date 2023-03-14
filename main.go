package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/vishvananda/netlink"
)

func main() {
	Listner("eno1")
}

/* List routes belonging to interface name(s) @ifs */
func Listner(ifs ...string) {
	for _, iface := range ifs {
		_, err := netlink.LinkByName(iface)
		if err != nil {
			log.Printf("failed to look up interface %s: %s", iface, err)
			continue
		}
	}

	fmt.Println("netlink listner (ignore link local addresses)...")

	ch := make(chan netlink.NeighUpdate)
	done := make(chan struct{})
	defer close(done)
	if err := netlink.NeighSubscribe(ch, done); err != nil {
		log.Fatal(err)
	}

	for data := range ch {
		ip := data.Neigh.IP.String()

		// ignore empty IPs
		if ip == "::" {
			continue
		}

		// ignore link local address
		if strings.HasPrefix(ip, "fe80") {
			continue
		}

		fmt.Printf("%+v\n", data)
	}
}
