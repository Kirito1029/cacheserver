package hashtable

import (
	"errors"
	"time"
)

func NewDLL() (dll *DLL) {
	dll = &DLL{head: nil}
	return
}

func (dll *DLL) Add(key string, value []byte, expiry uint64) {
	if dll.head != nil {
		dll.head.AddIfNot(key, value, expiry)
	} else {
		dll.head = &Node{
			key:    key,
			value:  value,
			expiry: time.Now().Add(time.Second * time.Duration(expiry)),
			next:   nil,
		}
	}
}

func (dll *DLL) Get(key string) (value []byte, err error) {
	if dll.head == nil {
		return value, errors.New("Empty")
	}
	return dll.head.Get(key)
}

func (dll *DLL) Expiry(key string, ttl uint64) (err error) {
	dll.head.SetExpiry(key, ttl)
	return
}

func (head *Node) SetExpiry(key string, ttl uint64) {
	th := head
	for th != nil {
		if th.key == key {
			th.expiry = time.Now().Add(time.Duration(ttl) * time.Second)
			return
		} else {
			th = th.next
		}
	}
}

func (head *Node) AddIfNot(key string, value []byte, expiry uint64) {
	th := head
	var exptime time.Time
	if expiry != 0 {
		exptime = time.Now().Add(time.Second * time.Duration(expiry))
	}
	for th != nil {
		if th.key == key {
			th.value = value
			th.expiry = exptime
			return
		} else {
			if th.next == nil {
				break
			}
			th = th.next
		}
	}

	th.next = &Node{
		key:    key,
		value:  value,
		expiry: exptime,
		next:   nil,
	}
}

func (head *Node) Get(key string) (value []byte, err error) {
	th := head
	for th != nil {
		if th.key == key {
			if !th.expiry.IsZero() && th.expiry.Before(time.Now()) {
				return value, errors.New("Empty")
			}
			value = th.value
			return
		} else {
			th = th.next
		}

	}
	return value, errors.New("Empty")

}
