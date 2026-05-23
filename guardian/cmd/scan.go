/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/bhavneet-singh/guardian/internal/scanner"
	"github.com/spf13/cobra"
)

var verbose bool
var jsonOutput bool
var packageFile string

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scans package.json",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

        pkg, err := scanner.ParsePackageJSON(args[0])
		
        if verbose {
			fmt.Println("verbose output enabled")

			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("Project:", pkg.Name)

			fmt.Println("\nDependencies")

			for name, version := range pkg.Name {
				fmt.Println(name, "=>", version)
			}
            return 
		}

		if jsonOutput {
			fmt.Println("json output enabled")
            output, err := json.MarshalIndent(pkg, " " , " ")
            if err != nil {
                fmt.Println(err)
                return
            }
            fmt.Println(string(output))
            return
		}

		if len(args) > 0 {
			fmt.Printf("Provide file path: %s\n", args[0])
		}


		if _, err := scanner.ParsePackageJSON(packageFile); err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)

	// Here you will define your flags and configuration settings.
	scanCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
	scanCmd.Flags().BoolVar(&jsonOutput, "json", false, "Enable json output")
	scanCmd.Flags().StringVarP(&packageFile, "file", "f", "package.json", "Path to package.json")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
