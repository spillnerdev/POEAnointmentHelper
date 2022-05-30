package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type Test[T any] func(T) bool

type Stash struct {
	Tabs []Tab `json:"tabs"`
}

type Tab struct {
	Name  string `json:"n"`
	Type  string `json:"type"`
	Id    int    `json:"i"`
	Items []Item `json:"items"`
}

type Item struct {
	BaseType    string   `json:"baseType"`
	EnchantMods []string `json:"enchantMods"`
	XPos        int      `json:"x"`
	YPos        int      `json:"y"`
	InventoryId string   `json:"inventoryId"`
}

func (i Item) String() string {
	return fmt.Sprintf("%s{%s X:%d Y:%d %s}", i.BaseType, i.EnchantMods, i.XPos, i.YPos, i.InventoryId)
}

const GetStashByIdUrl = "https://www.pathofexile.com/character-window/get-stash-items?accountName=%s&realm=pc&league=%s&tabIndex=%d"
const GetStashesUrl = "https://www.pathofexile.com/character-window/get-stash-items?accountName=%s&realm=pc&league=%s&tabs=1"

var client = &http.Client{}

var header = map[string]string{"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:86.0) Gecko/20100101 Firefox/86.0",
	"Host": "www.pathofexile.com"}

func GetUserStash(creds *Creds) (*Stash, error) {
	url := fmt.Sprintf(GetStashesUrl, creds.AccountName, creds.League)
	request := createRequest(url, creds)
	responseBody := doRequest(request)

	tabs := Stash{}
	if err := json.Unmarshal(responseBody, &tabs); err != nil {
		return nil, err
	}

	return &tabs, nil
}

func GetAccessories(creds *Creds, stash Tab) ([]Item, []Item) {
	url := fmt.Sprintf(GetStashByIdUrl, creds.AccountName, creds.League, stash.Id)
	request := createRequest(url, creds)
	responseBody := doRequest(request)

	err := ioutil.WriteFile("format.json", responseBody, 0644)
	if err != nil {
		return nil, nil
	}

	tab := Tab{}

	json.Unmarshal(responseBody, &tab)

	rings := filter(tab.Items, func(item Item) bool {
		return strings.HasSuffix(item.BaseType, "Ring") && item.EnchantMods != nil && len(item.EnchantMods) > 0
	})

	amulets := filter(tab.Items, func(item Item) bool {
		return strings.HasSuffix(item.BaseType, "Amulet") && len(item.EnchantMods) > 0
	})

	return rings, amulets
}

func createRequest(url string, creds *Creds) *http.Request {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(fmt.Errorf("Unable to create request: %w\n", err))
	}

	for k, v := range header {
		request.Header.Add(k, v)
	}
	request.AddCookie(&http.Cookie{
		Name:  "POESESSID",
		Value: creds.POESESSID,
	})

	return request
}

func doRequest(r *http.Request) []byte {
	response, err := client.Do(r)
	if err != nil {
		panic(fmt.Errorf("Unable to execute request: %w\n", err))
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		panic(fmt.Errorf("Unable to retrieve Tabs: %d\n", response.StatusCode))
	}

	all, err := io.ReadAll(response.Body)
	if err != nil {
		panic(fmt.Errorf("Unable to retrieve response body: %w\n", err))
	}
	return all
}

func anyMatch[T any](slice []T, test Test[T]) bool {

	for _, item := range slice {
		if test(item) {
			return true
		}
	}
	return false
}

func allMatch[T any](slice []T, test Test[T]) bool {
	for _, item := range slice {
		if !test(item) {
			return false
		}
	}
	return true
}

func filter[T any](slice []T, test Test[T]) []T {
	result := make([]T, len(slice))
	copy(result, slice)
	n := 0

	for _, item := range result {
		if test(item) {
			result[n] = item
			n++
		}

	}

	return result[:n]
}
