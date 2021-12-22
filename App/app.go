package app

import (
	"log"
	"os"
	"github.com/rivo/tview"
	"github.com/0xNathanW/chip-8/chip8"
)

type App struct {
	Chip8 		*chip8.Chip8		// The Chip8 system.
	App 		*tview.Application	// tview application object.
	Layout 		*tview.Grid  
	Opcodes 	*tview.TextView
	Info 		*tview.TextView
	SelectROM	*tview.List
	Pages 		*tview.Pages
	Display 	*tview.TextView
}



func NewApp() *App {
	app := &App{
		Chip8: chip8.NewSystem(),

		App: tview.NewApplication().
			EnableMouse(false),

		Layout: tview.NewGrid().
			SetRows(16, 16).
			SetColumns(128, 32).
			SetBorders(false),

		Opcodes: tview.NewTextView().
			SetScrollable(true).
			SetMaxLines(20),

		Info: tview.NewTextView().
			SetScrollable(false).
			SetWordWrap(true).
			SetText("Press 'p' to pause/unpause the emulator.\n" +
					"When paused:\n" +
					"Press <, > to decrease/increase the clock speed."),

		SelectROM: tview.NewList().
			ShowSecondaryText(false),

		Pages: tview.NewPages(),

		Display: tview.NewTextView().
			SetScrollable(false),

	}

	app.Pages.AddPage("Opcodes", app.Opcodes, true, false)
	app.Pages.AddPage("Roms", app.SelectROM, true, true)

	app.setupSelectROM()

	app.Opcodes.SetBorder(true).SetTitle("Opcodes")
	app.Info.SetBorder(true).SetTitle("Info")
	app.SelectROM.SetBorder(true).SetTitle("Select ROM")


	app.drawLayout()
	return app
}

func (a *App) drawLayout() {
 	a.Layout.AddItem(
		a.Display,
		0, 0, 2, 1,		// row, col, rowspan, colspan
		64, 128, false,		// min height, min width.
	).AddItem(
		a.Info,
		1, 1, 1, 1,
		64, 32, false,
	).AddItem(
		a.Pages,
		0, 1, 1, 1,
		64, 32, false,
	)
}

func (a *App) setupSelectROM() {
	files, err := os.ReadDir("./roms")
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		a.SelectROM.AddItem(file.Name(), "", '>', func() {
			a.Chip8.LoadROM(file.Name())
			log.Println("Loaded ROM:", file.Name())
			a.emulate()
		})
	}
}