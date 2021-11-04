package cmd

import (
    "github.com/spf13/cobra"
)

/**
 * @author Rancho
 * @date 2021/11/4
 */

var rootCmd = &cobra.Command{}

func init() {
    rootCmd.AddCommand(wordCmd)
}

func Execute() error {
    return rootCmd.Execute()
}