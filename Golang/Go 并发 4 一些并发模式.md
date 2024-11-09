## Table of Contents
  - [Immutable and Confinement](#Immutable-and-Confinement)
  - [The or-channel](#The-orchannel)
  - [Error Handling](#Error-Handling)
  - [Pipeline](#Pipeline)
    - [Channel & Pipeline](#Channel--Pipeline)
    - [Some Useful Generators](#Some-Useful-Generators)
    - [Fan-Out, Fan-In](#FanOut-FanIn)
  - [The or-done-channel](#The-ordonechannel)
  - [The tee-channel](#The-teechannel)
  - [The bridge-channel](#The-bridgechannel)
  - [Queuing](#Queuing)
  - [Heartbeats](#Heartbeats)
  - [Healing Unhealthy Goroutines](#Healing-Unhealthy-Goroutines)
  - [Replicated Requests](#Replicated-Requests)
  - [Rate Limiting](#Rate-Limiting)
  - [为什么说 Go 的并发好用](#%E4%B8%BA%E4%BB%80%E4%B9%88%E8%AF%B4-Go-%E7%9A%84%E5%B9%B6%E5%8F%91%E5%A5%BD%E7%94%A8)

## Immutable and Confinement

When working with concurrent code, there are a few different options for safe operation.

- Synchronization primitives for sharing memory (sync.Mutex)

- Synchronization via communicating (channels)  

However, there are a couple of other options that are *implicitly safe* within multiple concurrent processes.

- Immutable data
- Data protected by confinement

#### ➤ Immutable Data

Since the immutable objects cannot be modified after being created, then it’s inherently concurrent-safe. Each concurrent process may operate on the same data, but it may not modify it. If it wants to create new data, it must *create a new copy of the data with the desired modifications*. This allows not only a lighter cognitive load on the developer, but can also lead to faster programs if it leads to smaller critical sections. In Go, you can achieve this by writing code that utilizes copies of values instead of pointers to values in memory.   

#### ➤ Confinement

如果不共享数据、那么就不需要同步数据,  比如让每个协程有自己独享的一份 data

We know that accessing mutable, shared data between threads requires some level of synchronization, but one way to avoid doing this is to simply not share. Confinement is the simple yet powerful idea of *ensuring information is only ever available from one concurrent process*. When this is achieved, a concurrent program is implicitly safe and no synchronization is needed.



## The or-channel

At times you may find yourself wanting to combine one or more done channels into a single done channel that *closes if any of its component channels close*. You can combine these channels together using the or-channel pattern.  

```go
func or_channel_pattern() {
    // or() 函数返回聚合后的 channel, 当入参中的任意一个 channel 完成时, 返回值也会完成
    var or func(channels ...<-chan int) <-chan int
    or = func(channels ...<-chan int) <-chan int {

        // 递归终止条件: channels 长度为 0 或 1
        switch len(channels) {
        case 0:
            return nil
        case 1:
            return channels[0]
        }

        // 新开一个 goroutine 在合适的时候关闭 orDone
        orDone := make(chan int)
        go func() {
            defer close(orDone)
            switch len(channels) {
            case 2:
                select {
                case <-channels[0]:
                case <-channels[1]:
                }
            default:
                select {
                case <-channels[0]: // 当 012 任意一个 channel 完成时, 用于代表它们的 orDone 也会完成
                case <-channels[1]: // 因为会退出当前 select, 然后执行前面声明的 defer close(orDone)
                case <-channels[2]:
                case <-or(append(channels[3:], orDone)...):
                    // 总之把前三个 channel 聚合成一个 orDone,  当有 1000 个 channel 时
                    // channel 之间会聚合成 orDone,  orDone 之间又会聚合成更大的 orDone
                    // 虽然层层套娃,  但任意 channel 完成时,  与之关联的 orDone 也会完成
                }
            }
        }()
        return orDone
    }

    sleep := func(d time.Duration) <-chan int {
        c := make(chan int)
        go func() {
            defer close(c)
            time.Sleep(d)
        }()
        return c
    }

    start := time.Now()
    <-or(sleep(1*time.Second), sleep(2*time.Second))
    fmt.Printf("done after %v \n", time.Since(start))
}
```





## Error Handling

#### ➤ 生成 worker 的协程负责错误处理,   因为它比 worker 掌握更多的信息

The most fundamental question when thinking about error handling is, “Who should be responsible for handling the error?” With concurrent processes, this question becomes a little more complex.

Don’t put your goroutines in this awkward position. I suggest you separate your concerns: in general, your concurrent processes should send their errors to another part of your program that has complete information about the state of your program, and can make a more informed decision about what to do.

This is desirable because the goroutine that spawned the worker goroutine—in this case our main goroutine—has more context about the running program, and can make more intelligent decisions about what to do with errors.  

```go
func error_handling() {

    // 把结果和错误封装在一起
    type Result struct {
        Error    error
        Response *http.Response
    }

    // 子 goroutine 不关心错误处理, 把响应和错误都返给父 goroutine,  让它决定怎么做
    getUrls := func(done <-chan interface{}, urls ...string) <-chan Result {
        results := make(chan Result)
        go func() {
            defer close(results)
            for _, url := range urls {
                var result Result
                resp, err := http.Get(url)
                result = Result{Error: err, Response: resp}
                select {
                case <-done:
                    return
                case results <- result:
                }
            }
        }()
        return results
    }

    // 父 goroutine 依次处理子 goroutine 传过来的 Result 结构体
    // 父 goroutine 想怎么处理错误都行,  另外可以用 done 让子 goroutine 提前退出
    done := make(chan interface{})
    defer close(done)
    urls := []string{"https://www.baidu.com", "https://badhost", "https://www.bing.com"}
    for result := range getUrls(done, urls...) {
        if result.Error != nil {
            fmt.Printf("error: %v\n", result.Error)
            continue // 如果遇到错误想提前退出把这行改成 break 就好
        }
        fmt.Printf("Response: %v\n", result.Response.Status)
    }
}
```

#### ➤ 推荐把错误也打包返回, 而不是在 worker 中做错误处理

Errors should be considered first-class citizens when constructing values to return from goroutines. If your goroutine can produce errors, *those errors should be tightly coupled with your result type*, and passed along through the same lines of communication—just like regular synchronous functions.  

## Pipeline

A pipeline is nothing more than a series of things that take data in, perform an operation on it, and pass the data back out. We call each of these operations a stage of the pipeline. 

*By using a pipeline, you separate the concerns of each stage*, which provides numerous benefits. 

1. You can modify stages independent of one another, 
2. you can mix and match how stages are combined independent of modifying the stages, 
3. you can process each stage concurrent to upstream or downstream stages, 
4. and you can fan-out, or rate-limit portions of your pipeline.   

As mentioned previously, a stage is just something that takes data in, performs a transformation on it, and sends the data back out. Here is a function that could be considered a pipeline stage:  

```go
func pipeline_pattern() {

    multiply := func(values []int, multiplier int) []int {
        multipliedValues := make([]int, len(values))
        for i, v := range values {
            multipliedValues[i] = v * multiplier
        }
        return multipliedValues
    }

    add := func(values []int, additive int) []int {
        addedValues := make([]int, len(values))
        for i, v := range values {
            addedValues[i] = v + additive
        }
        return addedValues
    }

    // 然后把两个 stage 结合起来, 数据依次经过 multiply 和 add 两个阶段的处理
    ints := []int{1, 2, 3, 4}
    for _, v := range add(multiply(ints, 2), 1) {
        fmt.Println(v)
    }
}
```

Notice how each stage is taking a slice of data and *returning a slice of data*? These stages are performing what we call *batch processing*. This just means that they operate on chunks of data all at once instead of one discrete value at a time. There is another type of pipeline stage that performs stream processing. This means that the stage receives and *emits one element at a time*.  

Notice that for the original data to remain unaltered, each stage has to make a new slice of equal length to store the results of its calculations. That means that the memory footprint of our program at any one time is double the size of the slice we send into the start of our pipeline. Let’s convert our stages to be *stream oriented* and see what that looks like:  

```go
func stream_oriented_pipeline() {

    multiply := func(value, multiplier int) int {
        return value * multiplier
    }
    add := func(value, additive int) int {
        return value + additive
    }

    // 一次处理一个数据, 数据依次经过 multiply 和 add 两个阶段的处理
    ints := []int{1, 2, 3, 4}
    for _, v := range ints {
        fmt.Println(add(multiply(v, 2), 1))
    }
}
```



### Channel & Pipeline

Channels are uniquely suited to constructing pipelines in Go because they fulfill all of our basic requirements. They can receive and emit values, they can safely be used concurrently. Let’s take a moment and convert the previous example to utilize channels instead:  

```go
func 用_channel_实现_pipeline_模式() {
    // (1) 生成数字,  并通过 channel 发给下一个阶段
    generator := func(done <-chan interface{}, integers ...int) <-chan int {
        intStream := make(chan int)
        go func() {
            defer close(intStream)
            for _, i := range integers {
                select {
                case <-done: // 每一个 goroutine 都通过 done channel 确保没有泄露
                    return
                case intStream <- i:
                }
            }
        }()
        return intStream
    }
    // (2) 输入和输出都是一个 channel
    multiply := func(done <-chan interface{}, intStream <-chan int, multiplier int) <-chan int {
        multipliedStream := make(chan int)
        go func() {
            defer close(multipliedStream)
            for i := range intStream {
                select {
                case <-done:
                    return
                case multipliedStream <- i * multiplier: // 真正的处理逻辑
                }
            }
        }()
        return multipliedStream
    }
    // (3) for-range 循环会在上游 channel 关闭时退出,  done 用于提前退出
    add := func(done <-chan interface{}, intStream <-chan int, additive int) <-chan int {
        addedStream := make(chan int)
        go func() {
            defer close(addedStream)
            for i := range intStream {
                select {
                case <-done:
                    return
                case addedStream <- i + additive:
                }
            }
        }()
        return addedStream
    }
    // (4) 把 generator、done、stage 等元素组合起来
    done := make(chan interface{})
    defer close(done)
    intStream := generator(done, 1, 2, 3, 4)
    pipeline := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)
    for v := range pipeline {
        fmt.Println(v) // 如果在这里 break,  defer close(done) 确保了没有 goroutine 泄露
    }
    // (5) 注意每个阶段都并发执行、暂存了处理结果,  并且在等下一个阶段从我这取数据,
    //     若不是 for v := range pipeline {...} 不断从 pipeline 中拉数据,  整个流水钱跑不起来
}
```

*At each stage we can safely execute concurrently* because our inputs and outputs are safe in concurrent contexts. Each stage of the pipeline is executing concurrently. This means that any stage only need wait for its inputs, and to be able to send its outputs.   

每个 goroutine 是怎么退出的?  有以下两种途径:

- Ranging over the incoming channel. When the incoming channel is closed, the range will exit.
- The send sharing a select statement with the done channel  

*Regardless of what state the pipeline stage is in* (waiting on the incoming channel, or waiting on the send), closing the done channel will force the pipeline stage to terminate.  

### Some Useful Generators

In this basic example, we create a repeat generator to generate an infinite number of random numbers, but then only take the first 10. Let’s use it to generate 10 random numbers.

Because the repeat generator’s send blocks on the take stage’s receive, the repeat generator is very efficient. We only generate N+1 instances where N is the number we pass into the take stage.

```go
func repeat_and_take() {
    // This pipeline stage will call fn infinitely until you tell it to stop.
    repeatFn := func(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
        valueStream := make(chan interface{})
        go func() {
            defer close(valueStream)
            for {
                select {
                case <-done:
                    return
                case valueStream <- fn():
                }
            }
        }()
        return valueStream
    }
    // This pipeline stage will only take the first <num> items off of its incoming valueStream and then exit.
    take := func(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
        takeStream := make(chan interface{})
        go func() {
            defer close(takeStream)
            for i := 0; i < num; i++ {
                select {
                case <-done:
                    return
                case takeStream <- <-valueStream:
                }
            }
        }()
        return takeStream
    }

    // Let’s use it to generate 10 random numbers:
    done := make(chan interface{})
    defer close(done)
    f := func() interface{} { return rand.Int() }
    for num := range take(done, repeatFn(done, f), 10) {
        fmt.Println(num)
    }

    repeat := func(done <-chan interface{}, values ...interface{}) <-chan interface{} {
        valueStream := make(chan interface{})
        go func() {
            defer close(valueStream)
            for {
                for _, v := range values {
                    select {
                    case <-done:
                        return
                    case valueStream <- v:
                    }
                }
            }
        }()
        return valueStream
    }
    // Here’s a small example that introduces a toString pipeline stage.
    toString := func(done <-chan interface{}, valueStream <-chan interface{}) <-chan string {
        stringStream := make(chan string)
        go func() {
            defer close(stringStream)
            for v := range valueStream {
                select {
                case <-done:
                    return
                case stringStream <- v.(string):
                }
            }
        }()
        return stringStream
    }

    var message string
    for token := range toString(done, take(done, repeat(done, "I", "am."), 5)) {
        message += token
    }
    fmt.Printf("message: %s...", message)
}
```

You may be wondering why all of these generators and stages are receiving and sending on channels of `interface{}`. In the `take` stage, the concern is limiting our pipeline. None of these operations require information about the types they’re working on. Here’s a small example that introduces a toString pipeline stage.

### Fan-Out, Fan-In

Speaking of one stage being computationally expensive, how can we help mitigate this? Won’t it rate limit the entire pipeline? Sometimes, stages in your pipeline can be particularly computationally expensive. When this happens, upstream stages in your pipeline can become blocked while waiting for your expensive stages to complete. Not only that, but the pipeline itself can take a long time to execute as a whole. How can we address this?  

目前每个 stage 中只有一个 goroutine,  如果使用两个 goroutine 说不定能提升该 stage 的性能  
In fact, it turns out it can, and this pattern has a name: fan-out, fan-in. Fan-out is a term to describe the process of starting multiple goroutines to handle input from the pipeline, and fan-in is a term to describe the process of combining multiple results into one channel.  

解释下为什么这种思路可行:

1. 一个 stage 只对外部暴露两个接口,  分别是输入/输出 channel,  内部实现可随意更改
2. 输入/输出 channel 是并发安全的,  即使有两个 goroutine 从输入 channel 拉取内容也没问题

```go
func fan_out_fan_in() {
    numCPU := runtime.NumCPU()
    generator := func(done <-chan interface{}, urls ...string) <-chan string {
        urlStream := make(chan string)
        go func() {
            defer close(urlStream)
            for _, i := range urls {
                select {
                case <-done:
                    return
                case urlStream <- i:
                }
            }
        }()
        return urlStream
    }
    getUrl := func(done <-chan interface{}, input <-chan string) <-chan string {
        output := make(chan string)
        go func() {
            defer close(output)
            for url := range input {
                select {
                case <-done:
                    return
                case output <- get1024Bytes(url):
                }
            }
        }()
        return output
    }

    done := make(chan interface{})
    defer close(done)
    urlStream := generator(done, "https://baidu.com", "https://bilibili.com", "https://zhihu.com")

    // (1) fan-out,  多个 goroutine 从同一个 channel 拉取数据进行处理
    channels := make([]<-chan string, numCPU)
    for i := 0; i < numCPU; i++ {
        channels[i] = getUrl(done, urlStream)
    }

    // (2) fan-in,  把多个输出 channel 合并为一个
    fanIn := func(done <-chan interface{}, channels ...<-chan string) <-chan string {
        var wg sync.WaitGroup
        wg.Add(len(channels))
        multiplexedStream := make(chan string)

        // 为每个 channel 开一个 goroutine,  读出来然后转发到 multiplexedStream
        multiplex := func(c <-chan string) {
            defer wg.Done()
            for i := range c {
                select {
                case <-done:
                    return
                case multiplexedStream <- i:
                }
            }
        }
        for _, c := range channels {
            go multiplex(c)
        }
        // 再开一个 goroutine 等待 wait group 完成,  完成时关掉 multiplexedStream
        go func() {
            wg.Wait()
            close(multiplexedStream)
        }()
        return multiplexedStream
    }

    for v := range fanIn(done, channels...) {
        fmt.Println("\n------->\n", v)
    }
}
```

注意上面的实现打乱了顺序,  它一次性从输入 channel 中拿了 1、2、3、4 进行处理,  
因为有的快有的慢, 所以输出顺序不再是 1、2、3、4.  
Later, we’ll look at an example of a way to maintain order.  

## The or-done-channel

```go
func the_orDone_channel(done, myChan <-chan int) {
    // 像这样的循环,  只能等 myChan 被关闭才能退出,  不够灵活
    for val := range myChan {
        _ = val
    }

    // 有时候我们希望用 done channel 从循环提前退出,  就像这样
    var orDone func(<-chan int, <-chan int) <-chan int
    for val := range orDone(done, myChan) {
        _ = val
    }

    // 其中 orDone 的实现如下
    orDone = func(done, c <-chan int) <-chan int {
        valStream := make(chan int)
        go func() {
            defer close(valStream)
            for {
                select {
                case <-done:
                    return
                case v, ok := <-c:
                    if ok == false {
                        return
                    }
                    select {
                    case valStream <- v:
                    case <-done:
                    }
                }
            }
        }()
        return valStream
    }
}
```

## The tee-channel

Sometimes you may want to split values coming in from a channel so that you can send them off into two separate areas of your codebase. Imagine a channel of user commands: you might want to take in a stream of user commands on a channel, send them to something that executes them, and also send them to something that logs the commands for later auditing.  

You can pass it a channel to read from, and it will return two separate channels that will get the same value,  总之就是把一个 channel 变成多个等价的 channel,  它们都有相同的数据流 (影分身之术).

```go
func tee_channel() {
    tee := func(done, in <-chan int) (<-chan int, <-chan int) {
        out1 := make(chan int)
        out2 := make(chan int)
        go func() {
            defer close(out1)
            defer close(out2)
            for val := range orDone(done, in) {
                var out1, out2 = out1, out2 // 创建同名的局部变量,  初始值为外部变量
                for i := 0; i < 2; i++ {    // 循环两次、把同一数据发给两个 channel
                    select {
                    case <-done:
                    case out1 <- val:
                        out1 = nil // 如果第一次循环中把 out1 改成 nil
                    case out2 <- val: // 那么在第二次循环中只能往 out2 发送数据
                        out2 = nil // 因为 out1 变量已经变成了 nil
                    }
                }
            }
        }()
        return out1, out2
    }

    done := make(chan int)
    defer close(done)

    // 注意需要把 out1、out2 都读一次,  否则上面的两次循环会卡死
    out1, out2 := tee(done, nil)
    for val1 := range out1 {
        fmt.Printf("out1: %v, out2: %v\n", val1, <-out2)
    }
}
```

## The bridge-channel

通过 channel 传输的数据,  可以是 bool、int、string 等类型,  甚至可以是 `<-chan int` 类型  
这时候从 channel 读出的不是一个具体值,  而是另一个 channel,  所以处理起来比较麻烦,  可以化简一下.

```go
func a_channel_of_channels() {
    bridge := func(done <-chan int, chanStream <-chan <-chan int) <-chan int {
        valStream := make(chan int)
        go func() {
            defer close(valStream)
            for {
                // 从 chanStream 取一个 channel
                var stream <-chan int
                select {
                case maybeStream, ok := <-chanStream:
                    if ok == false {
                        return
                    }
                    stream = maybeStream
                case <-done:
                    return
                }
                // 然后等处理完这个 channel,  又会从 chanStream 取一个新的 channel
                for val := range orDone(done, stream) {
                    select {
                    case valStream <- val:
                    case <-done:
                    }
                }
            }
        }()
        return valStream
    }

    // bridge 函数的好处是消费者不必处理 <-chan int 类型,  而是处理更简单的 int 类型
    var channelChannels <-chan <-chan int
    valStream := bridge(nil, channelChannels)
    for val := range valStream {
        fmt.Println(val + 100)
    }
}
```

## Queuing

All this means is that once your stage has completed some work, it stores it in a temporary location in memory so that other stages can retrieve it later, and your stage doesn’t need to hold a reference to it. In the section on “Channels” on page 64, we discussed buffered channels, a type of queue.

While introducing queuing into your system is very useful, it’s usually one of the *last techniques* you want to employ when optimizing your program. Adding queuing prematurely can hide synchronization issues such as deadlocks and livelocks.

So what is queuing good for? Let’s begin to answer that question by addressing one of the common mistakes people make when trying to tune the performance of a system: introducing queues to try and address performance concerns. Queuing will almost never speed up the total runtime of your program; it will only allow the program to behave differently.  

*为什么给 stage 加 buffer 不会提升流水线性能?*

1. 因为流水线的整体性能取决于最慢的那个 stage
2. 就算 stage A 把所有结果都暂存/堆积到某个位置,  stage B 也只能以每秒一个的速度进行处理

*但是加 buffer 可以提高单个 stage 的速度*,  有如下好处

1. stage A 把结果堆积到某个位置,  就能不停地工作,  让 stage A 的速度不再受 stage B 影响
2. 虽然 10 个数据依旧需要 10 秒才能通过整个流水线,  但加 buffer 后只需 0 点几秒就能通过 stage A

So the answer to our question of the utility of introducing a queue isn’t that the runtime of one of stages has been reduced, but rather that the time it’s in a blocking state is reduced. This allows the stage to continue doing its job. In this example, users would likely experience lag in their requests, but they wouldn’t be denied service caused by timeout.  

In this way, *the true utility of queues is to decouple stages so that the runtime of one stage has no impact on the runtime of another*.  

Where should the queues be placed? Queuing should be implemented either:

- At the entrance to your pipeline.
- In stages where batching will lead to higher efficiency.





## Heartbeats

A heartbeat is a way to signal to its listeners that everything is well. There are two different types of heartbeats we’ll discuss in this section:

- Heartbeats that occur on a time interval.

- Heartbeats that occur at the beginning of a unit of work. (用于表示 goroutine 已开始工作)

  If you only care that the goroutine has started doing its work, this style of heartbeat is simple.  

```go
func 如何产生心跳(done <-chan int, pulseInterval time.Duration) (<-chan int, <-chan time.Time) {
    heartbeat := make(chan int)
    results := make(chan time.Time)
    go func() {
        defer close(heartbeat)
        defer close(results)
        pulse := time.Tick(pulseInterval)
        workGen := time.Tick(2 * pulseInterval)
        sendPulse := func() {
            select {
            case heartbeat <- 123:
            default: // 发送心跳是非阻塞的,  没人接收时会丢弃这次心跳
            }
        }
        sendResult := func(r time.Time) {
            for {
                select {
                case <-done:
                    return
                case <-pulse:
                    sendPulse() // 发送结果时也要尝试发送心跳,  别把心跳停了
                case results <- r:
                    return
                }
            }
        }
        for {
            select {
            case <-done:
                return
            case <-pulse:
                sendPulse() // 有心跳则发送心跳
            case r := <-workGen:
                sendResult(r) // 有结果则发送结果
            }
            time.Sleep(10 * time.Second)
        }
    }()
    // 注意心跳是在 select 语句等待时发送的,  如果 goroutine 一直在执行计算、就没机会发送心跳
    return heartbeat, results
}
```

Notice that because we might be sending out multiple pulses while we wait for input, or multiple pulses while waiting to send results, all the select statements need to be within for loops.   

```go
func 如何使用心跳() {
    done := make(chan int)
    time.AfterFunc(10*time.Second, func() { close(done) }) // 本协程最多 10 秒后关闭
    const timeout = 2 * time.Second
    heartbeat, results := 如何产生心跳(done, timeout/2)

    for {
        select {
        case _, ok := <-heartbeat:
            if ok == false {
                return
            }
            fmt.Println("pulse")
        case r, ok := <-results:
            if ok == false {
                return
            }
            fmt.Printf("results %v\n", r.Second())
        case <-time.After(timeout):
            // 已知每隔一秒产生一次心跳, 这里两秒过去了但还是没有心跳
            // 可以判定 worker goroutine 出问题了,  比如死锁、忘了关闭 channel 之类的
            fmt.Println("worker goroutine is not healthy!")
            return
        }
    }
}
```

Beautiful! Within two seconds our system realizes something is amiss with our goroutine and breaks the for-select loop. By using a heartbeat, we have successfully avoided a deadlock.

## Healing Unhealthy Goroutines

In long-lived processes such as daemons, it’s very common to have a set of long-lived goroutines. These goroutines are usually blocked, waiting on data to come to them through some means, so that they can wake up, do their work, and then pass the data on. Sometimes the goroutines are dependent on a resource that you don’t have very good control of. Maybe a goroutine receives a request to pull data from a web service, or maybe it’s monitoring an ephemeral file. The point is that it can be very easy for a goroutine to become stuck in a bad state from which it cannot recover without external help. If you separate your concerns, you might even say that it shouldn’t be the concern of a goroutine doing work to know how to heal itself from a bad state. In a long-running process, it can be useful to create a mechanism that ensures your goroutines remain healthy and restarts them if they become unhealthy. We’ll refer to this process of restarting goroutines as “healing.”2  

To heal goroutines, we’ll use our heartbeat pattern to check up on the liveliness of the goroutine we’re monitoring. The type of heartbeat will be determined by what you’re trying to monitor, but if your goroutine can become livelocked, make sure that the heartbeat contains some kind of information indicating that the goroutine is not only up, but doing useful work. In this section, for simplicity, we’ll only consider whether goroutines are live or dead.  

## Replicated Requests

For some applications, receiving a response as quickly as possible is the top priority. For example, maybe the application is servicing a user’s HTTP request. In these instances you can make a trade-off: you can replicate
the request to multiple handlers (whether those be goroutines, processes, or servers), and one of them will return faster than the other ones; you can then immediately return the result.

## Rate Limiting

Have you ever wondered why services put rate limits in place? Why not allow unfettered access to a system? The most obvious answer is that by rate limiting a system, you prevent entire classes of attack vectors against your system. If malicious users can access your system as quickly as their resources allow it, they can do all kinds of things.  

Malicious use isn’t the only reason. In distributed systems, a legitimate user could degrade the performance of the system for other users if they’re performing operations at a high enough volume.   

#### ➤ 限制最大并发数

Another technique that can be implemented with a buffered channel is backpressure. It is counterintuitive, but systems perform better overall when their components limit the amount of work they are willing to perform. 

```go
func 用_buffered_channel_限制并发数() {
    pg := New(10)
    wg := sync.WaitGroup{}
    wg.Add(20)
    for i := 0; i < 20; i++ {
        i := i
        go func() {
            defer wg.Done()
            err := pg.Process(func() {
                fmt.Println("处理", i)
            })
            if err != nil {
                fmt.Println(err)
            }
        }()
    }
    wg.Wait()
}

type PressureGauge struct {
    ch chan struct{} // 当 channel 中的数据类型不重要时,  使用 struct{}
}

func New(limit int) *PressureGauge {
    ch := make(chan struct{}, limit)
    for i := 0; i < limit; i++ {
        ch <- struct{}{}
    }
    return &PressureGauge{
        ch: ch,
    }
}

func (pg *PressureGauge) Process(f func()) error {
    select {
    case <-pg.ch: // 取一张执行凭据,  用完了再放回去
        f()
        pg.ch <- struct{}{}
        return nil
    default: // 总共就 10 个凭据,  所以限制了最大并发数为 10
        return errors.New("no more capacity")
    }
}
```



## 为什么说 Go 的并发好用

We have a function that calls three web services. 

1. We send data to two of those services, 
2. and then take the results of those two calls and send them to the third, returning the result. 
3. The entire process must take less than 50 milliseconds, or an error is returned.

By structuring our code with goroutines, channels, and select statements, we separate the individual steps, allow independent parts to run and complete in any order, and cleanly exchange data between the dependent parts. In addition, we make sure that no part of the program hangs, and we properly handle timeouts set both within this function and from earlier functions in the call history. 

*If you are not convinced that this is a better method for implementing concurrency,*  
*try to implement this in another language. You might be surprised at how difficult it is*.

```go
func main() {
    // We have a function that calls three web services. We send data to two of those services,
    // and then take the results of those two calls and send them to the third, returning the result.
    // The entire process must take less than 50 milliseconds, or an error is returned.
    result, err := GatherAndProcess(context.Background(), Input{})
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(result)
}

func GatherAndProcess(ctx context.Context, data Input) (COut, error) {
    // The first thing we do is set up a context that times out in 50 milliseconds.
    // After we create the context, we use a defer to make sure the context’s cancel
    // function is called.  You must call this function or resources leak.
    ctx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
    defer cancel()

    // Every channel is buffered, so that the goroutines that write to them can exit after
    // writing without waiting for a read to happen. (The errs channel has a buffer size of two,
    // because it could potentially have two errors written to it.)
    p := processor{
        outA: make(chan AOut, 1),
        outB: make(chan BOut, 1),
        inC:  make(chan CIn, 1),
        outC: make(chan COut, 1),
        errs: make(chan error, 2),
    }

    // Next, we call the launch method on processor to start three goroutines: one to call
    // getResultA, one to call getResultB, and one to call getResultC:
    p.launch(ctx, data)

    // After the goroutines are launched, we call the waitForAB method on processor:
    inputC, err := p.waitForAB(ctx)

    // We perform a standard nil check on the error.
    if err != nil {
        return COut{}, err
    }

    // If all is well, we write the inputC value to the p.inC
    // channel and then call the waitForC method on processor:
    p.inC <- inputC
    out, err := p.waitForC(ctx)
    return out, err
}

type processor struct {
    outA chan AOut
    outB chan BOut
    outC chan COut
    inC  chan CIn
    errs chan error
}

func (p *processor) launch(ctx context.Context, data Input) {
    // The goroutines for getResultA and getResultB are very similar. They call their
    // respective methods. If an error is returned, they write the error to the p.errs channel.
    // If a valid value is returned, they write the value to their channels.
    go func() {
        aOut, err := getResultA(ctx, data.A)
        if err != nil {
            p.errs <- err
            return
        }
        p.outA <- aOut
    }()
    go func() {
        bOut, err := getResultB(ctx, data.B)
        if err != nil {
            p.errs <- err
            return
        }
        p.outB <- bOut
    }()

    // Since the call to getResultC only happens if the calls to getResultA and getResultB
    // succeed and happen within 50 milliseconds, the third goroutine is slightly more complicated.
    go func() {
        select {
        // The first is triggered if the context is canceled.
        case <-ctx.Done():
            return
        // The second is triggered if the data for the call to getResultC is available.
        case inputC := <-p.inC:
            cOut, err := getResultC(ctx, inputC)
            if err != nil {
                p.errs <- err
                return
            }
            p.outC <- cOut
        }
    }()
}

func (p *processor) waitForAB(ctx context.Context) (CIn, error) {
    var inputC CIn
    count := 0
    // This uses a for-select loop to populate inputC.
    for count < 2 {
        select {
        case a := <-p.outA:
            inputC.A = a
            count++
        case b := <-p.outB:
            inputC.B = b
            count++
        case err := <-p.errs:
            // If an error was written to the p.errs channel, we return the error.
            return CIn{}, err
        case <-ctx.Done():
            // If the context has been canceled, we return an error to indicate the request is canceled.
            return CIn{}, ctx.Err()
        }
    }
    return inputC, nil
}

func (p *processor) waitForC(ctx context.Context) (COut, error) {
    select {
    case out := <-p.outC:
        // If getResultC completed successfully, we read its output from the p.outC channel and return it.
        return out, nil
    case err := <-p.errs:
        // If getResultC returned an error, we read the error from the p.errs channel and return it.
        return COut{}, err
    case <-ctx.Done():
        // Finally, if the context has been canceled, we return an error to indicate that.
        return COut{}, ctx.Err()
    }
}

type AOut struct{}
type BOut struct{}
type COut struct{}

type Input struct {
    A struct{}
    B struct{}
}

type CIn struct {
    A AOut
    B BOut
}

func getResultA(_ context.Context, _ struct{}) (AOut, error) { return AOut{}, nil }
func getResultB(_ context.Context, _ struct{}) (BOut, error) { return BOut{}, nil }
func getResultC(_ context.Context, _ CIn) (COut, error)      { return COut{}, nil }
```

