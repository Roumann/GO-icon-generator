package cmd

import "github.com/spf13/cobra"

var cmdSplash = &cobra.Command{
	Use:   "splash [logo_path] [background_color] [logo_size - Optional, between 0.5(50%) and 1.5(150%)]",
	Short: "Generate Splash screen",
	Long:  `Generate Splash screen for Android & IOS app.`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

	},
}