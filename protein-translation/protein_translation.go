package protein

import "errors"

var (
	ErrStop        = errors.New("Stop")         // Stop Condon Encountered
	ErrInvalidBase = errors.New("Invalid Base") // Invalid Condon
)

var codonRNA = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

//FromCodon returns the protein for provided input. ErrStop for STOP codons.  ErrInvalidBase for invalid codons
func FromCodon(input string) (string, error) {
	if c, ok := codonRNA[input]; ok {
		if c == "STOP" {
			return "", ErrStop
		}
		return c, nil
	}
	return "", ErrInvalidBase
}

//FromRNA returns proteins for given input.  ErrInvalidBase for invalid input
func FromRNA(input string) ([]string, error) {
	var proteins []string
	for i := 0; i+2 < len(input); i += 3 {
		c, ok := codonRNA[input[i:i+3]]
		if !ok {
			return proteins, ErrInvalidBase
		}
		if c == "STOP" {
			break
		}
		proteins = append(proteins, c)
	}
	return proteins, nil
}
