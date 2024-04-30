package utils

import (
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

// Helper function to create packet
func CreatePacket(sourceIP, destIP net.IP, sourcePort, destPort uint16) gopacket.Packet {
	ethernetLayer := layers.Ethernet{
		SrcMAC:       net.HardwareAddr{0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		DstMAC:       net.HardwareAddr{0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		EthernetType: layers.EthernetTypeIPv4,
	}

	ipLayer := layers.IPv4{
		SrcIP:    sourceIP,
		DstIP:    destIP,
		Protocol: layers.IPProtocolTCP,
	}

	tcpLayer := layers.TCP{
		SrcPort: layers.TCPPort(sourcePort),
		DstPort: layers.TCPPort(destPort),
	}

	tcpLayer.SetNetworkLayerForChecksum(&ipLayer)

	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: true,
	}

	if err := gopacket.SerializeLayers(buf, opts, &ethernetLayer, &ipLayer, &tcpLayer); err != nil {
		panic(err)
	}

	packet := gopacket.NewPacket(buf.Bytes(), layers.LayerTypeEthernet, gopacket.Default)

	return packet
}
