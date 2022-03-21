package main

import "math"


func DistanceCalc (x1, y1, x2, y2 float64) float64{
	
		if y1 != y2 && x1 != x2 {

		x := x2-x1
		y := y2-y1

		x= x*x
		y= y*y

		sum := x+y
		floatSum := float64(sum)

		SumSQRT:= math.Sqrt(floatSum)

		return (SumSQRT)

	} else if y1 == y2 && x1 != x2{

		return x2-x1
	
	}else if y1 != y2 && x1 == x2{

		return y2-y1

	} else {

		return 0
	}
}