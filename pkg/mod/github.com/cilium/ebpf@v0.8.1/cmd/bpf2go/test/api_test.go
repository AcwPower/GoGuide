package test

import (
	"testing"

	// Raise RLIMIT_MEMLOCK
	_ "github.com/cilium/ebpf/internal/testutils"
)

func TestLoadingSpec(t *testing.T) {
	spec, err := loadTest()
	if err != nil {
		t.Fatal("Can't load spec:", err)
	}

	if spec == nil {
		t.Fatal("Got a nil spec")
	}
}

func TestLoadingObjects(t *testing.T) {
	var objs testObjects
	if err := loadTestObjects(&objs, nil); err != nil {
		t.Fatal("Can't load objects:", err)
	}
	defer objs.Close()

	if objs.Filter == nil {
		t.Error("Loading returns an object with nil programs")
	}

	if objs.Map1 == nil {
		t.Error("Loading returns an object with nil maps")
	}
}
