package main

import (
	"bytes"
	"context"
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

var Mode map[string]string = map[string]string{
	"udp": "udp4",
}
const allIP string = "0.0.0.0"

func randomBytes() []byte{
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
	//var connection *icmp.PacketConn
	connection, err := icmp.ListenPacket(Mode["udp"], allIP)
	if err != nil {
		log.Println("Error Occured:", err.Error())
		panic(err.Error())
	}
	// error handle function needs
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
		log.Println("Error Occured:", err.Error())
		panic(err.Error())
	}
	// udp
	dest := &net.UDPAddr{IP: target.IP, Zone: target.Zone}

	msg := PingPacket()
	msgByte, err := msg.Marshal(nil)
	if err != nil {
		log.Println("Error Occured:", err.Error())
		panic(err.Error())
	}
	n, err := conn.WriteTo(msgByte, dest)
	if err != nil {
		log.Println("Error Occured During Seding Data:", err.Error())
		panic(err.Error())
	}

	resultChan := make(chan bool, 1)
	go func(){
		readBuffer := make([]byte, n)
		_, _, e := conn.ReadFrom(readBuffer)
		if e != nil {
			log.Println("Error Occured During Read Buffer:", e.Error())
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
	// too many ip addrs may cause hmm..
	resultChan := make(chan pingResult, len(addrs))
	wait := sync.WaitGroup{}
	wait.Add(len(addrs))
	for _, addr := range addrs {
		log.Println(addr)
		go func(addr string) {
			resultChan <- Ping(addr)
			wait.Done()
		}(addr)
	}
	result := make([]pingResult, len(addrs))
	wait.Wait()
	length := len(resultChan)
	for i := 0; i < length; i++ {
		temp := <- resultChan
		result[i] = temp
	}
	log.Println(result)
	return result
}

func TestPing(addr string) string{
	conn := UDPConnection()
	defer conn.Close()
	// ipv4
	conn.IPv4PacketConn().SetControlMessage(ipv4.FlagTTL, true)

	target, err := net.ResolveIPAddr("ip", addr)
	if err != nil {
		log.Println("Error Occured:", err.Error())
		panic(err.Error())
	}
	// udp
	dest := &net.UDPAddr{IP: target.IP, Zone: target.Zone}

	msg := PingPacket()
	msgByte, err := msg.Marshal(nil)
	if err != nil {
		log.Println("Error Occured:", err.Error())
		panic(err.Error())
	}
	n, err := conn.WriteTo(msgByte, dest)
	if err != nil {
		log.Println("Error Occured During Seding Data:", err.Error())
		panic(err.Error())
	}

	readBuffer := make([]byte, n)
	_, _, e := conn.ReadFrom(readBuffer)
	if e != nil {
		log.Println("Error Occured During Read Buffer:", e.Error())
	}


	return addr
}

func ContextPing(addr string, timeout int) pingResult {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	success := make(chan bool, 1)
	go func(){
		r := TestPing(addr)
		log.Println("doom")
		log.Println(r)
		success <- true

	}()
	var ok bool = false
	select{
		case <- ctx.Done():
			log.Println("failed", ctx.Err())
			ok = false
		case <- success:
			log.Println("success")
			ok = true
	}
	return pingResult{addr, ok}
}

func ContextMultiPing(addrs []string) []pingResult{
	mainCtx, cancel := context.WithTimeout(context.Background(), time.Second * 2)
	log.Println(mainCtx, cancel)

	//defer cancel()

	return nil
}

func main(){
	//log.Println(Ping("12.0.0.1"))
	//MultiPing([]string{"127.0.0.1", "100.50.204.153", "8.8.8.8"})
	//ContextMultiPing([]string{"127.0.0.1", "100.50.204.153"})
	ContextPing("123.124.123.1", 2)
}
