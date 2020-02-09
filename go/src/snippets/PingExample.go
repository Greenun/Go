package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"log"
	"math/rand"
	"net"
	"reflect"
	"time"
)
const Blank int = 0

func EchoPacket() *icmp.Message{
	rand.Seed(time.Now().UnixNano() & 0xff)
	randint := rand.Intn(100)
	//ipv6.ICMPTypeEchoRequest
	testData := append(testBytes(time.Now()), bytes.Repeat([]byte{0x01}, 8)...)
	fmt.Println(testData)

	return &icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: Blank,
		//Checksum: 0,
		Body: &icmp.Echo{
			ID: randint,
			Seq: 0,
			Data: testData,//[]byte{},
		},
	}
}

func main(){
	test := EchoPacket()
	fmt.Println(reflect.TypeOf(test))
	log.Println(test)
	j, e := json.Marshal(test.Body)
	fmt.Println("json: ", string(j), "error: ", e)
	//conn := &icmp.PacketConn{}
	conn, err := icmp.ListenPacket("udp4", "0.0.0.0")
	defer conn.Close()

	fmt.Println(conn, err)
	conn.IPv4PacketConn().SetControlMessage(ipv4.FlagTTL, true)
	fmt.Println(conn)
	//conn.IPv4PacketConn().SetControlMessage(ipv4.FlagTTL, true)
	//fmt.Println(ipv4.FlagTTL)

	json, err1 := test.Marshal(nil)
	log.Println(json, "||", err1)
	addr, err2 := net.ResolveIPAddr("ip", "172.217.31.142") // 8.8.8.8
	log.Println(addr, err2)
	log.Println(addr.IP, addr.Zone)
	dest := &net.UDPAddr{IP: addr.IP, Zone: addr.Zone}


	r, err := conn.WriteTo(json, dest)
	log.Println(r, "|", err)
	fmt.Println(r, err)

	readBuffer := make([]byte, 64)
	n, a, e := conn.ReadFrom(readBuffer)
	fmt.Println(n , "|||", e, "|||", a)
	if e == nil {
		fmt.Println(readBuffer[:n])
	}
}

func testBytes(t time.Time) []byte {
	nsec := t.Unix()
	temp := make([]byte, 8)
	for i := 0; i < 8; i++ {
		temp[i] = byte((nsec >> i) & 0xff)
	}
	return temp
}