package quoteslogic_test

import (
	"testing"

	"github.com/pulkit-tanwar/dispense-quotes/lib/quoteslogic"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDespenser(t *testing.T) {
	quotes := []quoteslogic.Quote{}

	quote1 := quoteslogic.Quote{"Of course, bad code can be cleaned up. But it’s very expensive.", "Robert C. Martin"}
	quote2 := quoteslogic.Quote{"Any fool can write code that a computer can understand. Good programmers write code that humans can understand.", "Martin Fowler"}
	quote3 := quoteslogic.Quote{"Programming is the art of telling another human being what one wants the computer to do.", "Donald Ervin Knuth"}

	quotes = append(quotes, quote1, quote2, quote3)

	md := quoteslogic.NewDespenser(quotes)

	require.NotNil(t, md)
	assert.Equal(t, 3, md.Count())
	assert.NotNil(t, md.RandomQuote())
	assert.Equal(t, "Of course, bad code can be cleaned up. But it’s very expensive.\n - Robert C. Martin", md.Get(0).String())
	assert.Equal(t, "Any fool can write code that a computer can understand. Good programmers write code that humans can understand.\n - Martin Fowler", md.Get(1).String())
	assert.Equal(t, "Programming is the art of telling another human being what one wants the computer to do.\n - Donald Ervin Knuth", md.Get(2).String())
}

func TestLoadQuotesFromJSONFile(t *testing.T) {
	filename := "../../quotes.json"

	d, err := quoteslogic.LoadQuotesFromJSONFile(filename)
	assert.Nil(t, err)

	assert.NotNil(t, d.RandomQuote())
}

func TestLoadQuotesFromInvalidFile(t *testing.T) {
	filename := "someInvalidFile"

	d, err := quoteslogic.LoadQuotesFromJSONFile(filename)
	assert.NotNil(t, err)
	assert.Nil(t, d)
	assert.Equal(t, "Error from ioutil.ReadFile: open someInvalidFile: no such file or directory", err.Error())
}

func TestLoadQuotesFromInvalidJSONFile(t *testing.T) {
	filename := "dispenser_test.go"

	d, err := quoteslogic.LoadQuotesFromJSONFile(filename)
	assert.NotNil(t, err)

	assert.Nil(t, d)
	assert.Equal(t, "Error from json.Unmarshal: invalid character 'p' looking for beginning of value", err.Error())
}

func TestDispenserRandomDistribution(t *testing.T) {
	q1 := quoteslogic.Quote{"Quote1", "Author1"}
	q2 := quoteslogic.Quote{"Quote2", "Author2"}

	quotesSlice := []quoteslogic.Quote{q1, q2}

	d := quoteslogic.NewDespenser(quotesSlice)

	require.NotNil(t, d)

	distribution := []int{0, 0}
	for i := 0; i < 100; i++ {
		q := d.RandomQuote()
		if q == q1 {
			distribution[0]++
		} else {
			distribution[1]++
		}
	}
	assert.True(t, distribution[0] > 40, "100 tests with 2 quotes should be close to 50 each, leaving a little wiggle room")
	assert.True(t, distribution[1] > 40, "100 tests with 2 quotes should be close to 50 each, leaving a little wiggle room")
}
