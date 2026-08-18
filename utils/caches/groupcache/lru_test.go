package groupcache_test

import (
	"testing"

	"github.com/imcrazytwkr/feedhub/utils/caches/groupcache"
)

func TestNewLRUValidatesSize(t *testing.T) {
	if _, err := groupcache.NewLRU[string, string](0); err != groupcache.ErrInvalidSize {
		t.Errorf("NewLRU(0) err = %v, want ErrInvalidSize", err)
	}
	if _, err := groupcache.NewLRU[string, string](-1); err != groupcache.ErrInvalidSize {
		t.Errorf("NewLRU(-1) err = %v, want ErrInvalidSize", err)
	}
	if _, err := groupcache.NewLRU[string, string](1); err != nil {
		t.Errorf("NewLRU(1) err = %v, want nil", err)
	}
}

func TestGetAddRoundTrip(t *testing.T) {
	c, err := groupcache.NewLRU[string, string](4)
	if err != nil {
		t.Fatalf("NewLRU(4) err = %v, want nil", err)
	}

	c.Add("a", "alpha")

	if got, ok := c.Get("a"); !ok || got != "alpha" {
		t.Errorf(`Get("a") = (%q, %v), want ("alpha", true)`, got, ok)
	}
	// a miss must return the zero value and ok=false
	if got, ok := c.Get("missing"); ok || got != "" {
		t.Errorf(`Get("missing") = (%q, %v), want ("", false)`, got, ok)
	}
	if got := c.Len(); got != 1 {
		t.Errorf("Len() = %d, want 1", got)
	}
}

func TestEvictsLeastRecentlyUsed(t *testing.T) {
	c, _ := groupcache.NewLRU[string, string](2)

	c.Add("a", "1")
	c.Add("b", "2")

	// touch "a"; "b" is now the least-recently-used entry
	if _, ok := c.Get("a"); !ok {
		t.Fatal(`Get("a") returned a miss`)
	}

	// inserting a third entry exceeds capacity and must evict the LRU entry ("b")
	c.Add("c", "3")

	if _, ok := c.Get("b"); ok {
		t.Error(`Get("b") ok = true; want false (should have been evicted as LRU)`)
	}
	if got, ok := c.Get("a"); !ok || got != "1" {
		t.Errorf(`Get("a") = (%q, %v), want ("1", true)`, got, ok)
	}
	if got, ok := c.Get("c"); !ok || got != "3" {
		t.Errorf(`Get("c") = (%q, %v), want ("3", true)`, got, ok)
	}
	if got := c.Len(); got != 2 {
		t.Errorf("Len() = %d, want 2 (bounded by capacity)", got)
	}
}

func TestRemoveKey(t *testing.T) {
	c, _ := groupcache.NewLRU[string, string](4)

	c.Add("a", "alpha")
	c.Remove("a")

	if _, ok := c.Get("a"); ok {
		t.Error(`Get("a") ok = true after Remove; want false`)
	}
	if got := c.Len(); got != 0 {
		t.Errorf("Len() = %d, want 0 after Remove", got)
	}

	// removing an absent key must be a no-op
	c.Add("b", "beta")
	c.Remove("missing")
	if got := c.Len(); got != 1 {
		t.Errorf("Len() = %d, want 1 (Remove of absent key must not change Len)", got)
	}
}

func TestRemoveOldest(t *testing.T) {
	c, _ := groupcache.NewLRU[string, string](4)

	c.Add("a", "1")
	c.Add("b", "2")

	// "a" is the least-recently-used (oldest) entry
	c.RemoveOldest()

	if _, ok := c.Get("a"); ok {
		t.Error(`Get("a") ok = true after RemoveOldest; want false (oldest evicted)`)
	}
	if got, ok := c.Get("b"); !ok || got != "2" {
		t.Errorf(`Get("b") = (%q, %v), want ("2", true)`, got, ok)
	}
	if got := c.Len(); got != 1 {
		t.Errorf("Len() = %d, want 1 after RemoveOldest", got)
	}
}

func TestClear(t *testing.T) {
	c, _ := groupcache.NewLRU[string, string](4)

	c.Add("a", "1")
	c.Add("b", "2")
	c.Clear()

	if got := c.Len(); got != 0 {
		t.Errorf("Len() = %d, want 0 after Clear", got)
	}
	if _, ok := c.Get("a"); ok {
		t.Error(`Get("a") ok = true after Clear; want false`)
	}
}
