package main
// Linux versions of this code need the xclip tool installed on Linux
import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func validateDepth(argTwo string) (n int) {
	n, err := strconv.Atoi(argTwo)
	if err == nil {
		return
	}
	testFloat, err := strconv.ParseFloat(argTwo, 64)
	if err == nil {
		n = int(math.Round(testFloat))
		return
	}
	n = 0
	return
}

func validateArgs(args []string) (master string, n int, result bool) {
	if len(args) == 2 {
		master = args[1]
		n = 1
		result = true
	} else if len(args) >= 3 && validateDepth(args[2]) > 0 {
		master = args[1]
		n = validateDepth(args[2])
		result = true
	} else if len(args) > 2 && validateDepth(args[2]) == 0 {
		master = "Invalid depth supplied! Program will terminate here!"
		n = 0
		result = false
	} else if len(args) == 2 {
		master = "No depth has been supplied! Program will terminate here!"
		n = 0
		result = false
	}
	return
}

func crunch(mstr string) (hashed string) {
	mstr = strings.TrimSpace(mstr)
	h := sha256.Sum256([]byte(mstr))
	hashed = fmt.Sprintf("%x\n", h)
	return
}

func copyToClipboard(content string) {
	cmd := exec.Command("xclip", "-selection", "clipboard")
	in, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer in.Close()
		_, err := io.WriteString(in, content)
		if err != nil {
			log.Fatal(err)
		}
	}()

	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	myArgs := os.Args
	if len(myArgs) == 1 {
		fmt.Print("Neither seed phrase nor depth nor flag has been supplied! Program will terminate here!\nFor help, run mastergen with one of these flags:\n", Help)
	} else if len(myArgs) > 1 {
		switch myArgs[1] {
		case "-d":
			fmt.Println(Doc)
		case "-v":
			fmt.Println(Version)
		case "-h":
			fmt.Println(Help)
		default:
			mstr, n, result := validateArgs(myArgs)
			if !result {
				fmt.Println(mstr)
			} else if result {
				for i := 1; i <= n; i++ {
					mstr = crunch(mstr)
				}
				finalOutput := string(mstr)
				fmt.Println(finalOutput)
				copyToClipboard(finalOutput)
			}
		}
	}
}
