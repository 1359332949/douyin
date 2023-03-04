package main

import (
	interact "github.com/1359332949/douyin/kitex_gen/interact/interact/interact/interactservice"
	"log"
)

func main() {
	svr := interact.NewServer(new(InteractServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
