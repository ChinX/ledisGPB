package db

import (
	"log"
	"gopkg.in/redis.v4"
	"github.com/chinx/ledisGPB/protos"
	"encoding/json"
	"github.com/golang/protobuf/proto"
)

func RunRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	pong, err := rdb.Ping().Result()
	log.Println(pong, err)
	//err = rdb.Set("hello", "world", 0).Err()
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//
	//val, err := rdb.Get("hello").Result()
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//log.Println(val)

	msg := &protos.Msg{
		MsgType: 1,
		MsgInfo: "I am chan",
		MsgFrom: "127.0.0.1",
	}

	inBytes, err := proto.Marshal(msg)
	if err != nil {
		log.Println(err)
		return
	}

	err = rdb.Set("buffer", inBytes, 0).Err()
	if err != nil {
		log.Println(err)
		return
	}

	outBytes, err := rdb.Get("buffer").Bytes()
	if err != nil {
		log.Println(err)
		return
	}

	msgInfo := &protos.Msg{}
	err = proto.Unmarshal(outBytes, msgInfo)
	if err != nil {
		log.Println("Unmarshaling error :", err)
		return
	}

	endBytes, _ := json.Marshal(msgInfo)
	log.Println(string(endBytes))
}
