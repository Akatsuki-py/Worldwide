package gbc

import (
	"fmt"
	"gbc/pkg/emulator/debug"
	"gbc/pkg/emulator/joypad"
	"image/color"
	"time"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	width        = 160
	height       = 144
	cyclePerLine = 114
)

const (
	debugWidth  = 1270.
	debugHeight = 740.
)

var (
	frames     = 0
	second     = time.Tick(time.Second)
	skipRender bool
	fps        = 0
)

func (g *GBC) Draw(screen *ebiten.Image) {
	display := g.video.Display()
	if g.Debug.Enable {
		dScreen := ebiten.NewImage(int(debugWidth), int(debugHeight))
		dScreen.Fill(color.RGBA{35, 27, 167, 255})
		{
			// debug screen
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Scale(2, 2)
			op.GeoM.Translate(float64(10), float64(25))
			dScreen.DrawImage(ebiten.NewImageFromImage(display), op)
		}

		// debug FPS
		title := fmt.Sprintf("GameBoy FPS: %d", fps)
		ebitenutil.DebugPrintAt(dScreen, title, 10, 5)

		// debug register
		ebitenutil.DebugPrintAt(dScreen, g.debugRegister(), 340, 5)
		ebitenutil.DebugPrintAt(dScreen, g.debugIOMap(), 490, 5)

		// debug Cartridge
		ebitenutil.DebugPrintAt(dScreen, g.Cartridge.Debug.String(), 680, 5)

		cpuUsageX := 340
		// debug history (optional)
		if g.Debug.history.Flag() {
			ebitenutil.DebugPrintAt(dScreen, g.Debug.history.History(), 340, 120)
			cpuUsageX = 540
		}
		// debug GBC Usage
		ebitenutil.DebugPrintAt(dScreen, "GBC", cpuUsageX, 120)
		g.Debug.monitor.GBC.DrawUsage(dScreen, cpuUsageX+2, 140, g.isBoost())

		bgMap := g.video.Debug.BGMap()
		if bgMap != nil {
			// debug BG
			ebitenutil.DebugPrintAt(dScreen, "BG map", 10, 320)
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(10), float64(340))
			dScreen.DrawImage(ebiten.NewImageFromImage(bgMap), op)
		}

		{
			// debug tiles
			ebitenutil.DebugPrintAt(dScreen, "Tiles", 200, 320)
			tile := g.video.GetTileData()
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Scale(2, 2)
			op.GeoM.Translate(float64(200), float64(340))
			dScreen.DrawImage(tile, op)
		}

		op := &ebiten.DrawImageOptions{}
		screen.DrawImage(dScreen, op)
		return
	}
	screen.ReplacePixels(display.Pix)
}

func (g *GBC) handleJoypad() {
	pad := g.Config.Joypad
	result := g.joypad.Input(pad.A, pad.B, pad.Start, pad.Select, pad.Threshold)
	if result != 0 {
		switch result {
		case joypad.Pressed: // Joypad Interrupt
			if g.Reg.IME && g.getJoypadEnable() {
				g.setJoypadFlag(true)
			}
		case joypad.Pause:
			p, b := &g.Debug.pause, &g.Debug.Break
			if !g.Debug.Enable {
				return
			}

			if b.On() {
				b.SetFlag(debug.BreakDelay)
				p.SetOff(30)
				return
			}

			if !p.Delay() {
				if p.On() {
					p.SetOff(30)
				} else {
					p.SetOn(30)
				}
			}
		}
	}
}
