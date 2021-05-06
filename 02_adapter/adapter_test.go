package adapter

import (
	"fmt"
	"testing"
)

//** 轉接器模式說明 **
//創建類別(struct)後，在該類別要做出行為(method)前先判斷是否需要經過轉換
//優點:簡單統一了行為(method)，並針對不需經過轉換接口例外者可自定義(method)
//缺點:只能針對已知者行為做出嚴謹的行為(method)規範，不需經過轉換接口的例外者難以正確規範

type computer struct {
	Name    string
	Message string
}

func (com *computer) Say() string {
	return fmt.Sprintf("我叫%v不是動物不需要翻譯,我想說: %v", com.Name, com.Message)
}

func TestAdapter(t *testing.T) {

	var adapter SayInterface
	msg := "Hello World!"

	adapter = &Animals{TType: &Snake{Name: "蛇一號"}, Message: msg}
	fmt.Println(adapter.Say())

	adapter = &Animals{TType: &Snake{Name: "蛇二號"}, Message: msg}
	fmt.Println(adapter.Say())

	adapter = &Animals{TType: &Frog{Name: "青蛙零號"}, Message: msg}
	fmt.Println(adapter.Say())

	adapter = &Animals{TType: nil, Message: msg}
	fmt.Println(adapter.Say())

	adapter = &computer{Name: "AI一號", Message: msg}
	fmt.Println(adapter.Say())

	// Trans: Hi I'm 蛇一號, Ss Ss Ss ~
	// Trans: Hi I'm 蛇二號, Ss Ss Ss ~
	// Trans: Hi I'm 青蛙零號, Ku Ku Ku ~
	// Hello World!
	// 我叫AI一號不是動物不需要翻譯,我想說: Hello World!

	//如果轉接器未正確規範容易發生邏輯錯誤,舉例如下：
	//adapter = nil
	//fmt.Println(adapter.Say()) //panic: runtime error: invalid memory address or nil pointer dereference

}
