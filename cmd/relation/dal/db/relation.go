
package db

import (
	"context"
	// "fmt"
	// "time"
	// "log"
	"github.com/1359332949/douyin/pkg/consts"
	// "github.com/1359332949/douyin/cmd/user/dal/db"
	// "github.com/1359332949/douyin/kitex_gen/relation"
	// "github.com/1359332949/douyin/kitex_gen/user"
	"github.com/1359332949/douyin/pkg/errno"
	"gorm.io/gorm"
	"github.com/cloudwego/kitex/pkg/klog"
)


type Relation struct {
	gorm.Model
	ID         int64     `gorm:"column:id;primary_key;AUTO_INCERMENT"`
	// FollowTime time.Time `gorm:"column:follow_time;default:CURRENT_TIMESTAMP;NOT NULL"`
	FromUserID int64     `gorm:"column:from_user_id;NOT NULL"`
	ToUserID   int64     `gorm:"column:to_user_id;NOT NULL"`
	// CreateTime int64  `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"`
	// UpdateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"`
}


func (r *Relation) TableName() string {
	return consts.RelationTableName
}


// GetRelation get relation info
func GetRelation(ctx context.Context, uid int64, tid int64) (*Relation, error) {
	relations := new(Relation)

	if err := DB.WithContext(ctx).First(&relations, "from_user_id = ? and to_user_id = ?", uid, tid).Error; err != nil {
		return nil, err
	}
	return relations, nil
}
// //根据id获取user
// // MGetUsers multiple get list of user info
// func MGetUsers(ctx context.Context, userIDs []int64) ([]*user.User, error) {
// 	res := make([]*user.User, 0)
// 	if len(userIDs) == 0 {
// 		return res, nil
// 	}
// 	// 从usr表中根据id查找到users的信息
// 	if err := DB.Table(consts.UserTableName).WithContext(ctx).Where("id in ?", userIDs).Find(&res).Error; err != nil {
// 		return nil, err
// 	}
// 	return res, nil
// }
// NewAction creates a new Relation
// uid关注tid，所以uid的关注人数加一，tid的粉丝数加一
func NewAction(ctx context.Context, uid int64, tid int64) error {
	relations,_ :=GetRelation(ctx,uid,tid)
		if relations != nil{
			return errno.RelationExistErr
		}

	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作
		// 1. 新增关注数据
		err := tx.Create(&Relation{FromUserID: uid, ToUserID: tid}).Error
		if err != nil {
			return err
		}

		// 2.改变 user 表中的 following count
		res := tx.Table(consts.UserTableName).Where("id = ?", uid).Update("follow_count", gorm.Expr("follow_count + ?", 1))
		if res.Error != nil {
			return res.Error
		}

		if res.RowsAffected != 1 {
			return errno.RelationActionErr
		}

		// 3.改变 user 表中的 follower count
		res = tx.Table(consts.UserTableName).Where("id = ?", tid).Update("follower_count", gorm.Expr("follower_count + ?", 1))
		if res.Error != nil {
			return res.Error
		}

		if res.RowsAffected != 1 {
			return errno.RelationActionErr
		}

		return nil
	})
	return err
}

// DelAction deletes a relation from the database.
func DelAction(ctx context.Context, uid int64, tid int64) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作
		relations := new(Relation)
		if err := tx.Where("from_user_id = ? AND to_user_id=?", uid, tid).First(&relations).Error; err != nil {
			return err
		}

		// 1. 删除关注数据
		err := tx.Unscoped().Delete(&relations).Error
		if err != nil {
			return err
		}
		// 2.改变 user 表中的 following count
		res := tx.Table(consts.UserTableName).Where("id = ?", uid).Update("follow_count", gorm.Expr("follow_count - ?", 1))
		if res.Error != nil {
			return res.Error
		}

		if res.RowsAffected != 1 {
			return errno.RelationActionErr
		}

		// 3.改变 user 表中的 follower count
		res = tx.Table(consts.UserTableName).Where("id = ?", tid).Update("follower_count", gorm.Expr("follower_count - ?", 1))
		if res.Error != nil {
			return res.Error
		}

		if res.RowsAffected != 1 {
			return errno.RelationActionErr
		}

		return nil
	})
	return err
}

// RelationFollowList returns the Following List.
func RelationFollowList(ctx context.Context, uid int64) ([]*Relation, error) {
	var RelationList []*Relation
	err := DB.WithContext(ctx).Where("from_user_id = ?", uid).Find(&RelationList).Error
	if err != nil {
		return nil, err
	}
	return RelationList, nil
	// userIDs :=make([]int64,0)
	// for _,u := range RelationList{
	// 	userIDs= append(userIDs,int64(u.ToUserID))
	// }
	// users, err := MGetUsers(ctx,userIDs)
	// if err != nil {
	// 	return nil, err
	// }
	// // log.Println(users)
	// return BuildUsers(ctx,uid,users)
}

