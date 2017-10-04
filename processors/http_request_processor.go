package processors

import (
	"fmt"
	"net/http"
	"scheduler/types"
	"strings"
	"time"
)

type HTTPProcessor struct {
	Name  string
	IsSSL bool
}

func (hp HTTPProcessor) Processing(sche types.Schedule) {
	fmt.Print("Executing schedule", sche, "...for ", hp.Name, strings.Repeat(".", 4))

	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}

	url := sche.URL
	if hp.IsSSL {
		url = strings.Replace(url, "http://", "https://", 1)
	}

	response, err := netClient.Get(url)
	if err != nil {
		fmt.Print("unknown response\n")
	} else {
		fmt.Print(response.StatusCode, "\n")
	}
}
