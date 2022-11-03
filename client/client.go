// client package contains the sdk to be used for accessing The One API.
// This is currently a partial implementation, with sorting, pagination supported but not all endpoints.
// They would all follow the same pattern implemented for the existing endpoints, Books, Movies and Characters.
// review the client_test.go file for usage examples.
package client

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type TheOneClient interface {
	GetBooks(opt *Options) (*BookData, error)
	GetBookById(id string) (*BookData, error)
	GetChaptersOfBook(id string, opt *Options) (*ChapterData, error)
	GetMovies(opt *Options) (*MovieData, error)
	GetMovieById(id string) (*MovieData, error)
	GetCharacters(opt *Options) (*CharacterData, error)
	GetCharacterById(id string) (*CharacterData, error)
}

type TheOneClientImpl struct {
	client http.Client
	token  string
}

type Options struct {
	SortOpt *Sort
	Limit   *int
	Page    *int
	Offset  *int
}

type Sort struct {
	Key   string
	Order SortOrder
}

type SortOrder string

const (
	DEFAULT_BASE_URL = "https://the-one-api.dev/v2"
	SortOrderAsc     = "asc"
	SortOrderDesc    = "desc"
	DefaultLimit     = 10
)

//TODO add warning and allow No accessToken
func NewClient(accessToken string) TheOneClient {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// TODO: These should be defined as Configurable via the client creation function
	c := http.Client{Transport: tr, Timeout: time.Duration(30) * time.Second}
	return TheOneClientImpl{client: c, token: accessToken}
}

//fetch does the actual http(s) request to the server, returning the body as a byte array
func (c TheOneClientImpl) fetch(path string, opt *Options) ([]byte, error) {

	url := buildUrl(DEFAULT_BASE_URL, path, opt)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error %s", err)
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.token)

	resp, err := c.client.Do(req)
	if err != nil {
		fmt.Printf("Error %s", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Error %s", err)
		return nil, err
	}

	fmt.Printf("Body : \n%s", body)
	return body, nil
}

func (c TheOneClientImpl) GetBooks(opt *Options) (*BookData, error) {
	path := "/book"
	body, err := c.fetch(path, opt)
	if err != nil {
		return nil, err
	}

	var data BookData
	if err = json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}
	return &data, nil
}

func (c TheOneClientImpl) GetBookById(id string) (*BookData, error) {
	path := "/book/" + id
	body, err := c.fetch(path, nil)
	if err != nil {
		return nil, err
	}

	var data BookData
	if err = json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}
	return &data, nil
}

func (c TheOneClientImpl) GetChaptersOfBook(id string, opt *Options) (*ChapterData, error) {
	path := "/book/" + id + "/chapter"
	body, err := c.fetch(path, opt)
	if err != nil {
		return nil, err
	}

	var data ChapterData
	if err = json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}
	return &data, nil
}

func (c TheOneClientImpl) GetMovies(opt *Options) (*MovieData, error) {
	path := "/movie"
	body, err := c.fetch(path, opt)
	if err != nil {
		return nil, err
	}

	var data MovieData
	if err = json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}
	return &data, nil
}

func (c TheOneClientImpl) GetMovieById(id string) (*MovieData, error) {
	path := "/movie/" + id
	body, err := c.fetch(path, nil)
	if err != nil {
		return nil, err
	}

	var data MovieData
	if err = json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}
	return &data, nil
}

func (c TheOneClientImpl) GetCharacters(opt *Options) (*CharacterData, error) {
	path := "/character"
	body, err := c.fetch(path, opt)
	if err != nil {
		return nil, err
	}

	var data CharacterData
	if err = json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}
	return &data, nil
}

func (c TheOneClientImpl) GetCharacterById(id string) (*CharacterData, error) {
	path := "/character/" + id
	body, err := c.fetch(path, nil)
	if err != nil {
		return nil, err
	}

	var data CharacterData
	if err = json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}
	return &data, nil
}

// buildUrl builds the complete url, including any options.
// Defaults to a limit of 10 unless otherwise specified.
func buildUrl(baseUrl string, path string, opt *Options) string {
	url := DEFAULT_BASE_URL + path + "?"
	if opt == nil {
		url += "limit=" + strconv.Itoa(DefaultLimit)
		return url
	}
	if opt.Limit == nil {
		url += "limit=" + strconv.Itoa(DefaultLimit)
	} else {
		url += "limit=" + strconv.Itoa(*opt.Limit)
	}
	if opt.Page != nil {
		url += "&page=" + strconv.Itoa(*opt.Page)
	}
	if opt.Offset != nil {
		url += "&offset=" + strconv.Itoa(*opt.Offset)
	}
	if opt.SortOpt != nil {
		url += fmt.Sprintf("&%s=%s", opt.SortOpt.Key, opt.SortOpt.Order)
	}

	return url
}
