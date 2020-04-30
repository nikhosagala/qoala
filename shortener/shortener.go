package shortener

import (
	"strings"
)

var shortenerDict = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789` // you can customize this, just make sure there's no duplicate chars
var shortenerBase = len(shortenerDict)

var containers []QoalaShortener

type QoalaShortener struct {
	id  int64
	url string
}

// Encode a url and return new shorten url
func Encode(url string) string {
	qs := QoalaShortener{
		id:  int64(len(containers) + 1),
		url: url,
	}
	shortenId := ShortenID(qs.id)
	containers = append(containers, qs)
	return "qoa.la/" + shortenId
}

// Decode shorten url and return long url
func Decode(url string) string {
	result := strings.Split(url, "/")
	shortenId := result[1]
	id := ExpandShortID(shortenId)
	qs := findById(id)
	return qs.url
}

func ShortenID(ID int64) string {
	result := ""
	for ID > 0 {
		mod := ID % int64(shortenerBase)
		result = shortenerDict[int(mod):int(mod)+1] + result
		ID /= int64(shortenerBase)
	}
	return result
}

func ExpandShortID(shortID string) int64 {
	var id int

	for i, char := range shortID {
		index := strings.Index(shortenerDict, string(char))
		id += index

		if i < len(shortID)-1 {
			id *= shortenerBase
		}
	}
	return int64(id)
}
