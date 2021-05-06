package prototype

import (
	"testing"
)

//** 原型模式說明 **
//可以透過接口載入原型範本,並且所輸出的struct具備複製(clone)功能
//優點:因輸出的struct具備複製(clone)功能,非常容易生產規格近似的struct
//缺點:複製前clone需要規劃餅設定好原型範本
func TestPrototype(t *testing.T) {

	frog := &Frog{Animal: Animal{ProtoName: "模範青蛙", Level: 11, Name: "青蛙一號"}}
	frog.Swimming() //I'm swimming, my name is 青蛙一號 (prototype:模範青蛙), weight is 0 ,level is 11

	snakeTemple := Animal{ProtoName: "模範蛇", Level: 99}
	snake1 := &Snake{Animal: snakeTemple}
	snake1.Set("蛇一號", 100)
	snake2 := &Snake{Animal: *snake1.Clone()}

	snake1.Scramble() //I'm scrambling, my name is 蛇一號 (prototype:模範蛇), weight is 100 ,level is 99
	snake2.Scramble() //I'm scrambling, my name is 模範蛇_clone (prototype:模範蛇), weight is 0 ,level is 99

}
