package app

import (
	"context"

	"git.woa.com/mojoma/sharpedge/pkg/sharpe"
	"github.com/spf13/cobra"
)

func NewServerCommand(ctx context.Context) *cobra.Command {
	command := &cobra.Command{
		Use:   "server",
		Short: "start the hippo API Server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return sharpe.NewServer().Run(ctx)
		},
	}
	return command
}
