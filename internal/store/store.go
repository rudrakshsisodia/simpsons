package store

import (
	"sync"
	"time"

	"github.com/rudrakshsisodia/simpsons/internal/model"
)

// Store holds all session metadata in memory with concurrent-safe access.
type Store struct {
	mu        sync.RWMutex
	sessions  map[string]*model.SessionMeta   // UUID -> meta
	byProject map[string][]string             // project path -> []UUID
	details   map[string]*model.SessionDetail // UUID -> detail (lazy loaded)

	scanMu      sync.RWMutex
	scanScanned int
	scanTotal   int

	history *model.HistoryStats

	cachedAnalytics *model.Analytics
	analyticsDirty  bool
}

// New creates a new empty Store.
func New() *Store {
	return &Store{
		sessions:       make(map[string]*model.SessionMeta),
		byProject:      make(map[string][]string),
		details:        make(map[string]*model.SessionDetail),
		analyticsDirty: true,
	}
}

// Add inserts a session into the store.
func (s *Store) Add(meta *model.SessionMeta) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.sessions[meta.UUID] = meta
	// Only append if not already in byProject
	found := false
	for _, id := range s.byProject[meta.ProjectPath] {
		if id == meta.UUID {
			found = true
			break
		}
	}
	if !found {
		s.byProject[meta.ProjectPath] = append(s.byProject[meta.ProjectPath], meta.UUID)
	}
	s.analyticsDirty = true
}

// Get returns a session by UUID, or nil if not found.
func (s *Store) Get(uuid string) *model.SessionMeta {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.sessions[uuid]
}

// AllSessions returns all sessions.
func (s *Store) AllSessions() []*model.SessionMeta {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]*model.SessionMeta, 0, len(s.sessions))
	for _, meta := range s.sessions {
		result = append(result, meta)
	}
	return result
}

// Projects returns a list of unique project paths.
func (s *Store) Projects() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	projects := make([]string, 0, len(s.byProject))
	for p := range s.byProject {
		projects = append(projects, p)
	}
	return projects
}

// SessionsByProject returns all sessions for a given project path.
func (s *Store) SessionsByProject(project string) []*model.SessionMeta {
	s.mu.RLock()
	defer s.mu.RUnlock()

	uuids := s.byProject[project]
	result := make([]*model.SessionMeta, 0, len(uuids))
	for _, uuid := range uuids {
		if meta, ok := s.sessions[uuid]; ok {
			result = append(result, meta)
		}
	}
	return result
}

// Analytics computes aggregate analytics from all sessions.
func (s *Store) Analytics() *model.Analytics {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.analyticsDirty && s.cachedAnalytics != nil {
		return s.cachedAnalytics
	}

	a := &model.Analytics{
		ModelsUsed:     make(map[string]int),
		ToolsUsed:      make(map[string]int),
		SessionsByDate: make(map[string]int),
		CostByDate:     make(map[string]float64),
	}

	projectSet := make(map[string]bool)

	for _, meta := range s.sessions {
		a.TotalSessions++
		a.TotalTokensIn += meta.TokensIn
		a.TotalTokensOut += meta.TokensOut
		a.TotalCacheRead += meta.CacheRead
		a.TotalCacheWrite += meta.CacheWrite
		a.TotalCostUSD += meta.CostUSD
		projectSet[meta.ProjectPath] = true

		for m, count := range meta.Models {
			a.ModelsUsed[m] += count
		}
		for tool, count := range meta.ToolUsage {
			a.ToolsUsed[tool] += count

			// Work mode classification
			switch tool {
			case "Read", "Grep", "Glob", "WebFetch", "WebSearch", "LS", "SemanticSearch":
				a.WorkModeExplore += count
			case "Write", "Edit", "StrReplace":
				a.WorkModeBuild += count
			case "Bash", "Agent", "TaskCreate", "TaskUpdate":
				a.WorkModeTest += count
			}
		}

		if !meta.StartTime.IsZero() {
			date := meta.StartTime.Format("2006-01-02")
			a.SessionsByDate[date]++
			a.CostByDate[date] += meta.CostUSD
		}
	}

	a.ActiveProjects = len(projectSet)
	s.cachedAnalytics = a
	s.analyticsDirty = false
	return a
}

// TodaySpend returns the total estimated cost for sessions started today (cached).
func (s *Store) TodaySpend() float64 {
	today := time.Now().Format("2006-01-02")
	a := s.Analytics()
	return a.CostByDate[today]
}

// GetDetail returns a cached session detail by UUID, or nil if not cached.
func (s *Store) GetDetail(uuid string) *model.SessionDetail {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.details[uuid]
}

// SetDetail caches a session detail by UUID.
func (s *Store) SetDetail(uuid string, detail *model.SessionDetail) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.details[uuid] = detail
}

// SetScanProgress updates the background scan progress.
func (s *Store) SetScanProgress(scanned, total int) {
	s.scanMu.Lock()
	defer s.scanMu.Unlock()
	s.scanScanned = scanned
	s.scanTotal = total
}

// ScanProgress returns the current scan progress.
func (s *Store) ScanProgress() (scanned, total int) {
	s.scanMu.RLock()
	defer s.scanMu.RUnlock()
	return s.scanScanned, s.scanTotal
}

// SetHistoryStats stores the computed history statistics.
func (s *Store) SetHistoryStats(stats *model.HistoryStats) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.history = stats
}

// HistoryStats returns the stored history statistics, or nil if not yet loaded.
func (s *Store) HistoryStats() *model.HistoryStats {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.history
}
