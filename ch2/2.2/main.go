package main

import (
	"fmt"
	"os"
	"strconv"
	"unitconverter/weightconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		}

		fmt.Printf("%s = %s, %s\n",
			weightconv.Kilogram(t).String(),
			weightconv.KGToLB(weightconv.Kilogram(t)).String(),
			weightconv.KGToOz(weightconv.Kilogram(t)).String(),
		)
		fmt.Printf("%s = %s, %s\n",
			weightconv.Ounces(t).String(),
			weightconv.OzToKG(weightconv.Ounces(t)).String(),
			weightconv.OzToLb(weightconv.Ounces(t)).String(),
		)
		fmt.Printf("%s = %s, %s\n",
			weightconv.Pounds(t).String(),
			weightconv.LBToKG(weightconv.Pounds(t)).String(),
			weightconv.LbToOz(weightconv.Pounds(t)).String(),
		)
	}
}
