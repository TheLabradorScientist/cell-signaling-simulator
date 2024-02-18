package main

import "github.com/hajimehoshi/ebiten/v2"

type SceneCreatorFunc func() State
type SceneConstructorMap map[string]SceneCreatorFunc

type State interface {
	Init()
	Update(*Game)
	Draw(*ebiten.Image)
	Finish()
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
	if s.state != nil {
		s.state.Finish()
	}
	s.state = s.s_map[s_name]()
	s.state.Init()
}

func (s StateMachine) update(g *Game) {
	s.state.Update(g)
}

func (s StateMachine) draw(screen *ebiten.Image) {
	s.state.Draw(screen)
}
