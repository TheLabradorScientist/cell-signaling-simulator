package main

func updateInfo() string {
	switch scene {
	case "Signal Reception":
		info = "WELCOME TO THE SIGNAL\nRECEPTION\nSTAGE!\n"
		info += "This simulator will\nguide you through the\ncomplete cell signaling\n"
		info += "pathway from reception\nthrough translation!\nClick the play "
		info += "button\nor select a level\nto begin."
	case "Signal Transduction":
		info = "WELCOME TO THE SIGNAL\nTRANSDUCTION\nSTAGE!\n"
		info += "This simulator will\nguide you through the\ncomplete cell signaling\n"
		info += "pathway from reception\nthrough translation!\nClick the play "
		info += "button\nor select a level\nto begin."
	case "Transcription":
		info = "WELCOME TO THE mRNA\nTRANSCRIPTIONY\nSTAGE!\n"
		info += "This simulator will\nguide you through the\ncomplete cell signaling\n"
		info += "pathway from reception\nthrough translation!\nClick the play "
		info += "button\nor select a level\nto begin."
	case "Translation":
		info = "WELCOME TO THE PROTEIN\nTRANSLATION\nSTAGE!\n"
		info += "This simulator will\nguide you through the\ncomplete cell signaling\n"
		info += "pathway from reception\nthrough translation!\nClick the play "
		info += "button\nor select a level\nto begin."
	default:
		info = ""
	}
	return info
}