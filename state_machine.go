package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type SceneCreatorFunc func(g *Game)
type SceneConstructorMap map[string]SceneCreatorFunc

type State interface {
	Init()
	Update(*Game)
	Draw(*Game, *ebiten.Image)
}

type StateMachine struct {
	state State
	s_map SceneConstructorMap
}

func newStateMachine(s_map SceneConstructorMap) StateMachine {
	return StateMachine{
		state: MainMenu{},
		s_map: s_map,
	}
}

func (s StateMachine) changeState(g *Game, s_name string) {
	s.s_map[s_name](g)
	s.state.Init()
}

func (s StateMachine) update(g *Game) {
	s.state.Update(g)
}

func (s StateMachine) draw(g *Game, screen *ebiten.Image) {
	s.state.Draw(g, screen)
}