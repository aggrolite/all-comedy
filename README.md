# gomedian
Generic interface for finding comedian tour dates.

## example

```Go
package main

import (
	"fmt"

	"github.com/aggrolite/gomedian"
)

func main() {

	// Where is Joe Rogan performing?
	joeRogan := &gomedian.Comedian{
		Name: "Joe Rogan",
		Url:  "http://joerogan.net/tour/",
		XPath: &gomedian.XPath{
			// XPath for individual event nodes (defaults to root of document)
			EventNodes: ".//div[@class='main']//div[@class='event-excerpt']",
			// XPath for event fields
			Date:      ".//div[@class='date']/h4",
			TicketUrl: ".//div[@class='details']//li[@class='tickets']/a/@href",
			VenueName: ".//div[@class='venue']/h3/a",
		},
	}

	joeRoganEvents, _ := joeRogan.GetEvents()
	for _, e := range *joeRoganEvents {
		fmt.Printf("Venue: %v\n", e.VenueName)
		fmt.Printf("Date: %v\n", e.Date)
		fmt.Printf("Tickets: %v\n\n", e.TicketUrl)
	}

	// What about Joey Diaz?
	joeyDiaz := &gomedian.Comedian{
		Name: "Joey Diaz",
		Url:  "http://joeydiaz.net/tour/",
		XPath: &gomedian.XPath{
			EventNodes: ".//div[@id='post-entries']//table[@class='gigpress-table upcoming']/tbody/tr",
			Date:       "./td[@class='gigpress-date']",
			TicketUrl:  ".//a[@class='gigpress-tickets-link']/@href",
			VenueName:  "./td[@class='gigpress-venue']",
		},
	}

	joeyDiazEvents, err := joeyDiaz.GetEvents()
	if err != nil {
		panic(err)
	}

	for _, e := range *joeyDiazEvents {
		fmt.Printf("Venue: %v\n", e.VenueName)
		fmt.Printf("Date: %v\n", e.Date)
		fmt.Printf("Tickets: %v\n\n", e.TicketUrl)
	}
}
```

## TODO

I'm thinking `gomedian` could support an extraction function which does not require XPath.
This would make the library more useful to users who are not familiar with XPath. :-)
