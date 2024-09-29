package main

import "rnGen/cmd"

func main() {
	cmd.Execute()

	
	// // Create a new RGBA image
	// m := image.NewRGBA(image.Rect(0, 0, 640, 480))

	// // Define blue color
	// blue := color.RGBA{0, 0, 255, 255}

	// // Draw the blue color onto the image
	// draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.Point{0, 0},  draw.Src)

	// // Create a new file to save the image
	// file, err := os.Create("blue_image.png")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	// // Encode the image as PNG and save it to the file
	// if err := png.Encode(file, m); err != nil {
	// 	panic(err)
	// }
}
