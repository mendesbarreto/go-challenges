package hackerearth

import (
	"github.com/mendesbarreto/go-challenges/pkg/utils"
	"os"
	"path"
	"testing"
)

func TestPrintResul(t *testing.T) {
	t.Run("PrintFavouriteSingerCountResult", func(t *testing.T) {

		currentFolderPath, _ := os.Getwd()
		file, err := os.Open(path.Join(currentFolderPath, "favourite_singles_play_list.text"))

		if err != nil {
			t.Error("Error open input file", err)
		}

		//catpuredOutput := utils.CaputreOuput( func () { PrintResult({1,2,3}) })
	})
}
