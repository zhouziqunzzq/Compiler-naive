package main

import (
	"fmt"
	"github.com/zhouziqunzzq/Compiler/GrammarAnalyzer/Expression/LL1Analyzer"
	"github.com/zhouziqunzzq/Compiler/GrammarAnalyzer/Expression/RecursiveAnalyzer"
	"github.com/zhouziqunzzq/Compiler/Scanner"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args[1:]) < 1 {
		fmt.Println("No input file specified")
		return
	}
	args := os.Args[1:]
	ft := 0
	if len(args) == 2 {
		if args[0] == "RA" || args[0] == "LL1" {
			ft = 1
		} else {
			fmt.Printf("Invalid argment \"%v\"\n", args[0])
			return
		}
	}

	buf, err := ioutil.ReadFile(args[ft])
	if err != nil {
		fmt.Println("Fail to open input file")
		return
	}
	// Add extra \n to the end of file to prevent scanner's bug
	content := string(buf) + "\n"
	switch args[0] {
	case "RA":
		fmt.Println("Performing grammar analysis using Recursive Analyzer...")
		testRecursiveAnalyzer(content)
	case "LL1":
		fmt.Println("Performing grammar analysis using LL(1) Analyzer...")
		testLL1Analyzer(content)
	default:
		fmt.Println("Performing lexical analysis...")
		testScanner(content)
	}
}

func testScanner(content string) {
	scanner := Scanner.NewScanner(&content)
	for scanner.LastToken.Type != Scanner.ERROR && scanner.LastToken.Type != Scanner.END {
		scanner.Next()
		if scanner.LastToken.Type == Scanner.ERROR {
			fmt.Printf("ERROR: Invalid token found at %v.\n", scanner.CurIndex)
			if len(scanner.Buffer) > 0 {
				fmt.Printf("Invalid token: %v\n", string(scanner.Buffer[len(scanner.Buffer)-1]))
			}
			fmt.Printf("DEBUG INFO: State: %v\n", scanner.State)
		} else if scanner.LastToken.Type == Scanner.COMMENT {
			fmt.Printf("Comment ignored.\n")
		} else if scanner.LastToken.Type == Scanner.END {
			fmt.Printf("END.\n")
			scanner.PrintTables()
		} else {
			fmt.Printf("%v - < %v  %v >\n",
				Scanner.TokenTypeName[scanner.LastToken.Type], scanner.LastToken.Word, scanner.LastToken.ID)
		}
	}
}

func testRecursiveAnalyzer(content string) {
	ra := RecursiveAnalyzer.NewRecursiveAnalyzer(&content)
	fmt.Println("==================================================")
	if ra.Analyze() == true {
		fmt.Println("Valid Expression.")
	} else {
		fmt.Println("Invalid Expression.")
		fmt.Printf("ERROR: Invalid token after \"%v\" found.\n", ra.S.LastToken.Word)
	}
}

func testLL1Analyzer(content string) {
	la := LL1Analyzer.NewLL1Analyzer(&content)
	fmt.Println("==================================================")
	if la.Analyze() == true {
		fmt.Println("Valid Expression.")
	} else {
		fmt.Println("Invalid Expression.")
		fmt.Printf("ERROR: Invalid token after \"%v\" found.\n", la.S.LastToken.Word)
		//fmt.Printf("LastTokenType: %v\n", Scanner.TokenTypeName[la.S.LastToken.Type])
		//fmt.Printf("StackEmpty: %v\n", la.AS.IsEmpty())
		/*fmt.Printf("Stack: ")
		for i := 0; i < la.AS.Top; i++ {
			fmt.Printf("%v ", la.AS.Stack[i])
		}*/
	}
}
