package hashtable

import (
	"cache-server/pkg/hash"
	"errors"
	"fmt"
	"time"
)

const size = 10

type (
	MHashTable interface {
		Set(key string, value []byte)
		Get(key string) []byte
		Expire(key string) error
	}
	HashTable struct {
		table map[int]*DLL
		MHashTable
	}
	DLL struct {
		head *Node
	}
	Node struct {
		key    string
		value  []byte
		expiry time.Time
		next   *Node
	}
)

func init() {

}

func New() (htable *HashTable) {
	htable = &HashTable{}
	htable.table = make(map[int]*DLL, size)
	return
}

func (htable *HashTable) Set(key string, value []byte, expiry uint64) (err error) {
	hsh := int(hash.Hash(key))
	fmt.Println(key, hsh%size)
	dlist := htable.table[hsh%size]
	if dlist == nil {
		dlist = &DLL{
			head: nil,
		}
		htable.table[hsh%size] = dlist
	}
	dlist.Add(key, value, expiry)
	return
}

func (htable *HashTable) Get(key string) (value []byte, err error) {
	hsh := int(hash.Hash(key))
	fmt.Println(key, hsh%size)
	dlist := htable.table[hsh%size]
	if dlist == nil {
		return value, errors.New("Empty")
	}
	return dlist.Get(key)
}

func (htable *HashTable) Expire(key string, ttl uint64) (err error) {
	fmt.Println()
	hsh := int(hash.Hash(key))
	fmt.Println(key, hsh%size)
	dlist := htable.table[hsh%size]
	if dlist == nil {
		return errors.New("Empty")
	}
	return dlist.Expiry(key, ttl)
}
