package worker

import (
	"sync"
	"testing"
)

func TestSpawnWorkers(t *testing.T) {
	count := 10
	var wg sync.WaitGroup
	ch1, ch2 := SpawnWorkers(count, &wg, 1)

	s1 := cap(ch1)
	if cap(ch1) != count {
		t.Errorf("expected chanel 1 size: %d, got: %d", count, s1)
	}

	s2 := cap(ch2)
	if cap(ch2) != count {
		t.Errorf("expected chanel 2 size: %d, got: %d", count, s2)
	}

	close(ch1)
	close(ch2)
}
