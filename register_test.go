package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

type RegisterUserDTO struct {
	GoGoID   uint64 `json:"go_go_id" form:"go_go_id"`
	Nickname string `json:"nickname" form:"nickname"`
	Password string `json:"password" form:"password"`
	Sex      string `json:"sex" form:"sex"`
	Age      string `json:"age" form:"age"`
}

func BenchmarkRegister_test(b *testing.B) {
	var wg sync.WaitGroup
	client := &http.Client{}
	url := "http://localhost:8081/register"
	registerUser := RegisterUserDTO{
		GoGoID:   123456121,
		Nickname: "测试88",
		Password: "123456",
		Sex:      "男",
		Age:      "20",
	}
	reqBody, _ := json.Marshal(&registerUser)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	startTime := time.Now()
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp, err := client.Do(req)
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(body))
		}()
	}
	wg.Wait()
	elapsed := time.Since(startTime)
	log.Printf("100 requests took %v seconds", elapsed.Seconds())
}
