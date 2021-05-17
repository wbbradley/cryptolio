package show

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/logrusorgru/aurora"
	"github.com/shopspring/decimal"
)

func Symbol(s string) string {
	return aurora.Magenta(s).String()
}

func Cash(amount float64) string {
	if amount < 0 {
		return aurora.Red(CashNoColor(amount)).String()
	} else {
		return aurora.Green(CashNoColor(amount)).String()
	}
}

func CashNoColor(amount float64) string {
	return fmt.Sprintf("$%0.2f", amount)
}

func CashDelta(amount float64) string {
	if amount < 0 {
		return aurora.Red(CashDeltaNoColor(amount)).String()
	} else {
		return aurora.Green(CashDeltaNoColor(amount)).String()
	}
}

func CashDeltaNoColor(amount float64) string {
	if amount < 0 {
		return fmt.Sprintf("-$%0.2f", -amount)
	} else {
		return fmt.Sprintf("+$%0.2f", amount)
	}
}

func CashDecimal(amount decimal.Decimal) string {
	a, _ := amount.Float64()
	return Cash(a)
}

func Cash32(amount float32) string {
	if amount < 0 {
		return aurora.Red(fmt.Sprintf("$%0.2f", amount)).String()
	} else {
		return aurora.Green(fmt.Sprintf("$%0.2f", amount)).String()
	}
}

func Percent(amount float64) string {
	return aurora.Yellow(fmt.Sprintf("%0.2f%%", 100.0*amount)).String()
}

func PercentDecimal(amount decimal.Decimal) string {
	a, _ := amount.Float64()
	return aurora.Yellow(fmt.Sprintf("%.03f%%", 100.0*a)).String()
}

func Percent32(amount float32) string {
	return aurora.Yellow(fmt.Sprintf("%0.2f%%", 100.0*amount)).String()
}

func Shares(s float64) string {
	return aurora.BrightBlue(decimal.NewFromFloat(s).String()).String()
}

func SharesDecimal(s decimal.Decimal) string {
	return aurora.BrightBlue(s.String()).String()
}

func Int32(i int32) string {
	return aurora.BrightBlue(fmt.Sprintf("%d", i)).String()
}

func Int64(i int64) string {
	return aurora.BrightBlue(fmt.Sprintf("%d", i)).String()
}

func Bool(b bool) string {
	if b {
		return "true"
	} else {
		return "false"
	}
}

func Json(o interface{}) string {
	b, _ := json.Marshal(o)
	return fmt.Sprintf("%s", string(b))
}

func Date(t time.Time) string {
	return t.Local().Format("2006-01-02")
}

func RFC3339(t time.Time) string {
	return t.Local().Format(time.RFC3339)
}
