package tools

import (
	"bytes"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"
)

type pingResult struct {
	ipAddr string
	success bool
}

var mode map[string]string = map[string]string{
	"udp": "udp4",
}

const allIP string = "0.0.0.0"

func randomBytes() []byte {
	sec := time.Now().Unix()
	rb := make([]byte, 8)
	for i := 0; i < 8; i++ {
		rb[i] = byte((sec >> (i + 1)) & 0xff)
	}
	return rb
}

func PingPacket() *icmp.Message {
	rand.Seed(time.Now().UnixNano() & 0xff)
	randint := rand.Intn(100)
	dummyData := append(randomBytes(), bytes.Repeat([]byte{0x11}, 8)...)

	return &icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &icmp.Echo{
			ID: randint,
			Seq: 0,
			Data: dummyData,
		},
	}
}

func UDPConnection() *icmp.PacketConn {
	connection, err := icmp.ListenPacket(mode["udp"], allIP)
	if err != nil {
		log.Println("Error Occured:", err.Error())
		panic(err.Error())
	}
	// recover
	//defer func(){
	//	if r := recover(); r != nil {
	//		log.Println(r)
	//	}
	//}()
	//
	return connection
}

func Ping(addr string) pingResult{
	conn := UDPConnection()
	defer conn.Close()
	// ipv4
	conn.IPv4PacketConn().SetControlMessage(ipv4.FlagTTL, true)

	target, err := net.ResolveIPAddr("ip", addr)
	if err != nil {
		log.Println("Error Occurred:", err.Error())
		panic(err.Error())
	}

	// udp
	dest := &net.UDPAddr{IP: target.IP, Zone: target.Zone}

	msg := PingPacket()
	msgByte, err := msg.Marshal(nil)
	if err != nil {
		log.Println("Error Occurred:", err.Error())
		panic(err.Error())
	}

	n, err := conn.WriteTo(msgByte, dest)
	if err != nil {
		log.Println("Error Occurred During Sending Data:", err.Error())
		panic(err.Error())
	}

	resultChan := make(chan bool, 1)
	go func(){
		readBuffer := make([]byte, n)
		_, _, e := conn.ReadFrom(readBuffer)
		if e != nil {
			log.Println("Error Occurred During Read Buffer:", e.Error())
		} else {
			resultChan <- true
		}
	}()
	var ret bool
	select {
	case <- resultChan:
		ret = true
	case <- time.After(time.Second):
		ret = false
	}
	return pingResult{addr, ret}
}

func MultiPing(addrs []string) []pingResult {
	pingResults := make(chan pingResult, len(addrs))
	wait := sync.WaitGroup{}
	wait.Add(len(addrs))
	for _, addr := range addrs {
		go func(ad string) {
			pingResults <- Ping(ad)
			wait.Done()
		}(addr)
	}
	result := make([]pingResult, len(addrs))
	wait.Wait()
	length := len(pingResults)
	for i := 0; i < length; i++ {
		temp := <- pingResults
		result[i] = temp
	}
	return result
}