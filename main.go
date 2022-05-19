package main

import (
	"fmt"
	"image/color"
	"image/png"
	"os"
	"sort"
)

// fatalf panic with the formatted message.
func fatalf(str string, args ...interface{}) {
	panic(fmt.Sprintf(str, args...))
}

// Unicode for various gradient of colors.
const empty = "  "
const light = "░░"  // U+2591
const medium = "▒▒" // U+2592
const dark = "▓▓"   // U+2593
const full = "██"   // U+2588
var invertedList = []string{empty, full, dark, medium, light}
var normalList = []string{empty, light, medium, dark, full}

// SortableColor implements the sort.Interface for colors, sorting on the
// light value of colors.
type SortableColor []color.Color

func (s SortableColor) Len() int {
	return len(s)
}

func (s SortableColor) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// toGray convert a given color to its light value.
func toGray(c color.Color) float32 {
	r, g, b, a := color.NRGBAModel.Convert(c).RGBA()

	fr := float32(r) * 0.3
	fg := float32(g) * 0.59
	fb := float32(b) * 0.11
	fa := float32(a) / float32(0xff)

	return (fr + fg + fb) * fa
}

func (s SortableColor) Less(i, j int) bool {
	return toGray(s[i]) < toGray(s[j])
}

func main() {
	usage := `Usage: %s [--inverted|-i] <image>

Convert a pixel art image to unicode characters. The image must be a PNG
file with transparency, and it must have at most %d colors.

By default, the unicode characters are sorted in a way better fitting for
dark-themed terminals. To invert the colors, use --inverted.
`
	charList := normalList

	if argc := len(os.Args); argc != 2 && argc != 3 {

		fmt.Printf(usage, os.Args[0], len(charList))
		return
	} else if val := os.Args[1]; argc == 3 && (val == "--inverted" || val == "-i") {
		charList = invertedList
	} else if argc == 3 {
		fmt.Printf("Invalid argument '%s'! Expected --inverted (or nothing).\n\n", val)
		fmt.Printf(usage, os.Args[0], len(charList))
		return
	}

	fp, err := os.Open(os.Args[len(os.Args) - 1])
	if err != nil {
		fatalf("Failed to open the input file: %+v", err)
	}
	defer fp.Close()

	img, err := png.Decode(fp)
	if err != nil {
		fatalf("Failed to decode the image: %+v", err)
	}

	// List the range of colors in the image
	colors := make(map[color.Color]string)
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			c := img.At(x, y)
			colors[c] = ""
			if max := len(charList); len(colors) > max {
				fatalf("Found more than %d colors in the picture!", max)
			}
		}
	}

	// Sort every color in the image
	sorted := SortableColor{}
	for k, _ := range colors {
		sorted = append(sorted, k)
	}
	sort.Stable(sorted)

	// Assign the equivalent unicode string for each color
	for i := range sorted {
		colors[sorted[i]] = charList[i]
	}

	// Convert the image to ascii
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			c := img.At(x, y)
			fmt.Print(colors[c])
		}
		fmt.Println("")
	}
}
