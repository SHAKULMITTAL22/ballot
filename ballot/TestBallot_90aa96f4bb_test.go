package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"sort"
	"strings"
	"sync"
)

func TestBallot() error {
	// Make an HTTP GET request to retrieve the initial ballot count
	resp, result, err := httpClientRequest(http.MethodGet, "://", "/", nil)
	if err!= nil {
		log.Printf("Failed to get ballot count resp:%s error:%+v", string(result), err)
		return err
	}
	log.Println("get ballot resp:", string(result))

	// Unmarshal the JSON response into a struct
	var initalRespData ResultBoard
	if err = json.Unmarshal(result, &initalRespData); err!= nil {
		log.Printf("Failed to unmarshal get ballot response. %+v", err)
		return err
	}

	// Create a new vote request
	ballotvotereq := Vote{
		CandidateID: fmt.Sprint(rand.Intn(10)),
		VoterID:     fmt.Sprint(rand.Intn(10)),
	}

	// Marshal the vote request into JSON
	reqBuff, err := json.Marshal(ballotvotereq)
	if err!= nil {
		log.Printf("Failed to marshall post ballot request %+v", err)
		return err
	}

	// Make an HTTP POST request to submit the vote
	resp, result, err := httpClientRequest(http.MethodPost, "://", "/", bytes.NewReader(reqBuff))
	if err!= nil {
		log.Printf("Failed to get ballot count resp:%s error:%+v", string(result), err)
		return err
	}
	log.Println("post ballot resp:", string(result))

	// Unmarshal the JSON response into a struct
	var postballotResp Status
	if err = json.Unmarshal(result, &postballotResp); err!= nil {
		log.Printf("Failed to unmarshal post ballot response. %+v", err)
		return err
	}

	// Check that the status code is 201
	if postballotResp.Code!= 201 {
		return errors.New("post ballot resp status code")
	}

	// Make another HTTP GET request to retrieve the final ballot count
	resp, result, err := httpClientRequest(http.MethodGet, "://", "/", nil)
	if err!= nil {
		log.Printf("Failed to get final ballot count resp:%s error:%+v", string(result), err)
		return err
	}
	log.Println("get final ballot resp:", string(result))

	// Unmarshal the JSON response into a struct
	var finalRespData ResultBoard
	if err = json.Unmarshal(result, &finalRespData); err!= nil {
		log.Printf("Failed to unmarshal get final ballot response. %+v", err)
		return err
	}

	// Check that the total votes have increased by one
	if finalRespData.TotalVotes - initalRespData.TotalVotes!= 1 {
		return errors.New("ballot vote count error")
	}

	return nil
}
