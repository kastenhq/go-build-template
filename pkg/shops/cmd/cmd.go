package cmd

import (
	"github.com/spf13/cobra"

	"github.com/thockin/go-build-template/pkg/shops"
)

func Root() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "shops",
		Short: "TODO",
		Long: `TODO
TODO
TODO`,
		SilenceUsage: false,
		RunE:         rootRun,
	}
	return rootCmd
}

func rootRun(cmd *cobra.Command, args []string) error {
	cmd.SilenceUsage = true
	return shops.Shops(cmd.Context())
}
