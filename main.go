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
	seedSignal = rand.Intn(4) + 1
	signal     Signal
	receptorA  Receptor
	receptorB  Receptor
	receptorC  Receptor
	receptorD  Receptor
	template   [5]string
)

type Game struct {
	defaultFont            Font
	switchedMenuToPlasma   bool
	switchedPlasmaToMenu   bool
	switchedPlasmaToCyto1  bool
	switchedCyto1ToNucleus bool
	switchedNucleusToCyto2 bool
}

func MenuToPlasma(g *Game) {
	g.switchedMenuToPlasma = true
	g.switchedPlasmaToMenu = false
	g.switchedPlasmaToCyto1 = false
	g.switchedCyto1ToNucleus = false
	g.switchedNucleusToCyto2 = false
}

func PlasmaToMenu(g *Game) {
	g.switchedMenuToPlasma = false
	g.switchedPlasmaToMenu = true
	g.switchedPlasmaToCyto1 = false
	g.switchedCyto1ToNucleus = false
	g.switchedNucleusToCyto2 = false
}

func PlasmaToCyto1(g *Game) {
	g.switchedMenuToPlasma = false
	g.switchedPlasmaToMenu = false
	g.switchedPlasmaToCyto1 = true
	g.switchedCyto1ToNucleus = false
	g.switchedNucleusToCyto2 = false
}

func Cyto1ToNucleus(g *Game) {
	g.switchedMenuToPlasma = false
	g.switchedPlasmaToMenu = false
	g.switchedPlasmaToCyto1 = false
	g.switchedCyto1ToNucleus = true
	g.switchedNucleusToCyto2 = false
}

func NucleusToCyto2(g *Game) {
	g.switchedMenuToPlasma = false
	g.switchedPlasmaToMenu = false
	g.switchedPlasmaToCyto1 = false
	g.switchedCyto1ToNucleus = false
	g.switchedNucleusToCyto2 = true
}

func init() {
	startBg, _, err = ebitenutil.NewImageFromFile("MenuBg.png")
	if err != nil {
		log.Fatal(err)
	}
	plasmaBg, _, err = ebitenutil.NewImageFromFile("PlasmaBg.png")
	if err != nil {
		log.Fatal(err)
	}
	playbutton = newButton("PlayButton.png", newRect(400, 300, 232, 129), MenuToPlasma)

	switch seedSignal {
	case 1:
		signal = newSignal("signalA.png", newRect(100, 100, 91, 87))
		signal.signalType = "signalA"
		// PLACEHOLDER IN CASE WE DO NOT GET TIME TO CODE RANDOM CODONS
		//template := []string{"TAC", "GTC", "CGG", "ACA", "UGA"}

	case 2:
		signal = newSignal("signalB.png", newRect(100, 100, 91, 87))
		signal.signalType = "signalA"
	case 3:
		signal = newSignal("signalC.png", newRect(100, 100, 91, 87))
		signal.signalType = "signalA"
	case 4:
		signal = newSignal("signalD.png", newRect(100, 100, 91, 87))
		signal.signalType = "signalA"
	}

	receptorA = newReceptor("receptorA.png", newRect(0, 500, 233, 284), "receptorA")
	receptorB = newReceptor("receptorB.png", newRect(250, 500, 233, 284), "receptorB")
	receptorC = newReceptor("receptorC.png", newRect(500, 500, 233, 284), "receptorC")
	receptorD = newReceptor("receptorD.png", newRect(750, 500, 233, 284), "receptorD")
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
		if !signal.is_dragged {
			receptorA.update()
			receptorB.update()
			receptorC.update()
			receptorD.update()
		}
		if receptorA.is_touching_signal {
			if matchSR(signal.signalType, receptorA.receptorType) {
				PlasmaToCyto1(g)
			}
		}
		if receptorB.is_touching_signal {
			if matchSR(signal.signalType, receptorB.receptorType) {
				PlasmaToCyto1(g)
			}
		}
		if receptorC.is_touching_signal {
			if matchSR(signal.signalType, receptorC.receptorType) {
				PlasmaToCyto1(g)
			}
		}
		if receptorD.is_touching_signal {
			if matchSR(signal.signalType, receptorD.receptorType) {
				PlasmaToCyto1(g)
			}
		}
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
		receptorA.draw(screen)
		receptorB.draw(screen)
		receptorC.draw(screen)
		receptorD.draw(screen)
		signal.draw(screen)
		g.defaultFont.drawFont(screen, "WELCOME TO THE PLASMA MEMBRANE! \n Drag the signal to the matching receptor to enter the cell!", 100, 100, color.White)
		if signal.is_dragged {
			signal.draw(screen)
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
