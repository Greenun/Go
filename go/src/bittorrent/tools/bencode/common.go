package bencode

import (
	"sort"
)

type Dictionary struct {
	data map[string]interface{}
}

type sortedKey []string

func (s sortedKey) Len() int { return len(s) }

func (s sortedKey) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortedKey) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (d Dictionary) Keys() []string{
	keySet := make(sortedKey, len(d.data))
	idx := 0
	for k, _ := range d.data {
		keySet[idx] = k
		idx++
	}
	sort.Sort(keySet)
	return keySet
}

func (d Dictionary) Values() []interface{} {
	valueSet := make([]interface{}, len(d.data))
	idx := 0
	for _, v := range d.data {
		valueSet[idx] = v
		idx++
	}
	return valueSet
}

func (d Dictionary) getValue(key string) interface{}{
	return d.data[key]
}

func ConvertToBytes(plain string) []byte{
	return []byte(plain)
}

func ConvertToString(bs []byte) string{
	return string(bs)
}