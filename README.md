# gomedian
Generic interface for finding comedian tour dates.

## example

```Go
package main

import (
	"fmt"
)

func main() {
	joeRogan := &Comedian{
		Name: "Joe Rogan",
		Url:  "http://joerogan.net/tour/",
		XPath: &XPath{
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
}
```
