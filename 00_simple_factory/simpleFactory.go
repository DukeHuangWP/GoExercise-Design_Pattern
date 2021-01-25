package simplefactory

import "fmt"

//宣告接口(以動物為範例)
type AnimalInterface interface {
	SetName(name string) string
	Creep() string
	Swimming() string
}

//創建動物通用接口
func NewAnimal(animalType int) AnimalInterface {
	switch animalType {
	case 1:
		return &snake{}
	case 2:
		return &frog{}
	default:
		panic(fmt.Sprintf("%v is unknow type !\n", animalType))
	}
}

//每個蛇都可以取名子
type snake struct {
	name string
}

//每個青蛙都可以取名子
type frog struct {
	name string
}

//幫蛇取名
func (animal *snake) SetName(name string) string {
	animal.name = name
	return fmt.Sprintf("Hi, I'm a snake and my name is %s", name)
}

//幫青蛙取名
func (animal *frog) SetName(name string) string {
	animal.name = name
	return fmt.Sprintf("Hi, I'm a frog and my name is %s", name)
}

//蛇跟青蛙都會爬行
func (animal *snake) Creep() string {
	return fmt.Sprintf("Hi, I'm a snake named %s, now I'm creeping", animal.name)
}

//蛇跟青蛙都會爬行
func (animal *frog) Creep() string {
	return fmt.Sprintf("Hi, I'm a frog named %s, now I'm creeping", animal.name)
}

//蛇不會游泳
func (animal *snake) Swimming() string {
	return fmt.Sprintf("Hi, I'm a snake named %s, and I can't swimming", animal.name)
}

//青蛙會游泳
func (animal *frog) Swimming() string {
	return fmt.Sprintf("Hi, I'm a frog named %s, now I'm swimming", animal.name)
}
