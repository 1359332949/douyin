package main

import (
	comment "github.com/1359332949/douyin/kitex_gen/comment/comment/comment/commentservice"
	"log"
)

func main() {
	svr := comment.NewServer(new(CommentServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
