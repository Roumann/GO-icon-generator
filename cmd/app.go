package cmd

import (
	"fmt"
	"image"
	"image/color"
	"os"
	"strconv"
	"strings"

	"github.com/disintegration/imaging"
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
	filePath := args[0]
	bgColor := args[1]
	padding, _ := cmd.Flags().GetFloat32("padding")

	if padding < 0.50 || padding > 1.50 {
		fmt.Println("Padding should be between 0.5 and 1.5")
		os.Exit(1)
	}

	icon, err := imaging.Open(filePath)
	if err != nil {
		fmt.Println("Failed to open file:", err)
		os.Exit(1)
	}

	// for _, size := range config.AppSizes {
		
	// }

	err = genBackgroudXml(bgColor)
	if err != nil {
		fmt.Println("Error generating ic_launcher_background.xml file")
		os.Exit(1)
	}

	err = genAnyDipXml()
	if err != nil {
		fmt.Println("Error generating ic_launcher.xml/ic_launcher_round.xml files")
		os.Exit(1)
	}


	err = genPlayStoreIc(icon, bgColor)
	if err != nil {
		fmt.Println("Error generating Play store icon.")
		os.Exit(1)
	}
}

func genPlayStoreIc(icon image.Image, bgColor string) (err error){
	dir := "android/app/src/main"
	filePath := dir + "/ic_launcher_playstore.png"

	rgbaClr, err := hexToRGBA(bgColor)
	if err != nil{
		fmt.Println("Error Converting color to RGBA")
		os.Exit(1)
	}

	background := imaging.New(512, 512, rgbaClr)
	foreground := imaging.Resize(icon,335,335,imaging.Lanczos)
	final := imaging.OverlayCenter(background, foreground, 1)

	err = imaging.Save(final, filePath)
	if err != nil {
		err = os.MkdirAll(dir, 0777)
		if err != nil {
			return err
		}
		err = imaging.Save(final, filePath)
		if err != nil{
			return err
		}
	}
	
	return nil
}


func genBackgroudXml(bgColor string) (err error){
	dir := "android/app/src/main/res/values"
	filePath := dir + "/ic_launcher_background.xml"

	xml := fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
<resources>
    <color name="ic_launcher_background">%s</color>
</resources>`, bgColor)
 
	err = os.WriteFile(filePath, []byte(xml), 0777)
	if err != nil {
		err := os.MkdirAll(dir, 0777)
		if err != nil {
			return err
		}
		err = os.WriteFile(filePath, []byte(xml), 0777)
		if err != nil {
			return err
		}
	}
	return nil
}

func genAnyDipXml() (err error){
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
				return err
			}
			err = os.WriteFile(filePath, []byte(xml), 0777)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func hexToRGBA(s string) (clr color.Color, err error){
	// #ff 57 22
	// RR GG BB

	// RGBA STRUCTURE
	// RR GG BB AA=(255 solid color)

	//trim # if there is one
	hex := strings.TrimPrefix(s, "#")
	
	//Select character by groups of 2 for R G B channel
	// Slice the string, 16 = 0-9 A-F hexadecimal char. only, 8 = return type 0-255 number
	r, err := strconv.ParseUint(hex[0:2], 16, 8)
	if err != nil {
		return nil, err
	}

	g, err := strconv.ParseUint(hex[2:4], 16, 8)
	if err != nil {
		return nil, err
	}

	b, err := strconv.ParseUint(hex[4:6], 16, 8)
	if err != nil {
		return nil, err
	}

	//feed it to color.RGBA{} and set A to 255
	return color.RGBA{uint8(r), uint8(g), uint8(b), 255}, err

}