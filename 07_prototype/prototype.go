package prototype

import "fmt"

//動物的原型接口,負責載入範本與複製動物
type AnimalPrototype interface {
	Set(name string, weight int) AnimalPrototype
	Clone() *Animal
}

//動物的共同屬性
type Animal struct {
	ProtoName string //原型範本名稱(方便追蹤範本來源)
	Level     int //該動物的等級
	Name   string //該動物的名字
	Weight int //該動物的重量
}

//每個複製動物都可以再取名字跟設定重量
func (this *Animal) Set(name string, weight int) AnimalPrototype {
	this.Name = name
	this.Weight = weight
	return this
}

//複製該動物
func (this *Animal) Clone() *Animal {
	protoName := this.ProtoName
	if protoName == "" {
		return nil
	}

	return &Animal{Name: protoName+"_clone", ProtoName: protoName, Level: this.Level}
}

//蛇是動物的一種
type Snake struct {
	Animal
}

//此條蛇開始爬行
func (this *Snake) Scramble() {
	fmt.Printf("I'm scrambling, my name is %v (prototype:%v), weight is %v ,level is %v \n", this.Name, this.ProtoName, this.Weight, this.Level)
	return
}

//青蛙是動物的一種
type Frog struct {
	Animal
}

//此青蛙開始游泳
func (this *Frog) Swimming() {
	fmt.Printf("I'm swimming, my name is %v (prototype:%v), weight is %v ,level is %v \n", this.Name, this.ProtoName, this.Weight, this.Level)
	return
}
