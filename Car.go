package main

import "errors"

type Car struct{
	name string
	price int
}

func getCar(name string, price int) (*Car, error){
	if name == ""{
		return nil, errors.New("missing name")
	}
	return &Car{
		name : name,
		price:price}, nil
}
