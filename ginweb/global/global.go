package global

// CmdLiveOpenPlatformDanmuData 弹幕数据
type CmdLiveOpenPlatformDanmuData struct {
	RoomID                 int    `json:"room_id"`
	Uid                    int    `json:"uid"`
	Uname                  string `json:"uname"`
	Msg                    string `json:"msg"`
	MsgID                  string `json:"msg_id"`
	FansMedalLevel         int    `json:"fans_medal_level"`
	FansMedalName          string `json:"fans_medal_name"`
	FansMedalWearingStatus bool   `json:"fans_medal_wearing_status"`
	GuardLevel             int    `json:"guard_level"`
	Timestamp              int    `json:"timestamp"`
	UFace                  string `json:"uface"`
	EmojiImgUrl            string `json:"emoji_img_url"`
	DmType                 int    `json:"dm_type"`
}

// CmdLiveOpenPlatformSendGiftData 礼物数据
type CmdLiveOpenPlatformSendGiftData struct {
	RoomID                 int    `json:"room_id"`
	Uid                    int    `json:"uid"`
	Uname                  string `json:"uname"`
	Uface                  string `json:"uface"`
	GiftID                 int    `json:"gift_id"`
	GiftName               string `json:"gift_name"`
	GiftNum                int    `json:"gift_num"`
	Price                  int    `json:"price"`
	Paid                   bool   `json:"paid"`
	FansMedalLevel         int    `json:"fans_medal_level"`
	FansMedalName          string `json:"fans_medal_name"`
	FansMedalWearingStatus bool   `json:"fans_medal_wearing_status"`
	GuardLevel             int    `json:"guard_level"`
	Timestamp              int    `json:"timestamp"`
	MsgID                  string `json:"msg_id"`
	AnchorInfo             struct {
		Uid   int    `json:"uid"`
		Uname string `json:"uname"`
		Uface string `json:"uface"`
	} `json:"anchor_info"`
	GiftIcon  string `json:"gift_icon"`
	ComboGift bool   `json:"combo_gift"`
	ComboInfo struct {
		ComboBaseNum int    `json:"combo_base_num"`
		ComboCount   int    `json:"combo_count"`
		ComboID      string `json:"combo_id"`
		ComboTimeout int    `json:"combo_timeout"`
	} `json:"combo_info"`
}

// CmdLiveOpenPlatformSuperChatData SC数据
type CmdLiveOpenPlatformSuperChatData struct {
	RoomID                 int    `json:"room_id"`
	Uid                    int    `json:"uid"`
	Uname                  string `json:"uname"`
	Uface                  string `json:"uface"`
	MessageID              int    `json:"message_id"`
	Message                string `json:"message"`
	MsgID                  string `json:"msg_id"`
	Rmb                    int    `json:"rmb"`
	Timestamp              int    `json:"timestamp"`
	StartTime              int    `json:"start_time"`
	EndTime                int    `json:"end_time"`
	GuardLevel             int    `json:"guard_level"`
	FansMedalLevel         int    `json:"fans_medal_level"`
	FansMedalName          string `json:"fans_medal_name"`
	FansMedalWearingStatus bool   `json:"fans_medal_wearing_status"`
}

// CmdLiveOpenPlatformSuperChatDelData SC删除数据
type CmdLiveOpenPlatformSuperChatDelData struct {
	RoomID     int    `json:"room_id"`
	MessageIds []int  `json:"message_ids"`
	MsgID      string `json:"msg_id"`
}

// CmdLiveOpenPlatformGuardData 付费大航海数据
type CmdLiveOpenPlatformGuardData struct {
	UserInfo struct {
		Uid   int    `json:"uid"`
		Uname string `json:"uname"`
		Uface string `json:"uface"`
	} `json:"user_info"`
	GuardLevel             int    `json:"guard_level"`
	GuardNum               int    `json:"guard_num"`
	GuardUnit              string `json:"guard_unit"`
	FansMedalLevel         int    `json:"fans_medal_level"`
	FansMedalName          string `json:"fans_medal_name"`
	FansMedalWearingStatus bool   `json:"fans_medal_wearing_status"`
	Timestamp              int    `json:"timestamp"`
	RoomID                 int    `json:"room_id"`
	MsgID                  string `json:"msg_id"`
}

// CmdLiveOpenPlatformLikeData 点赞数据
type CmdLiveOpenPlatformLikeData struct {
	Uname                  string `json:"uname"`
	Uid                    int    `json:"uid"`
	Uface                  string `json:"uface"`
	Timestamp              int    `json:"timestamp"`
	LikeText               string `json:"like_text"`
	FansMedalWearingStatus bool   `json:"fans_medal_wearing_status"`
	FansMedalName          string `json:"fans_medal_name"`
	FansMedalLevel         int    `json:"fans_medal_level"`
	MsgID                  string `json:"msg_id"`
	RoomID                 int    `json:"room_id"`
}

// CmdLiveOpenPlatformAuthData 鉴权数据
type CmdLiveOpenPlatformAuthData struct {
	Code int64 `json:"code,omitempty"`
}
