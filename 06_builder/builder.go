package builder

import "fmt"

//Director導演者,負責統籌所有builder
type Director struct {
	Snake SnakeBuilder
	Frog  FrogBuilder
}

//蛇的專屬builder
type SnakeBuilder interface {
	SetName(name string) SnakeBuilder
	SetLength(Length int) SnakeBuilder
	Build() *Snake
}

//每個蛇都有名子跟長度
type Snake struct {
	Name   string
	Length int
}

//每個蛇都可以取名子
func (this *Snake) SetName(name string) SnakeBuilder {
	this.Name = name
	return this
}

//每個蛇都可以量長度
func (this *Snake) SetLength(Length int) SnakeBuilder {
	this.Length = Length
	return this
}

//創建一條蛇
func (this *Snake) Build() *Snake {
	return this
}

//此條蛇開始爬行
func (this *Snake) Scramble() {
	fmt.Printf("I'm scrambling, my name is %v and %v cm\n", this.Name, this.Length)
	return
}

//青蛙的專屬builder
type FrogBuilder interface {
	SetName(name string) FrogBuilder
	SetHight(hight int) FrogBuilder
	Build() *Frog
}

//每個青蛙都有名子跟身高
type Frog struct {
	Name   string
	Height int
}

//每個青蛙都可取名字
func (this *Frog) SetName(name string) FrogBuilder {
	this.Name = name
	return this
}

//每個青蛙都量高度
func (this *Frog) SetHight(hight int) FrogBuilder {
	this.Height = hight
	return this
}

//創建一個青蛙
func (this *Frog) Build() *Frog {
	return this
}

//此青蛙開始游泳
func (this *Frog) Swimming() {
	fmt.Printf("I'm swimming, my name is %v and %v cm\n", this.Name, this.Height)
	return
}
