package bins

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

func (b *BinList) PrintBinList() {
	fmt.Println(b.Bins)
}

func (b *Bin) PrintBin() {
	fmt.Println(b.id, b.private, b.createdAt, b.name)
}

func CreateNewBin(binName string) (*Bin, error) {
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

func AddToBinList(bin *Bin) *BinList {
	binlist := make([]Bin, 0, 1)
	binlist = append(binlist, *bin)

	newList := &BinList{}
	newList.Bins = binlist
	return newList

}

func GetBinName(prompt string) string {
	var binName string
	fmt.Print(prompt + ": ")
	fmt.Scanln(&binName)
	return binName
}
