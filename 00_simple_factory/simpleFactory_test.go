package simplefactory

import (
	"fmt"
	"testing"
)

//** 簡單工廠模式說明 **
//透過interface創建各自類別(struct)，做出不同行為(method)對應
//優點:結構非常簡單,易於批次複製與生產
//缺點:當行為(method)需要修改時，每個類別(struct)的行為全部都需要作修改
func TestSimpleFactory(t *testing.T) {
	api1 := NewAnimal(1)
	fmt.Println(api1.SetName("S1"))
	fmt.Println(api1.Creep())
	fmt.Println(api1.Swimming())

	api2 := NewAnimal(1)
	fmt.Println(api2.SetName("S2"))
	fmt.Println(api2.Creep())
	fmt.Println(api2.Swimming())

	api3 := NewAnimal(2)
	fmt.Println(api3.SetName("F0"))
	fmt.Println(api3.Creep())
	fmt.Println(api3.Swimming())

	// api0 := NewAnimal(0)
	// fmt.Println(api0.SetName(""))
	// 簡單工廠模式還需要針對未知類別(struct)處理
}
