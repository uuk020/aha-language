package main

import "testing"

func TestEven(t *testing.T) {
	if Even(10) {
		t.Log(" 10 是偶数")
		t.Fail()
	}
	if Even(7) {
		t.Log(" 7 不是偶数")
		t.Fail()
	}
}

func TestOdd(t *testing.T) {
	if !Odd(11) {
		t.Log(" 11 是奇数")
		t.Fail()
	}
	if Odd(10) {
		t.Log("10 不是奇数")
		t.Fail()
	}
}
