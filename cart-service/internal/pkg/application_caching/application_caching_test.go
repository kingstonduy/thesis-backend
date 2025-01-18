package application_caching

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Define some test structs to use as values
type TestStruct1 struct {
	Name  string
	Value int
}

type TestStruct2 struct {
	ID    int
	Items []string
}

func TestCachingRepository(t *testing.T) {
	t.Run("Test Set and Get with TestStruct1", func(t *testing.T) {
		repo := NewClient()

		// Create a TestStruct1 value
		originalValue := TestStruct1{Name: "Test Name", Value: 42}

		// Set the value
		err := repo.Set(context.Background(), "key1", originalValue)
		assert.NoError(t, err)

		// Get the value back
		var retrievedValue TestStruct1
		err = repo.Get(context.Background(), "key1", &retrievedValue)
		assert.NoError(t, err)

		// Assert that the values are equal
		assert.Equal(t, originalValue, retrievedValue)
	})

	t.Run("Test Set and Get with TestStruct2", func(t *testing.T) {
		repo := NewClient()

		// Create a TestStruct2 value
		originalValue := TestStruct2{ID: 101, Items: []string{"Item1", "Item2", "Item3"}}

		// Set the value
		err := repo.Set(context.Background(), "key2", originalValue)
		assert.NoError(t, err)

		// Get the value back
		var retrievedValue TestStruct2
		err = repo.Get(context.Background(), "key2", &retrievedValue)
		assert.NoError(t, err)

		// Assert that the values are equal
		assert.Equal(t, originalValue, retrievedValue)
	})

	t.Run("Test KeyNotFound Error", func(t *testing.T) {
		repo := NewClient()

		// Attempt to get a non-existent key
		var retrievedValue TestStruct1
		err := repo.Get(context.Background(), "nonexistentKey", &retrievedValue)

		// Assert that the error is ErrKeyNotFound
		assert.Equal(t, ErrKeyNotFound, err)
	})

	t.Run("Test Set and Get with a simple type", func(t *testing.T) {
		repo := NewClient()

		// Set a simple value (string)
		originalValue := "Hello, world!"
		err := repo.Set(context.Background(), "key3", originalValue)
		assert.NoError(t, err)

		// Get the value back
		var retrievedValue string
		err = repo.Get(context.Background(), "key3", &retrievedValue)
		assert.NoError(t, err)

		// Assert that the values are equal
		assert.Equal(t, originalValue, retrievedValue)
	})

	t.Run("Test Set and Get with an int", func(t *testing.T) {
		repo := NewClient()

		// Set an int value
		originalValue := 123
		err := repo.Set(context.Background(), "key4", originalValue)
		assert.NoError(t, err)

		// Get the value back
		var retrievedValue int
		err = repo.Get(context.Background(), "key4", &retrievedValue)
		assert.NoError(t, err)

		// Assert that the values are equal
		assert.Equal(t, originalValue, retrievedValue)
	})

	t.Run("Test Set with error in Marshal", func(t *testing.T) {
		repo := NewClient()

		// Attempt to set an invalid value (e.g., a channel which can't be marshaled)
		invalidValue := make(chan int)
		err := repo.Set(context.Background(), "key5", invalidValue)

		// Assert that an error is returned
		assert.Error(t, err)
	})

	t.Run("Test Get with wrong type", func(t *testing.T) {
		repo := NewClient()

		// Set a string value
		originalValue := "Some value"
		err := repo.Set(context.Background(), "key6", originalValue)
		assert.NoError(t, err)

		// Try to get the value into an incorrect type (e.g., an int)
		var retrievedValue int
		err = repo.Get(context.Background(), "key6", &retrievedValue)

		// Assert that an error is returned due to type mismatch
		assert.Error(t, err)
	})
}
