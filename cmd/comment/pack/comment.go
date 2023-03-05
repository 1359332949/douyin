
package pack

import (
	"github.com/1359332949/douyin/cmd/comment/dal/db"
	"github.com/1359332949/douyin/kitex_gen/comment"
	"github.com/1359332949/douyin/kitex_gen/user"
)

// Comment pack comment info
func Comment(c *db.Comment, u *user.User) *comment.Comment {
	if u == nil {
		return nil
	}

	return &comment.Comment{
		Id: int64(c.ID), 
		User: u,
		// VideoId: int64(c.VideoId),
		Content: string(c.Content),
		CreateDate: string(c.CreateDate),
		
		}
}