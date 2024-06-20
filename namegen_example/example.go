package namegen_example

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/mtariq99/namegen/namegen"
)

// ExampleNewGenerator demonstrates how to create a new name generator with a seed.
func ExampleNewGenerator() {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	generator := namegen.NewGenerator(r.Int63())

	fmt.Println(generator.GetUniqueName())
	// Output: (example output, varies depending on seed and name list)
}

// ExampleNamegen_GetUniqueName demonstrates how to generate unique names using the name generator.
func ExampleNamegen_GetUniqueName() {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	generator := namegen.NewGenerator(r.Int63())

	name1 := generator.GetUniqueName()
	name2 := generator.GetUniqueName()

	fmt.Println(name1)
	fmt.Println(name2)
	// Output: (example output, varies depending on seed and name list)
}

