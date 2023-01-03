package konnakolliner

import (
	"strconv"
	"unicode"
)

func Highlight(line string, thalam, gathi int) (string, error) {
	result := ""
	hlSol := "" // Highlighted sol
	i := 0
	metSol := false
	for _, c := range line {
		if !unicode.IsSpace(c) {
			metSol = true
			if i%gathi == 0 {
				if hlSol == "" {
					hlSol = "|" + strconv.Itoa(i%thalam+1) + "|("
				}
				hlSol += string(c)
			} else {
				result += string(c)
			}
		} else if hlSol != "" {
			result += hlSol + ")" + string(c)
			hlSol = ""
			metSol = false
			i += 1
		} else {
			if metSol {
				i += 1
				metSol = false
			}
			result += string(c)
		}
	}
	if hlSol != "" {
		result += hlSol + ")"
	}
	return result, nil
}
