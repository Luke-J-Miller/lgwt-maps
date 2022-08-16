package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertString(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		assertError(t, err, ErrNotFound)
	})

}
func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := make(Dictionary)
		key := "test"
		value := "this is just a test"
		err := dictionary.Add(key, value)
		assertError(t, err, nil)
		assertDefinitiion(t, dictionary, key, value)
	})
	t.Run("existing word", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		dictionary := make(Dictionary)
		dictionary[key] = value
		err := dictionary.Add(key, value)
		assertError(t, err, ErrWordExists)
		assertDefinitiion(t, dictionary, key, value)
	})
}
func TestUpdate(t *testing.T) {
	t.Run("update word", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		dictionary := Dictionary{key: value}
		newValue := "new value"
		dictionary.Update(key, newValue)
		assertDefinitiion(t, dictionary, key, newValue)
	})
}
func TestDelete(t *testing.T) {
	key := "test"
	value := "this is just a test"
	dictionary := Dictionary{key: value}
	dictionary.Delete(key)
	_, err := dictionary.Search(key)
	if err != ErrNotFound {
		t.Errorf("could not find %s, nothing deleted", key)
	}
}
func assertString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
	if got == nil && want != nil {
		t.Fatal("expected to get an error")
	}
}
func assertDefinitiion(t *testing.T, d Dictionary, key, value string) {
	t.Helper()
	got, err := d.Search(key)
	want := value
	if err != nil {
		t.Fatal("should find added word", err)
	}
	assertString(t, got, want)
}
