package cmd

import (
	"fmt"
	"fuzztester-cli/gemini"
	"io"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
)

var (
	urlFlag  string
	bodyPath string
	Model    *genai.GenerativeModel
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "You can pass your testing endpoint URL with -a and request body file location with -b.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("URL:", urlFlag)
		fmt.Println("Body Path:", bodyPath)

		file, err := os.Open(bodyPath)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		byteBody, err := io.ReadAll(file)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		fmt.Println("READ:", string(byteBody))

		prompt := "I have this JSON data: " + string(byteBody) + ` give me fuzz testing options for me in one JSON.
GENERATE ALL POSSIBILITIES. In one remove name field, in two remove age field and LIKE THIS
Your example will be an array of objects same with this below:

{
    "description": "I changed this and this fields with these fields",
    "my json body content given you",
}

INSTEAD OF "my json body content given you", YOU MUST WRITE MY GIVEN JSON LIKE THIS EXAMPLE:
{
	"description": "added number to name field",
	"name": 999,
	"age": 18
}

AND GENERATE ALL POSSIBILITIES WITH GIVEN BODY
`

		response, err := gemini.GenerateJSON([]byte(prompt))
		if err != nil {
			log.Fatal(err)
		}

		pp.Println(response)
	},
}

func init() {
	testCmd.Flags().StringVarP(&urlFlag, "address", "a", "", "URL to fuzz test")
	testCmd.MarkFlagRequired("address")

	testCmd.Flags().StringVarP(&bodyPath, "body", "b", "", "JSON file inside request body")
	testCmd.MarkFlagRequired("body")
	rootCmd.AddCommand(testCmd)
}

func SetModel(m *genai.GenerativeModel) {
	Model = m
}
