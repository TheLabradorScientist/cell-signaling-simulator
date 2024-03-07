//MUST UPDATE ALL FUNCTIONS WITH NEW CODE

package main

import (
	//"fmt"

	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func setAllSwitchedFalse(g *Game) {
	g.switchedToPlasma = false
	g.switchedToMenu = false
	g.switchedToCyto1 = false
	g.switchedToNucleus = false
	g.switchedToCyto2 = false
	g.switchedToAbout   = false
	g.switchedToLevelSelect = false
}


func ToPlasma(g *Game) {
	setAllSwitchedFalse(g)
	g.switchedToPlasma = true
}

func ToMenu(g *Game) {
	setAllSwitchedFalse(g)
	g.switchedToMenu = true
}

func ToCyto1(g *Game) {
	setAllSwitchedFalse(g)
	g.switchedToCyto1 = true
}

func ToNucleus(g *Game) {
	setAllSwitchedFalse(g)
	g.switchedToNucleus = true
}

func ToCyto2(g *Game) {
	setAllSwitchedFalse(g)
	g.switchedToCyto2 = true
}

func ToLevelSelect(g *Game) {
	setAllSwitchedFalse(g)	
	g.switchedToLevelSelect = true
}


func ToAbout(g *Game) {
	setAllSwitchedFalse(g)
	g.switchedToAbout = true
}

func SwitchVol(g *Game) {
	if audioPlayer.IsPlaying() {
		audioPlayer.Pause()
		//var vol_image, _, err = ebitenutil.NewImageFromFile(loadFile("volButtonOff.png"))
		//if err != nil {
		//	fmt.Println("Error parsing date:", err)
		//}
		//volButton.image = vol_image
	} else { 
		audioPlayer.Play() 
		//var vol_image, _, err = ebitenutil.NewImageFromFile(loadFile("volButtonOn.png"))
		//if err != nil {
		//	fmt.Println("Error parsing date:", err)
		//}
		//volButton.image = vol_image
	}
}
/*
func ToMenu(g *Game) {
	g.stateMachine.changeState(g, "menu")
}

func ToLevel1(g *Game) {
	g.stateMachine.changeState(g, "level1")
}

func ToInfo(g *Game) {
	g.stateMachine.changeState(g, "info")
}

func ToLevelSelect(g *Game) {
	g.stateMachine.changeState(g, "level_select")
}

func ToLevel2(g *Game) {
	g.stateMachine.changeState(g, "level2")
}

func ToLevel3(g *Game) {
	g.stateMachine.changeState(g, "level3")
}

func ToLevel4(g *Game) {
	g.stateMachine.changeState(g, "level4")
}
*/