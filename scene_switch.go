package main

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
