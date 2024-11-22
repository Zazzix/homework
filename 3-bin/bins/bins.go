package bins

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Bin struct {
	Id        string    `json:"Id"`
	Private   bool      `json:"Private"`
	CreatedAt time.Time `json:"CreatedAt"`
	Name      string    `json:"Name"`
}

type BinList struct {
	Bins []Bin `json:"Bins"`
}

func (b *BinList) PrintBinList() {
	fmt.Println(b.Bins)
}

func (b *Bin) PrintBin() {
	fmt.Println(b.Id, b.Private, b.CreatedAt, b.Name)
}

func CreateNewBin(binName string, binList *BinList) error {
	if binName == "" {
		return errors.New("INVALID_BIN_NAME")
	}
	bin := Bin{
		Id:        "123",
		Private:   false,
		CreatedAt: time.Now(),
		Name:      binName,
	}

	binList.AddToBinList(bin)

	return nil
		
}

// func AddToBinList(bin *Bin) *BinList {
// 	binlist := make([]Bin, 0, 1)
// 	binlist = append(binlist, *bin)

// 	newList := &BinList{}
// 	newList.Bins = binlist
// 	return newList

// }

func (binList *BinList) AddToBinList(bin Bin) {
	binList.Bins = append(binList.Bins, bin)
}

func GetBinName(prompt string) string {
	var binName string
	fmt.Print(prompt + ": ")
	fmt.Scanln(&binName)
	return binName
}

func (binlist *BinList) ToBytes() ([]byte, error) {
	file, err := json.Marshal(binlist)
	if err != nil {
		return nil, err
	}

	return file, nil
}
