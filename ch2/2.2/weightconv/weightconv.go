// Package weightconv performs conversion between different weight measurements
package weightconv

import "fmt"

type Kilogram float64
type Pounds float64
type Ounces float64

const (
	LbInKg float64 = 2.204623
	OzInLb float64 = 16
)

func (kg Kilogram) String() string {
	return fmt.Sprintf("%.4g kg", kg)
}

func (lb Pounds) String() string {
	return fmt.Sprintf("%.4g lb", lb)
}

func (oz Ounces) String() string {
	return fmt.Sprintf("%.4g oz", oz)
}
