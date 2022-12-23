package main

import (
	"fmt"
)

type KeyValueStore struct {
	store map[string]string
}

func (kv *KeyValueStore) Store(key, value string) {
	kv.store[key] = value
}

func (kv *KeyValueStore) Get(key string) string {
	return kv.store[key]
}

func (kv *KeyValueStore) Remove(key string) {
	delete(kv.store, key)
}

func NewKeyValue() *KeyValueStore {
	return &KeyValueStore{
		store: make(map[string]string),
	}
}

func main() {
	kv := NewKeyValue()
	kv.Store("aobakwe", "programmer")
	fmt.Println(kv.Get("aobakwe"))
}
