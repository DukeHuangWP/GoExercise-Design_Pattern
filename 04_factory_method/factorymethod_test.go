package factoryMethod

import (
	"fmt"
	"testing"
)

//** 工廠方法模式說明 **
//透過interface創建類別(struct)與定義行為(method)，每個類別依據需求設定不同屬性(struct-field)
//優點:結構非常明瞭，可以快速建立類別，並簡化定義出所有類別的行為
//缺點:由於行為(method)統一，定義屬性(struct-field)需要事先完善規劃
func TestFactoryMethod(t *testing.T) {
	ai1, err := NewAnimal(1, "蛇一號")
	if err != nil {
		fmt.Println(err)
	} else {
		DoCreepNSwimming(ai1)
	}

	ai2, err := NewAnimal(1, "蛇二號")
	if err != nil {
		fmt.Println(err)
	} else {
		DoCreepNSwimming(ai2)
	}

	ai3, err := NewAnimal(2, "青蛙零號")
	ai3.SetName("我反悔了改叫青蛙九九號")
	if err != nil {
		fmt.Println(err)
	} else {
		DoCreepNSwimming(ai3)
	}

	ai4, err := NewAnimal(2, "")
	if err != nil {
		fmt.Println(err)
	} else {
		DoCreepNSwimming(ai4)
	}

	_, err = NewAnimal(0, "unknow")
	if err != nil {
		fmt.Printf("%v\n", err) //Wrong gun type passed
	} // 工廠方法模式還需要針對未知類別(struct)處理

	// Hi, I'm a snake named 蛇一號, now I'm creeping
	// Hi, I'm a snake named 蛇一號, and I can't swimming
	// Hi, I'm a snake named 蛇二號, now I'm creeping
	// Hi, I'm a snake named 蛇二號, and I can't swimming
	// Hi, I'm a frog named 我反悔了改叫青蛙九九號, now I'm creeping
	// Hi, I'm a frog named 我反悔了改叫青蛙九九號, now I'm swimming
	// Nameless!
	// 0 is unknow type !

	fmt.Println("------------------------")
}

func DoCreepNSwimming(ai AnimalInterface) {
	fmt.Println(ai.Creep())
	fmt.Println(ai.Swimming())
}
