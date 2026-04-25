package model

import "testing"

func TestAnalyticsCacheHitRate(t *testing.T) {
	// total = 800 (fresh) + 200 (write) + 100 (read) = 1100; hit = 100/1100 ≈ 0.0909
	a := &Analytics{TotalTokensIn: 800, TotalCacheWrite: 200, TotalCacheRead: 100}
	rate := a.CacheHitRate()
	if rate < 0.090 || rate > 0.092 {
		t.Errorf("expected ~0.0909, got %f", rate)
	}

	// high cache-read session should not show 100%: total = 50+500+450 = 1000; hit = 450/1000 = 0.45
	a3 := &Analytics{TotalTokensIn: 50, TotalCacheWrite: 500, TotalCacheRead: 450}
	rate3 := a3.CacheHitRate()
	if rate3 < 0.449 || rate3 > 0.451 {
		t.Errorf("expected ~0.45, got %f", rate3)
	}

	a2 := &Analytics{}
	if a2.CacheHitRate() != 0 {
		t.Errorf("expected 0 for empty analytics")
	}
}
