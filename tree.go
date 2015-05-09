package main

import (
	"io/ioutil"
	"net/http"

	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/html"
)

func newTree(url string) (*html.HtmlDocument, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	doc, err := gokogiri.ParseHtml(content)
	if err != nil {
		return nil, err
	}

	return doc, nil
}
