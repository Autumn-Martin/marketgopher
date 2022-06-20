package cmd

import (
	"log"
	"marketgopher/aggregator"
	"marketgopher/output"
)

func Start() {
	err := aggregator.ComputeMarketData()
	if err != nil {
		log.Fatalf("%s %s", output.Red("Error:"), err)
	}
}
