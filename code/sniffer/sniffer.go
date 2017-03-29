package sniffer

import (
	"Sniffer/code/base"
	"Sniffer/code/rule"
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// Sniffer is used to grab pictures from website
type Sniffer struct {
	path   string
	start  int
	end    int
	rule   rule.Rule
	client *http.Client
}

// New a Sniffer object
func New(path string, start int, end int, rule rule.Rule, client *http.Client) *Sniffer {
	return &Sniffer{path: path, start: start, end: end, rule: rule, client: client}
}

// Run Sniffer
func (s *Sniffer) Run() {
	// check storage directory
	if !strings.HasSuffix(s.path, "/") {
		s.path += "/"
	}
	if !base.IsDirExists(s.path) {
		os.MkdirAll(s.path, 0777)
		fmt.Println("create storage directory:\t" + s.path)
	} else {
		fmt.Println("the storage directory already exists:\t" + s.path)
	}

	// run
	for i := s.start; i <= s.end; i++ {
		res, err := s.getLinks(i)
		if err != nil {
			continue
		}
		for _, imageLink := range res {
			s.downloadImage(imageLink)
		}
	}
}

// Get links of pictures from page
func (s *Sniffer) getLinks(pageNumber int) ([]string, error) {
	var res []string
	doc, err := goquery.NewDocument(s.rule.URLRule() + s.rule.PageRule(pageNumber))
	if err != nil {
		return res, err
	}
	s.rule.ImageRule(doc, func(imageLink string) {
		if strings.HasPrefix(imageLink, "//") {
			res = append(res, "http:"+imageLink)
		} else {
			res = append(res, imageLink)
		}
	})
	return res, nil
}

// Download image
func (s *Sniffer) downloadImage(imageLink string) {
	fmt.Println("download picture:\t" + imageLink)
	arr := strings.Split(imageLink, "/")
	imageName := arr[len(arr)-1]
	if base.IsFileExists(s.path + imageName) {
		fmt.Println("pass:\tfile already exists")
		return
	}
	resp, err := s.client.Get(imageLink)
	if err != nil {
		fmt.Println("failed\twhen send request")
		return
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("failed:\twhen parse data")
		return
	}
	imageFile, err := os.OpenFile(s.path+imageName, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println("failed:\twhen create file")
		return
	}
	defer imageFile.Close()
	_, err = io.Copy(imageFile, bytes.NewReader(data))
	if err != nil {
		fmt.Println("failed:\twhen generate picture")
		return
	}
	fmt.Println("success")
}
