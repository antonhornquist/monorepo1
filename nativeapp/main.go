package main

import (
	"log"
	"os"
	"fmt"
	"flag"
	"net/http"

	"image"
	//"image/color"
	_ "image/jpeg"
	"golang.org/x/image/draw"
	"path/filepath"
	// "io/ioutil"

	"gioui.org/app"

	"gioui.org/io/key"
	"gioui.org/io/system"

	"gioui.org/op"
	//"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
)

const thumbnailsFolder = "images"

var keyHandler = new(int)
var currentImage = 0
var images = make([]image.Image, 0)

func resizeImage(src image.Image) image.Image {
	// Set the expected size that you want:
	height := 1000
	factor := src.Bounds().Max.Y/height
	dst := image.NewRGBA(image.Rect(0, 0, src.Bounds().Max.X/factor, height))

	// Resize:
	// TODO draw.NearestNeighbor faster draw.CatmullRom better quality
	draw.NearestNeighbor.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)

	return dst
}

func getRemoteImage(presentationPort int, id string) image.Image {
	requestURL := fmt.Sprintf("http://localhost:%d/thumbnails/%s", presentationPort, id)
	res, err := http.Get(requestURL)
	if err != nil {
		log.Fatal(fmt.Sprintf("HTTP request error: %v\n", err))
	}

	image, _, err := image.Decode(res.Body)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error decoding image: %v\n", err))
	}
	return image
}

func main() {
	var (
		// presentationPort = flag.Int("port", 80, "Listen port.")
		_ = flag.Int("port", 80, "Listen port.")
	)

	flag.Parse()

	/*
	ids := []string{"6o3vy", "v2seq", "dwsy6", "8fskr"}
	for _, id := range ids {
		images = append(images, getRemoteImage(*presentationPort, id))
	}

	matches, err := filepath.Glob(filepath.Join(thumbnailsFolder, "*"))
	if err != nil {
		log.Fatal(err)
	}

	for _, match := range matches {
		log.Printf(match)
		image, err := getImageFromThumbnail(match)
		if err != nil {
			log.Fatal(err)
		}
		// images = append(images, resizeImage(image))
		images = append(images, image)
	}
	*/

	go func() {
		scale := 0.5
		w := app.NewWindow(
			app.Title("Photos"),
			app.Size(unit.Dp(1366*scale), unit.Dp(768*scale)),
		)
		var ops op.Ops
		for e := range w.Events() {
			switch e := e.(type) {
			case system.DestroyEvent:
				os.Exit(0)
			case system.FrameEvent:
				ops.Reset()

				for _, e := range e.Queue.Events(keyHandler) {
					// log.Printf("%v", e)
					if x, ok := e.(key.Event); ok {
						switch x.State {
						case key.Press:
							if x.Name == key.NameLeftArrow {
								currentImage = currentImage - 1
								if currentImage < 0 {
									currentImage = len(images)-1
								}
							} else if x.Name == key.NameRightArrow {
								currentImage = (currentImage + 1) % len(images)
							}
						}
					}
				}
				key.InputOp{
					Tag:  keyHandler,
					Keys: "[" + key.NameLeftArrow + "," + key.NameRightArrow + "]",
				}.Add(&ops)

				if len(images) > 0 {
					paint.NewImageOp(images[currentImage]).Add(&ops)
					paint.PaintOp{}.Add(&ops)
				}

				/*
				rect := clip.Rect{Min: image.Pt(50, 25), Max: image.Pt(100, 100)}.Push(&ops)
				paint.ColorOp{
					Color: color.NRGBA{A: 0xff, R: 0xff},
				}.Add(&ops)
				paint.PaintOp{}.Add(&ops)
				rect.Pop()
				*/

				e.Frame(&ops)
			}
		}
	}()
	app.Main()
}

func getImageFromThumbnail(filePath string) (image.Image, error) {
	f, err := os.Open(filepath.Join(filePath))
	defer f.Close()
	if err != nil {
		return nil, err
	}
	image, _, err := image.Decode(f)
	return image, err
}
