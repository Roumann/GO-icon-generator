package cmd

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"os"
	"rnGen/cmd/config"
	"strconv"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
)

var cmdApp = &cobra.Command{
	Use:   "app [icon_path] [background_color] [icon_padding - Optional, between 0.2(20%) and 2(200%)]",
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
	var baseSize float32 = 48
	filePath := args[0]
	bgColor := args[1]
	padding, _ := cmd.Flags().GetFloat32("padding")

	if padding < 0.20 || padding > 2 {
		fmt.Println("Padding should be between 0.2 and 2")
		os.Exit(1)
	}

	rgbaClr, err := hexToRGBA(bgColor)
	if err != nil{
		fmt.Println("Error Converting color to RGBA")
		os.Exit(1)
	}

	icon, err := imaging.Open(filePath)
	if err != nil {
		fmt.Println("Failed to open file:", err)
		os.Exit(1)
	}


	for _, size := range config.AppSizes {
		dir := "android/app/src/main/res/"+size.Name

		//Create transparent backgrounds
		bgSquare := imaging.New(int(baseSize*size.Scale), int(baseSize*size.Scale), color.NRGBA{0, 0, 0, 0})
		bgCircle := makeCircleSmooth(bgSquare, float64(padding))
		bgForeground := imaging.New(int(108 * size.Scale), int(108 * size.Scale), color.NRGBA{0, 0, 0, 0})

		//Create color background for Icon
		midSquare := imaging.New(int((baseSize*size.Scale)-size.MidPadding), int((baseSize*size.Scale)-size.MidPadding), rgbaClr)
		midCircle := makeCircleSmooth(midSquare, 2)

		// create foreground Icon
		foreground := imaging.Resize(
			icon,
			int((baseSize*size.Scale-size.MidPadding)*padding),
			int((baseSize*size.Scale-size.MidPadding)*padding),
			imaging.Lanczos)
		
		finalBgSquare := imaging.OverlayCenter(midSquare, foreground, 1)
		finalBgSquare = imaging.OverlayCenter(bgSquare, finalBgSquare,  1)
		err = imaging.Save(finalBgSquare, dir + "/ic_launcher.png")
		if err != nil {
			err = os.MkdirAll(dir, 0777)
			if err != nil {
				fmt.Println("Error Creating directory", err)
				os.Exit(1)
			}
			err = imaging.Save(finalBgSquare, dir + "/ic_launcher.png")
			if err != nil{
				fmt.Println("Error saving square icon", err)
				os.Exit(1)
			}
		}

		bgCircle = imaging.OverlayCenter(bgCircle, midCircle, 1)
		bgCircle = imaging.OverlayCenter(bgCircle, foreground, 1)
		err = imaging.Save(bgCircle, dir + "/ic_launcher_round.png")
		if err != nil {
			fmt.Println("Error saving round icon", err)
			os.Exit(1)	
		}

		bgForeground = imaging.OverlayCenter(bgForeground, foreground, 1)
		err = imaging.Save(bgForeground, dir +  "/ic_launcher_foreground.png")
		if err != nil {
			fmt.Println("Error saving foreground icon", err)
			os.Exit(1)	
		}

		fmt.Println("✔", size.Name)
		
	}

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

	err = genPlayStoreIc(icon, rgbaClr, padding)
	if err != nil {
		fmt.Println("Error generating Play store icon.")
		os.Exit(1)
	}
}

func genPlayStoreIc(icon image.Image, rgbaClr color.Color, padding float32) (err error){
	dir := "android/app/src/main"
	filePath := dir + "/ic_launcher_playstore.png"

	background := imaging.New(512, 512, rgbaClr)
	// Resize the icon to the required size minus padding
	foreground := imaging.Resize(
		icon,
		int(512*padding),
		int(512*padding),
		imaging.Lanczos)

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

func makeCircleSmooth(src image.Image, factor float64) image.Image {
	d := src.Bounds().Dx()
	if src.Bounds().Dy() < d {
		d = src.Bounds().Dy()
	}
	dst := imaging.CropCenter(src, d, d)
	r := float64(d) / 2
	center := r - 0.5
	for x := 0; x < d; x++ {
		for y := 0; y < d; y++ {
			xf := float64(x)
			yf := float64(y)
			delta := math.Sqrt((xf-center)*(xf-center)+(yf-center)*(yf-center)) + factor - r
			switch {
			case delta > factor:
				dst.SetNRGBA(x, y, color.NRGBA{0, 0, 0, 0})
			case delta > 0:
				m := 1 - delta/factor
				c := dst.NRGBAAt(x, y)
				c.A = uint8(float64(c.A) * m)
				dst.SetNRGBA(x, y, c)
			}
		}
	}
	return dst
}
