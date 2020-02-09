package bencode

import "testing"

func TestDecodeData(t *testing.T) {
	x, r := DecodeData("d1:ad2:id20:abcdefghij01234567896:target20:mnopqrstuvwxyz123456e1:q9:find_node1:t2:aa1:y1:qe")
	t.Log("r:", r)
	t.Log("x:", x)
}

func TestDecodeStream(t *testing.T) {
	x := DecodeStream([]byte("d1:ad2:id20:abcdefghij01234567896:target20:mnopqrstuvwxyz123456e1:q9:find_node1:t2:aa1:y1:qe"))
	t.Log("x: ", x)
}