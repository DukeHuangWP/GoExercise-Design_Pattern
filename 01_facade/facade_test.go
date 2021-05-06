package facade

import (
	"fmt"
	"testing"
)

//** 外觀模式說明 **
//透過interface接口統一所有子接口(或另稱Module)的行為，
//優點:使用上因具備統一接口使用極為簡易，適合批次性但細節繁瑣的工作處理，未來擴充僅需要擴充子接口(或另稱Module)即可
//缺點:由於行為(method)統一並作批次處理，難以針對子接口的個別屬性(struct-field)做出個別設定
func TestFacadeAPI(t *testing.T) {
	animals := NewAnimal()
	fmt.Println(animals.AllDoCreep())
	fmt.Println(animals.AllDoSwimming())

	// Hi, we're named 蛇團體A, now we're creeping
	// Hi, we're named 青蛙團體B, now we're creeping
	// Hi, we're named 蛇團體A, now we can't swimming
	// Hi, we're named 青蛙團體B, now we're swimming
}
