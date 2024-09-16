package global

import (
	"log"
	"sync"
)

type Message struct {
	DisplayLikeCount     string        `json:"displayLikeCount"`
	DisplayWatchingCount string        `json:"displayWatchingCount"`
	GiftFeeds            []GiftFeed    `json:"giftFeeds"`
	CommentFeeds         []CommentFeed `json:"commentFeeds"`
	LikeFeeds            []LikeFeed    `json:"likeFeeds"`
	// 其他字段...
}
type GiftFeed struct {
	BatchSize      int    `json:"batchSize"`
	ComboCount     int    `json:"comboCount"`
	DeviceHash     string `json:"deviceHash"`
	ExpireDuration string `json:"expireDuration"`
	GiftID         int    `json:"giftId"`
	MergeKey       string `json:"mergeKey"`
	Rank           int    `json:"rank"`
	StyleType      string `json:"styleType"`
	User           User   `json:"user"`
}

type User struct {
	PrincipalID string `json:"principalId"`
	UserName    string `json:"userName"`
}

type CommentFeed struct {
	Content    string `json:"content"`
	DeviceHash string `json:"deviceHash"`
	ShowType   string `json:"showType"`
	User       User   `json:"user"`
}

type LikeFeed struct {
	User       User   `json:"user"`
	DeviceHash string `json:"deviceHash"`
}

type GlobalData struct {
	message Message
	lock    sync.RWMutex
}

var globalData struct {
	message Message
	lock    sync.RWMutex
}

func SetMessage(msg Message) {
	globalData.lock.Lock()
	defer globalData.lock.Unlock()
	globalData.message = msg
	log.Printf("写入成功，messagelen=%v|%v|%v\n", len(msg.GiftFeeds), len(msg.CommentFeeds), len(msg.LikeFeeds))
}

func GetMessage() Message {
	globalData.lock.RLock()
	defer globalData.lock.RUnlock()
	return globalData.message
}
