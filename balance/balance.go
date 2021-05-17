package balance

import (
	"cryptolio/cli"
	"cryptolio/show"
	"fmt"
	"os"
	"strings"

	cmc "github.com/miguelmota/go-coinmarketcap/pro/v1"
)

var (
	equity         *float64
	numComponents  *int
	excludeListCsv *string
)

func init() {
	cmd := cli.Register("balance", "Construct a portfolio of the highest market cap cryptocurrencies", makePortfolio)
	equity = cmd.Arg("equity", "Total equity to use in portfolio.").Required().Float64()
	numComponents = cmd.Flag("num-components", "Number of components in portfolioi.").Short('n').Default("10").Int()
	excludeListCsv = cmd.Flag("exclude", "Coins to exclude (use base currency symbol, like USDT,DOGE).").Default("USDT,DOGE").String()
}

type Component struct {
	Symbol    string
	Name      string
	PriceUSD  float64
	MarketCap float64
}

func buildExcludeList(excludeListCsv string) map[string]bool {
	m := map[string]bool{}
	xs := strings.Split(strings.ToUpper(excludeListCsv), ",")
	for _, x := range xs {
		m[x] = true
	}
	return m
}

func makePortfolio() {
	client := cmc.NewClient(&cmc.Config{
		// You can get a key at https://coinmarketcap.com/api/
		ProAPIKey: os.Getenv("CMC_PRO_API_KEY"),
	})

	listings, err := client.Cryptocurrency.LatestListings(&cmc.ListingOptions{
		Limit: *numComponents * 2,
	})
	if err != nil {
		panic(err)
	}
	excludeList := buildExcludeList(*excludeListCsv)
	totalMarketCap := 0.0
	components := []Component{}
	for _, listing := range listings {
		if len(components) == *numComponents {
			break
		}
		if excludeList[listing.Symbol] {
			continue
		}
		components = append(components, Component{
			Symbol:    listing.Symbol,
			Name:      listing.Name,
			PriceUSD:  listing.Quote["USD"].Price,
			MarketCap: listing.Quote["USD"].MarketCap,
		})
		totalMarketCap += listing.Quote["USD"].MarketCap
	}
	fmt.Printf("You should take your %s and divide it as follows\n", show.Cash(*equity))
	chunks := []string{}
	for _, c := range components {
		fraction := c.MarketCap / totalMarketCap
		dollars := fraction * *equity
		coins := dollars / c.PriceUSD
		fmt.Printf("%s(%s): Spend %s on %0.6f %s.\n", show.Symbol(c.Symbol), show.Percent(fraction), show.Cash(dollars), coins, c.Name)
		chunks = append(chunks, fmt.Sprintf("%s%s*%.06f", c.Symbol, "USD", coins))
	}
	fmt.Printf("TradingView: %s\n", strings.Join(chunks, "+"))
}
