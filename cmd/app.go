package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var cmdApp = &cobra.Command{
	Use:   "app [icon_path] [background_color] [icon_padding - Optional, between 0.5(50%) and 1.5(150%)]",
	Short: "Generate app icons",
	Long:  ` Generate app icons for Android & IOS app.`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Print: " + strings.Join(args, " "))
	},
}