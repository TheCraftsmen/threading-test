//libadd.go
package main

import (
    "C"
    "fmt"
    "net/http"
    "time"
    "net"
    "sync"
    "runtime"
    "io/ioutil"
)

//export add
func add() {
    start := time.Now()
    runtime.GOMAXPROCS(runtime.NumCPU())
    buffer := make([]byte, 4096)
    buffer = buffer[:runtime.Stack(buffer, true)]
    var netTransport = &http.Transport{
        Dial: (&net.Dialer{
            Timeout: 120 * time.Second,
        }).Dial,
        TLSHandshakeTimeout: 120 * time.Second,
        MaxIdleConns: 100,
        IdleConnTimeout: 90 * time.Second,
        ExpectContinueTimeout: 120 * time.Second,
    }
    var netClient = &http.Client{
        Timeout: time.Second * 120,
        Transport: netTransport,
    }   
    var wg sync.WaitGroup
    for j := 0; j < 1000; j++ {
        wg.Add(1)
        go func() {
            req, _ := http.NewRequest("GET", "https://api.mercadolibre.com/sites", nil) 
            resp, err := netClient.Do(req)
            runtime.Gosched()
            if err != nil {
                panic(err)
            }
            defer resp.Body.Close()
            ioutil.ReadAll(resp.Body)
            resp.Close = true
            fmt.Println(resp.StatusCode, j)
            wg.Done()
        }()

    }
    wg.Wait()
    elapsed := time.Since(start)
    fmt.Println("Listo!", runtime.NumGoroutine(), elapsed)
    return
}

func main() {
}
