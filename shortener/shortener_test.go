package shortener_test

import (
	"qoala/shortener"
	"testing"
)

func TestExpandShortID(t *testing.T) {
	testcases := []struct {
		id int64
	}{
		{123},
		{342},
	}

	for _, testcase := range testcases {
		shortenId := shortener.ShortenID(testcase.id)
		if id := shortener.ExpandShortID(shortenId); id != testcase.id {
			t.Errorf("id from shortenId and testcase id not match, shortenId %d testcaseId %d", id, testcase.id)
		}

	}
}

func TestEncodeDecode(t *testing.T) {
	testcases := []struct {
		url string
	}{
		{"https://grabkios.qoala.app?agentId=908123123&phoneNumber=085716028129"},
		{"https://offline.qoalaplus.app?agentId=908123123&phoneNumber=085716028129"},
		{"https://integration.qoala.my?agentId=908123123&phoneNumber=085716028129"},
	}
	for _, testcase := range testcases {
		encode := shortener.Encode(testcase.url)
		decode := shortener.Decode(encode)
		if testcase.url != decode {
			t.Errorf("url from decode not same with testcase, decode %s testcase %s", decode, testcase.url)
		}
	}
}
