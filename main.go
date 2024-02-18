package main

import (
	"image/color"
	_ "image/png"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 960
	screenHeight = 1000
)

var (
	err        error
	scene      string = "Main Menu"
	startBg    *ebiten.Image
	plasmaBg   *ebiten.Image
	cytoBg_1   *ebiten.Image
	nucleusBg  *ebiten.Image
	cytoBg_2   *ebiten.Image
	playbutton Button
	s_machine  StateMachine
	min = 1
    max = 4
    seedSignal = rand.Intn(max - min) + min
	signal     Signal
	receptorA  Receptor
	receptorB  Receptor
	receptorC  Receptor
	receptorD  Receptor
	
)

type Game struct {
	defaultFont         Font
	switchedMenuToPlasma bool
	switchedPlasmaToMenu bool
	switchedPlasmaToCyto1 bool
	switchedCyto1ToNucleus bool
	switchedNucleusToCyto2 bool
}

func MenuToPlasma(g *Game) {
	g.switchedMenuToPlasma = true
	g.switchedPlasmaToMenu = false
}

func PlasmaToMenu(g *Game) {
	g.switchedMenuToPlasma = false
	g.switchedPlasmaToMenu = true
}

func init() {
	startBg, _, err = ebitenutil.NewImageFromFile("MenuBg.png")
	if err != nil {
		log.Fatal(err)
	}
	plasmaBg, _, err = ebitenutil.NewImageFromFile("PlasmaBg.png")

	playbutton = newButton("PlayButton.png", newRect(400, 300, 232, 129), MenuToPlasma)

	switch seedSignal {
	case 1:
		signal = newSignal("signalA.png", newRect(100, 100, 91, 87))
	case 2:
		signal = newSignal("signalB.png", newRect(100, 100, 91, 87))
	case 3:
		signal = newSignal("signalC.png", newRect(100, 100, 91, 87))
	case 4:
		signal = newSignal("signalD.png", newRect(100, 100, 91, 87))
	}

	receptorA = newReceptor("receptorA.png", newRect(0, 500, 233, 284))
	receptorB = newReceptor("receptorB.png", newRect(250, 500, 233, 284))
	receptorC = newReceptor("receptorC.png", newRect(500, 500, 233, 284))
	receptorD = newReceptor("receptorD.png", newRect(750, 500, 233, 284))
}

func (g *Game) Update() error {
	switch scene {
	case "Main Menu":
		ebiten.SetWindowTitle("Main Menu")
		ebiten.SetWindowSize(screenWidth, screenHeight)
		playbutton.on_click(g)
	case "Signal Reception":
		ebiten.SetWindowTitle("Signal Reception")
		ebiten.SetWindowSize(screenWidth, screenHeight)
		signal.on_click(g)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch scene {
	case "Main Menu":
		screen.DrawImage(startBg, nil)
		playbutton.draw(screen)
		//ebitenutil.DebugPrint(screen, transcribe("CAT"))
		//ebitenutil.DebugPrint(screen, translate("UAG"))
		if g.switchedMenuToPlasma {
			scene = "Signal Reception"
		}
	case "Signal Reception":
		screen.DrawImage(plasmaBg, nil)
		signal.draw(screen)
		receptorA.draw(screen)
		receptorB.draw(screen)
		receptorC.draw(screen)
		receptorD.draw(screen)
		g.defaultFont.drawFont(screen, "WELCOME TO THE PLASMA MEMBRANE! \n Drag the signal to the matching receptor to enter the cell!", 100, 100, color.White)
		if g.switchedPlasmaToMenu {
			scene = "Main Menu"
			if signal.is_dragged {
				signal.draw(screen)
			}
		}
	case "Signal Transduction":
		screen.DrawImage(cytoBg_1, nil)
	
		g.defaultFont.drawFont(screen, "WELCOME TO THE CYTOPLASM! \n Click when each kinase overlaps to follow the phosphorylation cascade!!", 100, 100, color.White)

	case "Transcription":
		screen.DrawImage(nucleusBg, nil)

		g.defaultFont.drawFont(screen, "WELCOME TO THE NUCLEUS! \n Match each codon on the DNA template to the corresponding RNA codon to transcribe a new mRNA molecule!!!", 100, 100, color.White)

	case "Translation":
		screen.DrawImage(cytoBg_2, nil)

		g.defaultFont.drawFont(screen, "FINALLY, BACK TO THE CYTOPLASM! \n Match each codon from your mRNA template to its corresponding amino acid to synthesize your protein!!!!", 100, 100, color.White)
	
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	game := &Game{}
	game.defaultFont = newFont("ThaleahFat.ttf", 25)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
