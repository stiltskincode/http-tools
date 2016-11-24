package cmd

import (
	"sync"
	"net/http"
	"github.com/Sirupsen/logrus"
	"time"
	"io"
	"bytes"
	"math/rand"
)

const (
	postfixLength = 20
	chunkSize = 1000
)
var (
 	wg sync.WaitGroup
	LogChannel = make(chan BenchmarkLog)
)

//33 126

func random(min, max int) int {
	return rand.Intn(max - min) + min
}

func GetRandomBytes(size int64, begin, end int) []byte {
	var data []byte

	var i int64
	for i =  0; i < size; i++ {
		data = append(data, byte(random(begin, end)))
	}

	return data
}

func GetRandomPrintableBytes(size int64) []byte {
	return GetRandomBytes(size, 33, 126)
}


type BenchmarkJob struct{
	Url string
	Method string
	ReqestsNum int
	Number int
}

func (b *BenchmarkJob) NewRequest()(*http.Request, error){
	var body io.Reader = nil

	if b.Method == "PUT"  || b.Method == "POST" {
		body = bytes.NewReader(GetRandomBytes(chunkSize, 0, 126))
	}

	return http.NewRequest(b.Method, b.Url, body)
}


func (b *BenchmarkJob) Start(){
	go func() {
		var error error;
		start := time.Now()

		for i := 0; i < b.ReqestsNum; i++ {
			if request, error := b.NewRequest(); error == nil {
				if response, error := http.DefaultClient.Do(request); error == nil {
					elapsed := time.Since(start)
					blog := BenchmarkLog{
						Fields: logrus.Fields{
							"threadId": b.Number,
							"requestId": i,
							"duration": elapsed.String(),
							"status" : response.Status,
							"statusCode" : response.StatusCode,
						},
						Message: "Multi-thread response from host",
					}
					LogChannel <- blog
				}
			}

			if error != nil {
				wg.Done()
				return
			}
		}
		wg.Done()
	}()
}

func NewBenchmarkJob(url, method string, threadId,  reqNum int, randomPostfix bool) BenchmarkJob{
	if randomPostfix {
		url += "/" +  string(GetRandomPrintableBytes(postfixLength)[:postfixLength])
	}
	return BenchmarkJob{
		Url: url,
		Method: method,
		ReqestsNum: reqNum,
		Number:threadId,
	}
}

func HttpEndpointBenchmark(method, urlStr string, thrNum, reqNum int, randomPostfix bool){
	wg.Add(thrNum)
	for i:=1; i <= thrNum; i++ {
		b := NewBenchmarkJob(urlStr, method, i, reqNum, randomPostfix)
		b.Start()
	}

	logDispatch()
	wg.Wait()
}

