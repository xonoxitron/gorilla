package main

import (
	"github.com/xonoxitron/gorilla/scheduler"
	"github.com/xonoxitron/gorilla/storage"
)

func main() {
	storage.Setup()
	scheduler.StartJob()
}
