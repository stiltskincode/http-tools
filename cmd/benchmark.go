package cmd

import (
	"sync"
	"net/http"
	"github.com/Sirupsen/logrus"
	"time"
)

var (
 	wg sync.WaitGroup
	LogChannel = make(chan BenchmarkLog)
)

type BenchmarkJob struct{
	Url string
	Method string
	ReqestsNum int
	Number int
}

func (b *BenchmarkJob) Start(){
	go func() {
		var error error;
		start := time.Now()

		for i := 0; i < b.ReqestsNum; i++ {
			if request, error := http.NewRequest(b.Method, b.Url, nil); error == nil {
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

func NewBenchmarkJob(url, method string, threadId,  reqNum int) BenchmarkJob{
	return BenchmarkJob{
		Url: url,
		Method: method,
		ReqestsNum: reqNum,
		Number:threadId,
	}
}

func HttpEndpointBenchmark(method, urlStr string, thrNum, reqNum int){
	wg.Add(thrNum)
	for i:=1; i <= thrNum; i++ {
		b := NewBenchmarkJob( urlStr, method, i, reqNum)
		b.Start()
	}

	logDispatch()
	wg.Wait()
}

