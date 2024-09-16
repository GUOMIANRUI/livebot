package savetofile

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/vtb-link/bianka/live"
)

func ReadFromFile(filename string, data interface{}) error {
	// 从文件中读取数据
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	// 解析 JSON 数据
	err = json.Unmarshal(content, data)
	if err != nil {
		return err
	}

	return nil
}

func SaveToFile(data interface{}, filename string) error {
	// 将数据序列化为 JSON 字符串
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	// 保存 JSON 数据到文件
	err = ioutil.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

// 实时增量保存 danmuData 数据到 json 文件
func SaveDanmuData(danmuData *live.CmdLiveOpenPlatformDanmuData) {
	messagetmps := make([]live.CmdLiveOpenPlatformDanmuData, 0)
	if _, err := os.Stat("danmuData.json"); !os.IsNotExist(err) {

		// 读取文件中的数据
		// messagetmp := &live.CmdLiveOpenPlatformDanmuData{}
		file, err := os.Open("danmuData.json")
		if err != nil {
			log.Println("读取文件时发生错误：", err)
			return
		}
		defer file.Close()

		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			log.Println("读取文件时发生错误：", err)
			return
		}

		err = json.Unmarshal(bytes, &messagetmps)
		if err != nil {
			log.Println("解析文件数据时发生错误：", err)
			return
		}

		// messagetmp.RoomID = danmuData.RoomID
		// messagetmp.Uname = danmuData.Uname
		// messagetmp.Msg = danmuData.Msg
		// messagetmp.FansMedalLevel = danmuData.FansMedalLevel
		// messagetmp.FansMedalName = danmuData.FansMedalName
		// messagetmp.FansMedalWearingStatus = danmuData.FansMedalWearingStatus
		// messagetmp.GuardLevel = danmuData.GuardLevel
		// messagetmp.Timestamp = danmuData.Timestamp
		// messagetmp.UFace = danmuData.UFace
		// messagetmp.EmojiImgUrl = danmuData.EmojiImgUrl

		messagetmps = append(messagetmps, *danmuData)
		err = SaveToFile(messagetmps, "danmuData.json")
		if err != nil {
			log.Println("保存数据到文件时发生错误：", err)
		} else {
			log.Println("数据已成功保存到文件")
		}
	} else {
		// 文件不存在，直接保存数据
		messagetmps = append(messagetmps, *danmuData)
		err := SaveToFile(messagetmps, "danmuData.json")
		if err != nil {
			log.Println("保存数据到文件时发生错误：", err)
		} else {
			log.Println("数据已成功保存到文件")
		}
	}
}

func SaveGiftData(giftData *live.CmdLiveOpenPlatformSendGiftData) {
	messagetmps := make([]live.CmdLiveOpenPlatformSendGiftData, 0)
	if _, err := os.Stat("giftData.json"); !os.IsNotExist(err) {
		// 读取文件中的数据
		// messagetmp := &live.CmdLiveOpenPlatformSendGiftData{}
		file, err := os.Open("giftData.json")
		if err != nil {
			log.Println("读取文件时发生错误：", err)
			return
		}
		defer file.Close()

		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			log.Println("读取文件时发生错误：", err)
			return
		}

		err = json.Unmarshal(bytes, &messagetmps)
		if err != nil {
			log.Println("解析文件数据时发生错误：", err)
			return
		}

		// messagetmp.Uname = giftData.Uname
		// messagetmp.GiftID = giftData.GiftID
		// messagetmp.GiftName = giftData.GiftName
		// messagetmp.FansMedalLevel = giftData.FansMedalLevel
		// messagetmp.FansMedalName = giftData.FansMedalName
		// messagetmp.FansMedalWearingStatus = giftData.FansMedalWearingStatus
		// messagetmp.GuardLevel = giftData.GuardLevel

		// messagetmp.AnchorInfo.Uid = giftData.AnchorInfo.Uid
		// messagetmp.AnchorInfo.Uname = giftData.AnchorInfo.Uname
		// messagetmp.AnchorInfo.Uface = giftData.AnchorInfo.Uface

		// messagetmp.ComboGift = giftData.ComboGift
		// messagetmp.ComboInfo.ComboBaseNum = giftData.ComboInfo.ComboBaseNum
		// messagetmp.ComboInfo.ComboCount = giftData.ComboInfo.ComboCount
		// messagetmp.ComboInfo.ComboID = giftData.ComboInfo.ComboID
		// messagetmp.ComboInfo.ComboTimeout = giftData.ComboInfo.ComboTimeout

		messagetmps = append(messagetmps, *giftData)
		err = SaveToFile(messagetmps, "giftData.json")
		if err != nil {
			log.Println("保存数据到文件时发生错误：", err)
		} else {
			log.Println("数据已成功保存到文件")
		}
	} else {
		// 文件不存在，直接保存数据
		messagetmps = append(messagetmps, *giftData)
		err := SaveToFile(messagetmps, "giftData.json")
		if err != nil {
			log.Println("保存数据到文件时发生错误：", err)
		} else {
			log.Println("数据已成功保存到文件")
		}
	}
}

