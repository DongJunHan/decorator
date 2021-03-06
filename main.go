package main

import(
	"fmt"
	"WEB-INF/decorator/lzw"
	"WEB-INF/decorator/cipher"
)

type Component interface {
	Operator(string)
}

var sentData string
type SendComponent struct {}

func (self *SendComponent) Operator(data string){
	//Send data
	sentData = data
}

type ZipComponent struct{
	com Component
}

func (self *ZipComponent) Operator(data string){
	zipData, err := lzw.Write([]byte(data))
	if err != nil {
		panic(err)
	}

	self.com.Operator(string(zipData))

}

type EncryptComponent struct{
	key string
	com Component
}

func (self *EncryptComponent) Operator(data string){
	encryptData, err := cipher.Encrypt([]byte(data),self.key)
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(encryptData))

}

func main(){
	sender := &EncryptComponent{key : "abcde",
		com : &ZipComponent{ 
			com : &SendComponent{}}}

	sender.Operator("Hello world")
	fmt.Println(sentData)
}


