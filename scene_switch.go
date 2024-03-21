package main

func ToPlasma(g *Game) {
	scene = "Signal Reception"
	g.stateMachine.changeState(g, scene)
}

func ToMenu(g *Game) {
	scene = "Main Menu"
	g.stateMachine.changeState(g, scene)
}

func ToCyto1(g *Game) {
	scene = "Signal Transduction"
	g.stateMachine.changeState(g, scene)
}

func ToNucleus(g *Game) {
	scene = "Transcription"
	g.stateMachine.changeState(g, scene)
}

func ToCyto2(g *Game) {
	scene = "Translation"
	g.stateMachine.changeState(g, scene)
}

func ToLevelSelect(g *Game) {
	scene = "Level Selection"
	g.stateMachine.changeState(g, scene)
}


func ToAbout(g *Game) {
	scene = "About"
	g.stateMachine.changeState(g, scene)
}


func (g *Game) reset() {
	// Set length of all sprite arrays to 0
	// Set seed signal to random integer 
	// Set template to random codons
	
}