// RelationFollowerList returns the Follower List.
func RelationFollowerList(ctx context.Context, tid int64) ([]*Relation, error) {
	var RelationList []*Relation
	err := DB.WithContext(ctx).Where("to_user_id = ?", tid).Find(&RelationList).Error
	if err != nil { 
		return nil, err
	}
	return RelationList, nil
	// userIDs :=make([]int64,0)
	// for _,u := range RelationList{
	// 	userIDs= append(userIDs,int64(u.FromUserID))
	// }
	// users, err := MGetUsers(ctx,userIDs)
	// if err != nil {
	// 	return nil, err
	// }
	// return BuildUsers(ctx,tid,users)
}

// 朋友：相互关注者->粉丝和关注者的交集
// RelationFriendList returns the Follower List.
func RelationFriendList(ctx context.Context, id int64) ([]*Relation, []*Relation, error) {
	var LRelationList []*Relation //关注者
	var RRelationList []*Relation //粉丝
	err := DB.WithContext(ctx).Where("from_user_id = ?", id).Find(&LRelationList).Error
	if err != nil {
		return nil, nil, err
	}
	err = DB.WithContext(ctx).Where("to_user_id = ?", id).Find(&RRelationList).Error
	if err != nil {
		return nil, nil, err
	}
	return LRelationList, RRelationList, nil
	// LuserIDs :=make([]int64,0)
	// for _,u := range LRelationList{
	// 	LuserIDs= append(LuserIDs,int64(u.ToUserID))
	// 	log.Println(LuserIDs)
	// }

	// RuserIDs :=make([]int64,0)
	// for _,u := range RRelationList{
	// 	RuserIDs= append(RuserIDs,int64(u.FromUserID))
	// 	log.Println(RuserIDs)
	// }
	// userIDs :=make([]int64,0)

	// m := make(map[int64]int)
	// for _,v :=range LuserIDs{
	// 	m[v]++
	// }
	// for _,v :=range RuserIDs{
	// 	if m[v]==1{
	// 		userIDs = append(userIDs,v)
	// 	}
	// }
	// log.Println(userIDs)
	// users, err := MGetUsers(ctx,userIDs)
	// if err != nil {
	// 	return nil, err
	// }
	// return BuildFriendUsers(ctx,id,users)
}


// 根据当前用户id和目标用户id获取关注信息
func QueryRelationByIds(ctx context.Context, currentId int64, userIds []int64) (map[int64]*Relation, error) {
	var relations []*Relation
	err := DB.WithContext(ctx).Where("from_user_id = ? AND to_user_id IN ?", currentId, userIds).Find(&relations).Error
	if err != nil {
		klog.Error("query relation by ids " + err.Error())
		return nil, err
	}
	relationMap := make(map[int64]*Relation)
	for _, relation := range relations {
		relationMap[relation.ToUserId] = relation
	}
	return relationMap, nil
}

// 增加当前用户的关注总数，增加其他用户的粉丝总数，创建关注记录
func Create(ctx context.Context, currentId int64, toUserId int64) error {
	relationRaw := &Relation{
		UserId:   currentId,
		ToUserId: toUserId,
	}
	DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Table("user").Where("id = ?", currentId).Update("follow_count", gorm.Expr("follow_count + ?", 1)).Error
		if err != nil {
			klog.Error("add user follow_count fail " + err.Error())
			return err
		}

		err = tx.Table("user").Where("id = ?", toUserId).Update("follower_count", gorm.Expr("follower_count + ?", 1)).Error
		if err != nil {
			klog.Error("add user follower_count fail " + err.Error())
			return err
		}

		err = tx.Table("relation").Create(&relationRaw).Error
		if err != nil {
			klog.Error("create relation record fail " + err.Error())
			return err
		}

		return nil
	})
	return nil
}

// 减少当前用户的关注总数，减少其他用户的粉丝总数，删除关注记录
func Delete(ctx context.Context, currentId int64, toUserId int64) error {
	var relationRaw *Relation
	DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Table("user").Where("id = ?", currentId).Update("follow_count", gorm.Expr("follow_count - ?", 1)).Error
		if err != nil {
			klog.Error("sub user follow_count fail " + err.Error())
			return err
		}

		err = tx.Table("user").Where("id = ?", toUserId).Update("follower_count", gorm.Expr("follower_count - ?", 1)).Error
		if err != nil {
			klog.Error("sub user follower_count fail " + err.Error())
			return err
		}

		err = tx.Table("relation").Where("user_id = ? AND to_user_id = ?", currentId, toUserId).Delete(&relationRaw).Error
		if err != nil {
			klog.Error("delete relation record famain " + err.Error())
			return err
		}
		return nil
	})
	return nil
}

// 通过用户id，查询该用户关注的用户，返回两者之间的关注记录
func QueryFollowById(ctx context.Context, userId int64) ([]*Relation, error) {
	var relations []*Relation
	err := DB.WithContext(ctx).Table("relation").Where("user_id = ?", userId).Find(&relations).Error
	if err != nil {
		klog.Error("query follow by id fail " + err.Error())
		return nil, err
	}
	return relations, nil
}

// 通过用户id，查询该用户的粉丝， 返回两者之间的关注记录
func QueryFollowerById(ctx context.Context, userId int64) ([]*Relation, error) {
	var relations []*Relation
	err := DB.WithContext(ctx).Table("relation").Where("to_user_id = ?", userId).Find(&relations).Error
	if err != nil {
		klog.Error("query follower by id fail " + err.Error())
		return nil, err
	}
	return relations, nil
}

