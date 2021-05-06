package builder

import (
	"fmt"
	"testing"
)

//** 建造者模式說明 **
//透過Director(struct)透過接口選擇子Builder設定屬性,並將物件輸出
//優點:結構分工細膩可自由將建造過程細部拆解或是一次性輸出
//缺點:Director分類構造較為繁瑣
func TestBuilder1(t *testing.T) {

	director := &Director{Snake: &Snake{}}
	snake := director.Snake.SetName("蛇一號").SetLength(1234).Build()
	fmt.Println(snake.Name)
	snake.Scramble()

	director = &Director{Frog: &Frog{}}
	frog := director.Frog.SetName("青蛙一號").SetHight(5678).Build()
	fmt.Println(frog.Name)
	frog.Swimming()

	// 	蛇一號
	// I'm scrambling, my name is 蛇一號 and 1234 cm
	// 青蛙一號
	// I'm swimming, my name is 青蛙一號 and 5678 cm

}
