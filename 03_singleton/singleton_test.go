package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestSingleton(t *testing.T) {
	ins1 := CallMeeting("tester1")
	ins2 := CallMeeting("tester2")
	if ins1 != ins2 {
		t.Fatal("instance is not equal")
	}
	CloseMeeting("test0")
} //會議室的指標必須是同一個，即單體模式所指向的目標必須是唯一

//** 單體模式說明 **
//創建唯一公用struct後，該公用struct支援並行處理並且不會造成deadlock
//優點:通過非常簡單的上鎖機制，提高並行處理的安全性與效率
//缺點:想不到...
func TestParallelSingleton(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	for index := 0; index < 20; index++ {
		go CallMeeting("使用者" + fmt.Sprint(rand.Intn(100)))
		go CallMeeting("使用者" + fmt.Sprint(rand.Intn(100)))
	}

	// 使用者95 : 我到會議室發現沒人使用，並開始了一場會議！
	// 使用者48 : 我到會議室門口發現會議已經開始！，我發現門口沒人排隊於是就直接進入
	// 使用者36 : 我到會議室門口發現會議已經開始！，我依照順序在門口排隊等待進入！
	// 使用者85 : 我到會議室門口發現會議已經開始！，我發現門口沒人排隊於是就直接進入
	// ...
	select {}
}

