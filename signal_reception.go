package main

// Signal stores 4 different shapes
func matchSR(signal string, receptor string) bool {
	return signalReceptor[signal] == receptor
}

var signalReceptor = map[string]string{
	"signalA.png":"receptorA.png", 
	"signalB.png":"receptorB.png", 
	"signalC.png":"receptorC.png", 
	"signalD.png":"receptorD.png", 
}