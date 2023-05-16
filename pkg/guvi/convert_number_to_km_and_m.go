package guvi

import "fmt"

//You are given a number A in Kilometers. Convert this into B: Meters and C: Centi-Metres.
//Input Description:
//A number "A" representing some distance in kilometer is provided to you as the input.
//Output Description:
//Convert and print this value in meters and centimeters.
//

func ConvertKilometersToMeters(kilometers float32) float32 {
	return kilometers * 1000
}

func ConvertKilometersToCentimeters(kilometers float32) float32 {
	return kilometers * 100000
}

func PrintResult(kilometers float32) {
	fmt.Println(ConvertKilometersToMeters(kilometers))
	fmt.Println(ConvertKilometersToCentimeters(kilometers))
}
