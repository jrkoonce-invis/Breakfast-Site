package data

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type item struct {
	ID          uint16  `json:"id"          validate:"-"`
	Name        string  `json:"name"        validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float32 `json:"price"       validate:"required,gt=0"`
}

func (item *item) validate() error {
	validate := validator.New()
	return validate.Struct(item)
}

var myItems = []*item{
	&item{1, "test", "test", 1.00},
}

// returns the highest ID value that exists + 1
func getNextID() uint16 {
	return myItems[len(myItems)-1].ID + 1
}

// This adjusts the IDs of the items that may have been affected by deletion
func adjustItems() {
	for i := 1; i < len(myItems)+1; i++ {
		myItems[i-1].ID = uint16(i)
	}
}

// GetData : Gets the existing data
func GetData(log *logrus.Logger) (string, error) {
	data, err := json.Marshal(myItems)
	return string(data), err
}

// PostData : Adds an item to the existing data
func PostData(given []byte) error {
	nextID := getNextID()

	var item = item{}
	err := json.Unmarshal(given, &item)
	if err != nil {
		return err
	}

	err2 := item.validate()
	if err2 != nil {
		return err2
	}

	item.ID = nextID
	myItems = append(myItems, &item)

	return nil
}

// ReplaceDatum : Replaces an item of the existing data
func ReplaceDatum(given []byte, id uint16) error {
	var item = item{}
	err := json.Unmarshal(given, &item)
	if err != nil {
		return err
	}

	err2 := item.validate()
	if err2 != nil {
		return err2
	}

	item.ID = id
	myItems[id-1] = &item

	return nil
}

// DeleteDatum : Deletes an item on the existing data
func DeleteDatum(id uint16) {
	myItems = append(myItems[:id-1], myItems[(id-1)+1:]...)
	adjustItems() // This adjusts the IDs of the items that may have been affected
}
