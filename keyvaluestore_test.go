package main

import "testing"

func TestStore(t *testing.T) {
	kv := NewKeyValue()
	kv.Store("key", "value")
	value := kv.Get("key")

	if value != "value" {
		t.Errorf("Expected value 'value', got %s", value)
	}
}

func TestDelete(t *testing.T) {
	kv := NewKeyValue()
	kv.Store("key", "value")
	kv.Remove("key")
	value := kv.Get("key")

	if value != "" {
		t.Errorf("Expected value '', got %s", value)
	}
}
