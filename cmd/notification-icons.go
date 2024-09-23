package notifications

import (
	"fmt"
	"image/color"
	"os"
	"strconv"

	"github.com/disintegration/imaging"
)

type iconValues struct {
	name       string
	multiplier float32
}

func GenerateIcons() {
	var baseSize int = 24
	var padding float32 = 0.75

	icons := []iconValues{
		{"mdpi", 1},
		{"hdpi", 1.5},
		{"xhdpi", 2},
		{"xxhdpi", 3},
		{"xxxhdpi", 4},
	}

	// Get the input arguments - icon path
	inputArgs := os.Args[1:]
	if len(inputArgs) != 1 && len(inputArgs) != 2 {
		fmt.Println("Usage: go run main.go <icon_path> <icon_padding - Optional, between 0.5(50%) and 1.5(150%)>")
		os.Exit(1)
	}

	// Open the icon file
	iconPath := inputArgs[0]
	icon, err := imaging.Open(iconPath)
	if err != nil {
		fmt.Println("🔴 Failed to open file:", err)
		os.Exit(1)
	}

	if len(inputArgs) == 2 {
		padding64, err := strconv.ParseFloat(inputArgs[1], 32)
		if err != nil {
			fmt.Println("🔴 Failed to parse padding: Invalid number")
			os.Exit(1)
		}
		if padding64 < 0.50 || padding64 > 1.50 {
			fmt.Println("🔴 Padding should be between 0.5 and 1.5")
			os.Exit(1)
		}
		padding = float32(padding64)
	}

	for _, iconConfig := range icons {
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
		err = imaging.Save(finalImg, "android/app/src/main/res/drawable-"+iconConfig.name+"/ic_stat_notification_icon.png")
		if err != nil {
			err := os.MkdirAll("android/app/src/main/res/drawable-"+iconConfig.name, os.ModePerm)
			if err != nil {
				fmt.Println("🔴 Failed to create folder - "+iconConfig.name, err)
				os.Exit(1)
			}
			err = imaging.Save(finalImg, "android/app/src/main/res/drawable-"+iconConfig.name+"/ic_stat_notification_icon.png")
			if err != nil {
				fmt.Println("🔴 Failed to save file - "+iconConfig.name, err)
				os.Exit(1)
			}
		}

		fmt.Println("🟢", iconConfig.name)
	}

	fmt.Println("\n🎉 Notification icons generated successfully. ")
	fmt.Println("\nPaste this code into your android-manifest.xml:")
	fmt.Println("\n<meta-data \n android:name=\"com.google.firebase.messaging.default_notification_icon\" \n android:resource=\"@drawable/ic_stat_notification_icon\" />")
	fmt.Println("")

}
