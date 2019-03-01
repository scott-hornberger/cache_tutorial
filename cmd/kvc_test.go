package cmd

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/spf13/cobra"
)

func TestKvcCommandRunner(t *testing.T) {
	runner := &kvcCommandRunner{
		cache: &MockKVCache{},
	}
	testCmd := &cobra.Command{Use: "kvc"}

	t.Run("it can handle 'put <key> <value>'", func (t *testing.T) {
		args := []string{"test-key", "test-value"}
		err := runner.Put(testCmd, args)
		assert.NoError(t, err)
	})
}
// Mock cache implementation to use with testing
type MockKVCache struct {}
func (mock *MockKVCache) Put(k,v string) error { return nil }
func (mock *MockKVCache) Read(k string) (string, error) { return "", nil }
func (mock *MockKVCache) Update(k,v string) error {	return nil }
func (mock *MockKVCache) Delete(k string) error { return nil }