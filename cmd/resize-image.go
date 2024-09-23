package generate

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"path/filepath"

// 	"github.com/disintegration/imaging"
// 	"github.com/manifoldco/promptui"
// )

// // GetFilesFromDir retrieves files from the current directory
// func GetFilesFromDir() ([]string, error) {
//     var files []string
//     err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
//         if err != nil {
//             return err
//         }
//         if !info.IsDir() && (filepath.Ext(path) == ".jpg" || filepath.Ext(path) == ".png") {
//             files = append(files, path)
//         }
//         return nil
//     })
//     return files, err
// }

// // ResizeImage keeps aspect ratio while resizing to fit within 200x200 pixels
// func ResizeImage(inputPath, outputPath string) error {
//     src, err := imaging.Open(inputPath)
//     if err != nil {
//         return fmt.Errorf("failed to open image: %v", err)
//     }

//     // Resize while keeping the aspect ratio
//     resized := imaging.Fit(src, 200, 200, imaging.Lanczos)

//     err = imaging.Save(resized, outputPath)
//     if err != nil {
//         return fmt.Errorf("failed to save image: %v", err)
//     }

//     return nil
// }

// func main() {
//     // Get list of images in current directory
//     files, err := GetFilesFromDir()
//     if err != nil {
//         log.Fatalf("Failed to retrieve files: %v", err)
//     }

//     // If no image files are found, exit
//     if len(files) == 0 {
//         log.Println("No image files found in the current directory.")
//         return
//     }

//     // Prompt user to select a file
//     prompt := promptui.Select{
//         Label: "Select Image to Resize",
//         Items: files,
//     }

//     _, selectedFile, err := prompt.Run()
//     if err != nil {
//         log.Fatalf("Prompt failed: %v", err)
//     }

//     // Resize the selected image
//     outputPath := "resized_" + filepath.Base(selectedFile)
//     err = ResizeImage(selectedFile, outputPath)
//     if err != nil {
//         log.Fatalf("Error resizing image: %v", err)
//     }

//     fmt.Printf("Image resized successfully. Saved as %s\n", outputPath)
// }

