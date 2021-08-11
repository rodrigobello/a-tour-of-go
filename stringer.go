package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

func (address IPAddr) String() string {
	stringAddress := make([]string, len(address))
	for i, ip := range address {
		stringAddress[i] = strconv.Itoa(int(ip))
	}
	return strings.Join(stringAddress, ".")
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
