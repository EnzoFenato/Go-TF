package main

import "time"

func horariosSobrepostos(
	inicio1 string,
	fim1 string,
	inicio2 string,
	fim2 string,
) bool {

	i1, _ := time.Parse("15:04", inicio1)
	f1, _ := time.Parse("15:04", fim1)

	i2, _ := time.Parse("15:04", inicio2)
	f2, _ := time.Parse("15:04", fim2)

	return i1.Before(f2) &&
		f1.After(i2)
}
