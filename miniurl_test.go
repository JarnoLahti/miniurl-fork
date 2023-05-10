package miniurl_test

import (
	"fmt"
	"github.com/JarnoLahti/miniurl-fork"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashLength(t *testing.T) {
	const (
		input          = "test_input"
		expectedLength = 32
	)

	output := miniurl.Hash(input)
	assert.Len(t, output, expectedLength)
}

func FuzzHash(f *testing.F) {
	f.Add("https://example.com")
	f.Fuzz(func(t *testing.T, input string) {
		output1 := miniurl.Hash(input)
		output2 := miniurl.Hash(input)
		assert.Equal(t, output1, output2)
		assert.Len(t, output1, 32)
		assert.Len(t, output2, 32)
	})
}

func ExampleHash() {
	const input = "test_string"
	output := miniurl.Hash(input)
	fmt.Println(output)
	// output:
	// 3474851a3410906697ec77337df7aae4
}

func BenchmarkHash(b *testing.B) {
	const input = "test_string"
	for n := 0; n < b.N; n++ {
		miniurl.Hash(input)
	}
}
