package db

import (
	"context"
	"github.com/1359332949/douyin/kitex_gen/relation"
	"github.com/1359332949/douyin/kitex_gen/user"
	// "log"
)

// db中的user relation是和数据库交互的数据结构
// service中需要回传的是kitex_gen/relation 中定义的user和frienduser

func BuildUsers(ctx context.Context, uid int64 ,users []*user.User)  ([]*user.User,error) {
	relationUsers := make([]*user.User,0)
	for _,u := range users{
		isfollow := false
		relations,_ :=GetRelation(ctx,uid,int64(u.Id))
		if relations != nil{
			isfollow=true
		}
		relationUsers = append(relationUsers,&user.User{
			Id : int64(u.Id),
			Name: u.Name,
			FollowCount: int64(u.FollowCount),
			FollowerCount: int64(u.FollowerCount),
			IsFollow: isfollow, 
		})
	}
	// log.Println(relationUsers)
	return relationUsers,nil
}

func BuildFriendUsers(ctx context.Context, uid int64 ,users []*user.User)  ([]*relation.FriendUser,error) {
	relationUsers := make([]*relation.FriendUser,0)
	for _,u := range users{
		relationUsers = append(relationUsers,&relation.FriendUser{
			Id : int64(u.Id),
			Name: u.Name,
			FollowCount: int64(u.FollowCount),
			FollowerCount: int64(u.FollowerCount),
			IsFollow: true,
			Avatar:"https://p.qqan.com/up/2020-6/2020060316583052133.jpg",
			Message:"测试下好用不",
			MsgType:1,
		})
	}
	return relationUsers,nil
}