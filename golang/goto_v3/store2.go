package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"sync"
)

type URLStore2 struct {
	urls map[string]string
	mu  sync.RWMutex
	save chan record2
}

type record2 struct {
	Key, URL string
}

const saveQueueLength = 1000

func NewURLStore2(filename string) *URLStore2 {
	s := &URLStore2{
		urls: make(map[string]string),
		save: make(chan record2, saveQueueLength),
	}
	if err := s.load2(filename); err != nil {
		log.Println("Error loading URLStore2: ", err)
	}
	go s.saveLoop(filename)
	return s
}

func (s *URLStore2) Set2(key, url string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, present := s.urls[key]
	if present {
		return false
	}
	s.urls[key] = url
	return true
}

func (s *URLStore2) Get2(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.urls[key]
}

func (s *URLStore2) Count2() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.urls)
}

func (s *URLStore2) Put2(url string) string {
	for {
		key := genKey2(s.Count2())
		if ok := s.Set2(key, url); ok {
			s.save <- record2{key, url}
			return key
		}
	}
}

func (s *URLStore2) load2(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		log.Println("Error opening URLStore2: ", err)
		return err
	}
	defer f.Close()
	d := json.NewDecoder(f)
	for err == nil {
		var r record2
		if err = d.Decode(&r); err == nil {
			s.Set2(r.Key, r.URL)
		}
	}
	if err == io.EOF {
		return nil
	}
	return err
}

func (s *URLStore2) saveLoop(filename string) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("URLStore2: ", err)
	}
	defer f.Close()
	e := json.NewEncoder(f)
	for {
		r := <-s.save
		if err := e.Encode(r); err != nil {
			log.Println("URLStore2: ", err)
		}
	}
}
