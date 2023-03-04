package main

import (
	relation "github.com/1359332949/douyin/kitex_gen/relation/relation/relation/relationservice"
	"log"
)

func main() {
	svr := relation.NewServer(new(RelationServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
