package main

import (
	"errors"
	"time"
)

//Car is a test struct for writing test case.
type Car struct{
	name string
	price int
}

func getCar(name string, price int) (*Car, error){
	if name == ""{
		return nil, errors.New("missing name")
	}
	time.Sleep(1 * time.Second)
	return &Car{
		name : name,
		price:price}, nil
}