func SaveSuperChatData(superChatData *live.CmdLiveOpenPlatformSuperChatData) {
	messagetmps := make([]live.CmdLiveOpenPlatformSuperChatData, 0)
	if _, err := os.Stat("superChatData.json"); !os.IsNotExist(err) {
		// 读取文件中的数据
		// messagetmp := &live.CmdLiveOpenPlatformSuperChatData{}
		file, err := os.Open("superChatData.json")
		if err != nil {
			log.Println("读取文件时发生错误：", err)
			return
		}
		defer file.Close()

		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			log.Println("读取文件时发生错误：", err)
			return
		}

		err = json.Unmarshal(bytes, &messagetmps)
		if err != nil {
			log.Println("解析文件数据时发生错误：", err)
			return
		}

		// messagetmp.RoomID = superChatData.RoomID
		// messagetmp.Uname = superChatData.Uname
		// messagetmp.Message = superChatData.Message
		// messagetmp.FansMedalLevel = superChatData.FansMedalLevel
		// messagetmp.FansMedalName = superChatData.FansMedalName
		// messagetmp.FansMedalWearingStatus = superChatData.FansMedalWearingStatus
		// messagetmp.GuardLevel = superChatData.GuardLevel
		// messagetmp.Timestamp = superChatData.Timestamp
		// messagetmp.Uface = superChatData.Uface
		// messagetmp.StartTime = superChatData.StartTime
		// messagetmp.EndTime = superChatData.EndTime

		messagetmps = append(messagetmps, *superChatData)
		err = SaveToFile(messagetmps, "superChatData.json")
		if err != nil {
			log.Println("保存数据到文件时发生错误：", err)
		} else {
			log.Println("数据已成功保存到文件")
		}
	} else {
		// 文件不存在，直接保存数据
		messagetmps = append(messagetmps, *superChatData)
		err := SaveToFile(messagetmps, "superChatData.json")
		if err != nil {
			log.Println("保存数据到文件时发生错误：", err)
		} else {
			log.Println("数据已成功保存到文件")
		}
	}
}

func SaveSuperChatDelData(superChatDelData *live.CmdLiveOpenPlatformSuperChatDelData) {
	if _, err := os.Stat("superChatDelData.json"); !os.IsNotExist(err) {
		// 读取文件中的数据
		messagetmp := &live.CmdLiveOpenPlatformSuperChatDelData{}
		file, err := os.Open("superChatDelData.json")
		if err != nil {
			log.Println("读取文件时发生错误：", err)
			return
		}
		defer file.Close()

		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			log.Println("读取文件时发生错误：", err)
			return
		}

		err = json.Unmarshal(bytes, messagetmp)
		if err != nil {
			log.Println("解析文件数据时发生错误：", err)
			return
		}

		messagetmp.RoomID = superChatDelData.RoomID
		messagetmp.MessageIds = superChatDelData.MessageIds
		messagetmp.MsgID = superChatDelData.MsgID

		err = SaveToFile(messagetmp, "superChatDelData.json")
		if err != nil {
			log.Println("保存数据到文件时发生错误：", err)
		} else {
			log.Println("数据已成功保存到文件")
		}
	} else {
		// 文件不存在，直接保存数据
		err := SaveToFile(superChatDelData, "superChatDelData.json")
		if err != nil {
			log.Println("保存数据到文件时发生错误：", err)
		} else {
			log.Println("数据已成功保存到文件")
		}
	}
}

