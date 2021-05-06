package prototype

import (
	"fmt"
	"testing"
)

//** 建造者模式說明 **
//透過Director(struct)透過接口選擇子Builder設定屬性,並將物件輸出
//優點:結構分工細膩可自由將建造過程細部拆解或是一次性輸出
//缺點:Director分類構造較為繁瑣
func TestPrototype(t *testing.T) {

	temple := &Snake{ProtoName: "模範蛇", Level: 99}
	snake1,_ := temple.Clone()
	snake1.Set("蛇一號",100)
	fmt.Println(snake1.ProtoName)

	snake2,_ := snake1.Clone()
	fmt.Println(snake2.ProtoName)

	snake1.Scramble()
	snake2.Scramble()
	// 範蛇_clone
	// 模範蛇_clone_clone
	// I'm scrambling, my name is 蛇一號, length is 100 ,level is 99 
	// I'm scrambling, my name is 模範蛇_clone_clone, length is 0 ,level is 99 
}