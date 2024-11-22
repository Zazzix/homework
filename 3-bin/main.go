package main

import (
	"fmt"

	"github.com/bin/bins"
	"github.com/bin/storage"
)

func main() {
	binName := bins.GetBinName("Введите название Bin")

	binList := storage.GetBinList()

	err := bins.CreateNewBin(binName, binList)

	if err != nil {
		fmt.Println(err)
		return
	}

	storage.SaveBin(binList)
}
