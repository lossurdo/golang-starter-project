package cmd

import (
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/levigross/grequests"
	"github.com/spf13/cobra"
)

type Holiday struct {
	Date string `json:"date"`
	Name string `json:"name"`
	Type string `json:"type"`
}

var requestCmd = &cobra.Command{
	Use:   "request",
	Short: "Example of 'grequests' usage",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Calling API...")

		isPost, _ := cmd.Flags().GetBool("post")

		if isPost {
			postRequest()
		} else {
			getRequest()
		}
	},
}

func postRequest() {
	log.Info("POST request")

	url := os.Getenv("URL_TWO")
	urlPath := "/post"

	payload := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}
	options := &grequests.RequestOptions{
		Headers: headers,
		JSON:    payload,
	}

	resp, err := grequests.Post(url+urlPath, options)
	if err != nil {
		panic(err)
	}

	log.Infof("Status Code: %v\n", resp.StatusCode)
	log.Infof("Body: %v\n", resp.String())
}

func getRequest() {
	log.Info("GET request")

	url := os.Getenv("URL_ONE")
	year := time.Now().Year()
	urlPath := fmt.Sprintf("/api/feriados/v1/%v", year)

	headers := map[string]string{
		"Content-Type": "application/json",
	}
	options := &grequests.RequestOptions{
		Headers: headers,
	}

	resp, err := grequests.Get(url+urlPath, options)
	if err != nil {
		panic(err)
	}

	var holidays []Holiday
	if err := resp.JSON(&holidays); err != nil {
		panic(err)
	}

	for _, holiday := range holidays {
		log.Infof("Date: %v\tName: %v\n", holiday.Date, holiday.Name)
	}
}

func init() {
	requestCmd.Flags().Bool("post", false, "POST request (default: GET)")
	rootCmd.AddCommand(requestCmd)
}
