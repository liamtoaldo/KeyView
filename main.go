package main

import (
	"log"
	//"strings"
	"bytes"
	_ "image/png"
	"image"
	"image/color"

	//ebiten (first time using this library)
	"github.com/hajimehoshi/ebiten"
	//"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/examples/keyboard/keyboard"
	kbimg "github.com/hajimehoshi/ebiten/examples/resources/images/keyboard"
)

const (
	screenWidth = 320
	screenHeight = 200
)

var keyboardImage *ebiten.Image

func init() {
	
	//getting image from ebiten website
	
	img, _, err := image.Decode(bytes.NewReader(kbimg.Keyboard_png))
	if err != nil {
		log.Fatal(err)
	}

	keyboardImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

}

type Game struct {
	pressed []ebiten.Key
}


//I created this function to update the screen everytime a key gets pressed

func (g *Game) Update(screen *ebiten.Image) error {
	g.pressed = nil

	// Fill the screen with transparent color
    screen.Fill(color.NRGBA{0, 0, 0, 0})

	//I used a for loop to update the whole process
	for k := ebiten.Key(0); k <= ebiten.KeyMax; k++ {
		if ebiten.IsKeyPressed(k) {
			g.pressed = append(g.pressed, k)
		}
	}
	return nil
}

//basically the function to draw the keyboard image in the window etc.

func (g *Game) Draw(screen *ebiten.Image) {
	const (
		offsetX = 24
		offsetY = 40
	)

	//Draws the gray keyboard image (when not clicked)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(offsetX, offsetY)
	op.ColorM.Scale(0.5, 0.5, 0.5, 1)
	screen.DrawImage(keyboardImage, op)

	//Draws the highlighted white keys (when clicked)
	op = &ebiten.DrawImageOptions{}
	for _, p := range g.pressed {
		op.GeoM.Reset()
		r, ok := keyboard.KeyRect(p)
		if !ok {
			continue
		}
		op.GeoM.Translate(float64(r.Min.X), float64(r.Min.Y))
		op.GeoM.Translate(offsetX, offsetY)
		screen.DrawImage(keyboardImage.SubImage(r).(*ebiten.Image), op)

	}

	//joins the strings when key is pressed simultaneously

	/* TODO */
	
}


//Setting the window size

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Keyboard Viewer")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}

}