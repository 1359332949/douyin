package main

import (
	message "github.com/1359332949/douyin/kitex_gen/message/message/message/messageservice"
	"log"
)

func main() {
	svr := message.NewServer(new(MessageServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
