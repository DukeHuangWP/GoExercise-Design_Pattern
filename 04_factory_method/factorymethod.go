package factorymethod

import (
	"fmt"
)

type iGun interface {
	setPower(power int)
	getName() string
	getPower() int
}

type gun struct {
	name  string
	power int
}

func (g *gun) getName() string {
	return g.name
}

func (g *gun) setPower(power int) {
	g.power = power
}

func (g *gun) getPower() int {
	return g.power
}

type maverick struct {
	gun
}

type ak47 struct {
	gun
}

func NewGun(gunType string) (iGun, error) {
	switch gunType {

	case "ak47":
		return &ak47{
			gun: gun{
				name:  "AK47 gun",
				power: 4,
			},
		}, nil

	case "maverick 88":
		return &maverick{
			gun: gun{
				name:  "Maverick gun",
				power: 5,
			},
		}, nil

	default:
		return nil, fmt.Errorf("Wrong gun type passed")
	}
}
