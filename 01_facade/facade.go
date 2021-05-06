package facade

import "fmt"

//宣告動物行為接口，定義批次執行行為
type AnimalDoInterface interface {
	AllDoCreep() string    //所有動物作爬行行為
	AllDoSwimming() string //所有動物作游泳行為
}

//定義動物行為接口中的動物成員
type AnimalModule struct {
	snake SnakeModule //蛇模組
	frog  FrogModule  //青蛙模組
}

func (module *AnimalModule) AllDoCreep() string {
	snakes := module.snake.Creep()
	frogs := module.frog.Creep()
	return fmt.Sprintf("%s\n%s", snakes, frogs)
}

func (module *AnimalModule) AllDoSwimming() string {
	snakes := module.snake.Swimming()
	frogs := module.frog.Swimming()
	return fmt.Sprintf("%s\n%s", snakes, frogs)
}

//創建動物通用接口
func NewAnimal() AnimalDoInterface {
	return &AnimalModule{
		snake: &SnakeModuleType{Name: "蛇團體A"},
		frog:  &FrogModuleType{Name: "青蛙團體B"},
	}
}

//蛇模組統一接口
type SnakeModule interface {
	Creep() string    //該動物團作爬行行為
	Swimming() string //該動物團作游泳行為
	Nothing() string
}

type SnakeModuleType struct {
	Name string //蛇團體名稱
}

//蛇跟青蛙都會爬行
func (animal *SnakeModuleType) Creep() string {
	return fmt.Sprintf("Hi, we're named %s, now we're creeping", animal.Name)
}

//蛇不會游泳
func (animal *SnakeModuleType) Swimming() string {
	return fmt.Sprintf("Hi, we're named %s, now we can't swimming", animal.Name)
}

//蛇可以作其他動物無法做到的事情
func (_ *SnakeModuleType) Nothing() string {
	return fmt.Sprint("s s s s s s s !")
}

//青蛙模組統一接口
type FrogModule interface {
	Creep() string    //該動物團作爬行行為
	Swimming() string //該動物團作游泳行為
}

type FrogModuleType struct {
	Name string //青蛙團體名稱
}

//蛇跟青蛙都會爬行
func (animal *FrogModuleType) Creep() string {
	return fmt.Sprintf("Hi, we're named %s, now we're creeping", animal.Name)
}

//青蛙會游泳
func (animal *FrogModuleType) Swimming() string {
	return fmt.Sprintf("Hi, we're named %s, now we're swimming", animal.Name)
}
