package main

import (
	"log"

	"github.com/chinx/ledisGPB/db"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	db.RunLevel(db.ProtoKind)
	db.RunRedis()
}
