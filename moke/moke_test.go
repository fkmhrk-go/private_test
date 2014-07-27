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

// テストのパッケージと同じパッケージなので、
// 小文字始まりの関数も呼べる
func TestInternalHello(t *testing.T) {
	v := internalHello("Mike")
	if v != "My name is Mike" {
		t.Errorf("Assertin failed %s", v)
		return
	}
}
