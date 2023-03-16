package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/vishvananda/netlink"
	"github.com/vishvananda/netlink/nl"
	"golang.org/x/sys/unix"
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

		// Ignore RTM_NEWNEIGH entries with States PROBE, STALE, INCOMPLETE, FAILED stc.
		if (data.Type == unix.RTM_NEWNEIGH) && (data.Neigh.State != netlink.NUD_REACHABLE) {
			continue
		}

		// Here we get entries of RTM_DELNEIGH and RTM_NEWNEIGH + REACHABLE state
		fmt.Printf("%s,%+v\n", time.Now().Format(time.RFC3339), data)
	}
}
