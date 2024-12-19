package cmd

import (
	"fmt"
	"time"

	"github.com/lossurdo/golang-starter-project/utils"
	"github.com/spf13/cobra"
)

var dateCmd = &cobra.Command{
	Use:   "date",
	Short: "A Date usage examples",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Date: %s\n", utils.GetDateTime(utils.DateTimeFormat{DDMMYYYY: true}, time.Now()))
		fmt.Printf("Date: %s\n", utils.GetDateTime(utils.DateTimeFormat{YYYYMMDD: true, AddTime: true}, time.Now()))
		fmt.Printf("Date: %s\n", utils.GetDateTime(utils.DateTimeFormat{DDMMYYYY: true, AddTimeAndSeconds: true}, time.Now()))
		fmt.Printf("Date: %s\n", utils.GetDateTime(utils.DateTimeFormat{DDMMYYYY: true, Dashed: true, AddTimeAndSeconds: true}, time.Now()))
		fmt.Printf("Date: %s\n", utils.GetDateTime(utils.DateTimeFormat{YYYYMMDD: true, Dashed: true, AddTimeAndSeconds: true}, time.Now()))

		fmt.Printf("Time: %s\n", utils.GetTime(utils.DateTimeFormat{AddTime: true}, time.Now()))
		fmt.Printf("Time: %s\n", utils.GetTime(utils.DateTimeFormat{AddTimeAndSeconds: true}, time.Now()))
	},
}

func init() {
	rootCmd.AddCommand(dateCmd)
}
