package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/chinx/ledisGPB/protos"
	"github.com/golang/protobuf/proto"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

const (
	protoKind = "proto"
	jsonKind  = "json"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	kind := protoKind

	var inBytes []byte

	ro := &opt.ReadOptions{}
	wo := &opt.WriteOptions{}

	db, err := leveldb.OpenFile(fmt.Sprintf("E:/chan/leveldb/%s", kind), nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	msg := &protos.Msg{
		MsgType: 1,
		MsgInfo: "I am chan",
		MsgFrom: "127.0.0.1",
	}

	if kind == jsonKind {
		inBytes, err = json.Marshal(msg)
	} else {
		inBytes, err = proto.Marshal(msg)
	}

	if err != nil {
		log.Println("Marchaling error :", err)
		return
	}

	err = db.Put([]byte("msg"), inBytes, wo)
	if err != nil {
		log.Println("Proto putting error :", err)
		return
	}

	outBytes, err := db.Get([]byte("msg"), ro)
	if err != nil {
		log.Println("Proto getting error :", err)
		return
	}

	msgInfo := &protos.Msg{}
	if kind == jsonKind {
		err = json.Unmarshal(outBytes, msgInfo)
	} else {
		err = proto.Unmarshal(outBytes, msgInfo)
	}

	if err != nil {
		log.Println("Unmarshaling error :", err)
		return
	}

	endBytes, _ := json.Marshal(msgInfo)
	log.Println(string(endBytes))

}
