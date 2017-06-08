package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/kkimu/generate-binary-file-from-protobuf/myproto"
)

func main() {
	data := generateRequest()
	writeBinaryFile("request.bin", binary.BigEndian, data)
}

func generateRequest() []byte {
	req := &myproto.Request{
		Address:  "test@test.com",
		Password: "password",
		Name:     "name",
		Age:      30,
	}

	data, err := proto.Marshal(req)
	if err != nil {
		fmt.Println("marshaling error: ", err)
		return nil
	}
	return data
}

func writeBinaryFile(filename string, order binary.ByteOrder, data interface{}) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, order, data)

	if err != nil {
		fmt.Println("err:", err)
		return
	}

	file, err2 := os.Create(filename)
	if err2 != nil {
		fmt.Println("file create err:", err2)
		return
	}

	_, err3 := file.Write(buf.Bytes())
	if err3 != nil {
		fmt.Println("file write err:", err3)
		return
	}

	fmt.Println("file write ok.")
}
