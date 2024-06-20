package namegen

import (
	"nameGen/consts"
	"testing"
	"time"
)

func TestGetUniqueName(t *testing.T) {
	seed := time.Now().UnixNano()
	generator := NewGenerator(seed)

	// Generate a name and check if it's in the list
	name := generator.GetUniqueName()
	namesList := consts.ListNamesArr()

	found := false
	for _, n := range namesList {
		if n == name {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Generated name %s is not in the list", name)
	}
}

func TestDifferentSeedsGenerateDifferentNames(t *testing.T) {
	seed1 := time.Now().UnixNano()
	// Ensuring different seeds
	seed2 := seed1 + 1 

	generator1 := NewGenerator(seed1)
	generator2 := NewGenerator(seed2)

	name1 := generator1.GetUniqueName()
	name2 := generator2.GetUniqueName()

	if name1 == name2 {
		t.Errorf("Different seeds generated the same name: %s", name1)
	}
}

func TestSameSeedGeneratesSameNames(t *testing.T) {
	seed := time.Now().UnixNano()

	generator1 := NewGenerator(seed)
	generator2 := NewGenerator(seed)

	name1 := generator1.GetUniqueName()
	name2 := generator2.GetUniqueName()

	if name1 != name2 {
		t.Errorf("Same seed generated different names: %s and %s", name1, name2)
	}
}

func TestRandomnessWithinRange(t *testing.T) {
	seed := time.Now().UnixNano()
	generator := NewGenerator(seed)

	namesList := consts.ListNamesArr()

	// Generate multiple names to check the range
	for i := 0; i < 1000; i++ { 
		name := generator.GetUniqueName()
		found := false
		for _, n := range namesList {
			if n == name {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Generated name %s is not in the list", name)
		}
	}
}
