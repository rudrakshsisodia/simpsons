package model

import (
	"fmt"
	"strings"
)

// modelPricing holds per-million-token prices for a Claude model.
type modelPricing struct {
	InputPerM       float64 // USD per 1M input tokens
	OutputPerM      float64 // USD per 1M output tokens
	CacheWritePerM  float64 // USD per 1M cache-write tokens (typically 1.25x input)
	CacheReadPerM   float64 // USD per 1M cache-read tokens (typically 0.1x input)
}

// pricingTable maps model name prefixes to their pricing.
// Prices are in USD per million tokens.
var pricingTable = []struct {
	prefix  string
	pricing modelPricing
}{
	// Claude 4 family
	{"claude-opus-4", modelPricing{15.00, 75.00, 18.75, 1.50}},
	{"claude-sonnet-4", modelPricing{3.00, 15.00, 3.75, 0.30}},
	{"claude-haiku-4", modelPricing{0.80, 4.00, 1.00, 0.08}},
	// Claude 3.5 family
	{"claude-3-5-sonnet", modelPricing{3.00, 15.00, 3.75, 0.30}},
	{"claude-3-5-haiku", modelPricing{0.80, 4.00, 1.00, 0.08}},
	// Claude 3 family
	{"claude-3-opus", modelPricing{15.00, 75.00, 18.75, 1.50}},
	{"claude-3-sonnet", modelPricing{3.00, 15.00, 3.75, 0.30}},
	{"claude-3-haiku", modelPricing{0.25, 1.25, 0.31, 0.03}},
}

// lookupPricing returns the pricing for a given model name.
// Falls back to Sonnet pricing if the model is unknown.
func lookupPricing(model string) modelPricing {
	lower := strings.ToLower(model)
	for _, entry := range pricingTable {
		if strings.Contains(lower, entry.prefix) {
			return entry.pricing
		}
	}
	// Default to Sonnet pricing for unknown models.
	return modelPricing{3.00, 15.00, 3.75, 0.30}
}

// ComputeCost returns the estimated USD cost for a single assistant message.
func ComputeCost(model string, inputTokens, outputTokens, cacheWrite, cacheRead int) float64 {
	p := lookupPricing(model)
	const perM = 1_000_000.0
	return float64(inputTokens)*p.InputPerM/perM +
		float64(outputTokens)*p.OutputPerM/perM +
		float64(cacheWrite)*p.CacheWritePerM/perM +
		float64(cacheRead)*p.CacheReadPerM/perM
}

// FormatCost formats a USD cost value for display.
// Shows cents for small amounts, dollars for larger ones.
func FormatCost(usd float64) string {
	if usd == 0 {
		return "-"
	}
	if usd < 0.01 {
		return "<$0.01"
	}
	if usd < 1.00 {
		return fmt.Sprintf("$%.2f", usd)
	}
	return fmt.Sprintf("$%.2f", usd)
}
