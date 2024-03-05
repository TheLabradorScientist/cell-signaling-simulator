package main

// Signal stores 4 different shapes
func matchSR(signalType string, receptorType string) bool {
	return signalReceptor[signalType] == receptorType
}

var signalReceptor = map[string]string{
	"signalA": "receptorA",
	"signalB": "receptorB",
	"signalC": "receptorC",
	"signalD": "receptorD",
}
