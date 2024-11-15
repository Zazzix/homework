package main

import (
	"fmt"
	"github.com/bin/bins"
)

func main() {
	binName := bins.GetBinName("Введите название Bin")
	newBin, err := bins.CreateNewBin(binName)
	if err != nil {
		fmt.Println("Пустое название Bin")
		return
	}
	newList := bins.AddToBinList(newBin)
	newList.PrintBinList()
	newBin.PrintBin()
}
