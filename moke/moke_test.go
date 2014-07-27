package moke

import (
	"testing"
)

func TestHello(t *testing.T) {
	v := Hello()
	if v != "My name is Moke" {
		t.Errorf("Assertin failed %s", v)
		return
	}
}

func TestInternalHello(t *testing.T) {
	v := ExFunc("Mike")
	if v != "My name is Mike" {
		t.Errorf("Assertin failed %s", v)
		return
	}
}
