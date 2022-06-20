package aggregator

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"marketgopher/calc"
	"marketgopher/models"
	"marketgopher/output"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func ComputeMarketData() error {
	cmd := exec.Command("stdoutinator")

	// Capture end signal & results so we can print even when the program ends early
	var markets map[int]models.Market
	earlySignal := make(chan os.Signal, 1)
	signal.Notify(earlySignal, syscall.SIGINT, syscall.SIGTERM)

	// Redirect stdoutinator's trade output and errors for processing when it runs
	tradesPipe, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout

	if err != nil {
		return err
	}

	// Create a channel for processed data
	marketsChannel := make(chan map[int]models.Market)

	// Create a scanner that scans stdout line-by-line
	scanner := bufio.NewScanner(tradesPipe)

	// Run scanner in a goroutine
	go func() {
		markets := make(map[int]models.Market)

		for scanner.Scan() {
			// Avoid scanning output that is not valid json, like comments
			if json.Valid(scanner.Bytes()) {
				var trade models.Trade
				err := json.Unmarshal(scanner.Bytes(), &trade)

				if err != nil {
					log.Fatalf("%s %s", output.Red("Error:"), err)
				} else {
					fmt.Printf("Processing trade: %d \n", trade.ID)
					// Copy non-changed values & update continuously
					market := markets[trade.MarketID]
					markets[trade.MarketID] = processMarketUpdate(market, trade)
				}

			}

		}
		marketsChannel <- markets

	}()

	fmt.Println("Aggregating stdoutinator's market data...")

	// Run stdoutinator
	err = cmd.Start()
	if err != nil {
		return err
	}

	// Wait for signal that processing is done
	markets = <-marketsChannel

	go func() {
		<-earlySignal
		// Go ahead and print what we have so far with a warning
		err := printMarketData(markets)
		if err != nil {
			log.Fatalf("%s %s", output.Red("Error:"), err)
		}

		log.Fatal(output.Red("Warning: Program exited early. Processing may not be complete."))
	}()

	// Close the channel
	close(marketsChannel)

	// Print results
	err = printMarketData(markets)
	if err != nil {
		return err
	}

	return nil
}

func processMarketUpdate(market models.Market, trade models.Trade) models.Market {
	// Add trade to market data
	market.ID = trade.MarketID
	market.TotalTrades++
	market.TotalPrice = market.TotalPrice + trade.Price
	market.TotalVolume = market.TotalVolume + trade.Volume

	if trade.IsBuy {
		market.TotalBuys++
	}

	market.PercentBuys = calc.PercentBuys(market.TotalBuys, market.TotalTrades)
	market.AvgPrice = calc.AvgPrice(market.TotalPrice, market.TotalVolume)
	market.AvgVolume = calc.AvgVolume(market.TotalVolume, market.TotalTrades)
	market.VolumeWeightedAvgPrice = calc.VolumeWeightedAvgPrice(market.AvgPrice, market.TotalVolume)

	return market
}

func printMarketData(marketMap map[int]models.Market) error {
	for _, value := range marketMap {
		market, err := json.Marshal(value)

		if err != nil {
			return err
		}

		fmt.Println(string(market))
	}
	return nil
}
