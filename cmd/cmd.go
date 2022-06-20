package cmd

import (
	"fmt"
	"log"
	"marketgopher/output"
	"os"
)

const (
	Welcome string = `Welcome to MarketGopher, a CLI tool to compute aggregate market data.

MarketGopher listens to standard output for trades, parses each trade as it comes in, 
continuously computes aggregate metrics for each market, and outputs the final result.

Not financial advice. DYOR & use at your own risk.` + "\n\n" + Help

	Help string = `Usage: marketgopher [command]
Available commands:
  start  Start computing aggregate market data for stdout trades`
)

func Execute() {
	log.SetFlags(0) // Excludes timestamps

	numArgs := len(os.Args) - 1

	switch numArgs {
	case 0:
		fmt.Println(Welcome)
	case 1:
		directRequest()
	case 2:
		fmt.Println(Help)
	}
}

func directRequest() {
	request := os.Args[1]

	switch request {
	case "-h", "--help":
		fmt.Println(Help)
	case "start":
		Start()
	default:
		log.Fatal(output.Yellow(`Unsupported command.`) + "\n" + Help)
	}
}
