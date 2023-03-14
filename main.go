package main

import (
	"fmt"
	"log"

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

	fmt.Println("netlink listner...")

	ch := make(chan netlink.AddrUpdate)
	done := make(chan struct{})
	defer close(done)
	if err := netlink.AddrSubscribe(ch, done); err != nil {
		log.Fatal(err)
	}

	for data := range ch {
		fmt.Printf("%+v\n", data)
	}
}
