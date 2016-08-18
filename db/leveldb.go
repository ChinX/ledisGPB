package db

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
	ProtoKind LevelKind = "proto"
	JsonKind LevelKind = "json"
)

type LevelKind string

func RunLevel(kind LevelKind) {
	var inBytes []byte

	ro := &opt.ReadOptions{}
	wo := &opt.WriteOptions{}

	ldb, err := leveldb.OpenFile(fmt.Sprintf("E:/chan/leveldb/%s", kind), nil)
	if err != nil {
		panic(err)
	}
	defer ldb.Close()

	msg := &protos.Msg{
		MsgType: 1,
		MsgInfo: "I am chan",
		MsgFrom: "127.0.0.1",
	}

	if kind == JsonKind {
		inBytes, err = json.Marshal(msg)
	} else {
		inBytes, err = proto.Marshal(msg)
	}

	if err != nil {
		log.Println("Marchaling error :", err)
		return
	}

	err = ldb.Put([]byte("msg"), inBytes, wo)
	if err != nil {
		log.Println("Proto putting error :", err)
		return
	}

	outBytes, err := ldb.Get([]byte("msg"), ro)
	if err != nil {
		log.Println("Proto getting error :", err)
		return
	}

	msgInfo := &protos.Msg{}
	if kind == JsonKind {
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
