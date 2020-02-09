package node

import (
	"math/rand"
	"time"
)

var pool = []rune("abcdefghijklmnopqrstuvwxyz0123456789")
const tl = 2

func randomNodeId() {
	// generate random node id

}

func randomInfo() {
	// generate random info hash

}

func randomTid() string{
	// generate random transaction id
	tid := make([]rune, 2)
	rand.Seed(time.Now().Unix())
	for i := 0; i < tl; i++ {
		tid[i] = pool[rand.Intn(len(pool))]
	}
	return string(tid)
}