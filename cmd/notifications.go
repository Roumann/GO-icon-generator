package cmd

import (
	"fmt"
	"image/color"
	"os"
	"rnGen/cmd/config"

	"github.com/disintegration/imaging"
	termClr "github.com/fatih/color"
	"github.com/spf13/cobra"
)

var cmdNotif = &cobra.Command{
	Use:   "notif",
	Short: "Generate notification icons",
	Long:  `Generate notification icons for Android notifications.`,
	Args:  cobra.MinimumNArgs(1),
	Run: notificationIcons,
}

func init(){
	rootCmd.AddCommand(cmdNotif)
	cmdNotif.Flags().Float32P("padding", "p", 0.75, "Padding for the icon")
}

func notificationIcons(cmd* cobra.Command, args []string) {
	green := termClr.New(termClr.FgGreen).SprintFunc()
	red := termClr.New(termClr.FgRed).SprintFunc()

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

	for _, size := range config.NotifSizes {
		dirPath := "android/app/src/main/res/drawable-" + size.Name
		filePath := dirPath + "/ic_stat_notification_icon.png"

		// Create a transparent background for the icon in the required size.
		background := imaging.New(
			int(float32(baseSize)*size.Scale),
			int(float32(baseSize)*size.Scale),
			color.NRGBA{0, 0, 0, 0})

		// Resize the icon to the required size minus padding
		iconResized := imaging.Resize(
			icon,
			int((float32(22)*size.Scale)*padding),
			int((float32(22)*size.Scale)*padding),
			imaging.Lanczos)

		// Combine the icon and the background
		finalImg := imaging.OverlayCenter(background, iconResized, 1)

		// Save the final image
		err = imaging.Save(finalImg, filePath)
		if err != nil {
			err := os.MkdirAll(dirPath, os.ModePerm)
			if err != nil {
				fmt.Println(red("Failed to create folder - "+size.Name, err))
				os.Exit(1)
			}
			err = imaging.Save(finalImg, filePath)
			if err != nil {
				fmt.Println(red("Failed to save file - "+size.Name, err))
				os.Exit(1)
			}
		}

		fmt.Println(green("✔"), size.Name)
	}

	fmt.Println(green("\n🎉 Notification icons generated successfully. "))
	fmt.Println(`
Paste this code into your android-manifest.xml:

<meta-data
  android:name="com.google.firebase.messaging.default_notification_icon"
  android:resource="@drawable/ic_stat_notification_icon" /> 
 `)
	
}

