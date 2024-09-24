package cmd

import (
	"github.com/fatih/color"

	"github.com/spf13/cobra"
)

var colorCmd = &cobra.Command{
	Use:   "color",
	Short: "A color package example",
	Run: func(cmd *cobra.Command, args []string) {
		color.HiRed("Red colored text")
		color.HiGreen("Green colored text")
		color.HiYellow("Yellow colored text")
		color.HiBlue("Blue colored text")
		color.HiMagenta("Magenta colored text")
		color.HiCyan("Cyan colored text")
		color.HiWhite("White colored text")
	},
}

func init() {
	rootCmd.AddCommand(colorCmd)
}
