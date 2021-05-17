package balance

import (
	"cryptolio/cli"
	"cryptolio/show"
	"fmt"
	"os"

	cmc "github.com/miguelmota/go-coinmarketcap/pro/v1"
)

var (
	equity *float64
)

func init() {
	cmd := cli.Register("balance", "Construct a portfolio of the highest market cap cryptocurrencies", makePortfolio)
	equity = cmd.Arg("equity", "Total equity to use in portfolio.").Required().Float64()
}

type Component struct {
	Symbol    string
	Name      string
	PriceUSD  float64
	MarketCap float64
}

func makePortfolio() {
	client := cmc.NewClient(&cmc.Config{
		// You can get a key at https://coinmarketcap.com/api/
		ProAPIKey: os.Getenv("CMC_PRO_API_KEY"),
	})

	listings, err := client.Cryptocurrency.LatestListings(&cmc.ListingOptions{
		Limit: 10,
	})
	if err != nil {
		panic(err)
	}

	totalMarketCap := 0.0
	components := []Component{}
	for _, listing := range listings {
		components = append(components, Component{
			Symbol:    listing.Symbol,
			Name:      listing.Name,
			PriceUSD:  listing.Quote["USD"].Price,
			MarketCap: listing.Quote["USD"].MarketCap,
		})
		totalMarketCap += listing.Quote["USD"].MarketCap
	}
	fmt.Printf("You should take your %s and divide it as follows\n", show.Cash(*equity))
	for _, c := range components {
		fraction := c.MarketCap / totalMarketCap
		dollars := fraction * *equity
		coins := dollars / c.PriceUSD
		fmt.Printf("%s: Spend %s on %0.6f %s.\n", show.Symbol(c.Symbol), show.Cash(dollars), coins, c.Name)
	}
}
