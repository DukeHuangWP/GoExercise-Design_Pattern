package builder

import "fmt"

type AnimalType struct {
	Name string
	High int
}

//Builder 是生成器接口
type AnimalBuilder interface {
	setName() //幫動物取名字(注意繼承方式與工廠方法有所不同)
	setHigh()
	BuildAnimal() *AnimalType
}

type AnimalAppraiser struct {
	Animal AnimalBuilder
}

func (appraiser *AnimalAppraiser) Build() *AnimalType {
	appraiser.Animal.setName()
	appraiser.Animal.setHigh()
	return appraiser.Animal.BuildAnimal()
}

func Appraise(builder AnimalBuilder) (appraiser *AnimalAppraiser, err error) {
	if builder == nil {
		return nil, fmt.Errorf("Builder is nil")
	}

	appraiser = &AnimalAppraiser{Animal: builder}
	appraiser.Build()
	return appraiser, nil
}

type Snake struct {
	Name   string
	Length int
}

func (animal *Snake) setName() {
	if animal.Name == "" {
		animal.Name = "無名蛇"
	}
}

func (animal *Snake) setHigh() {
	if animal.Length <= 0 {
		animal.Length = 1
	}
}

func (animal *Snake) BuildAnimal() *AnimalType {
	return &AnimalType{
		Name: animal.Name,
		High: animal.Length,
	}
}

// func (animal *Snake) SaySanke() string {
// 	return fmt.Sprintf("Hi, I'm a snake and my name is %s", animal.Name)
// }

type Frog struct {
	Name   string
	Length int
}

func (animal *Frog) setName() {
	if animal.Name == "" {
		animal.Name = "無名青蛙"
	}
}

func (animal *Frog) setHigh() {
	if animal.Length <= 0 {
		animal.Length = 1
	}
}

func (animal *Frog) BuildAnimal() *AnimalType {
	return &AnimalType{
		Name: animal.Name,
		High: animal.Length,
	}
}

// func (animal *Frog) SayFrog() string {
// 	return fmt.Sprintf("Hi, I'm a frog and my name is %s", animal.Name)
// }
