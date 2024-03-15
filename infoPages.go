package main

func updateInfo() string {
	switch scene {
	case "Signal Reception":
		info = "WELCOME TO THE SIGNAL\nRECEPTION STAGE!\n" +
		"The cell signaling pathway begins\nwhen a signaling molecule" +
		" (ligand)\napproaches the outside of the cell's\nplasma" +
		" membrane. Receptors embedded\nin the plasma membrane are " +
		"SHAPE\nSPECIFIC to the ligands they bind."
	case "Signal Transduction":
		info = "WELCOME TO THE SIGNAL\nTRANSDUCTION STAGE!\n" +
		"The phosphorylated TK1 travels through\nthe cytoplasm to bind" +
		"with\nand activate TK2. Notice that the\nkinase phosphorylates" +
		" by transferring\nthe 3rd phosphate group of\nan ATP molecule" +
		"to TK2;\nthe phosphate group on TK1\nremains bound."
	case "Transcription":
		info = "WELCOME TO THE mRNA\nTRANSCRIPTION STAGE!\n" +
		"The activated TFA enters the nucleus\nand binds to the DNA" +
		" template strand,\nallowing RNA polymerase to bind\nto the template.\n" +
		"RNA polymerase then 'reads' the\ntemplate strand from 3' to 5',\n" +
		"synthesizing a new mRNA molecule with\ncomplementary bases from 5' to 3'."
	case "Translation":
		info = "WELCOME TO THE PROTEIN\nTRANSLATION STAGE!\n" +
		"The complete mRNA molecule exits the\nnucleus and travels to the\n" +
		"cytoplasm, where a ribosome finds the 5'\nguanosine cap and scans for\n" +
		"the first start codon. The ribosome then\nforms peptide bonds between\n" +
		"amino acids from tRNA with complementary\ncodons to the mRNA transcript."
	default:
		info = ""
	}
	return info
}
