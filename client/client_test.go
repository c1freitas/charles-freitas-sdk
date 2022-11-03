// client_tests package contains some rudimentray tests for invoking the client SDK.
// This can also be used as a guide for how to use the API in your code.
//
// In order to run the tests you will need an API token. This can be found by logging
// into your account at https://the-one-api.dev/login
//
//   $ THE_ONE_TOKEN="<YOUR-TOKEN>" go test client/*
//
package client_test

import (
	"os"
	"testing"

	"github.com/c1freitas/charles-freitas-sdk/client"
	"github.com/google/go-cmp/cmp"
)

func TestClientGetBooksAndChapters(t *testing.T) {

	client := client.NewClient(os.Getenv("THE_ONE_TOKEN"))

	bookData, err := client.GetBooks(nil)
	if err != nil {
		t.Errorf("Error: %v", err)
		t.Fail()
	}
	if len(bookData.Books) == 0 {
		t.Error("Should have returned at least one book")
	}
	book := bookData.Books[0]
	newBook, err := client.GetBookById(book.Id)

	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if len(newBook.Books) != 1 {
		t.Error("Should have returned exactly one book")
	}

	chaptersData, err := client.GetChaptersOfBook(book.Id, nil)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if len(chaptersData.Chapters) == 0 {
		t.Error("Should have returned at least one Chapter")
	}
	chapter := chaptersData.Chapters[0]

	if len(chapter.ChapterName) == 0 {
		t.Error("ChapterName should be populated")
	}
}

func TestClientMovies(t *testing.T) {

	client := client.NewClient(os.Getenv("THE_ONE_TOKEN"))

	movieData, err := client.GetMovies(nil)
	if err != nil {
		t.Errorf("Error: %v", err)
		t.Fail()
	}

	if len(movieData.Movies) == 0 {
		t.Error("Should have returned at least one Movie")
	}
	item := movieData.Movies[0]
	newMovie, err := client.GetMovieById(item.Id)

	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if len(newMovie.Movies) != 1 {
		t.Error("Should have returned exactly one Movie")
	}
}

func TestCharacters(t *testing.T) {

	client := client.NewClient(os.Getenv("THE_ONE_TOKEN"))

	charData, err := client.GetCharacters(nil)
	if err != nil {
		t.Errorf("Error: %v", err)
		t.Fail()
	}

	if len(charData.Characters) == 0 {
		t.Error("Should have returned at least one Character")
	}
	item := charData.Characters[0]
	newItem, err := client.GetCharacterById(item.Id)

	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if len(newItem.Characters) != 1 {
		t.Error("Should have returned exactly one Character")
	}
	if len(newItem.Characters[0].Name) == 0 {
		t.Error("Character new should be populated")
	}
}

func TestOptions(t *testing.T) {

	c := client.NewClient(os.Getenv("THE_ONE_TOKEN"))
	testLimit := 2
	Opt := client.Options{Limit: &testLimit}

	charData, err := c.GetCharacters(&Opt)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if len(charData.Characters) != 2 {
		t.Errorf("Should have returned 2 but returned: %v", len(charData.Characters))
	}

	testLimit = 1
	Opt = client.Options{Limit: &testLimit, SortOpt: &client.Sort{Key: "name", Order: client.SortOrderAsc}}
	sortCharDataAsc, err := c.GetCharacters(&Opt)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if len(sortCharDataAsc.Characters) != 1 {
		t.Error("Should have only returned one Character")
	}
	Opt.SortOpt.Order = client.SortOrderDesc
	sortCharDataDesc, err := c.GetCharacters(&Opt)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if len(sortCharDataDesc.Characters) != 1 {
		t.Error("Should have only returned one Character")
	}

	if cmp.Equal(sortCharDataAsc.Characters[0], sortCharDataDesc.Characters[0]) {
		t.Error("Error objects should not be equal")
	}
}
