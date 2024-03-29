package main

import (
	"crypto/sha256"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

//func to validate that the 3rd argument submitted by the user is in fact a valid int.
// Takes argTwo, a string. If the string is an int, it converts to type int, sets n to
// it, and returns n. If the string is a float, it rounds it, converts it to an int
// and then sets n to it and returns it.
func validateDepth(argTwo string) (n int) {
	// if argTwo is an int, return n
	n, err := strconv.Atoi(argTwo)
	if err == nil {
		return
	} 
	// if argTwo is not an int, test to see if it is a float
	testFloat, err := strconv.ParseFloat(argTwo, 64)
	if err == nil {
		// if there is no error, round the float and turn it into an int and return it
		n = int(math.Round(testFloat))
		return
	}
	// if it neither an int or a float, it's test, so set n to 0 and return that
	n = 0
	return
}

// func to validate input. Takes the OS args and returns the master phrase string, the 
// depth as an int, and a result bool.
func validateArgs(args []string) (master string, n int, result bool ) {
	// the valid return consists of len(args) >= 3, with a valid int in args[2]
	if len(args) == 2 {
		master = args[1]
		n = 1
		result = true
	} else if len(args) >= 3 && validateDepth(args[2]) > 0 {
		master = args[1]
		n = validateDepth(args[2])
		result = true
		// if there is no valid int in args[2]
		} else if len(args) > 2 && validateDepth(args[2]) == 0 {
			master = "Invalid depth supplied! Program will terminate here!"
			n = 0
			result = false
			// if there is no in/valid depth in args[2] supplied
			} else if len(args) == 2 {
				master = "No depth has been supplied! Program will terminate here!"
				n = 0
				result = false
				} 
	return
}


func crunch(mstr string) (hashed string) {
	// trimspace from the string or func will hash it with the whitespace
	mstr = strings.TrimSpace(mstr)
	// run the sha256.Sum256() function of mstr string turned into a []array
	h := sha256.Sum256([]byte(mstr))
	// return a string literal of the hashed string without the initial 0x
	hashed = fmt.Sprintf("%x\n", h)
	return
}

func main() {
	// take arguments as an array from the OS
	myArgs := os.Args
	// check the length of the array. If the array is empty, inform the user and exit.
	if len(myArgs) == 1 {
		fmt.Print("Neither seed phrase nor depth nor flag has been supplied! Program will terminate here!\nFor help, run mastergen with one of these flags:\n", Help)
	// if the array is not empty, check for flags from doc.go package
	} else if len(myArgs) > 1 {
		switch myArgs[1] {
		case "-d":
			fmt.Println(Doc)
		case "-v":
			fmt.Println(Version)
		case "-h":
			fmt.Println(Help)
		// if the array is not empty but there are no flags, validate the input
		default:
			mstr, n, result := validateArgs(myArgs)
			// if result is false, an error has occurred, inform the user
			if !result {
				fmt.Println(mstr)
			} else if result {
				// else if the result if valid, get the hash of your inputs using a 
				// for-loop from 1 to <= n.
					for i := 1; i <= n; i++ {
						// keep crunch(mstr)
					mstr = crunch(mstr)	
				}
				// print out the stringified final output of crunch(master)
				fmt.Println(string(mstr))
			}
		}
	}
}