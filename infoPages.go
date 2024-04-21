package main

func updateInfo() string {
	switch scene {
	case "Signal Reception":
		info = "\nWELCOME TO SIGNAL RECEPTION!\n\n" +
			"The cell signaling pathway begins\n" +
			"when a signaling molecule (ligand)\n" +
			"approaches the outside of the plasma\n" +
			"membrane. Receptors embedded in the\n" +
			"plasma membrane are SHAPE SPECIFIC\n" +
			"to the ligands they bind."
	case "Signal Transduction":
		info = "\nWELCOME TO SIGNAL TRANSDUCTION!\n\n" +
			"The phosphorylated TK1 travels through\n" +
			"the cytoplasm to bind and activate TK2.\n" +
			"Notice that the kinase phosphorylates by\n" +
			"transferring the 3rd phosphate group of\n" +
			"a new ATP molecule to TK2; the phosphate\n" +
			"group on TK1 remains bound."
	case "Transcription":
		info = "\nWELCOME TO mRNA TRANSCRIPTION!\n\n" +
			"The activated TFA enters the nucleus\n" +
			"and binds to the DNA template strand,\n" +
			"allowing RNA polymerase to bind to the\n" +
			"promoter region of the template strand.\n" +
			"RNA polymerase then reads the template\n" +
			"strand from 3' to 5', synthesizing a\n" +
			"new mRNA molecule with complementary\n" +
			"bases from 5' to 3' until it reaches\n" +
			"the terminator sequence."
	case "Translation":
		info = "\nWELCOME TO PROTEIN TRANSLATION!\n\n" +
			"The complete mRNA molecule exits\n" +
			"the nucleus and travels to the\n" +
			"cytoplasm, where a ribosome finds\n" +
			"the 5' guanosine cap and scans\n" +
			"for the first start codon.\n" +
			"The ribosome forms peptide bonds\n" +
			"between amino acids from tRNA\n" +
			"with complementary anticodons\n" +
			"to the mRNA transcript."
	default:
		info = ""
	}
	return info
}
