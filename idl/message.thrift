namespace go message
enum ErrCode {
    SuccessCode                = 0
    ServiceErrCode             = 90001
    ParamErrCode               = 90002
    MessageIsNullErrCode    = 90003
    AuthorizationFailedErrCode = 90004
}
struct BaseResp {
    1: i32 status_code
    2: string status_msg
    3: i64 service_time
}


struct Message {
    1: i64 id                  // 消息id
    2: i64 to_user_id          // 该消息接收者的id
    3: i64 from_user_id        // 该消息发送者的id
    4: string content         // 消息内容
    5: i64 create_time      // 消息创建时间
}


struct MessageChatRequest {
    1: i64 from_user_id          // 用户id
    2: string token       
    3: i64 to_user_id        // 对方用户id
}

struct MessageChatResponse {
    1: i32 status_code
    2: string status_msg
    3: list<Message> messages
    4: i64 create_time
}

struct MessageActionRequest {
    1: i64 from_user_id           // 用户鉴权token
    2: string token
    3: i64 to_user_id         // 对方用户id
    4: i64 action_type       // 1-发送消息
    5: string content                // 消息内容
}

struct MessageActionResponse {
    1: i32 status_code
    2: string status_msg
}



service MessageService {
    
    MessageChatResponse MessageChat(1: MessageChatRequest req)               // 消息记录
    MessageActionResponse MessageAction(1: MessageActionRequest req)         // 发送消息
}
