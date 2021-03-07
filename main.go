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
var receiveData string

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

type DecryptComponent struct{
	key string
	com Component
}

func (self *DecryptComponent) Operator(data string){
	decryptData,err := cipher.Decrypt([]byte(data),self.key)
	if err != nil {
		panic(err)
	}

	self.com.Operator(string(decryptData))
}

type UnzipComponent struct{
	com Component
}

func (self *UnzipComponent) Operator(data string){
	unzipData,err := lzw.Read([]byte(data))
	if err != nil{
		panic(err)
	}

	self.com.Operator(string(unzipData))
}

type ReadComponent struct{}
func (self *ReadComponent) Operator(data string){
	receiveData = data
}

func main(){
	//프로토콜 순서상 압축을 하고 암호화를 함.
	//따라서 받는쪽에서는 복호화를 하고 압축을 풀어야함.
	sender := &EncryptComponent{key : "abcde",
		com : &ZipComponent{ 
			com : &SendComponent{},
		},
	}

	sender.Operator("Hello world")
	fmt.Println(sentData)

	receiver := &UnzipComponent{
		com : &DecryptComponent{
			key : "abcde",
			com : &ReadComponent{},
		},
	}

	receiver.Operator(sentData)
	fmt.Println(receiveData)
}


