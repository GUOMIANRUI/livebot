package global

import (
	"fmt"
	"sync"
)

var (
	modulenameids = make(map[string][]int)
	trafficlimits = make(map[string]int32)
	lineIDs2      = make(map[int]string)
	openaikey     = ""
	proxyURL      = "http://127.0.0.1:7890"

	mu sync.Mutex // 用于保护 modulenameids 和 trafficlimits 的互斥锁
)

func GetOpenaikey() string {
	mu.Lock()
	defer mu.Unlock()
	return openaikey
}

func GetproxyURL() string {
	mu.Lock()
	defer mu.Unlock()
	return proxyURL
}

// GetModuleNameIDs 返回 modulenameids 的副本
func GetModuleNameIDs() map[string][]int {
	mu.Lock()
	defer mu.Unlock()

	copy := make(map[string][]int)
	for k, v := range modulenameids {
		copy[k] = append([]int(nil), v...)
	}
	return copy
}

// SetModuleNameIDs 设置 modulenameids
func SetModuleNameIDs(new map[string][]int) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Println("正在修改modulenameids...")
	modulenameids = new
}

// GetTrafficLimits 返回 trafficlimits 的副本
func GetTrafficLimits() map[string]int32 {
	mu.Lock()
	defer mu.Unlock()

	copy := make(map[string]int32)
	for k, v := range trafficlimits {
		copy[k] = v
	}
	return copy
}

// SetTrafficLimits 设置 trafficlimits
func SetTrafficLimits(new map[string]int32) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Println("正在修改trafficlimits...")
	trafficlimits = new
}

// GetLineIDs2 返回 lineIDs2 的副本
func GetLineIDs2() map[int]string {
	mu.Lock()
	defer mu.Unlock()

	copy := make(map[int]string)
	for k, v := range lineIDs2 {
		copy[k] = v
	}
	return copy
}

// SetLineIDs2 设置 lineIDs2
func SetLineIDs2(new map[int]string) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Println("正在修改lineIDs2...")
	lineIDs2 = new
}
