package prototype

import "fmt"

//蛇的專屬builder
type SnakeBuilder interface {
	Set(name string, length int) SnakeBuilder
	Clone() (*Snake, error)
}

//每個蛇都有名子跟長度
type Snake struct {
	ProtoName string
	Level     int

	Name   string
	Length int
}

//每個蛇都可以取名子
func (this *Snake) Set(name string, length int) SnakeBuilder {
	this.Name = name
	this.Length = length
	return this
}

//複製一條蛇的範本
func (this *Snake) Clone() (*Snake, error) {
	protoName := this.ProtoName
	if protoName == "" {
		return nil, fmt.Errorf("原型名稱不可為空")
	}

	protoName += "_clone"
	return &Snake{Name: protoName, ProtoName: protoName, Level: this.Level}, nil
}

//此條蛇開始爬行
func (this *Snake) Scramble() {
	fmt.Printf("I'm scrambling, my name is %v, length is %v ,level is %v \n", this.Name, this.Length, this.Level)
	return
}