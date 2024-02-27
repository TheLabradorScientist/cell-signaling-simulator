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