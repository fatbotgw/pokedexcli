package pokecache

import (
    // "fmt"
    "testing"
    "time"
)

func TestAddGet(t *testing.T) {
    const interval = 5 * time.Millisecond

    cache := NewCache(interval)

    key := "https://example.com"
    val := []byte("testdata")

    cache.Add(key, val)
    got, ok := cache.Get(key)

    if !ok {
        t.Errorf("expected to find key")
        return
    }

    if string(got) != string(val) {
        t.Errorf("expected %q, got %q", string(val), string(got))
    }
}

func TestReapLoop(t *testing.T) {
    const interval = 5 * time.Millisecond

    cache := NewCache(interval)

    key := "https://example.com"
    val := []byte("testdata")

    cache.Add(key, val)
    got, ok := cache.Get(key)

    if !ok {
        t.Errorf("expected to find key")
        return
    }

    if string(got) != string(val) {
        t.Errorf("expected %q, got %q", string(val), string(got))
    }

    time.Sleep(15 * time.Millisecond)

    got, ok = cache.Get(key)

    if ok {
        t.Errorf("didn't expect to find key, should have been reaped")
        return
    }
}