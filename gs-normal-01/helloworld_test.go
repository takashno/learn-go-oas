package main

import (
	"os"
	"testing"
)

// mainのテスト
func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

// messageのテスト
func TestMessage(t *testing.T) {
	if "Hello World !!" != message() {
		t.Fatal("fail test")
	}
}
