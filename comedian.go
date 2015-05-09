package main

import (
	"reflect"
	"strings"
)

type Event struct {
	Date      string
	TicketUrl string
	Time      string
	VenueAddr string
	VenueName string
}

type XPath struct {
	Date       string
	EventNodes string
	TicketUrl  string
	Time       string
	VenueName  string
	VenueAddr  string
}

type Comedian struct {
	Name  string
	Url   string
	XPath *XPath
}

func (c *Comedian) GetEvents() (*[]*Event, error) {

	tree, err := newTree(c.Url)
	if err != nil {
		return nil, err
	}

	eventContainerXps := c.XPath.EventNodes
	if eventContainerXps == "" {
		eventContainerXps = "."
	}
	eventNodes, err := tree.Root().Search(eventContainerXps)
	if err != nil {
		return nil, err
	}

	var events []*Event
	for _, node := range eventNodes {

		event := new(Event)

		// ihavenoideawhatimdoing.jpg
		eElem := reflect.ValueOf(event).Elem()
		xVal := reflect.ValueOf(*c.XPath)
		for i := 0; i < xVal.NumField(); i++ {
			field := xVal.Type().Field(i)

			// EventsNode is used for searching
			if field.Name == "EventNodes" {
				continue
			}

			// Skip empty XPath expressions
			xps := xVal.Field(i)
			if xps.String() == "" {
				continue
			}

			// Extract text content
			results, _ := node.Search(xps.String())
			if err != nil {
				continue // FIXME
			}

			// Create Event's type field
			if results != nil {
				// TODO handle missing content
				text := results[0].Content()
				textVal := reflect.ValueOf(strings.TrimSpace(text))
				eElem.FieldByName(field.Name).Set(textVal)
			}
		}
		events = append(events, event)
	}

	tree.Free()

	return &events, nil
}
