package fda

import (
	"fmt"
	"io/ioutil"
	"net/http"
	// "net/url"

	// "strconv"
	// "strings"
	// "bytes"
	// "encoding/json"
)

// Request Struct
// =================
// search=patient.drug.openfda.pharm_class_epc:"nonsteroidal+anti-inflammatory+drug"
// count=patient.reaction.reactionmeddrapt.exact

type Series struct {
	Search, Count string
}

func (s *Series) Request() string {
	// u, err := url.Parse("https://api.fda.gov/drug/event.json")
	// check(err)

	// q := u.Query()
	// q.Set("Search", s.Search)
	// q.Set("Count", s.Count)
	// u.RawQuery = q.Encode()
	// fmt.Println("Query: ", u.String())
	// resp, err := http.Get(u.String())

	resp, err := http.Get("https://api.fda.gov/drug/event.json?search=patient.drug.openfda.pharm_class_epc:\"nonsteroidal+anti-inflammatory+drug\"&count=patient.reaction.reactionmeddrapt.exact")
	check(err)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	output := string(body)
	fmt.Println("\n", output)
	return output
}

func check(err error) {
	if err != nil {
		panic(err)
	}

}
