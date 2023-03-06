
package consts

const (
	
	UserTableName   = "user"
	VideoTableName   = "video"
	FavoriteTableName = "favorite"
	CommentTableName = "comment"
	RelationTableName   = "relation"
	MessageTableName   = "message"
	
	SecretKey       = "secret key"
	IdentityKey     = "id"
	Total           = "total"
	Messages        = "messages"
	Videos			= "videos"

	ApiServiceName  = "api"
	UserServiceName = "user"
	VideoServiceName = "video"
	FavoriteServiceName = "favorite"
	CommentServiceName = "comment"
	RelationServiceName = "relation"
	MessageServiceName = "message"

	MySQLDefaultDSN = "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local"
	TCP             = "tcp"

	UserServiceAddr = ":9000"
	VideoServiceAddr = ":9001"
	FavoriteServiceAddr = ":9002"
	CommentServiceAddr = ":9003"
	RelationServiceAddr = ":9004"
	MessageServiceAddr = ":9005"

	ExportEndpoint  = ":4317"
	ETCDAddress     = "127.0.0.1:2379"
	DefaultLimit    = 10
	VideoSavePath   = "/home/guo/go/impl/douyin/pkg/video"
	CoverPath = "/home/guo/go/impl/douyin/pkg/img"
	StatusOK = 200

	//favorite actiontype,1是点赞，2是取消点赞
	Like   = 1
	Unlike = 2
)
