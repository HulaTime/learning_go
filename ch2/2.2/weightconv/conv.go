package weightconv

func KGToLB(kg Kilogram) Pounds {
	return Pounds(float64(kg) * LbInKg)
}

func KGToOz(kg Kilogram) Ounces {
	return Ounces(KGToLB(kg) * Pounds(OzInLb))
}

func LBToKG(lb Pounds) Kilogram {
	return Kilogram(float64(lb) / LbInKg)
}

func LbToOz(lb Pounds) Ounces {
	return Ounces(lb * Pounds(OzInLb))
}

func OzToKG(oz Ounces) Kilogram {
	return Kilogram(OzToLb(oz) / Pounds(LbInKg))
}

func OzToLb(oz Ounces) Pounds {
	return Pounds(oz / Ounces(OzInLb))
}
