package main

import "github.com/google/gopacket"
import "fmt"

func ParsePacket(packet gopacket.Packet) {
	fmt.Println(packet)
}
