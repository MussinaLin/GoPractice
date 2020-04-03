package main

import (
	"log"
	"reflect"
	"testing"
)
func TestGetCar(t *testing.T){
	c, err := getCar("", 100)
	if c == nil{
		t.Log("c is nil")
	}
	if err != nil{
		t.Log("got error:", err)
	}
}

//func TestGetCarWithAssert(t *testing.T){
//	c, err := getCar("", 100)
//	assert.NotNil(t, err)
//}

func Test_getCar(t *testing.T) {
	type args struct {
		name  string
		price int
	}
	tests := []struct {
		name    string
		args    args
		want    *Car
		wantErr bool
	}{
		{
			name:"test data 1",
			args:args{name:"BMW", price:100},
			want:&Car{name:"BMW", price:100},
			wantErr:false,
		},
		{
			name:"test data 2",
			args:args{name:"BME", price:200},
			want:&Car{name:"BME", price:200},
			wantErr:false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			log.Println(tt.args)
			got, err := getCar(tt.args.name, tt.args.price)
			if (err != nil) != tt.wantErr {
				t.Errorf("getCar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCar() got = %v, want %v", got, tt.want)
			}
		})
	}
}