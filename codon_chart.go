package main

import "strings"

var codonChart = map[string]string{
	"UUU": "Phe", "UUC": "Phe", "UUA": "Leu", "UUG": "Leu",
	"CUU": "Leu", "CUC": "Leu", "CUA": "Leu", "CUG": "Leu",
	"AUU": "Ile", "AUC": "Ile", "AUA": "Ile", "AUG": "Met",
	"GUU": "Val", "GUC": "Val", "GUA": "Val", "GUG": "Val",

	"UCU": "Ser", "UCC": "Ser", "UCA": "Ser", "UCG": "Ser",
	"CCU": "Pro", "CCC": "Pro", "CCA": "Pro", "CCG": "Pro",
	"ACU": "Thr", "ACC": "Thr", "ACA": "Thr", "ACG": "Thr",
	"GCU": "Ala", "GCC": "Ala", "GCA": "Ala", "GCG": "Ala",

	"UAU": "Tyr", "UAC": "Tyr", "UAA": "STOP", "UAG": "STOP",
	"CAU": "His", "CAC": "His", "CAA": "Gln", "CAG": "Gln",
	"AAU": "Asn", "AAC": "Asn", "AAA": "Lys", "AAG": "Lys",
	"GAU": "Asp", "GAC": "Asp", "GAA": "Glu", "GAG": "Glu",

	"UGU": "Cys", "UGC": "Cys", "UGA": "STOP", "UGG": "Trp",
	"CGU": "Arg", "CGC": "Arg", "CGA": "Arg", "CGG": "Arg",
	"AGU": "Ser", "AGC": "Ser", "AGA": "Arg", "AGG": "Arg",
	"GGU": "Gly", "GGC": "Gly", "GGA": "Gly", "GGG": "Gly",
}

func transcribe(codon string) string {
	transcription := []string{}
	var n string
	for i := 0; i < len(codon); i++ {
		n = string(codon[i])
		if n == "A" {
			transcription = append(transcription, "U")
		}
		if n == "T" {
			transcription = append(transcription, "A")
		}
		if n == "G" {
			transcription = append(transcription, "C")
		}
		if n == "C" {
			transcription = append(transcription, "G")
		}
	}
	var result string = strings.Join([]string(transcription), "")
	return result
}

func translate(codon string) string {
	result := codonChart[codon]
	return result
}