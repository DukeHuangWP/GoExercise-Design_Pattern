package factoryMethod

import "fmt"

//宣告接口(以動物為範例)，每個type-Method必須符合該接口宣告的格式
type AnimalInterface interface {
	SetName(name string) string //幫動物取名字
	Creep() string              //該動物作爬行行為
	Swimming() string           //該動物作游泳行為
}

//創建動物通用接口,並寫入錯誤提示
func NewAnimal(animalType int, animalName string) (AnimalInterface,error) {
	if animalName == "" {
		return nil,fmt.Errorf("Nameless!")
	}

	switch animalType {
	case 1:
		return &snake{Name:animalName},nil
	case 2:
		return &frog{Name:animalName},nil
	default:
		return nil,fmt.Errorf("%v is unknow type !", animalType)
	}
}

//每個蛇都可以取名子
type snake struct {
	Name string
}

//每個青蛙都可以取名子
type frog struct {
	Name string
}

//幫蛇取名
func (animal *snake) SetName(name string) string {
	animal.Name = name
	return fmt.Sprintf("Hi, I'm a snake and my name is %s", name)
}

//幫青蛙取名
func (animal *frog) SetName(name string) string {
	animal.Name = name
	return fmt.Sprintf("Hi, I'm a frog and my name is %s", name)
}

//蛇跟青蛙都會爬行
func (animal *snake) Creep() string {
	return fmt.Sprintf("Hi, I'm a snake named %s, now I'm creeping", animal.Name)
}

//蛇跟青蛙都會爬行
func (animal *frog) Creep() string {
	return fmt.Sprintf("Hi, I'm a frog named %s, now I'm creeping", animal.Name)
}

//蛇不會游泳
func (animal *snake) Swimming() string {
	return fmt.Sprintf("Hi, I'm a snake named %s, and I can't swimming", animal.Name)
}

//青蛙會游泳
func (animal *frog) Swimming() string {
	return fmt.Sprintf("Hi, I'm a frog named %s, now I'm swimming", animal.Name)
}
