package builder

import "fmt"

//导演着角色
type AnimalBuilder struct {
	Animal AnimalDirector
}

type AnimalDirector interface {
	SetName(name string) AnimalDirector
	SetHight(hight int) AnimalDirector
	Build() AnimalType
}

//每個蛇都是動物的一種
type Snake struct {
	Animal AnimalType
}

func (this *Snake) SetName(name string) AnimalDirector {
	this.Animal.Name = name
	return this
}

func (this *Snake) SetHight(hight int) AnimalDirector {
	this.Animal.Height = hight
	return this
}

func (this *Snake) Build() AnimalType {
	//return &CarType{}
	return this.Animal
}

//每個青蛙都是動物的一種
type Frog struct {
	Animal AnimalType
}

func (this *Frog) SetName(name string) AnimalDirector {
	this.Animal.Name = name
	return this
}

func (this *Frog) SetHight(hight int) AnimalDirector {
	this.Animal.Height = hight
	return this
}

func (this *Frog) Build() AnimalType {
	//return &CarType{}
	return this.Animal
}

//所有動物都有一個專屬名字跟身高
type AnimalType struct {
	Name string
	Height int
}

func (this *AnimalType) Drive() error {
	fmt.Printf("A %s %v\n", this.Name, this.Height)
	return nil
}
