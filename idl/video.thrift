namespace go video
include "user.thrift"
enum ErrCode {
    SuccessCode                = 0
    ServiceErrCode             = 10001
    ParamErrCode               = 10002
    UserAlreadyExistErrCode    = 10003
    AuthorizationFailedErrCode = 10004
}
struct BaseResp {
    1: i32 status_code
    2: string status_msg
    3: i64 service_time
}

struct Video {
    1: i64 id;
    2: user.User author;
    3: string play_url;
    4: string cover_url;
    5: i64 favorite_count;
    6: i64 comment_count;
    7: bool is_favorite;
    8: string title;
}


struct FeedRequest {
    1: i64 latest_time; // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
    2: string token; // 可选参数，登录用户设置
    3: i64 user_id
}

// 例如当前请求的latest_time为9:00，那么返回的视频列表时间戳为[8:55,7:40, 6:30, 6:00]
// 所有这些视频中，最早发布的是 6:00的视频，那么6:00作为下一次请求时的latest_time
// 那么下次请求返回的视频时间戳就会小于6:00

struct FeedResponse {
    1: i32 status_code; // 状态码，0-成功，其他值-失败
    2: string status_msg; // 返回状态描述
    3: list<Video> video_list; // 视频列表
    4: i64 next_time; // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

struct VideoIdRequest{
    1:i64 video_id ;
    2:i64 search_id ;
}


struct PublishActionRequest {
    1: i64 user_id;
    2: string file_url;
    3: string cover_url;
    4: string title;
}

struct PublishActionResponse {
    1: i32 status_code;
    2: string status_msg;
}

struct PublishListRequest {
    1: i64 user_id;
    2: string token;
}

struct PublishListResponse {
    1: i32 status_code;
    2: string status_msg;
    3: list<Video> video_list;
}

struct QueryVideoByVideoIdsRequest {
    1: list<i64> video_ids
}

struct QueryVideoByVideoIdsResponse {
    1: i32 status_code;
    2: string status_msg;
    3: list<Video> video_list;
}

service VideoService {
    
    PublishActionResponse PublishAction(1: PublishActionRequest req)
    PublishListResponse PublishList(1: PublishListRequest req)
    FeedResponse GetVideoFeed (1:FeedRequest req)
    QueryVideoByVideoIdsResponse QueryVideoByVideoIds (1: QueryVideoByVideoIdsRequest req)
    
}