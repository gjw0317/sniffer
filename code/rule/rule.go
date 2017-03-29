package rule

import (
	"github.com/PuerkitoBio/goquery"
)

// Rule interface is used for custom extension functions
type Rule interface {
	URLRule() (url string)
	PageRule(pageNumber int) (page string)
	ImageRule(doc *goquery.Document, f func(image string))
}
