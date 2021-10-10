package Tests

// "fmt"
// "github.com/hajimehoshi/ebiten/v2"
// "github.com/0xNathanW/CHIP-8/CHIP8"
// "github.com/hajimehoshi/ebiten/ebitenutil"

// func Test() *CHIP8.Chip8 {
// 	inst := CHIP8.Initialise()
// 	// fmt.Println("------- Memory Before Load--------")
// 	// fmt.Println(inst.Memory)
// 	// fmt.Println("-----------------------------------")
// 	inst.LoadROM()
// 	// fmt.Println("------- Memory After Load--------")
// 	// fmt.Println(inst.Memory)
// 	// fmt.Println("-----------------------------------")

// 	for {
// 		fmt.Println("----------- New Cycle ----------------")
// 		inst.Cycle()
// 		fmt.Println("Program Counter: ", inst.PC)
// 		fmt.Println("Registers: ", inst.V)
// 		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>> GFX <<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
// 		TestDisplayOutput(inst)
// 	}
// }

// func TestDisplayOutput(c *CHIP8.Chip8) {
// 	for y := 0; y < int(len(c.Display)); y++ {
// 		for x := 0; x < int(len(c.Display[0])); x++ {
// 			pixel := c.Display[y][x]
// 			if pixel == 1 {
// 				fmt.Print("██")
// 			} else {
// 				fmt.Print("--")
// 			}
// 		}
// 		fmt.Print("\n")
// 	}
// }

// func TestKeypresses(c *CHIP8.Chip8() {

// }

// func FPSempty(g *Game) {
// 	type Game struct{}

// 	func (g *Game) Update() error {
// 		return nil
// 	}

// 	func (g *Game) Draw(screen *ebiten.Image) {
// 		ebitenutil.DebugPrint(screen, "Hello, World!")
// 	}

// 	func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
// 		return 320, 240
// 	}

// 	ebiten.SetWindowSize(640, 480)
// 	ebiten.SetWindowTitle("Hello, World!")
// 	if err := ebiten.RunGame(&Game{}); err != nil {
// 		log.Fatal(err)
// 	}

// }

// func get_win_size(out syscall.Handle) coord {
// 	err := get_console_screen_buffer_info(out, &tmp_info)
// 	if err != nil {
// 		panic(err)
// 	}

// 	min_size := get_win_min_size(out)

// 	size := coord{
// 		x: tmp_info.window.right - tmp_info.window.left + 1,
// 		y: tmp_info.window.bottom - tmp_info.window.top + 1,
// 	}

// 	if size.x < min_size.x {
// 		size.x = min_size.x
// 	}

// 	if size.y < min_size.y {
// 		size.y = min_size.y
// 	}

// 	return size
// }
