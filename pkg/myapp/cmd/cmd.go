package cmd

import (
	"github.com/spf13/cobra"
)

func Root() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "myapp",
		Short: "TODO",
		Long: `TODO
TODO
TODO`,
	}
	return rootCmd
}
