package bf

import (
	"fmt"
	"io/ioutil"
	"os"
)

func CompileString(input string) string {
	ret := "#include <stdio.h>\nint main() {\nchar array[30000] = {0}; char *ptr = array;\n"
	for _, v := range input {
		switch string(v) {
		case "+":
			ret += "++*ptr;\n"
		case "-":
			ret += "--*ptr;\n"
		case "<":
			ret += "--ptr;\n"
		case ">":
			ret += "++ptr;\n"
		case ".":
			ret += "putchar(*ptr);\n"
		case ",":
			ret += "*ptr = getchar();\n"
		case "[":
			ret += "while (*ptr) {\n"
		case "]":
			ret += "}\n"
		}
	}

	ret += "return 0; \n}"
	return ret
}

func CompileFile(inputFile string) {
	fileBytes, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading %v: %v\n", inputFile, err)
		os.Exit(1)
	}
	str := string(fileBytes)

	cStr := CompileString(str)

	fmt.Println(cStr)

}
