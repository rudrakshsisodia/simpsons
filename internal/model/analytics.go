package model

// Analytics holds aggregated analytics computed from all sessions.
type Analytics struct {
	TotalSessions    int
	TotalTokensIn    int64
	TotalTokensOut   int64
	TotalCacheRead   int64
	TotalCacheWrite  int64
	ActiveProjects   int
	ModelsUsed       map[string]int // model → total message count
	ToolsUsed        map[string]int // tool → total call count
	SessionsByDate   map[string]int // "2026-03-03" → count
	WorkModeExplore  int            // Read, Grep, Glob, WebFetch, WebSearch calls
	WorkModeBuild    int            // Write, Edit calls
	WorkModeTest     int            // Bash, Agent, Task calls
	TotalCostUSD     float64            // estimated total cost across all sessions
	CostByDate       map[string]float64 // "2026-04-25" → cost USD
}

// CacheHitRate returns the fraction of all input tokens served from cache (0.0–1.0).
// denominator = fresh input + cache writes + cache reads (total tokens sent to model context).
func (a *Analytics) CacheHitRate() float64 {
	total := a.TotalTokensIn + a.TotalCacheWrite + a.TotalCacheRead
	if total == 0 {
		return 0
	}
	return float64(a.TotalCacheRead) / float64(total)
}
