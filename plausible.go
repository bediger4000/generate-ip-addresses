package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	ch := make(chan []rune, 5)

	go generate(os.Args[1], ch)

	for address := range ch {
		// fmt.Printf("candidate %q\n", string(address))
		if validAddress(address) {
			fmt.Printf("%q\n", string(address))
		}
	}
}

func generate(str string, ch chan []rune) {
	runes := []rune(str)
	address := []rune{}
	realgenerate(runes, address, ch)
	close(ch)
}

func realgenerate(runes []rune, address []rune, ch chan []rune) {
	if len(runes) == 0 {
		cpyaddress := make([]rune, len(address))
		copy(cpyaddress, address)
		ch <- cpyaddress
		return
	}

	nextaddress := make([]rune, len(address)+1)
	copy(nextaddress, address)

	nextaddress[len(address)] = runes[0]
	realgenerate(runes[1:], nextaddress, ch)
	if len(address) > 0 && nextaddress[len(address)-1] != '.' {
		nextaddress[len(address)] = '.'
		realgenerate(runes, nextaddress, ch)
	}
}

func validAddress(address []rune) bool {
	octetCount := 0
	dotCount := 0
	var current []rune

	for _, r := range address {
		if r == '.' {
			dotCount++
			if dotCount > 3 {
				// log.Printf("Found at least %d dots in %q\n", dotCount, string(address))
				return false
			}
			if len(current) > 0 {
				octetCount++
				if octetCount > 4 {
					// log.Printf("Found at least %d octets in %q\n", octetCount, string(address))
					return false
				}
			}
			if len(current) > 1 && current[0] == '0' {
				// log.Printf("Leading 0 in %q from %q\n", string(current), string(address))
				return false
			}
			if !checkValue(current) {
				// log.Printf("Incorrect value in %q from %q\n", string(current), string(address))
				return false
			}

			current = current[:0]
			continue
		}
		current = append(current, r)
	}
	if dotCount != 3 {
		// log.Printf("Found %d dots in %q\n", dotCount, string(address))
		return false
	}
	if len(current) > 0 {
		octetCount++
	}
	if octetCount != 4 {
		// log.Printf("Found %d octets in %q\n", octetCount, string(address))
		return false
	}

	if len(current) > 0 && current[0] == '0' {
		// log.Printf("Leading 0 in %q from %q\n", string(current), string(address))
		return false
	}

	return checkValue(current)
}

func checkValue(current []rune) bool {
	n, err := strconv.Atoi(string(current))
	if err != nil {
		log.Print(err)
		return false
	}
	if n > 255 || n < 0 {
		return false
	}
	return true
}
