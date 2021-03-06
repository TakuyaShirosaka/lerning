package main

import (
    "fmt"
    "testing"
    "time"
    "sync"
    "strconv"
    "strings"
)

func main() {
    result := testing.Benchmark(func(b *testing.B) { run("1,Java", "2,Lambda", "3,Golang", "4,JavaScript", "5,VBA") })
    fmt.Println(result.T)
}

func run(name ...string) {
    fmt.Println("Start!")
    // WaitGroupを作成する
    wg := new(sync.WaitGroup)

    // channelを処理の数分だけ作成する
    isFin := make(chan bool, len(name))

    for _, v := range name {
        // 処理1つに対して、1つ数を増やす（この例の場合は5になる）
        wg.Add(1)
        // サブスレッドに処理を任せる
        go process(v, isFin, wg)
    }

    // wg.Doneが5回通るまでブロックし続ける
    wg.Wait()
    close(isFin)
    fmt.Println("Finish!")
}

func process(name string, isFin chan bool, wg *sync.WaitGroup) {
    // wgの数を1つ減らす（この関数が終了した時）
    defer wg.Done()
    
    prefixnum, _:= strconv.ParseInt(strings.Split(name,",")[0], 10, 64)
    fmt.Println(prefixnum)
    time.Sleep(2 * time.Second)
    fmt.Println("    ",name)
    isFin <- true
}