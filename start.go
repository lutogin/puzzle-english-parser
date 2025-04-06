package main

import (
	"fmt"
	peClient "go-pe-parser/src/client"
	"go-pe-parser/src/config"
	outputCsv "go-pe-parser/src/output"
	peParser "go-pe-parser/src/parser"
	utils "go-pe-parser/src/utils"
	"log"
	"os"
	"strconv"
)

func start(cfg *config.Config) {
	// Get cookies from user input
	fmt.Println("Please enter your wp_logged_in_cookie cookies string [example: wp_logged_in_cookie=some@email.com...]:")

	cookies, err := utils.GetPrompt()
	if err != nil {
		log.Fatal(err)
	}

	if err := utils.ValidateCookies(cookies); err != nil {
		log.Fatal(err)
	}

	// Create new client and make request
	client, err := peClient.NewPuzzleEnglishClient(cookies, cfg)
	if err != nil {
		log.Fatal("Error initializing client:", err)
	}

	parser := peParser.NewHTMLParser()
	fmt.Println("Starting to process pages...")

	wordsPerPage, err := strconv.Atoi(cfg.APP.Config["wordsPerPage"])
	if err != nil {
		log.Fatal("Error parsing wordsPerPage:", err)
	}

	for i := 0; true; i++ {
		utils.ShowProgress(i)

		// Get the HTML content
		htmlContent, err := client.GetDictionaryPage(i)
		if err != nil {
			log.Fatal("Error getting dictionary page:", err)
		}

		// Parse the HTML content
		words, err := parser.ParseDictionaryPage(htmlContent)
		if err != nil {
			log.Fatal("Error parsing HTML:", err)
		}

		// Save to CSV
		err = outputCsv.SaveToCSV(words, cfg)
		if err != nil {
			log.Fatal("Error saving to CSV:", err)
		}

		// Check if there are no more words
		if len(words) < wordsPerPage {
			fmt.Println("Completed processing all pages!")
			break
		}
	}

	workDir, err := os.Getwd()
	fmt.Println(fmt.Sprintf("Data has been saved to %s/words.csv", workDir))

	separator := cfg.APP.Config["csvSeparator"]
	fmt.Println(fmt.Sprintf("Separator for CSV: \"%s\"", separator))

	utils.WaitForKeyPress()
}
