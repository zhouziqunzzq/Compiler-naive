package main

import (
	"fmt"
	"github.com/zhouziqunzzq/Compiler/Scanner"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args[1:]) < 1 {
		fmt.Println("No input file specified!")
		return
	}
	args := os.Args[1:]
	buf, err := ioutil.ReadFile(args[0])
	if err != nil {
		fmt.Println("Fail to open input file!")
		return
	}
	content := string(buf)
	scanner := Scanner.NewScanner(&content)
	/*str := scanner.GetContent()
	fmt.Print(string(str))*/
	for scanner.LastToken.Type != Scanner.ERROR && scanner.LastToken.Type != Scanner.END {
		scanner.Next()
		if scanner.LastToken.Type == Scanner.ERROR {
			fmt.Printf("ERROR: Invalid token found at %v.\n", scanner.CurIndex)
		} else if scanner.LastToken.Type == Scanner.END {
			fmt.Printf("END.\n")
			scanner.PrintTables()
		} else {
			fmt.Printf("%v - < %v  %v >\n",
				Scanner.TokenTypeName[scanner.LastToken.Type], scanner.LastToken.Word, scanner.LastToken.ID)
		}
	}
}
