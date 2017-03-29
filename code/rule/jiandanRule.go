package rule

import (
	"github.com/PuerkitoBio/goquery"
	"strconv"
)

// JandanRule is an implementation of Rule
type JandanRule struct{}

// NewJandanRule create a new JandanRule object
func NewJandanRule() Rule {
	return &JandanRule{}
}

// UrlRule implements the method in interface of Rule
func (p *JandanRule) URLRule() (url string) {
	return "http://jandan.net/ooxx/"
}

// PageRule implements the method in interface of Rule
func (p *JandanRule) PageRule(pageNumber int) (page string) {
	return "page-" + strconv.Itoa(pageNumber)
}

// ImageRule implements the method in interface of Rule
func (p *JandanRule) ImageRule(doc *goquery.Document, f func(image string)) {
	doc.Find("a.view_img_link").Each(func(i int, s *goquery.Selection) {
		if img, exist := s.Attr("href"); exist {
			f(img)
		}
	})
}
