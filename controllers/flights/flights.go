package flights

import (
	"log"
	"math"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// FlightList maps the json request structure: {"flights": [["SRC1", "DST1"],[SRC2,DST2]...]}
type FlightList struct {
	Flights [][]string `json:"flights"`
}

func SortFlights(c *fiber.Ctx) error {
	var list FlightList

	if err := c.BodyParser(&list); err != nil {
		log.Printf("Bad request: %v", err)
		return c.SendStatus(http.StatusBadRequest)
	}

	start, end := list.Sorted()
	return c.JSON([]string{start, end})
}

// Sorted sorts the flight records, returning the starting and ending airport.
func (f *FlightList) Sorted() (start, end string) {
	tracker := make(map[string]int)

	for _, flight := range f.Flights {
		source := flight[0]
		dest := flight[1]
		tracker[source] -= 1
		tracker[dest] += 1
	}

	min := math.MaxInt32 // keeps track of minimum tracking count on the tracker map
	max := math.MinInt32 // keeps track of maximum tracking count on the tracker map

	for k, v := range tracker {
		if v > max {
			end = k
			max = v
		}

		if v < min {
			start = k
			min = v
		}
	}

	return start, end
}
