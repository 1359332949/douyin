package main

import (
	video "github.com/1359332949/douyin/kitex_gen/video/video/video/userservice"
	"log"
)

func main() {
	svr := video.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
