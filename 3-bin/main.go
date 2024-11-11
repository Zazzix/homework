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

func (binlist *BinList) addBin(bin *Bin) {
	binlist.Bins = append(binlist.Bins, *bin)

}

func createNewBin(BinName string) (*Bin, error) {
	if BinName == "" {
		return nil, errors.New("INVALID_BIN_NAME")
	}
	mockBin := &Bin{
		id:        "123",      // временно
		private:   false,      // временно
		createdAt: time.Now(), // временно
		name:      BinName,
	}

	return mockBin, nil
}

func main() {
	BinName := BinName("Введите название BIN")

	mockBin, err := createNewBin(BinName)
	if err != nil {
		fmt.Println("Нет имени BIN")
	}

	fmt.Print(mockBin)

}

func BinName(prompt string) string {
	var BinName string
	fmt.Print(prompt + ": ")
	fmt.Scanln(&BinName)
	return BinName
}
