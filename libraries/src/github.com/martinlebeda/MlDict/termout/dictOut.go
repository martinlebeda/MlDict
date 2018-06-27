package termout

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/martinlebeda/MlDict/service"
)

var styleHeader = color.New(color.FgCyan).Add(color.Bold)
var styleKey = color.New(color.Bold)

func PrintDicts(dicts []string) {
	for _, dict := range dicts {
		fmt.Println(dict)
	}
	fmt.Println("")
	fmt.Println("number of dictionaries: ", len(dicts))

}

func PrintResult(result service.Data) {
	if len(result) == 0 {
		fmt.Println("Nenalezeno...")
	} else {
		for _, dict := range result {
			styleHeader.Println(" *", dict.Dict, "* ")

			for _, term := range dict.Terms {
				styleKey.Print(term.Term)
				fmt.Println(" - ", term.Explanation)
			}

			fmt.Println("")
		}
	}
}
