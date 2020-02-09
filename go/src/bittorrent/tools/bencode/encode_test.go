package bencode

import (
	"testing"
)

func TestEncode(t *testing.T) {

}

func TestEncodeDict(t *testing.T) {
	var x Dictionary = Dictionary{map[string]interface{}{"t":"aa", "y":14, "q":"find_node", "a": map[string]interface{}{"id":"abcdefghij0123456789", "target":"mnopqrstuvwxyz123456"}}}
	t.Log(EncodeDict(x))
	// "d1:ad2:id20:abcdefghij01234567896:target20:mnopqrstuvwxyz123456e1:q9:find_node1:t2:aa1:y1:qe"
}