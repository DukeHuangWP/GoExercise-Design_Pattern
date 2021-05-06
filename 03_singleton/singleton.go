package main

import (
	"fmt"
	"sync"
	"time"
)

var locker = &sync.Mutex{}

type Meeting struct {
	LastTime int64
}

var MeetingRoom *Meeting

func CallMeeting(name string) *Meeting {
	if MeetingRoom == nil {
		locker.Lock()
		defer locker.Unlock()
		if MeetingRoom == nil {
			fmt.Printf("%v : 我到會議室發現沒人使用，並開始了一場會議！\n", name)
			MeetingRoom = &Meeting{LastTime: time.Now().Unix()}
		} else {
			fmt.Printf("%v : 我到會議室門口發現會議已經開始！，我依照順序在門口排隊等待進入！\n", name)
		}
	} else {
		fmt.Printf("%v : 我到會議室門口發現會議已經開始！，我發現門口沒人排隊於是就直接進入\n", name)
	}
	return MeetingRoom
}

func CloseMeeting(name string) *Meeting {
	MeetingRoom = nil
	fmt.Printf("%v : 我關閉了會議室的使用！\n", name)
	return MeetingRoom
}
