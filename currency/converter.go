package currency

import "fmt"

type Converter struct {
	service RateService
}

func NewConverter() *Converter {

	return &Converter{
		service: &DefaultRateService{},
	}

}

func (c *Converter) GetRate(base string, target string) float64 {

	rate, err := c.service.GetExchangeRate(base, target)
	if err != nil {
		fmt.Printf("error getting exchange rate from: %s to: %s: %s\n", base, target, err)
		return 0
	}

	return rate
}

func (c *Converter) Convert(base string, target string, amount float64) float64 {

	rate, err := c.service.GetExchangeRate(base, target)
	if err != nil {
		fmt.Printf("error getting exchange rate from: %s to: %s: %s\n", base, target, err)
		return 0
	}

	return amount * rate
}
