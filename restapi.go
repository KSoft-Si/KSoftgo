package ksoftgo

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var (
	ErrUnauthorized = errors.New("HTTP request was unauthorized")
	ErrRatelimited  = errors.New("too many requests")
)

func (s *KSession) PostForm(urlStr string, data url.Values) (err error) {
	if s.Debug {
		log.Printf("REQUEST %8s :: %s\n", "POST", urlStr)
		log.Printf("REQUEST  PAYLOAD :: [%s]\n", data.Encode())
	}
	req, err := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode()))

	if s.Token != "" {
		req.Header.Set("authorization", "Bearer "+s.Token)
	}
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", s.UserAgent)

	if s.Debug {
		for k, v := range req.Header {
			log.Printf("REQUEST  HEADER :: [%s] = %+v\n", k, v)
		}
	}

	resp, err := s.Client.Do(req)

	defer func() {
		err2 := resp.Body.Close()
		if err2 != nil {
			log.Println("error closing response body")
		}
	}()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if s.Debug {
		log.Printf("RESPONSE  STATUS :: %s\n", resp.Status)
		for k, v := range resp.Header {
			log.Printf("RESPONSE  HEADER :: [%s] = %+v\n", k, v)
		}
		log.Printf("RESPONSE  BODY :: [%s]\n\n\n", response)
	}

	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusCreated:
	case http.StatusNoContent:
	case 429:
		s.log(1, "Rate Limiting %s", urlStr)
		err = ErrRatelimited
	case http.StatusUnauthorized:
		if s.Token != "" {
			s.log(1, ErrUnauthorized.Error())
		}
		fallthrough
	default:
		err = newRestError(req, resp, response)
	}
	return
}

func (s *KSession) request(method, urlStr string, b []byte) (response []byte, err error) {
	if s.Debug {
		log.Printf("REQUEST %8s :: %s\n", method, urlStr)
		log.Printf("REQUEST  PAYLOAD :: [%s]\n", string(b))
		return
	}

	req, err := http.NewRequest(method, urlStr, bytes.NewBuffer(b))

	if s.Token != "" {
		req.Header.Set("authorization", "Bearer "+s.Token)
	}

	if b != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("User-Agent", s.UserAgent)

	if s.Debug {
		for k, v := range req.Header {
			log.Printf("REQUEST  HEADER :: [%s] = %+v\n", k, v)
		}
	}

	resp, err := s.Client.Do(req)

	defer func() {
		err2 := resp.Body.Close()
		if err2 != nil {
			log.Println("error closing response body")
		}
	}()

	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if s.Debug {
		log.Printf("RESPONSE  STATUS :: %s\n", resp.Status)
		for k, v := range resp.Header {
			log.Printf("RESPONSE  HEADER :: [%s] = %+v\n", k, v)
		}
		log.Printf("RESPONSE  BODY :: [%s]\n\n\n", response)
	}

	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusCreated:
	case http.StatusNoContent:
	case 429:
		s.log(1, "Rate Limiting %s", urlStr)
		err = ErrRatelimited
	case http.StatusUnauthorized:
		if s.Token != "" {
			s.log(1, ErrUnauthorized.Error())
		}
		fallthrough
	default:
		err = newRestError(req, resp, response)
	}
	return
}
