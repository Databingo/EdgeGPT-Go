package cmd

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"github.com/pavel-one/EdgeGPT-Go"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
)

var gpt *EdgeGPT.GPT

var rootCmd = &cobra.Command{
	Use:   "EdgeGPT-Go",
	Short: "CLI for using edge bing",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		gpt = newChat()
		reader := bufio.NewReader(os.Stdin)

		color.Green("Hello, I am a chatbot for speak with edge bing!")

		for {
			fmt.Print("You:\n\t")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			if input == "exit" || input == "q" || input == "quiet" {
				color.Yellow("Goodbye!")
				break
			}

			ask(input)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func ask(input string) {
	mw, err := gpt.AskAsync(input)
	if err != nil {
		log.Fatalln(err)
	}

	go mw.Worker()

	for range mw.Chan {
	}

	fmt.Print("Bot:\n\t" + mw.Answer.GetAnswer())
}

func newChat() *EdgeGPT.GPT {
	s := EdgeGPT.NewStorage()

	gpt, err := s.GetOrSet("cli")
	if err != nil {
		log.Fatalln(err)
	}

	return gpt
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
