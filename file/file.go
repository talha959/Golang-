package Files

type fileOperation struct {
	vowels, Puntutations, newlines, words, spaces float64
}

func Files(createChannel chan fileOperation, data string) {
	operation := fileOperation{}

	for _, char := range data {
		if string(char) == "a" || string(char) == "e" || string(char) == "i" || string(char) == "o" || string(char) == "u" ||
			string(char) == "A" || string(char) == "E" || string(char) == "I" || string(char) == "O" || string(char) == "U" {
			operation.vowels++
		}
		if string(char) == "!" || string(char) == "@" || string(char) == "#" || string(char) == "$" || string(char) == "%" ||
			string(char) == "^" || string(char) == "&" || string(char) == "*" || string(char) == "(" || string(char) == ")" ||
			string(char) == "-" || string(char) == "_" || string(char) == "+" || string(char) == "=" || string(char) == "{" ||
			string(char) == "}" || string(char) == "[" || string(char) == "]" || string(char) == "|" || string(char) == "\\" ||
			string(char) == ":" || string(char) == ";" || string(char) == "'" || string(char) == "<" || string(char) == ">" ||
			string(char) == "," || string(char) == "." || string(char) == "/" || string(char) == "?" {
			operation.Puntutations++
		}
		if string(char) == "\n" {
			operation.newlines++
		}
		if string(char) == " " {
			operation.spaces++
		}
		operation.words++
	}

	createChannel <- operation
}
