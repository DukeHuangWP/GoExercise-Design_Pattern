package abstractFactory

import "fmt"

//[父接口]宣告接口(以動物為範例)，將個動物接口共同屬性如名字統一在type-Method當中處理
type AnimalInterface interface {
	SetName(name string) AnimalDoInterface
}

//所有動物都有一個專屬名字
type AnimalType struct {
	Name string
}

//幫動物取一個專屬名字
func (animal *AnimalType) SetName(name string) string {
	animal.Name = name
	return fmt.Sprintf("Hi, I'm a snake and my name is %s", name)
}

//創建動物通用接口,並寫入錯誤提示
func NewAnimal(animalType int, animalName string) (AnimalDoInterface, error) {
	if animalName == "" {
		return nil, fmt.Errorf("Nameless!")
	}

	switch animalType {
	case 1:
		return &snake{AnimalType{Name: animalName}}, nil
	case 2:
		return &frog{AnimalType{Name: animalName}}, nil
	default:
		return nil, fmt.Errorf("%v is unknow type !", animalType)
	}
}

//[子接口]宣告動物行為接口，每個type-Method必須符合該接口宣告的格式
type AnimalDoInterface interface {
	SetName(name string) string //幫動物取名字(注意繼承方式與工廠方法有所不同)
	Creep() string              //該動物作爬行行為
	Swimming() string           //該動物作游泳行為
}

//每個蛇都是動物的一種
type snake struct {
	AnimalType
}

//每個青蛙都是動物的一種
type frog struct {
	AnimalType
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
