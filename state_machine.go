package main

import "github.com/hajimehoshi/ebiten/v2"

type SceneCreatorFunc func() State
type SceneConstructorMap map[string]SceneCreatorFunc

type State interface {
	Init()
	Update(*Game)
	Draw(*ebiten.Image)
}

type StateMachine struct {
	state State
	s_map SceneConstructorMap
}

func newStateMachine(s_map SceneConstructorMap) StateMachine {
	return StateMachine{
		state: nil,
		s_map: s_map,
	}
}

func (s StateMachine) changeState(s_name string) {
	s.state = s.s_map[s_name]()
	s.state.Init()
}

func (s StateMachine) update(g *Game) {
	s.state.Update(g)
}

func (s StateMachine) draw(screen *ebiten.Image) {
	s.state.Draw(screen)
}




func MenuToPlasma(g *Game) {
	g.switchedMenuToPlasma = true
	g.switchedPlasmaToMenu = false
	g.switchedPlasmaToCyto1 = false
	g.switchedCyto1ToNucleus = false
	g.switchedNucleusToCyto2 = false

	g.switchedMenuToLevelSelect = false
	g.switchedLevelSelectToPlasma = false
	g.switchedLevelSelectToCyto1 = false
	g.switchedLevelSelectToNucleus = false
	g.switchedLevelSelectToCyto2 = false
	g.switchedLevelSelectToMenu = false 
}

func PlasmaToMenu(g *Game) {
	g.switchedMenuToPlasma = false
	g.switchedPlasmaToMenu = true
	g.switchedPlasmaToCyto1 = false
	g.switchedCyto1ToNucleus = false
	g.switchedNucleusToCyto2 = false

	g.switchedMenuToLevelSelect = false
	g.switchedLevelSelectToPlasma = false
	g.switchedLevelSelectToCyto1 = false
	g.switchedLevelSelectToNucleus = false
	g.switchedLevelSelectToCyto2 = false
	g.switchedLevelSelectToMenu = false 
}

func PlasmaToCyto1(g *Game) {
	g.switchedMenuToPlasma = false
	g.switchedPlasmaToMenu = false
	g.switchedPlasmaToCyto1 = true
	g.switchedCyto1ToNucleus = false
	g.switchedNucleusToCyto2 = false

	g.switchedMenuToLevelSelect = false
	g.switchedLevelSelectToPlasma = false
	g.switchedLevelSelectToCyto1 = false
	g.switchedLevelSelectToNucleus = false
	g.switchedLevelSelectToCyto2 = false
	g.switchedLevelSelectToMenu = false 
}

func Cyto1ToNucleus(g *Game) {
	g.switchedMenuToPlasma = false
	g.switchedPlasmaToMenu = false
	g.switchedPlasmaToCyto1 = false
	g.switchedCyto1ToNucleus = true
	g.switchedNucleusToCyto2 = false

	g.switchedMenuToLevelSelect = false
	g.switchedLevelSelectToPlasma = false
	g.switchedLevelSelectToCyto1 = false
	g.switchedLevelSelectToNucleus = false
	g.switchedLevelSelectToCyto2 = false
	g.switchedLevelSelectToMenu = false 
}

func NucleusToCyto2(g *Game) {
	g.switchedMenuToPlasma = false
	g.switchedPlasmaToMenu = false
	g.switchedPlasmaToCyto1 = false
	g.switchedCyto1ToNucleus = false
	g.switchedNucleusToCyto2 = true
	g.switchedLevelSelectToMenu = false 

	g.switchedMenuToLevelSelect = false
	g.switchedLevelSelectToPlasma = false
	g.switchedLevelSelectToCyto1 = false
	g.switchedLevelSelectToNucleus = false
	g.switchedLevelSelectToCyto2 = false
	g.switchedLevelSelectToMenu = false 
}

func MenuToLevelSelect(g *Game) {
	g.switchedMenuToPlasma = false
	g.switchedPlasmaToMenu = false
	g.switchedPlasmaToCyto1 = false
	g.switchedCyto1ToNucleus = false
	g.switchedNucleusToCyto2 = false
	
	g.switchedMenuToLevelSelect = true
	g.switchedLevelSelectToPlasma = false
	g.switchedLevelSelectToCyto1 = false
	g.switchedLevelSelectToNucleus = false
	g.switchedLevelSelectToCyto2 = false
	g.switchedLevelSelectToMenu = false 
}

func LevelSelectToPlasma(g *Game) {
	g.switchedMenuToPlasma = false
	g.switchedPlasmaToMenu = false
	g.switchedPlasmaToCyto1 = false
	g.switchedCyto1ToNucleus = false
	g.switchedNucleusToCyto2 = false
	
	g.switchedMenuToLevelSelect = false
	g.switchedLevelSelectToPlasma = true
	g.switchedLevelSelectToCyto1 = false
	g.switchedLevelSelectToNucleus = false
	g.switchedLevelSelectToCyto2 = false
	g.switchedLevelSelectToMenu = false  
}

func LevelSelectToCyto1(g *Game) {
	g.switchedMenuToPlasma = false
	g.switchedPlasmaToMenu = false
	g.switchedPlasmaToCyto1 = false
	g.switchedCyto1ToNucleus = false
	g.switchedNucleusToCyto2 = false
	
	g.switchedMenuToLevelSelect = false
	g.switchedLevelSelectToPlasma = false
	g.switchedLevelSelectToCyto1 = true
	g.switchedLevelSelectToNucleus = false
	g.switchedLevelSelectToCyto2 = false
	g.switchedLevelSelectToMenu = false  
}

func LevelSelectToNucleus(g *Game) {
	g.switchedMenuToPlasma = false
	g.switchedPlasmaToMenu = false
	g.switchedPlasmaToCyto1 = false
	g.switchedCyto1ToNucleus = false
	g.switchedNucleusToCyto2 = false
	
	g.switchedMenuToLevelSelect = false
	g.switchedLevelSelectToPlasma = false
	g.switchedLevelSelectToCyto1 = false
	g.switchedLevelSelectToNucleus = true
	g.switchedLevelSelectToCyto2 = false
	g.switchedLevelSelectToMenu = false  
}

func LevelSelectToCyto2(g *Game) {
	g.switchedMenuToPlasma = false
	g.switchedPlasmaToMenu = false
	g.switchedPlasmaToCyto1 = false
	g.switchedCyto1ToNucleus = false
	g.switchedNucleusToCyto2 = false
	
	g.switchedMenuToLevelSelect = false
	g.switchedLevelSelectToPlasma = false
	g.switchedLevelSelectToCyto1 = false
	g.switchedLevelSelectToNucleus = false
	g.switchedLevelSelectToCyto2 = true
	g.switchedLevelSelectToMenu = false  
}

func LevelSelectToMenu(g *Game) {
	g.switchedMenuToPlasma = false
	g.switchedPlasmaToMenu = false
	g.switchedPlasmaToCyto1 = false
	g.switchedCyto1ToNucleus = false
	g.switchedNucleusToCyto2 = false
	
	g.switchedMenuToLevelSelect = false
	g.switchedLevelSelectToPlasma = false
	g.switchedLevelSelectToCyto1 = false
	g.switchedLevelSelectToNucleus = false
	g.switchedLevelSelectToCyto2 = false
	g.switchedLevelSelectToMenu = true  
}