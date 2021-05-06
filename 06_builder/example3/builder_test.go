package builder

import (
	"testing"
)

func TestBuilder1(t *testing.T) {
	dr := AnimalBuilder{&Snake{}}
	adCar := dr.Animal.SetName("XXX").SetHight(9999).Build()
	adCar.Drive()

	// bwCar := dr.Builder.SetType("sporting").AddBrand("宝马").PaintColor("red").Build()
	// bwCar.Drive()

}
