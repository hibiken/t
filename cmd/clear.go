package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Deletes all todos",
	Args:  cobra.NoArgs,
	Run: func(_ *cobra.Command, _ []string) {
		const prompt = "Are you sure that you want to delete all todos? [Yes/No] (default \"No\") "
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print(prompt)
		for scanner.Scan() {
			answer := scanner.Text()
			if contains([]string{"Yes", "yes", "y"}, answer) {
				break
			}
			if contains([]string{"No", "no", "n", ""}, answer) {
				return
			}
			fmt.Print("\n" + prompt)
		}
		if err := writeToFile([]*Todo{}, filepath); err != nil {
			printErrorAndExit(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
}
