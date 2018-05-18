package kvstore

import (
	"errors"
	"fmt"
)

var cache = make(map[string]interface{})

// Write writes to the global diskv db
func Put(k string, v interface{}) {
	cache[k] = v
}

// Read reads from the global diskv db
func Get(k string) (interface{}, error) {
	v, ok := cache[k]
	if (ok != true) {
		return nil, errors.New(fmt.Sprintf("Key %s not exists", k))
	}
	return v, nil

}

// Erase deletes a key from the global discv db
func Delete(k string) error {
	_, ok := cache[k]
	if (ok == false) {
		return errors.New(fmt.Sprintf("Key %s not exists", k))
	}
	cache[k] = nil
	return nil
}