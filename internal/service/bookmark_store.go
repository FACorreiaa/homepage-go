package service

import (
	"io/fs"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

// BookmarkStore serves a BookmarkIndex and transparently rebuilds it in the
// background once it is older than the TTL, so vault files synced onto the
// server (rsync timer) become visible without a restart. Requests always get
// the current index immediately — never blocked on a rescan.
type BookmarkStore struct {
	vaultFS    fs.FS
	ttl        time.Duration
	mu         sync.RWMutex
	idx        BookmarkIndex
	refreshing atomic.Bool
}

func NewBookmarkStore(vaultFS fs.FS) *BookmarkStore {
	ttl := 5 * time.Minute
	if v := os.Getenv("BOOKMARK_INDEX_TTL"); v != "" {
		if d, err := time.ParseDuration(v); err == nil && d > 0 {
			ttl = d
		} else if secs, err := strconv.Atoi(v); err == nil && secs > 0 {
			ttl = time.Duration(secs) * time.Second
		}
	}
	return &BookmarkStore{
		vaultFS: vaultFS,
		ttl:     ttl,
		idx:     BuildBookmarkIndex(vaultFS),
	}
}

// Index returns the current index, kicking off a background rebuild if stale.
func (s *BookmarkStore) Index() BookmarkIndex {
	s.mu.RLock()
	idx := s.idx
	stale := time.Since(idx.BuiltAt) > s.ttl
	s.mu.RUnlock()

	if stale && s.refreshing.CompareAndSwap(false, true) {
		go func() {
			defer s.refreshing.Store(false)
			fresh := BuildBookmarkIndex(s.vaultFS)
			s.mu.Lock()
			s.idx = fresh
			s.mu.Unlock()
		}()
	}
	return idx
}
