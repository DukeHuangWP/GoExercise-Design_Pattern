package builder

import (
	"fmt"
	"testing"
)

func TestBuilder1(t *testing.T) {
	var builder AnimalBuilder
	builder = &Snake{Name: "", Length: 0}
	snake := builder.BuildAnimal()
	fmt.Println("--------動物鑑定前--------")
	fmt.Println("動物名稱：", snake.Name)
	fmt.Println("動物身高：", snake.High)

	animal, err := Appraise(builder)
	if err == nil {
		appraised := animal.Build()
		fmt.Println("--------動物鑑定後--------")
		fmt.Println("動物名稱：", appraised.Name)
		fmt.Println("動物身高：", appraised.High)
	}
}