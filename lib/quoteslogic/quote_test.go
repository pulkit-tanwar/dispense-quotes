package quoteslogic_test

import (
	"encoding/json"
	"testing"

	. "github.com/pulkit-tanwar/dispense-quotes/lib/quoteslogic"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestQuote(t *testing.T) {
	q := NewQuote("Go is fun!", "Jack Black.")
	assert.Equal(t, "Go is fun!", q.Quote)
	assert.Equal(t, "Jack Black.", q.Author)
}

func TestString(t *testing.T) {
	q := NewQuote("Go is fun.", "Jack Black")
	assert.Equal(t, "Go is fun.\n - Jack Black", q.String())
}

func TestQuoteAsJson(t *testing.T) {
	q := NewQuote("Go is fun", "Jack Black")
	expected := `{ "quote": "Go is fun", "author": "Jack Black" }`

	bytes, err := json.Marshal(q)
	require.Nil(t, err)
	assert.JSONEq(t, expected, string(bytes))

}
