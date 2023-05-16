package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
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
	wg.Add(1)

	go worker(&wg)

	wg.Wait()
	fmt.Println("All done!")
}
func worker(wg *sync.WaitGroup) {
	for i := 0; i < 100000; i++ {
		testRegister()
	}
	wg.Done()
}
func testRegister() {
	client := &http.Client{}
	url := "http://localhost:8081/register"
	rand.Seed(time.Now().UnixNano())
	registerUser := RegisterUserDTO{
		GoGoID:   uint64(rand.Int63n(100000000)),
		Nickname: "测试" + fmt.Sprint(rand.Int63n(100000000)),
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
}
