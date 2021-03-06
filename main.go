package main

type Component interface {
	Operator(string)
}

var sentData string
type SendComponent struct {}

func (self *SendComponent) Operator(data string){
	//Send data
	SendData = data
}

type ZipComponent struct{
	com Component
}

func (self *ZipComponent) Operator(data string){
	
}
