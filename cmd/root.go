package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/spf13/cobra"
)

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

var rootCmd = &cobra.Command{
	Use:   "main",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {
		routine, _ := cmd.Flags().GetInt("routine")
		starttime := time.Now()
		path, _ := cmd.Flags().GetString("file")
		file, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Println("Error in opening file:", err)
			return
		}

		fmt.Println(string(file))

		createChannel := make(chan fileOperation)
		// for i := 0; i < routine; i++ {
		// 	go Files(createChannel, string(file[:len(file)/routine]))
		// 	go Files(createChannel, string(file[len(file)/routine:]))
		// }
		for i := 0; i < routine; i++ {
			start := i * len(file) / routine
			end := (i + 1) * len(file) / routine

			go Files(createChannel, string(file[start:end]))
		}

		vowel := float64(0)
		punctutation := float64(0)
		newline := float64(0)
		word := float64(0)
		space := float64(0)

		for i := 0; i < routine; i++ {
			result := <-createChannel
			vowel += result.vowels
			punctutation += result.Puntutations
			newline += result.newlines
			word += result.words
			space += result.spaces
		}
		excetiontime := time.Since(starttime)
		fmt.Println("vowels are ", vowel)
		fmt.Println("Puntutations are ", punctutation)
		fmt.Println("new lines are ", newline)
		fmt.Println("number of words are ", word)
		fmt.Println("number of free spaces are ", space)
		fmt.Println("Exceutiontime", excetiontime)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntP("routine", "r", 2, "Go routines")
	rootCmd.Flags().StringP("file", "f", "task.txt", "filesname")

}
