package cmd

import (
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var goroutinesCmd = &cobra.Command{
	Use:   "goroutines",
	Short: "Example of goroutines",
	Run: func(cmd *cobra.Command, args []string) {

		chanFlag, _ := cmd.Flags().GetBool("chan")

		if chanFlag {
			ch := make(chan string)

			go func() {
				ch <- "Hello, World! Gouroutines using channel..."
			}()

			msg := <-ch
			log.Println(msg)
		} else {
			var wg sync.WaitGroup
			wg.Add(1)
			go simpleGoroutines(&wg)
			wg.Wait()
		}

	},
}

func simpleGoroutines(wg *sync.WaitGroup) {
	log.Println("Hello, World! Goroutines using WaitGroup...")
	wg.Done()
}

func init() {
	goroutinesCmd.Flags().BoolP("chan", "c", true, "Example of goroutines with channel or WaitGroup")
	rootCmd.AddCommand(goroutinesCmd)
}
