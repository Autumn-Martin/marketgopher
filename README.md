# MarketGopher

Welcome to MarketGopher, a CLI tool to compute aggregate market data.

MarketGopher listens to standard output for trades, parses each trade as it comes in, continuously computes aggregate metrics for each market, and outputs the final result.

Not financial advice. DYOR & use at your own risk.

### Getting Started
This app relies on another app, Stdoutinator, to be installed locally. Stdoutinator generates the 10 million trade objects as JSON that MarketGopher parses for its MVP.

Please verify that you have installed Stdoutinator if you see this error:
> Error: exec: "stdoutinator": executable file not found in $PATH

To install MarketGopher from within the cloned repo:
```sh
go install
```

To start MarketGopher, run:
```sh
marketgopher start
```

### Goals

MarketGopher was built to explore low-level performance optimizations. Future goals could include:
- benchmarking alternative low-level solutions for comparison
- testing
- live data streaming and processing
- continuous streaming beyond 10 million trades
- exploring higher-level services for streaming events, like Kafka
