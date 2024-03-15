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
	scene = "Signal Reception"
	g.stateMachine.changeState(g, scene)
}

func ToMenu(g *Game) {
	setAllSwitchedFalse(g)
	g.switchedToMenu = true
	scene = "Main Menu"
	g.stateMachine.changeState(g, scene)
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
	scene = "Level Selection"
	g.stateMachine.changeState(g, scene)
}


func ToAbout(g *Game) {
	setAllSwitchedFalse(g)
	g.switchedToAbout = true
	scene = "About"
	g.stateMachine.changeState(g, "About")
}

/*
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