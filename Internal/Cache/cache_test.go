package pokecache

import (
	"testing"
	"time"
)

func TestCache(t *testing.T){
	inter := time.Millisecond*5000
	cache := NewCache(inter)
	if cache.cache == nil{
		t.Errorf("cache is nil")
	}
}

func TestCacheAddGet(t *testing.T){
	inter := time.Millisecond*5000
	cache := NewCache(inter)

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "one",
			inputVal: []byte("Hello world"),
		},
		{
			inputKey: "key",
			inputVal: []byte("that is a key"),
		},
		{
			inputKey: "",
			inputVal: []byte("???"),
		},
	}

	for _, c := range cases{
		cache.Add(c.inputKey, []byte(c.inputVal))
		actual, ok := cache.Get(c.inputKey)
		if !ok{
			t.Errorf("%v not found", c.inputKey)
			continue
		}
		if string(actual) != string(c.inputVal){
			t.Errorf("%v value doen't match", string(c.inputVal))
			continue
		}
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}