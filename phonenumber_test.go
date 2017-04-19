package phonenumber

import (
	"testing"
)

func TestJP(t *testing.T) {
	number := Parse("090 6135 3368", "jp")
	expected := "819061353368"
	if number != expected {
		t.Errorf("JP: got %v\nwant %v", number, expected)
	}
}

func TestCN(t *testing.T) {
	number := Parse("+8615948692360", "cn")
	expected := "8615948692360"
	if number != expected {
		t.Errorf("CN: got %v\nwant %v", number, expected)
	}
}

func TestUSA(t *testing.T) {
	number := Parse("(817) 569-8900", "usa")
	expected := "18175698900"
	if number != expected {
		t.Errorf("CN: got %v\nwant %v", number, expected)
	}
}
