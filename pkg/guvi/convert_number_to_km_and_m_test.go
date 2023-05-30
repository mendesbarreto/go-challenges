package guvi

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/mendesbarreto/go-challenges/pkg/utils"
)

//You are given a number A in Kilometers. Convert this into B: Meters and C: Centi-Metres.
//
//
//Input Description:
//A number "A" representing some distance in kilometer is provided to you as the input.
//
//Output Description:
//Convert and print this value in meters and centimeters.
//
//Sample Input :
//2
//Sample Output :
//2000
//200000
func TestPrintResult(t *testing.T) {
	t.Run("PrintResult", func(t *testing.T) {
		capturedOutput := utils.CaputreOuput(func() { PrintResult(2) })

		currentFolderPath, _ := os.Getwd()
		file, err := os.Open(path.Join(currentFolderPath, "convert_number_to_km_and_m.text"))

		defer file.Close()

		if err != nil {
			t.Error("Error opening file", err)
		}

		contents, err := ioutil.ReadAll(file)

		if err != nil {
			t.Error("Error reading file", err)
		}

		if capturedOutput != string(contents) {
			t.Error("Expected: \n", string(contents), len(string(contents)), " got: \n", capturedOutput, len(string(contents)))
		}
	})
}

func TestConvertKilometersToCentimeters(t *testing.T) {
	t.Run("should convert 2 kilometers to 200000 centimeters", func(t *testing.T) {
		if ConvertKilometersToCentimeters(2) != 200000 {
			t.Error("Expected 200000, got ", ConvertKilometersToCentimeters(2))
		}
	})

	t.Run("should convert 3 kilometers to 300000 centimeters", func(t *testing.T) {
		if ConvertKilometersToCentimeters(3) != 300000 {
			t.Error("Expected 300000, got ", ConvertKilometersToCentimeters(3))
		}
	})
}

func TestConvertKilometersToMeters(t *testing.T) {
	t.Run("should convert 2 kilometers to 2000 meters", func(t *testing.T) {
		if ConvertKilometersToMeters(2) != 2000 {
			t.Error("Expected 2000, got ", ConvertKilometersToMeters(2))
		}
	})

	t.Run("should convert 2 kilometers to 2000 meters", func(t *testing.T) {
		if ConvertKilometersToMeters(3) != 3000 {
			t.Error("Expected 2000, got ", ConvertKilometersToMeters(2))
		}
	})
}
