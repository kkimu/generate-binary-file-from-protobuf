package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"reflect"

	proto1 "github.com/golang/protobuf/proto"
	"github.com/techcampman/team-5_3/proto"
)

func main() {
	// uint8型(1バイト)の値を用意する

	data := GenerateRequest()
	WriteBinaryFile("SignUpRequest", binary.BigEndian, data)

}

func GenerateRequest() []byte {
	/*
		req := &proto.CreateTeamAndSignUpRequest{
			Address:  "aaaa@b.com1",
			Password: "password1",
			UserName: "kk1",
			TeamName: "namae1",
		}
	*/
	///*
	req := &proto.SignUpRequest{
		Address:  "bbbbb@b.com2",
		Password: "password2",
		UserName: "kk2",
		TeamName: "namae1",
	}
	//*/
	data, err := proto1.Marshal(req)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	return data
}

/*
バイナリデータに変換してファイルに出力する
*/
func WriteBinaryFile(filename string, order binary.ByteOrder, val interface{}) {
	// バイナリデータの格納用
	buf := new(bytes.Buffer)
	fmt.Println("buf", buf, " typeofbuf", reflect.TypeOf(buf))
	fmt.Println("buf.Bytes()", buf.Bytes(), " typeofbuf.Bytes()", reflect.TypeOf(buf.Bytes()))

	// valの値をバイナリデータに変換してbufに格納する
	err := binary.Write(buf, order, val)

	if err != nil {
		fmt.Println("err:", err)
		return
	}

	// ファイル作成
	file, err2 := os.Create(filename)
	if err2 != nil {
		fmt.Println("file create err:", err2)
		return
	}

	// バイナリデータをファイルに書き込み
	_, err3 := file.Write(buf.Bytes())
	if err3 != nil {
		fmt.Println("file write err:", err3)
		return
	}

	fmt.Println("file write ok.")
}
