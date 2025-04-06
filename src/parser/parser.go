package peParser

import (
	"html"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type HTMLParser struct{}

type WordPair struct {
	Word              string
	Translation       string
	Phrase            string
	PhraseTranslation string
}

func NewHTMLParser() *HTMLParser {
	return &HTMLParser{}
}

func (p *HTMLParser) ParseDictionaryPage(htmlContent string) ([]WordPair, error) {
	unescapedHTML := html.UnescapeString(htmlContent)
	// Wrap the unescaped HTML in a table to make it easier to parse. By default we don't have a table and td/tr tags inside the HTML, and goquery changes tr/rd tags to divs
	wrappedHTML := `<html><body><table class="dict__video__list-table">` + unescapedHTML + `</table></body></html>`
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(wrappedHTML))
	if err != nil {
		return nil, err
	}

	var words []WordPair

	// Try both selectors
	doc.Find("tr.dict__video__list-table__row").Each(func(i int, s *goquery.Selection) {
		word := strings.TrimSpace(s.Find("td.dict__video__list-table__col div.dict__video__list-table__word__main span.word-wrapper").Text())
		translation := strings.TrimSpace(s.Find("td.dict__video__list-table__col div.dict__video__list-table__word__translate").Text())
		phrase := strings.TrimSpace(s.Find("td.dict__video__list-table__col .dict__video__list-table__phrase__eng").Text())
		phraseTranslation := strings.TrimSpace(s.Find("td.dict__video__list-table__col .dict__video__list-table__phrase__rus").Text())

		if word != "" && translation != "" {
			words = append(words, WordPair{
				Word:              word,
				Translation:       translation,
				Phrase:            phrase,
				PhraseTranslation: phraseTranslation,
			})
		}
	})

	return words, nil
}
