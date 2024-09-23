package main

import (
	"fmt"
	generate "image-resizer/cmd"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	var cmdNotif = &cobra.Command{
		Use:   "notif [icon_path] [icon_padding - Optional, between 0.5(50%) and 1.5(150%)]",
		Short: "Generate notification icons",
		Long:  ` Generate notification icons for Android notifications.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var padding float32 = 0.75
			filePath := args[0]

			if len(args) == 2 {
				float64Val, err := strconv.ParseFloat(args[1], 32)
				if err != nil {
					fmt.Println("🔴 Failed to parse padding: Invalid number")
					os.Exit(1)
				}
				if float64Val < 0.50 || float64Val > 1.50 {
					fmt.Println("🔴 Padding should be between 0.5 and 1.5")
					os.Exit(1)
				}
				padding = float32(float64Val)
			}

			generate.NotificationIcons(filePath, padding)
		},
	}

	var cmdApp = &cobra.Command{
		Use:   "app [icon_path] [background_color] [icon_padding - Optional, between 0.5(50%) and 1.5(150%)]",
		Short: "Generate app icons",
		Long:  ` Generate app icons for Android & IOS app.`,
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Print: " + strings.Join(args, " "))
		},
	}

	var cmdSplash = &cobra.Command{
		Use:   "splash [logo_path] [background_color] [logo_size - Optional, between 0.5(50%) and 1.5(150%)]",
		Short: "Generate Splash screen",
		Long:  `Generate Splash screen for Android & IOS app.`,
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	var rootCmd = &cobra.Command{Use: "gen"}
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
	rootCmd.AddCommand(cmdNotif, cmdApp, cmdSplash)
	// Adds subcommands to the cmdEcho
	// cmdEcho.AddCommand(cmdTimes)
	rootCmd.Execute()
}
