package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"sync"
)

type URLStore1 struct {
	urls map[string]string
	mu sync.RWMutex
	file *os.File
}

type record struct {
	Key, URL string
}

func NewURLStore1(filename string) *URLStore1 {
	s := &URLStore1{urls: make(map[string]string)}
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("URLStore1:", err)
	}
	s.file = f
	if err := s.load1(); err != nil {
		log.Println("Error loading data in URLStore1: ", err)
	}
	return s
}

func (s *URLStore1) save1(key, url string) error {
	e := json.NewEncoder(s.file)
	return e.Encode(record{key, url})
}

func (s *URLStore1) load1() error {
	if _, err := s.file.Seek(0, 0); err != nil {
		return err
	}
	d := json.NewDecoder(s.file)
	var err error
	for err == nil {
		var r record
		if err = d.Decode(&r); err == nil {
			s.Set1(r.Key, r.URL)
		}
	}
	if err == io.EOF {
		return nil
	}
	return err
}

func (s *URLStore1) Set1(key, url string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, present := s.urls[key]
	if present {
		return false
	}
	s.urls[key] = url
	return true
}

func (s *URLStore1) Get1(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.urls[key]
}

func (s *URLStore1) Count1() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.urls)
}

func (s *URLStore1) Put1(url string) string {
	for {
		key := genKey1(s.Count1())
		if ok := s.Set1(key, url); ok {
			if err := s.save1(key, url); err != nil {
				log.Println("Error saving to URLStore1: ", err)
			}
		}
		return key
	}
}

