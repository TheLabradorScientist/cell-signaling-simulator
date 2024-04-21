package main

import (
	"math/rand"
	"strings"
)

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
	for i := 0; i < len(codon); i++ {
		transcription = append(transcription, baseToBase(string(codon[i])))
	}
	var result string = strings.Join([]string(transcription), "")
	return result
}

func baseToBase(base string) string {
	var newBase string
	switch base {
	case "A":
		newBase = "U"
	case "T":
		newBase = "A"
	case "G":
		newBase = "C"
	case "C":
		newBase = "G"
	case "U":
		newBase = "A"
	}
	return newBase
}


func translate(codon string) string {
	result := codonChart[codon]
	return result
}

func randomBase(nuclAcid string) string {
	seedSignal = rand.Intn(4) + 1
	switch seedSignal {
	case 1:
		return "A"
	case 2:
		if nuclAcid == "DNA" {
			return "T"
		} else {
			return "U"
		}
	case 3:
		return "G"
	case 4:
		return "C"
	default:
		return "A"
	}
}

func randomRNACodon(exception string) string {
	randCodon := ""
	for x := 0; x < 3; x++ {
		randCodon += randomBase("RNA")
	}
	if randCodon != exception {
		return randCodon
	} else {
		return randomRNACodon(exception)
	}
}

func randomDNACodon() string {
	exceptions := []string{"ATC", "ATT", "ACT"}
	randCodon := ""
	for x := 0; x < 3; x++ {
		randCodon += randomBase("DNA")
	}
	if !contains(exceptions, randCodon) {
		return randCodon
	} else {
		return randomDNACodon()
	}
}

func contains(list []string, T any) bool {
	for index := 0; index < len(list); index++ {
		if list[index] == T {
			return true
		}
	}
	return false
}

