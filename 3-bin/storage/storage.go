package storage

import (
	"encoding/json"
	"fmt"

	"github.com/bin/bins"
	"github.com/bin/file"
)

func SaveBin(binlist *bins.BinList) {
	data, err := binlist.ToBytes()
	if err != nil {
		fmt.Println(err)
		return
	}
	file.WriteFile(data, "binlist.json")
}

func GetBinList() *bins.BinList {
	data, err := file.ReadFile("binlist.json")
	newBinList :=
		&bins.BinList{
			Bins: []bins.Bin{},
		}
	if err != nil {
		fmt.Println(err)
		return newBinList
	}
	var binList bins.BinList
	err = json.Unmarshal(data, &binList)
	if err != nil {
		fmt.Println(err)
		return newBinList
	}

	return &binList
}