package main

func updateInfo() string {
	switch scene {
	case "Signal Reception":
		info = "WELCOME TO THE SIGNAL\nRECEPTION STAGE!\n"
		info += "The cell signaling pathway begins\nwhen a signaling molecule"
		info += " (ligand)\napproaches the outside of the cell's\nplasma"
		info += " membrane. Receptors embedded\nin the plasma membrane are "
		info += "SHAPE\nSPECIFIC to the ligands they bind."
	case "Signal Transduction":
		info = "WELCOME TO THE SIGNAL\nTRANSDUCTION STAGE!\n"
		info += "The phosphorylated TK1 travels through\nthe cytoplasm to bind"
		info += "with\nand activate TK2. Notice that the\nkinase phosphorylates"
		info += " by transferring\nthe 3rd phosphate group of\nan ATP molecule"
		info += "to TK2;\nthe phosphate group on TK1\nremains bound."
	case "Transcription":
		info = "WELCOME TO THE mRNA\nTRANSCRIPTION STAGE!\n"
		info += "The activated TFA enters the nucleus\nand binds to the DNA"
		info += " template strand,\nallowing RNA polymerase to bind\nto the template.\n"
		info += "RNA polymerase then 'reads' the\ntemplate strand from 3' to 5',\n"
		info += "synthesizing a new mRNA molecule with\ncomplementary bases from 5' to 3'."
	case "Translation":
		info = "WELCOME TO THE PROTEIN\nTRANSLATION STAGE!\n"
		info += "The complete mRNA molecule exits the\nnucleus and travels to the\n"
		info += "cytoplasm, where a ribosome finds the 5'\nguanosine cap and scans for\n"
		info += "the first start codon. The ribosome then\nforms peptide bonds between\n"
		info += "amino acids from tRNA with complementary\ncodons to the mRNA transcript."
	default:
		info = ""
	}
	return info
}
