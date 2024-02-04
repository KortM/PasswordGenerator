package main

import (
	"bytes"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Password generator (author: imarchenko) \nUse flag -help")
		return
	}
	help := flag.Bool("help", false, "Help")
	special := flag.Bool("special", false, "Special symbols")
	lower := flag.Bool("lower", false, "Lower Characters")
	upper := flag.Bool("upper", false, "Upper Characters")
	number := flag.Bool("number", false, "Numbers")
	var length int
	flag.IntVar(&length, "length", 8, "Password length")
	flag.Parse()

	if *help {
		printHelp()
		return
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	generatePassword(r, length, *special, *lower, *upper, *number)
}

// Refence
func printHelp() {
	fmt.Println("Keys:\n-length - use to specify the password length")
	fmt.Println("-special - use to include special characters")
	fmt.Println("-lower - use to include lower symbols [a-z]")
	fmt.Println("-upper - use to include upper symbols [A-Z]")
	fmt.Println("-number - use to include numbers [0-9]")
	fmt.Println("Example:\n.\\main.exe -length 8 -special -lower -upper -number\nResult: 34jv8FW&")
}

// Generate a password by selecting random Unicode characters
func generatePassword(r *rand.Rand, length int, special, alpha_lower, alhpa_upper, number bool) {
	var applyes []func(*rand.Rand) rune
	if special {
		applyes = append(applyes, getSpecialSymbol)
	}
	if alpha_lower {
		applyes = append(applyes, getLowerSymbol)
	}
	if alhpa_upper {
		applyes = append(applyes, getUpperSymbol)
	}
	if number {
		applyes = append(applyes, getNumberSybol)
	}
	var passwd bytes.Buffer

	for i := 0; i < length; i++ {
		num := r.Intn(len(applyes))
		passwd.WriteRune(applyes[num](r))
	}
	fmt.Println(passwd.String())

}

// Generate one lower symbol
func getLowerSymbol(r *rand.Rand) rune {
	code := r.Intn(int('z')-int('a')+1) + int('a')
	return rune(code)

}

// Generate one Upper symbol
func getUpperSymbol(r *rand.Rand) rune {
	code := r.Intn(int('Z')-int('A')+1) + int('A')
	return rune(code)
}

// Generate one Special symbol
func getSpecialSymbol(r *rand.Rand) rune {
	code := r.Intn(41-33+1) + 33
	return rune(code)
}

// Generate one number and return rune
func getNumberSybol(r *rand.Rand) rune {
	code := r.Intn(10)
	return []rune(strconv.Itoa(code))[0]
}
