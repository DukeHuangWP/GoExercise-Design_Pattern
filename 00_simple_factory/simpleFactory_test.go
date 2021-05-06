package simpleFactory

import (
	"fmt"
	"testing"
)

//** 簡單工廠模式說明 **
//透過interface創建各自類別(struct)，做出不同行為(type-method)對應
//優點:結構非常簡單,易於批次複製與生產
//缺點:當行為(method)需要修改時，每個類別(struct)的行為全部都需要作修改
func TestSimpleFactory(t *testing.T) {
	ai1 := NewAnimal(1)
	fmt.Println(ai1.SetName("蛇一號"))
	DoCreepNSwimming(ai1)

	ai2 := NewAnimal(1)
	fmt.Println(ai2.SetName("蛇二號"))
	DoCreepNSwimming(ai2)

	ai3 := NewAnimal(2)
	fmt.Println(ai3.SetName("青蛙零號"))
	DoCreepNSwimming(ai3)

	ai4 := NewAnimal(2)
	fmt.Println(ai4.SetName(""))
	DoCreepNSwimming(ai4)

	// Hi, I'm a snake and my name is 蛇一號
	// Hi, I'm a snake named 蛇一號, now I'm creeping
	// Hi, I'm a snake named 蛇一號, and I can't swimming
	// Hi, I'm a snake and my name is 蛇二號
	// Hi, I'm a snake named 蛇二號, now I'm creeping
	// Hi, I'm a snake named 蛇二號, and I can't swimming
	// Hi, I'm a frog and my name is 青蛙零號
	// Hi, I'm a frog named 青蛙零號, now I'm creeping
	// Hi, I'm a frog named 青蛙零號, now I'm swimming

	// api0 := NewAnimal(0)
	// fmt.Println(api0.SetName(""))
	// 簡單工廠模式還需要針對未知類別(struct)處理
}

func DoCreepNSwimming(ai AnimalInterface) {
	fmt.Println(ai.Creep())
	fmt.Println(ai.Swimming())
}
