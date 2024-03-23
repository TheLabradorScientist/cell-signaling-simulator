package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type SceneCreatorFunc func(g *Game)
type SceneConstructorMap map[string]SceneCreatorFunc

type State interface {
	Init(*Game)
	Update(*Game)
	Draw(*Game, *ebiten.Image)
}

type StateMachine struct {
	state State
	s_map SceneConstructorMap
}

func newStateMachine(s_map SceneConstructorMap) *StateMachine {
	return &StateMachine{
		state: nil,
		s_map: s_map,
	}
}

func (s *StateMachine) changeState(g *Game, s_name string) {
	s.s_map[s_name](g)
	s.state.Init(g)
	info = updateInfo()
	g.switchedScene = true
}

func (s *StateMachine) update(g *Game) {
	s.state.Update(g)
}

func (s *StateMachine) draw(g *Game, screen *ebiten.Image) {
	s.state.Draw(g, screen)
}

func (s *StateMachine) Scale(g *Game, screen *ebiten.Image) {
	for _, element := range g.state_array {
		element.scaleToScreen()
	} 
}