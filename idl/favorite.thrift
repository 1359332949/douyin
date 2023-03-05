namespace go favorite

struct BaseResp {
    1: i32 status_code
    2: string status_msg
    3: i64 service_time
}
struct FavoriteActionRequest {
    1: i64 user_id
    2: string token // 用户鉴权token
    3: i64 video_id  // 视频id
    4: i32 action_type   // 1-点赞，2-取消点赞
}

struct FavoriteActionResponse {
    1: i32 status_code   // 状态码，0-成功，其他值-失败
    2: string status_msg // 返回状态描述
}

struct FavoriteListRequest {
     1: i64 user_id //用户id
     2: string token //用户鉴权token
 }

struct FavoriteListResponse {
    1: i32 status_code //状态码，0-成功，其他值失败
    2:  string status_msg //返回状态描述
    3: list<Video> video_list //用户点赞视频列表
}


service FavoriteService {
    
    FavoriteActionResponse FavoriteAction(1:FavoriteActionRequest req) // 用户点赞
    FavoriteListResponse FavoriteList(1:FavoriteListRequest req) // 用户点赞列表
   
}

