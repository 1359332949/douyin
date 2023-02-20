namespace go api

enum ErrCode {
    SuccessCode                = 0
    ServiceErrCode             = 10001
    ParamErrCode               = 10002
    UserAlreadyExistErrCode    = 10003
    AuthorizationFailedErrCode = 10004

    MessageIsNullErrCode    = 90003
    
}


struct UserLogin {
    1: string username // 注册用户名，最长32个字符
    2: string password // 密码，最长32个字符
}

struct User {
    1: i64 id
    2: string name
    3: i64 follow_count // 关注总数
    4: i64 follower_count  // 粉丝总数
    5: bool is_follow  // true-已关注，false-未关注

}

struct Video {
    1: i64 video_id;
    2: User author;
    3: string play_url;
    4: string cover_url;
    5: i64 favorite_count;
    6: i64 comment_count;
    7: bool is_favorite;
    8: string title;
}

struct Message {
    1:required i64 id                  // 消息id
    2:required i64 to_user_id          // 该消息接收者的id
    3:required i64 from_user_id        // 该消息发送者的id
    4:required string content         // 消息内容
    5:optional i64 create_time      // 消息创建时间
}

struct FriendUser {
    1: i64 id // 用户id
    2: string name // 用户名称
    3: i64 follow_count // 关注总数
    4: i64 follower_count // 粉丝总数
    5: bool is_follow // true-已关注，false-未关注
    6: string avatar // 用户头像Url
    7: string message // 和该好友的最新聊天消息
    8: i64 msg_type // message消息的类型，0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息
}

struct LoginUserRequest {
    1: string username
    2: string password
}

struct LoginUserResponse {
    1: i32 status_code
    2: string status_msg
    3: i64 user_id
    4: string token
}

struct RegisterUserRequest {
    1: string username
    2: string password
}

struct RegisterUserResponse {
    1: i32 status_code
    2: string status_msg
    3: i64 user_id
    4: string token
}

struct UserInfoRequest {
    1: string user_id
    2: string token
}

struct UserInfoResponse {
    1: i32 status_code
    2: string status_msg
    3: User user

}

struct PublishActionRequest {
    1: i64 user_id;
    2: string token;
    3: binary data;
    4: string title;
}

struct PublishActionResponse {
    1: i32 status_code;
    2: string status_msg;
}

struct PublishListRequest {
    
    1: string token;
    2: string user_id;
}

struct PublishListResponse {
    1: i32 status_code;
    2: string status_msg;
    3: list<Video> video_list;
}




struct RelationActionRequest {
    1: i64 user_id // 用户id
    2: string token // 用户鉴权token
    3: i64 to_user_id // 对方用户id
    4: i32 action_type // 1-关注，2-取消关注
}

struct RelationActionResponse {
    1: i32 status_code // 状态码，0-成功，其他值-失败
    2: string status_msg // 返回状态描述
}
struct RelationFollowListRequest {
    1: i64 user_id // 用户id
    2: string token // 用户鉴权token
}
struct RelationFollowListResponse {
    1: i32 status_code // 状态码，0-成功，其他值-失败
    2: string status_msg // 返回状态描述
    3: list<User> user_list // 用户信息列表
}
struct RelationFollowerListRequest {
    1: i64 user_id // 用户id
    2: string token // 用户鉴权token
}
struct RelationFollowerListResponse {
    1: i32 status_code // 状态码，0-成功，其他值-失败
    2: string status_msg // 返回状态描述
    3: list<User> user_list // 用户列表
}
struct RelationFriendListRequest {
    1: i64 user_id // 用户id
    2: string token // 用户鉴权token
}
struct RelationFriendListResponse {
    1: i32 status_code // 状态码，0-成功，其他值-失败
    2: string status_msg // 返回状态描述
    3: list<FriendUser> user_list // 用户列表
}

struct MessageChatRequest {
    1:required i64 from_user_id          // 用户id
    2:required i64 to_user_id        // 对方用户id
}

struct MessageChatResponse {
    1: i32 status_code
    2: string status_msg
    3: list<Message> messages
    
}

struct MessageActionRequest {
    1:required i64 from_user_id           // 用户鉴权token
    2:required i64 to_user_id         // 对方用户id
    3:required i32 action_type       // 1-发送消息
    4:required string content                // 消息内容
}

struct MessageActionResponse {
    1: i32 status_code
    2: string status_msg
}

service UserService {
    LoginUserResponse LoginUser(1: LoginUserRequest req) (api.post="/douyin/user/login/")
    RegisterUserResponse RegisterUser(1: RegisterUserRequest req) (api.post="/douyin/user/register/")
    UserInfoResponse UserInfo(1: UserInfoRequest req) (api.get="/douyin/user/")
    PublishActionResponse PublishAction(1: PublishActionRequest req) (api.post="/douyin/publish/action/");
    PublishListResponse PublishList(1: PublishListRequest req) (api.get="/douyin/publish/list/");
}

service RelationService {
    RelationActionResponse RelationAction (1: RelationActionRequest req) (api.post="/douyin/relation/action/")
    RelationFollowListResponse RelationFollowList (1: RelationFollowListRequest req) (api.get="/douyin/relation/follow/list/")
    RelationFollowerListResponse RelationFollowerList (1: RelationFollowerListRequest req) (api.get="/douyin/relation/follower/list/")
    RelationFriendListResponse RelationFriendList (1: RelationFriendListRequest req) (api.get="/douyin/relation/friend/list/")
    MessageChatResponse MessageChat(1: MessageChatRequest req) (api.get="/douyin/message/chat/")               // 消息记录
    MessageActionResponse MessageAction(1: MessageActionRequest req) (api.post="/douyin/message/action/")         // 发送消息
}
