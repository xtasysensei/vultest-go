/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"sync"

	"github.com/spf13/cobra"
	"github.com/xtasysensei/vultest/cmd/handlers"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Initializes the type of scan to be run",
	Long: `This colect a url as parameter,determines and execute the
type of vulnerabulity to be scanned for.
For example:

vultest scan --type xss --url <URL>`,
	Run: scanURL,
}

func init() {
	rootCmd.AddCommand(scanCmd)

	// scanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	scanCmd.Flags().StringP("type", "t", "", "type of vulnerability to be scanned for")
	scanCmd.Flags().StringP("url", "u", "", "url to be scanned for possible vulnerability")
	scanCmd.Flags().IntP("depth", "d", 0, "depth of scan")
}

func scanURL(cmd *cobra.Command, args []string) {

	scanType, err := cmd.Flags().GetString("type")
	cobra.CheckErr(err)
	if scanType == "" {
		emptyScan := "No Scan type(xss, sqli) was provided"
		cobra.CheckErr(emptyScan)
	}

	targetUrl, err := cmd.Flags().GetString("url")
	cobra.CheckErr(err)
	if targetUrl == "" {
		emptyUrl := "No Url was provided"
		cobra.CheckErr(emptyUrl)

	}

	scanDepth, err := cmd.Flags().GetInt("depth")
	cobra.CheckErr(err)

	var wg sync.WaitGroup
	var mu sync.Mutex

	userAgentList := []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:90.0) Gecko/20100101 Firefox/90.0",
		// Add more user agents as needed
	}
	switch scanType {
	case "xss":
		fmt.Println("+--------------------------+")
		fmt.Println("| Xss Scanner")
		fmt.Println("+--------------------------+")

		wg.Add(1)
		defer wg.Done()
		handlers.XSSCrawler(targetUrl, scanDepth, &wg, &mu, userAgentList)
		wg.Wait()
	case "sqli":
		fmt.Println("+--------------------------+")
		fmt.Println("| SQLi Scanner")
		fmt.Println("+--------------------------+")

	default:
		invalidErr := `Invalid argument. use '--help' flag to check usage info`
		cobra.CheckErr(invalidErr)
	}
}
