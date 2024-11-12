package main

import (
	"errors"
	"fmt"
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

type BinList struct {
	Bins []Bin
}

func (b *BinList) printBinList() {
	fmt.Println(b.Bins)
}

func (b *Bin) printBin() {
	fmt.Println(b.id, b.private, b.createdAt, b.name)
}

func main() {
	binName := getBinName("Введите название Bin")
	newBin, err := createNewBin(binName)
	if err != nil {
		fmt.Println("Пустое название Bin")
		return
	}
	newList := addToBinList(newBin)
	newList.printBinList()
	newBin.printBin()
}

func createNewBin(binName string) (*Bin, error) {
	if binName == "" {
		return nil, errors.New("INVALID_BIN_NAME")
	}
	return &Bin{
		id:        "123",
		private:   false,
		createdAt: time.Now(),
		name:      binName,
	}, nil

}

func addToBinList(bin *Bin) *BinList {
	binlist := make([]Bin, 0, 1)
	binlist = append(binlist, *bin)

	newList := &BinList{}
	newList.Bins = binlist
	return newList

}

func getBinName(prompt string) string {
	var binName string
	fmt.Print(prompt + ": ")
	fmt.Scanln(&binName)
	return binName
}