func SaveGuardData(guardData *live.CmdLiveOpenPlatformGuardData) {
	messagetmps := make([]live.CmdLiveOpenPlatformGuardData, 0)
	if _, err := os.Stat("guardData.json"); !os.IsNotExist(err) {
		// 读取文件中的数据
		// messagetmp := &live.CmdLiveOpenPlatformGuardData{}
		file, err := os.Open("guardData.json")
		if err != nil {
			log.Println("读取文件时发生错误：", err)
			return
		}
		defer file.Close()

		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			log.Println("读取文件时发生错误：", err)
			return
		}

		err = json.Unmarshal(bytes, &messagetmps)
		if err != nil {
			log.Println("解析文件数据时发生错误：", err)
			return
		}

		// messagetmp.UserInfo = guardData.UserInfo
		// messagetmp.RoomID = guardData.RoomID
		// messagetmp.GuardLevel = guardData.GuardLevel
		// messagetmp.Timestamp = guardData.Timestamp
		// messagetmp.GuardNum = guardData.GuardNum
		// messagetmp.GuardUnit = guardData.GuardUnit
		// messagetmp.FansMedalLevel = guardData.FansMedalLevel
		// messagetmp.FansMedalName = guardData.FansMedalName
		// messagetmp.FansMedalWearingStatus = guardData.FansMedalWearingStatus
		// messagetmp.MsgID = guardData.MsgID

		messagetmps = append(messagetmps, *guardData)
		err = SaveToFile(messagetmps, "guardData.json")
		if err != nil {
			log.Println("保存数据到文件时发生错误：", err)
		} else {
			log.Println("数据已成功保存到文件")
		}
	} else {
		// 文件不存在，直接保存数据
		messagetmps = append(messagetmps, *guardData)
		err := SaveToFile(messagetmps, "guardData.json")
		if err != nil {
			log.Println("保存数据到文件时发生错误：", err)
		} else {
			log.Println("数据已成功保存到文件")
		}
	}
}

func SaveLikeData(likeData *live.CmdLiveOpenPlatformLikeData) {
	messagetmps := make([]live.CmdLiveOpenPlatformLikeData, 0)
	if _, err := os.Stat("likeData.json"); !os.IsNotExist(err) {
		// 读取文件中的数据
		// messagetmp := &live.CmdLiveOpenPlatformLikeData{}
		file, err := os.Open("likeData.json")
		if err != nil {
			log.Println("读取文件时发生错误：", err)
			return
		}
		defer file.Close()

		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			log.Println("读取文件时发生错误：", err)
			return
		}

		err = json.Unmarshal(bytes, &messagetmps)
		if err != nil {
			log.Println("解析文件数据时发生错误：", err)
			return
		}

		// messagetmp.RoomID = likeData.RoomID
		// messagetmp.Uname = likeData.Uname
		// messagetmp.Uid = likeData.Uid
		// messagetmp.Timestamp = likeData.Timestamp
		// messagetmp.Uface = likeData.Uface
		// messagetmp.LikeText = likeData.LikeText
		// messagetmp.MsgID = likeData.MsgID

		messagetmps = append(messagetmps, *likeData)
		err = SaveToFile(messagetmps, "likeData.json")
		if err != nil {
			log.Println("保存数据到文件时发生错误：", err)
		} else {
			log.Println("数据已成功保存到文件")
		}
	} else {
		// 文件不存在，直接保存数据
		messagetmps = append(messagetmps, *likeData)
		err := SaveToFile(messagetmps, "likeData.json")
		if err != nil {
			log.Println("保存数据到文件时发生错误：", err)
		} else {
			log.Println("数据已成功保存到文件")
		}
	}
}
