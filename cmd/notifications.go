package cmd

import (
	"fmt"
	"image/color"
	"os"

	"github.com/disintegration/imaging"
	coloredTerm "github.com/fatih/color"
	"github.com/spf13/cobra"
)

type iconValues struct {
	name       string
	multiplier float32
}

var icons = []iconValues{
	{"mdpi", 1},
	{"hdpi", 1.5},
	{"xhdpi", 2},
	{"xxhdpi", 3},
	{"xxxhdpi", 4},
}

var cmdNotif = &cobra.Command{
	Use:   "notif",
	Short: "Generate notification icons",
	Long:  `Generate notification icons for Android notifications.`,
	Args:  cobra.MinimumNArgs(1),
	Run: notificationIcons,
}


func notificationIcons(cmd* cobra.Command, args []string) {
	green := coloredTerm.New(coloredTerm.FgGreen).SprintFunc()
	red := coloredTerm.New(coloredTerm.FgRed).SprintFunc()

	var baseSize int = 24
	filePath := args[0]
	padding, _ := cmd.Flags().GetFloat32("padding")

	if padding < 0.50 || padding > 1.50 {
		fmt.Println(red("Padding should be between 0.5 and 1.5"))
		os.Exit(1)
	}

	icon, err := imaging.Open(filePath)
	if err != nil {
		fmt.Println(red("Failed to open file:", err))
		os.Exit(1)
	}

	for _, iconConfig := range icons {
		dirPath := "android/app/src/main/res/drawable-" + iconConfig.name
		filePath := dirPath + "/ic_stat_notification_icon.png"

		// Create a transparent background for the icon in the required size.
		background := imaging.New(
			int(float32(baseSize)*iconConfig.multiplier),
			int(float32(baseSize)*iconConfig.multiplier),
			color.NRGBA{0, 0, 0, 0})

		// Resize the icon to the required size minus padding
		iconResized := imaging.Resize(
			icon,
			int((float32(22)*iconConfig.multiplier)*padding),
			int((float32(22)*iconConfig.multiplier)*padding),
			imaging.Lanczos)

		// Combine the icon and the background
		finalImg := imaging.OverlayCenter(background, iconResized, 1)

		// Save the final image
		err = imaging.Save(finalImg, filePath)
		if err != nil {
			err := os.MkdirAll(dirPath, os.ModePerm)
			if err != nil {
				fmt.Println(red("Failed to create folder - "+iconConfig.name, err))
				os.Exit(1)
			}
			err = imaging.Save(finalImg, filePath)
			if err != nil {
				fmt.Println(red("Failed to save file - "+iconConfig.name, err))
				os.Exit(1)
			}
		}

		fmt.Println(green("✔"), iconConfig.name)
	}

	fmt.Println(green("\n🎉 Notification icons generated successfully. "))
	fmt.Println(`
Paste this code into your android-manifest.xml:

<meta-data
  android:name="com.google.firebase.messaging.default_notification_icon"
  android:resource="@drawable/ic_stat_notification_icon" /> 
 `)
	
}

func init(){
	rootCmd.AddCommand(cmdNotif)

	cmdNotif.Flags().Float32P("padding", "p", 0.75, "Padding for the icon")
}
