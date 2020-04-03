package main

import "testing"

func TestGetCar(t *testing.T){
	c, err := getCar("", 100)
	if c == nil{
		t.Log("c is nil")
	}
	if err != nil{
		t.Log("got error:", err)
	}
}
