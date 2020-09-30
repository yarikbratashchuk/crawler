package linkstorage

import (
	"testing"
	"time"
)

func TestInMemoryStorage(t *testing.T) {
	store := NewInMemory()

	go store.Add("link1", "link2", "link3")
	go store.Add("link1", "link2", "link4")

	time.Sleep(1 * time.Second)

	if l := len(store.links); l != 4 {
		t.Errorf("expected to have 4 links, but has %v", l)
	}

	for i := 0; i < 4; i++ {
		go func() {
			link, err := store.Get()
			if err != nil {
				t.Errorf("should not return: %v", err)
			}

			_ = store.Processed(link)
		}()
	}

	time.Sleep(1 * time.Second)

	_, err := store.Get()
	if err == nil {
		t.Error("should return ErrNoLinks")
	}

	for _, state := range store.links {
		if state != stateProcessed {
			t.Errorf("expected state: %v, got: %v", stateProcessed, state)
		}
	}
}
