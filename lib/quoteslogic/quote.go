package quoteslogic

import "fmt"

type Quote struct {
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

func NewQuote(quote, author string) Quote {
	return Quote{quote, author}
}

func (q Quote) String() string {
	return fmt.Sprintf("%s\n - %s", q.Quote, q.Author)
}
