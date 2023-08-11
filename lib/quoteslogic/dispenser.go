package quoteslogic

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/rand"
)

type Dispenser interface {
	Get(n int) Quote
	Count() int
	RandomQuote() Quote
}

// NewDespenser - A constructor function for memory dispenser
func NewDespenser(quotes []Quote) Dispenser {
	md := &MemoryDispenser{quotes}
	return md
}

// ====================================================================================

// MemoryDispenser struct implements the Dispenser interface
type MemoryDispenser struct {
	quotes []Quote
}

func (md *MemoryDispenser) Get(n int) Quote {
	return md.quotes[n]
}

func (md *MemoryDispenser) Count() int {
	return len(md.quotes)
}

func (md *MemoryDispenser) RandomQuote() Quote {
	return md.quotes[rand.Intn(len(md.quotes))]
}

// ====================================================================================

// LoadQuotesFromJSONFile - This function will read the quotes from a json file
func LoadQuotesFromJSONFile(filename string) (Dispenser, error) {
	quotes := []Quote{}

	byteArray, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.New("Error from ioutil.ReadFile: " + err.Error())
	}

	err = json.Unmarshal(byteArray, &quotes)
	if err != nil {
		return nil, errors.New("Error from json.Unmarshal: " + err.Error())
	}

	return NewDespenser(quotes), nil
}

// ====================================================================================
