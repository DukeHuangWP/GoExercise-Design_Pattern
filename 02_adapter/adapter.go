package adapter

import "fmt"

//宣告接口(以說話介面為範例)，作為使用該介面的入口
type SayInterface interface {
	Say() string
}

//所有動物說話的時候都配有翻譯機
type Animals struct {
	TType   AnimalInterface
	Message string
}

//所有動物說話都有一個專屬翻譯機
type AnimalInterface interface {
	LangTrans(orgMsg string) (newMsg string)
}

//所有動物說話都要先判斷是否需要經過翻譯,在判斷使用哪種動物的翻譯機
func (animal *Animals) Say() (newMsg string) {
	if animal.TType == nil {
		return animal.Message
	} else {
		return "Trans: " + animal.TType.LangTrans(animal.Message)
	}
}

//每個蛇都可以取名子
type Snake struct {
	Name string
}

//每個蛇說話時都有統一的翻譯機
func (animal *Snake) LangTrans(orgMsg string) (newMsg string) {

	switch orgMsg {
	case "Hello World!":
		return fmt.Sprintf("Hi I'm %s, Ss Ss Ss ~", animal.Name)
	case "Bye!":
		return "Bye Ss Ss Ss ~"
	default:
		return "Ss Ss Ss ~"
	}

}

//每個青蛙都可以取名子
type Frog struct {
	Name string
}

//每個青蛙說話時都有統一的翻譯機
func (animal *Frog) LangTrans(orgMsg string) (newMsg string) {

	switch orgMsg {
	case "Hello World!":
		return fmt.Sprintf("Hi I'm %s, Ku Ku Ku ~", animal.Name)
	case "Bye!":
		return "Bye Ku Ku Ku ~"
	default:
		return "Ku Ku Ku ~"
	}

}
