package factorymethod

import (
	"fmt"
	"testing"
)

//** 工廠方法模式說明 **
//透過interface創建類別(struct)與定義行為(method)，每個類別依據需求設定不同屬性(struct-field)
//優點:結構非常明瞭，可以快速建立類別，並簡化定義出所有類別的行為
//缺點:由於行為(method)統一，定義屬性(struct-field)需要事先完善規劃
func TestFactoryMethod(t *testing.T) {
	ak47, _ := NewGun("ak47")
	maverick, _ := NewGun("maverick 88")

	fmt.Printf("Gun: %v\n", ak47.getName())
	fmt.Printf("Power: %v\n", ak47.getPower())
	fmt.Printf("Gun: %v\n", maverick.getName())
	fmt.Printf("Power: %v\n", maverick.getPower())
	maverick.setPower(9)
	fmt.Printf("Power: %v\n", maverick.getPower())

}
