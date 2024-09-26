package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cmdApp = &cobra.Command{
	Use:   "app [icon_path] [background_color] [icon_padding - Optional, between 0.5(50%) and 1.5(150%)]",
	Short: "Generate app icons",
	Long:  `Generate app icons for Android & IOS app.`,
	Args:  cobra.MinimumNArgs(2),
	Run: appIcons,
		
}

func init(){
	rootCmd.AddCommand(cmdApp)
	cmdApp.Flags().Float32P("padding", "p", 0.75, "Padding for the icon")
}


func appIcons(cmd *cobra.Command, args []string) {
	// var baseSize int = 48
	// filePath := args[0]
	bgColor := args[1]
	// padding, _ := cmd.Flags().GetFloat32("padding")

	// if padding < 0.50 || padding > 1.50 {
	// 	fmt.Println(red("Padding should be between 0.5 and 1.5"))
	// 	os.Exit(1)
	// }

	// icon, err := imaging.Open(filePath)
	// if err != nil {
	// 	fmt.Println(red("Failed to open file:", err))
	// 	os.Exit(1)
	// }

	// for _, size := range config.Sizes {
		
	// }

	createBackgroudXml(bgColor)
	createAnyDipXml()
}


func createBackgroudXml(bgColor string){
	dir := "android/app/src/main/res/values"
	filePath := dir + "/ic_launcher_background.xml"

	xml := fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
<resources>
    <color name="ic_launcher_background">%s</color>
</resources>`, bgColor)
 
	err := os.WriteFile(filePath, []byte(xml), 0777)
	if err != nil {
		err := os.MkdirAll(dir, 0777)
		if err != nil {
			fmt.Println("Failed to create folder - "+filePath, err)
		}
		err = os.WriteFile(filePath, []byte(xml), 0777)
		if err != nil {
			fmt.Println("Failed to save file - "+filePath, err)
		}
	}
}

func createAnyDipXml(){
	dir := "android/app/src/main/res/mipmap-anydpi-v26"
	filePaths := []string{
		dir + "/ic_launcher.xml",
		dir + "/ic_launcher_round.xml",
	}

	xml := `<?xml version="1.0" encoding="utf-8"?>
 <adaptive-icon xmlns:android="http://schemas.android.com/apk/res/android">
	<background android:drawable="@color/ic_launcher_background"/>
	<foreground android:drawable="@mipmap/ic_launcher_foreground"/>
</adaptive-icon>`

	for _, filePath := range filePaths {
		err := os.WriteFile(filePath, []byte(xml), 0777)
		if err != nil {
			err := os.MkdirAll(dir, 0777)
			if err != nil {
				fmt.Println("Failed to create folder - "+filePath, err)
			}
			err = os.WriteFile(filePath, []byte(xml), 0777)
			if err != nil {
				fmt.Println("Failed to save file - "+filePath, err)
			}
		}
	}

}