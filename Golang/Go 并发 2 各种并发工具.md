## Table of Contents
  - [Goroutine](#Goroutine)
    - [概述](#%E6%A6%82%E8%BF%B0)
    - [协程/线程/调度器](#%E5%8D%8F%E7%A8%8B%E7%BA%BF%E7%A8%8B%E8%B0%83%E5%BA%A6%E5%99%A8)
    - [Goroutine 很轻量](#Goroutine-%E5%BE%88%E8%BD%BB%E9%87%8F)
    - [Goroutine Lifetimes](#Goroutine-Lifetimes)
  - [The Sync Package](#The-Sync-Package)
    - [WaitGroup](#WaitGroup)
    - [Mutex](#Mutex)
    - [RWMutex](#RWMutex)
    - [Cond](#Cond)
    - [Once](#Once)
    - [Atomic](#Atomic)
    - [Pool](#Pool)
  - [Channel](#Channel)
    - [概述](#%E6%A6%82%E8%BF%B0)
    - [基础使用](#%E5%9F%BA%E7%A1%80%E4%BD%BF%E7%94%A8)
    - [Buffered Channel](#Buffered-Channel)
    - [for-range and Channels](#forrange-and-Channels)
    - [Closing a Channel](#Closing-a-Channel)
    - [使用 channel 的几种危险情况](#%E4%BD%BF%E7%94%A8-channel-%E7%9A%84%E5%87%A0%E7%A7%8D%E5%8D%B1%E9%99%A9%E6%83%85%E5%86%B5)
    - [推荐用函数封装 channel](#%E6%8E%A8%E8%8D%90%E7%94%A8%E5%87%BD%E6%95%B0%E5%B0%81%E8%A3%85-channel)
  - [Select](#Select)
    - [概述](#%E6%A6%82%E8%BF%B0)
    - [同时性和随机性](#%E5%90%8C%E6%97%B6%E6%80%A7%E5%92%8C%E9%9A%8F%E6%9C%BA%E6%80%A7)
    - [select 可添加超时](#select-%E5%8F%AF%E6%B7%BB%E5%8A%A0%E8%B6%85%E6%97%B6)
    - [for-select 循环](#forselect-%E5%BE%AA%E7%8E%AF)
    - [关于 default 分支](#%E5%85%B3%E4%BA%8E-default-%E5%88%86%E6%94%AF)
  - [The Context Package](#The-Context-Package)
    - [概述](#%E6%A6%82%E8%BF%B0)
    - [取消请求](#%E5%8F%96%E6%B6%88%E8%AF%B7%E6%B1%82)
    - [可取消的任务](#%E5%8F%AF%E5%8F%96%E6%B6%88%E7%9A%84%E4%BB%BB%E5%8A%A1)
    - [封装长期运行的任务](#%E5%B0%81%E8%A3%85%E9%95%BF%E6%9C%9F%E8%BF%90%E8%A1%8C%E7%9A%84%E4%BB%BB%E5%8A%A1)
    - [理解 Context Chain](#%E7%90%86%E8%A7%A3-Context-Chain)
    - [几个 Context 的方法](#%E5%87%A0%E4%B8%AA-Context-%E7%9A%84%E6%96%B9%E6%B3%95)
    - [用 Context 存储键值对](#%E7%94%A8-Context-%E5%AD%98%E5%82%A8%E9%94%AE%E5%80%BC%E5%AF%B9)
    - [用 Middleware 设置请求元信息](#%E7%94%A8-Middleware-%E8%AE%BE%E7%BD%AE%E8%AF%B7%E6%B1%82%E5%85%83%E4%BF%A1%E6%81%AF)
  - [其他好用的工具](#%E5%85%B6%E4%BB%96%E5%A5%BD%E7%94%A8%E7%9A%84%E5%B7%A5%E5%85%B7)
    - [Using WaitGroups](#Using-WaitGroups)
    - [Using ErrGroup](#Using-ErrGroup)

## Goroutine

### 概述

In fact, every Go program has at least one goroutine: the main goroutine, which is automatically created and started when the process begins. Goroutines are lightweight processes managed by the Go runtime. When a Go program starts, the Go runtime creates a number of threads and launches a single goroutine to run your program.   

Put very simply, a goroutine is a function that is running concurrently (remember: not necessarily in parallel!) alongside other code. You can start one simply by placing the go keyword before a function.

### 协程/线程/调度器

#### ➤ Goroutine 是一种协程

Goroutines are not OS threads, they’re a higher level of abstraction known as coroutines. Coroutines are simply concurrent subroutines (functions, closures, or methods in Go) that are nonpreemptive—that is, they cannot be interrupted. Instead, coroutines *have multiple points throughout which allow for suspension or reentry*.  

What makes goroutines unique to Go are their deep integration with Go’s runtime. Goroutines don’t define their own suspension or reentry points; Go’s runtime observes the runtime behavior of goroutines and automatically suspends them when they block and then resumes them when they become unblocked.

#### ➤ 调度器把 Goroutine 分配到线程上执行、把阻塞的 Goroutine 换成其他可执行的 Goroutine

Note that concurrency is not a property of a coroutine: something must host several coroutines simultaneously and give each an opportunity to execute—otherwise, they wouldn’t be concurrent!   

Go’s mechanism for hosting goroutines is an implementation of what’s called an M:N scheduler, which means it maps M green threads to N OS threads (Green threads are threads that are managed by a language’s runtime). Goroutines are then scheduled onto the green threads. When we have more goroutines than green threads available, *the scheduler handles the distribution of the goroutines across the available threads* and ensures that when these goroutines become blocked, other goroutines can be run.   

### Goroutine 很轻量

All of the goroutines created by your program, including the initial one, are assigned to these threads automatically by the Go runtime scheduler, just as the operating system schedules threads across CPU cores. This might seem like extra work, since the underlying operating system already includes a scheduler that manages threads and processes, but it has several benefits:  

- Goroutine creation is faster than thread creation, because you aren’t creating an operating system–level resource.  
- Goroutine initial stack sizes are smaller than thread stack sizes and can grow as needed. This makes goroutines more memory efficient.
- Switching between goroutines is faster than switching between threads because it happens entirely within the process, avoiding operating system calls that are (relatively) slow.  
- The scheduler is able to optimize its decisions because it is part of the Go process. The scheduler works with the network poller, detecting when a goroutine can be unscheduled because it is blocking on I/O.  

These advantages allow Go programs to spawn hundreds, thousands, even tens of thousands of simultaneous goroutines. If you try to launch thousands of threads in a language with native threading, your program will slow to a crawl.  

### Goroutine Lifetimes

#### ➤ When you spawn goroutines, make it clear when or whether they exit. 

Concurrent code should be written such that the goroutine lifetimes are obvious. Typically this will mean keeping synchronization-related code constrained within the scope of a function and factoring out the logic into [synchronous functions](https://google.github.io/styleguide/go/decisions#synchronous-functions). If the concurrency is still not obvious, it is important to document when and why the goroutines exit. The important part is that the goroutine’s end is evident for subsequent maintainers.

```go
func (w *Worker) Run(ctx context.Context) error {
    for item := range w.q {
        // Good: process returns at latest when the context is cancelled.
        go process(ctx, item)

        // Bad: 任务执行到一半中断了, 没有正常取消, 造成资源泄露, 数据不一致
        go process(item)
    }
}
```

#### ➤ Never start a goroutine without knowing how it will stop

```go
ch := somefunction()
go func() {
    for range ch {} // 不清楚 ch 会不会被关闭, 如果忘了关, 这个 goroutine 会泄露
}()
```

When will this goroutine exit? It will only exit when `ch` is closed. When will that occur? It’s hard to say, `ch` is returned by `somefunction.` So, depending on the state of `somefunction,` `ch` might never be closed, causing the goroutine to quietly leak.



## The Sync Package

The sync package contains the concurrency primitives that are most useful for lowlevel memory access synchronization. If you’ve worked in languages that primarily handle concurrency through memory access synchronization, these types will likely already be familiar to you.  

### WaitGroup

WaitGroup is a great way to wait for a set of concurrent operations to complete when you either don’t care about the result of the concurrent operation, or you have other means of collecting their results. If neither of those conditions are true, I suggest you use channels and a select statement instead.  

```go
func waitGroup用法() {
    var wg sync.WaitGroup

    wg.Add(1)              // Add 了几就要调用几次 Done
    go func() {
        defer wg.Done()    // 使用 defer Done 确保函数结束时 Done 被调用
        time.Sleep(2 * time.Second)
        fmt.Println("joey doesn't share food!")
    }()

    wg.Add(1)
    go func() {
        defer wg.Done()
        time.Sleep(1 * time.Second)
        fmt.Println("how you doing?")
    }()

    wg.Wait()              // 调用 Wait 进行等待
}
```

You can think of a `sync.WaitGroup` like a concurrent-safe counter: calls to `Add` increment the counter by the integer passed in, and calls to `Done` decrement the counter by one. Calls to `Wait` block until the counter is zero.  

A `sync.WaitGroup` doesn’t need to be initialized, just declared, as its zero value is useful. `Done` is called within the goroutine. To ensure that it is called, even if the goroutine panics, we use a `defer`.  

Notice that the calls to `Add` are done outside the goroutines they’re helping to track. If we didn’t do this, we would have introduced a race condition, *because we have no guarantees about when the goroutines will be scheduled*; we could reach the call to `Wait` before either of the goroutines begin.  

You’ll notice that we don’t explicitly pass the `sync.WaitGroup`. If you pass the `sync.WaitGroup` to the goroutine function and don’t use a pointer, then the function has a copy and the call to `Done` won’t decrement the original `sync.WaitGroup`. By using a closure to capture the `sync.WaitGroup`, we are assured that every goroutine is referring to the same instance.  

### Mutex

Mutex stands for “mutual exclusion” and is a way to guard critical sections of your program. If you remember from Chapter 1, a critical section is an area of your program that requires exclusive access to a shared resource. A `sync.Mutex` provides a concurrent-safe way to express exclusive access to these shared resources. 

Whereas channels share memory by communicating, a Mutex shares memory by creating a convention developers must follow to synchronize access to the memory. You are responsible for coordinating access to this memory by guarding access to it with a mutex.

```go
func 用锁保护对共享内存的修改() {
    var count int
    var lock sync.Mutex
    var wg sync.WaitGroup

    add := func(delta int, useLock bool) {
        defer wg.Done()
        if useLock {
            lock.Lock()         // 访问共享内存前,  加互斥锁
            defer lock.Unlock() // 使用 defer Unlock 确保函数退出后锁会被释放,  否则可能产生死锁
        }

        for i := 0; i < 10_0000; i++ {
            count += delta
        }
    }

    wg.Add(2)
    go add(1, true)
    go add(-1, true)
    wg.Wait()
    fmt.Println(count) // 同时进行 10 万次加 1 和减 1, 如果不使用锁进行同步,  会发现结果并不是 0
} 
```

Critical sections are so named because they reflect a bottleneck in your program. It is somewhat expensive to enter and exit a critical section, and so generally people attempt to minimize the time spent in critical sections.  

### RWMutex

The `sync.RWMutex` is conceptually the same thing as a Mutex: it guards access to memory; however, RWMutex gives you a little bit more control over the memory. You can request a lock for reading, in which case you will be granted access unless the lock is being held for writing. This means that an arbitrary number of readers can hold a reader lock so long as nothing else is holding a writer lock.   

```go
func 使用读写锁() {
    var data int
    var wg sync.WaitGroup
    wg.Add(4)

    reader := func(lock sync.Locker) {
        defer wg.Done()
        lock.Lock()
        defer lock.Unlock()

        fmt.Println("----------------------> 获取读锁")
        fmt.Printf("data is %v \n", data)
        time.Sleep(2 * time.Second)
        fmt.Println("----------------------> 释放读锁")
    }

    writer := func(lock sync.Locker) {
        defer wg.Done()
        lock.Lock()
        defer lock.Unlock()

        fmt.Println("----------------------> 获取互斥锁\n修改数据中...")
        time.Sleep(2 * time.Second)
        data++
        fmt.Println("----------------------> 释放互斥锁")
    }

    var m sync.RWMutex      // (0) 可以用 m.Lock() 加互斥锁、用 m.RLock() 加读锁
    go writer(&m)           // (1) writer goroutine 使用 &m 把互斥锁传过去
    time.Sleep(time.Second) //
    go reader(m.RLocker())  // (2) 这里使用 m.RLocker() 把读锁传过去
    go reader(m.RLocker())  // (3) 可以发现写锁(即互斥锁)被释放后,  同时有三个 goroutine 获取了读锁
    go reader(m.RLocker())
    wg.Wait()
}
```

### Cond

#### ➤ Cond 表示一个用于等待的事件

The comment for the `sync.Cond` type really does a great job of describing its purpose: ①a rendezvous point for goroutines waiting for ②announcing the occurrence of an event.  

In that definition, an “event” is any arbitrary signal between two or more goroutines that carries no information other than the fact that it has occurred. Very often you’ll want to wait for one of these signals before continuing execution on a goroutine. 

#### ➤ 为什么用轮询模拟 Cond 不太好

If we were to look at how to accomplish this without the Cond type, one naive approach to doing this is to use an infinite loop:  

```go
for condition() == false {
    time.Sleep(1*time.Millisecond)
}
```

You have to figure out how long to sleep for: too long, and you’re artificially degrading performance; too short, and you’re unnecessarily consuming too much CPU time. It would be better if there were some kind of way for a goroutine to efficiently sleep until it was signaled to wake and check its condition. This is exactly what the Cond type does for us.

#### ➤ 一个例子

```go
func 使用Cond模拟恰饭() {
    var 饭好了 bool
    var log = fmt.Println
    condition := sync.NewCond(&sync.Mutex{})

    log("我: 点菜")
    go func() {
        time.Sleep(time.Second)
        log("厨师: 开始做菜")
        time.Sleep(time.Second)
        log("厨师: 疯狂翻炒....")
        time.Sleep(time.Second)
        log("厨师: 装到碗里")
        time.Sleep(time.Second)
        log("厨师: 搞定!")

        condition.L.Lock() // 修改共享内存前要加锁
        饭好了 = true
        condition.Signal() // 通知另一个协程某某事件发生了
        condition.L.Unlock()
    }()

    condition.L.Lock() //    访问共享内存前要加锁
    for 饭好了 == false { // 注意这里用一个 for 循环, 只要饭没好就一直等
        log("我: 等呀等")
        condition.Wait() // 这行有一点魔法:
        //                     (1) 调用 Wait() 会释放锁,  并挂起协程,  直到 condition.Signal()
        //                     (2) 协程被唤醒时会重新申请锁,  等拿到锁后,  才会从 Wait() 返回
    }
    condition.L.Unlock()

    log("我: 终于好了")
    log("我: 真香~")
}
```

This approach is much more efficient. Note that the call to Wait doesn’t just block, it suspends the current goroutine, allowing other goroutines to run on the OS thread.   

#### ➤ Wait 函数的执行细节

A few other things happen when you call `Wait`: 

- upon entering Wait, Unlock is called on the Cond variable’s Locker, 
- and upon exiting Wait, Lock is called on the Cond variable’s Locker. 

In my opinion, this takes a little getting used to; it’s effectively a hidden side effect of the method. It looks like we’re holding this lock the entire time while we wait for the condition to occur, but that’s not actually the case.   

#### ➤ [为什么要在循环中检查条件、然后调用 Wait ?](https://stackoverflow.com/questions/33186280/why-we-must-use-while-for-checking-race-condition-not-if)

(1) 存在一种名为 「 spurious wakeup 」的罕见现象,  处于 Wait 的线程会毫无理由地被唤醒.  
但根据 Cond.Wait() 的源码注释,  Golang 不存在此现象:  
Unlike in other systems, Wait cannot return unless awoken by Broadcast or Signal.

(2) 线程从 Wait 恢复时要经历这样的流程: 收到通知线程被唤醒 -> 线程重新申请锁 -> 拿到锁后恢复执行  
A 线程被唤醒并不意味着 A 线程能立刻拿到锁,  有可能 B 线程抢到了锁并修改了数据,  
所以 A 线程恢复执行后,  不能假设 「 被唤醒 -> 拿到锁 」这段时间内,  数据没有发生过变化

#### ➤ 其他注意事项

1. You need to make sure that `c.Broadcast` is called after your call to `c.Wait`.  
   ( 如果发通知时没人接,  这条通知会从世界上消失,  没有任何作用...
2. `Signal()` finds the goroutine that’s been waiting the longest and notifies that, whereas  
   `Broadcast()` sends a signal to all goroutines that are waiting.   

### Once

As the name implies, `sync.Once` is a type that utilizes some sync primitives internally to ensure that *only one call to Do ever calls the function passed in*—even on different goroutines. It may seem like the ability to call a function exactly once is a strange thing to encapsulate and put into the standard package, but it turns out that the need for this pattern comes up rather frequently.

```go
func 使用Once() {
    var wg sync.WaitGroup
    var once sync.Once
    rand.Seed(time.Now().UnixMilli())

    var 下单成功 = func() {
        fmt.Println("恭喜你抢到了产品")
    }
    var 抢购者 = func() {
        defer wg.Done()
        fmt.Println("我抢!")
        time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
        once.Do(下单成功) // 3 个协程都会执行这一行,  但下单成功只会执行一次
    }

    wg.Add(3)
    go 抢购者()
    go 抢购者()
    go 抢购者()
    wg.Wait()
}
```

#### ➤ 两个注意事项

```go
func 可以给Do传不同的函数() {
    var once sync.Once
    go func() { once.Do(func() { fmt.Println("恰饭") }) }()
    go func() { once.Do(func() { fmt.Println("打游戏") }) }()
    time.Sleep(time.Millisecond) // 要么恰饭、要么打游戏,  只有一个调用会成功
}

func 调用Do时会申请锁() {
    var once sync.Once
    f := func() {
        once.Do(func() {}) // (2) 再次申请同一个锁, 由于尚未执行完 f 并释放锁, 所以会死锁
    }
    once.Do(f) // (1) 申请 once 内部的锁、然后执行 f、等 f 执行完才会释放锁
} 
```

### Atomic

- Usually you don't want or need the stuff in the `sync/atmoic` package. Try channels or the `sync` package first.

- For a simple counter, you can use the `sync/atomic` package's functions to make atomic updates without the lock:

```go
// (1) 定义全局变量, 并且加上这个注释
// viewCount must be updated atomically
var viewCount int64

func TestAtomic(*testing.T) {
    // (2) 使用 atomic 的函数进行读写
    atomic.AddInt64(&viewCount, 1)
    atomic.AddInt64(&viewCount, 1)

    // (3) 使用 atomic 的函数进行读写 ( 没错, 读写都必须用函数, 否则就不是原子的 )
    v := atomic.LoadInt64(&viewCount)
    fmt.Println(v)

    // (4) 如果目的仅限于读写一个 int64, 那么可以用 atomic 包, 对比 mutex 它的开销更小、效率更高
    //     除了 Add, Load 还有 CompareAndSwap, Store, Swap 等函数
    atomic.CompareAndSwapInt64(&viewCount, 2, 3)
    atomic.StoreInt64(&viewCount, 4)
    atomic.SwapInt64(&viewCount, 5)
}
```

Another option for sharing state is `atomic.Value`:

```go
type Config struct {
    Timeout time.Duration
}

var config atomic.Value

func TestAtomicValue(t *testing.T) {
    config.Store(&Config{2 * time.Second}) // 修改配置就是让 config 指向另一个实例
    _ = config.Load().(*Config).Timeout    // 各个 goroutine 使用 Load 读取配置

    // 注意不能修改 config 所指的值, 这个修改不是原子的, 只能修改 config 的指向
    // Once Store() has been called, a Value must not be copied.
    // Note that storing different types in the same atomic.Value will cause a panic.
    // Store of an inconsistent type panics, as does Store(nil).
}
```



### Pool

`sync.Pool` is a concurrent-safe implementation of the object pool pattern. Since Pool resides in the sync package, we’ll briefly discuss why you might be interested in utilizing it.  

At a high level, a the pool pattern is a way to create and make available a fixed number, or pool, of things for use. It’s commonly used to *constrain the creation of things that are expensive* (e.g., database connections) so that only a fixed number of them are ever created, but an indeterminate number of operations can still request access to these things. In the case of Go’s sync.Pool, this data type can be safely used by multiple goroutines.  

Pool’s primary interface is its `Get` method. When called, `Get` will first check whether there are any available instances within the pool to return to the caller, and if not, call its `New` member variable to create a new one. When finished, callers call `Put` to place the instance they were working with back in the pool for use by other processes. Here’s a simple example to demonstrate:  

```go
func 使用Pool重用对象() {
    myPool := &sync.Pool{
        New: func() interface{} {
            fmt.Println("Creating new instance.")
            return struct{}{}
        },
    }
    instance1 := myPool.Get() // 目前没有空闲, 创建一个新实例
    instance2 := myPool.Get() // 目前没有空闲, 创建一个新实例
    myPool.Put(instance2)     // 把 instance2 还回去
    myPool.Get()              // 有空闲的资源,  直接返回,  不必创建新的
    fmt.Println(instance1)
} 
```

So when working with a Pool, just remember the following points:

- When instantiating sync.Pool, give it a New member variable that is thread-safe when called.
- When you receive an instance from Get, make no assumptions regarding the state of the object you receive back.
- Make sure to call Put when you’re finished with the object you pulled out of the pool. Otherwise, the Pool is useless. Usually this is done with defer.  

## Channel

### 概述

Channels are one of the two things that set apart Go’s concurrency model. They guide you into thinking about your code as a series of stages and making data dependencies clear, which makes it easier to reason about concurrency. 

Other languages rely on global shared state to communicate between threads. This mutable shared state
makes it hard to understand how data flows through a program, which in turn makes it difficult to understand whether two threads are actually independent.  

Like a river, a channel serves as a conduit for a stream of information; values may be passed along the channel, and then read out downstream. The disparate parts of your program don’t require knowledge of each other, only a reference to the same place in memory where the channel resides.  

### 基础使用

#### ➤ 创建 channel

Channels are a built-in type created using the make function:  `ch := make(chan int)`  
*Like maps, channels are reference types*. When you pass a channel to a function, you are really passing a pointer to the channel. Also like maps and slices, the zero value for a channel is `nil`.

#### ➤ 读写 channel

You read from a channel by placing the `<-` operator to the left of the channel variable, and you write to a channel by placing it to the right:  

```go
a := <-ch // 从 channel 读取值
ch <- b   // 把 b 的值写入 channel
```

Each value written to a channel *can only be read once*. If multiple goroutines are reading from the same channel, a value written to the channel *will only be read by one of them*.  

#### ➤ 只读或只写的 channel 变量

It is rare for a goroutine to read and write to the same channel. (要么只负责读、要么只负责写)  
When assigning a channel to a variable or field, or passing it to a function, use an arrow before the chan keyword (`ch <-chan int`) to indicate that the goroutine only reads from the channel. Use an arrow after the chan keyword (`ch chan<- int`) to indicate that the goroutine only writes to the channel. Doing so allows the Go compiler to ensure that a channel is only read from or written by a function. Go will implicitly convert bidirectional channels to unidirectional channels when needed.   

### Buffered Channel

#### ➤ 读写 channel 都会阻塞, 等待另一方发送/接收

By default channels are unbuffered. Every write to an open, unbuffered channel causes the writing goroutine to *pause until another goroutine reads from* the same channel. Likewise, a read from an open, unbuffered channel causes the reading goroutine to *pause until another goroutine writes to* the same channel. This means you cannot write to or read from an unbuffered channel without at least two concurrently running goroutines.  

#### ➤ 可选设置 buffer size

Go also has buffered channels. These channels buffer a limited number of writes without blocking. If the buffer fills before there are any reads from the channel, a subsequent write to the channel pauses the writing goroutine until the channel is read. Just as writing to a channel with a full buffer blocks, reading from a channel with an empty buffer also blocks.  

A buffered channel is created by specifying the capacity of the buffer: `ch := make(chan int, 10)`  
The built-in functions len and cap return information about a buffered channel. Use `len` to find out how many values are currently in the buffer and use `cap` to find out the maximum buffer size. The capacity of the buffer cannot be changed. Passing an unbuffered channel to both len and cap returns 0. This makes sense because, by definition, an unbuffered channel doesn’t have a buffer to store values.

#### ➤ 但一般不用 buffered channel

*Most of the time, you should use unbuffered channels*. Buffered channels can be useful in certain situations, but you should create them with care. Buffered channels can easily become a premature optimization and also hide deadlocks by making them more unlikely to happen.  

### for-range and Channels

You can also read from a channel using a for-range loop:

```go
for v := range ch {
    fmt.Println(v)
}
```

Unlike other for-range loops, there is only a single variable declared for the channel, which is the value. *The loop continues until the channel is closed*, or until a `break` or `return` statement is reached.

### Closing a Channel

#### ➤ 读写已经关闭的 channel 会发生什么?

When you are done writing to a channel, you close it using the built-in close function: `close(ch)`.  
Once a channel is closed, any attempts to write to the channel or close the channel again *will panic*. 

Interestingly, *attempting to read from a closed channel always succeeds*. If the channel is buffered and there are values that haven’t been read yet, they will be returned in order. If the channel is unbuffered or the buffered channel has no more values, the *zero value* for the channel’s type is returned. This is to allow support for multiple downstream reads from a single upstream writer on the channel.

#### ➤ 如何判断 channel 的返回值是不是 close 后返回的零值

This leads to a question that might sound familiar from our experience with maps: when we read from a channel, how do we tell *the difference between a zero value that was written and a zero value that was returned because the channel is closed?*  We have a familiar answer: we use the comma ok idiom to detect whether a channel has been closed or not:  `v, ok := <-ch`

If ok is set to true, then the channel is open. If it is set to false, the channel is closed. Any time you are reading from a channel that might be closed, use the comma ok idiom to ensure that the channel is still open.  

#### ➤ 调用 close(ch) 是因为有 goroutine 在等待 channel 被关闭

The responsibility for closing a channel lies with the goroutine that writes to the channel. Be aware that *closing a channel is only required if there is a goroutine waiting for the channel to close* (such as one using a for-range loop to read from the channel). Since a channel is just another variable, Go’s runtime can detect channels that are no longer used and garbage collect them.  

#### ➤ 可用 close(ch) 让多个 goroutine 结束等待,  实现 sync.Cond 的效果

Closing a channel is also one of the ways you can *signal multiple goroutines simultaneously*. If you have n goroutines waiting on a single channel, instead of writing n times to the channel to unblock each goroutine, you can simply close the channel. Since a closed channel can be read from an infinite number of times, it doesn’t matter how many goroutines are waiting on it. We discussed using the `sync.Cond` type to perform the same behavior. You can certainly use that, but as we’ve discussed, channels are composable, so this is my favorite way to unblock multiple goroutines at the same time.

### 使用 channel 的几种危险情况

1. channel 类型的零值是 nil,  而 nil channel 很危险
2. 从 channel 读取值要用 `v, ok := <-ch` 中的 ok 来判断 v 是不是 closed channel 返回的零值
3. 写入 closed channel 会 panic, 关闭同一个 channel 两次也会 panic

As mentioned earlier, the standard pattern is to make the writing goroutine responsible for closing the channel when there’s nothing left to write. When multiple goroutines are writing to the same channel, this becomes more complicated, *as calling close twice on the same channel causes a panic*. Furthermore, *if you close a channel in one goroutine, a write to the channel in another goroutine triggers a panic as well*. The way to address this is to use a `sync.WaitGroup`.

A `nil` channel is dangerous, but there are cases where it is useful. Reading from a `nil` channel will block a program. Writes to a `nil` channel will also block. What happens if we attempt to close a nil channel? Panic!

| 操作  | Unbuffered, open                 | Unbuffered, closed                                  | Buffered, open                      | Buffered, closed                                             | Nil          |
| ----- | -------------------------------- | --------------------------------------------------- | ----------------------------------- | ------------------------------------------------------------ | ------------ |
| Read  | Pause until something is written | Return zero value ( use comma ok to see if closed ) | Pause if buffer is empty            | Return a remaining value in the buffer. If the buffer is empty, return zero value ( use comma ok to see if closed ) | Hang forever |
| Write | Pause until something is read    | PANIC                                               | Pause if buffer is full             | PANIC                                                        | Hang forever |
| Close | Works                            | PANIC                                               | Works, remaining values still there | PANIC                                                        | PANIC        |

- Send and receive operations on a `nil` channel block forver.
- Sending to an closed channel causes a panic. 
- Receiving from a closed channel is safe (会返回零值).

### 推荐用函数封装 channel

The goroutine that owns a channel should:

1. Instantiate the channel.
2. Perform writes, or pass ownership to another goroutine.
3. Close the channel.
4. Ecapsulate the previous three things in this list and expose them via a reader channel.  

```go
func ChannelOwnership() {
    chanOwner := func() <-chan int {
        resultStream := make(chan int, 5) // (1) 负责创建 channel
        go func() {
            defer close(resultStream)     // (2) 负责关闭 channel
            for i := 0; i <= 5; i++ {
                resultStream <- i         // (3) 负责写入 channel
            }
        }()
        return resultStream               // (4) 封装上述逻辑并把 channel 返给 consumer
    }

    resultStream := chanOwner()
    for result := range resultStream { // 不断处理数据,  直到 channel 被关闭
        fmt.Printf("Received: %d\n", result)
    }
    fmt.Println("Done receiving!")
}
```

#### ➤ 推荐用函数封装 channel

Notice how the lifecycle of the resultStream channel is encapsulated within the `chanOwner` function. It’s very clear that the writes will not happen on a nil or closed channel, and that the close will always happen once. This removes a large swath of risk from our program. *I highly encourage you to do what you can in your programs to keep the scope of channel ownership small so that these things remain obvious*. If you have a channel as a member variable of a struct with numerous methods on it, it’s going to quickly become unclear how the channel will behave.  

The consumer function only has access to a read channel, and therefore only needs to know how it should handle blocking reads and channel closes. In this small example, we’ve taken the stance that it’s perfectly OK to block the life of the program until the channel is closed.  

## Select

### 概述

The select statement is the other thing that sets apart Go’s concurrency model. It is the control structure for concurrency in Go, and it elegantly solves a common problem: if you can perform two concurrent operations, which one do you do first? You can’t favor one operation over others, or you’ll never process some cases. This is called starvation.

The `select` keyword allows a goroutine to read from or write to one of a set of multiple channels.

```go
func 使用select(ch, ch2, ch3, ch4 chan int, x int) {
    select {
    case v := <-ch:                      // 从 ch 读取值
        fmt.Println(v)
    case v, ok := <-ch2:                 // 从 ch2 读取值, 可以检查 ok 值
        fmt.Println(v, ok)
    case ch3 <- x:                       // 把值写到 ch3
        fmt.Println("wrote", x)
    case <-ch4:                          // 从 ch4 读取值, 但忽略返回值
        fmt.Println("got value on ch4, but ignored it")
    }
}
```

Each case in a `select` is a read or a write to a channel. If a read or write is possible for a case, it is executed along with the body of the case. Like a switch, each case in a select creates its own block. Unlike switch blocks, `case` statements in a `select` block aren’t tested sequentially, and execution won’t automatically fall through if none of the criteria are met. If none of the channels are ready, the entire `select` statement blocks. 

### 同时性和随机性

#### ➤ What happens if multiple cases have channels that can be read or written?

*The select algorithm is simple: it picks randomly* from any of its cases that can go forward; order is unimportant; each has an equal chance of being selected. This is very different from a switch statement, which always chooses the first case that resolves to true. It also cleanly resolves the starvation problem, as no case is favored over another and all are checked at the same time.

#### ➤ 使用 channel 可能发生死锁

Another advantage of select choosing at random is that it prevents one of *the most common causes of deadlocks: acquiring locks in an inconsistent order*. If you have two goroutines that both access the same two channels, they must be accessed in the same order in both goroutines, or they will deadlock. This means that neither one can proceed because they are waiting on each other.

```go
func 使用channel可能死锁() {
    ch1 := make(chan int)
    ch2 := make(chan int)

    go func() {
        ch1 <- 111 // 等待 ch1 被读取,  可惜这是等不到的事
        _ = <-ch2  // 从 ch2 读取内容
    }()

    ch2 <- 222 // 等待 ch2 被读取
    _ = <-ch1  // 从 ch1 读取内容
}
```

#### ➤ 搭配 select 和 channel 可以避免死锁

If we wrap the channel accesses in the main goroutine in a select, we avoid deadlock. *Because a select checks if any of its cases can proceed*, the deadlock is avoided.

```go
func 搭配select_channel可避免死锁() {
    ch1 := make(chan int)
    ch2 := make(chan int)

    go func() {
        ch1 <- 123456 // 写入 ch1 (①)
    }()

    var v2 int
    select {         // 使用 select 同时等待 ch1、ch2 两个 channel
    case ch2 <- 789: // 等 ch2 被读取
    case v2 = <-ch1: // 等 ch1 被写入,  从 ① 可知此 case 会被选中
    }
    fmt.Println(v2)  // select 执行完一个 case 就会退出,  这里打印 123456
}
```

### select 可添加超时

The `time.After` function takes in a `time.Duration` argument and returns a channel that will send the current time after the duration you provide it. This offers a concise way to time out in select statements.

```go
func 给select加timeout(c <-chan int) {
    select {
    case <-c:
    case <-time.After(time.Second):
        // 一秒后超时
    }
}
```

### for-select 循环

Since `select` is responsible for communicating over a number of channels, it is often embedded within a for loop. This is so common that the combination is often referred to as a for-select loop. When using a for-select loop, *you must include a way to exit the loop*. We’ll see one way to do this in “The Done Channel Pattern” on page 215.

```go
func select经常在一个无限循环中(done, ch chan int) {
    for {
        select {
        case <-done: // 必须包含用于退出无限循环的 case
            return
        case v := <-ch:
            fmt.Println(v)
        }
    }
}
```

### 关于 default 分支

Just like `switch` statements, a `select` statement can have a default clause. Also just like switch, default is selected when there are no cases with channels that can be read or written. If you want to implement *a nonblocking read or write on a channel, use a select with a default*. The following code does not wait if there’s no value to read in `ch`; it immediately executes the body of the `default`.

```go
func 通过default实现非阻塞的select(ch chan int) {
    select {
    case v := <-ch:
        fmt.Println("read from ch:", v)
    default:
        fmt.Println("no available case,  run this default branch")
    }
}
```

Having a default case inside a for-select loop is almost always the wrong thing to do. It will be triggered every time through the loop when there’s nothing to read or write for any of the cases. This makes your for loop run constantly, which uses a great deal of CPU.  

#### ➤ 非阻塞的 for-select 循环

- for-select-default 循环可理解为:  
- 看一眼 channel 中有没有数据,  没有的话也不等了,  先干点其他事,  等会再回来看看  

```go
func 非阻塞的_for_select(done <-chan int) {
    for {
        select {
        case <-done:
            return
        default:
            // ① 可以在这里插入代码
        }
        // ② 也可以在这里插入代码
    }
}
```

## The Context Package

### 概述

The context package serves two primary purposes:  
(1) To provide an API for canceling.  
(2) To provide a data-bag for transporting request-scoped data through your callgraph. 

### 取消请求

#### ➤ 有的时候需要支持取消,  结束那些没有意义的工作

Imagine that you have a request that spawns several goroutines, each one calling a different HTTP service. If one service returns an error that prevents you from returning a valid result, there is no point in continuing to process the other goroutines. In Go, this is called cancellation and the context provides the mechanism for implementation.  

#### ➤ context.WithCancel 返回新的 Context,  附带一个取消函数

To create a cancellable context, use the `context.WithCancel` function. It takes in a `context.Context` as a parameter and returns a `context.Context` and a `context.CancelFunc`. The returned context.Context is not the same context that was passed into the function. Instead, it is a child context that wraps the passed-in parent context.Context. A context.CancelFunc is a function that cancels the context, telling all of the code that’s listening for potential cancellation that it’s time to stop processing.

#### ➤ 同时调用两个服务,  若遇到错误,  则取消另一个调用

First, our `callBoth` function creates a cancellable context and a cancellation function from the passed-in context. It is important to remember that any time you create a cancellable context, you must call the `cancel` function. It is fine to call it more than once; every invocation after the first is ignored. We use a `defer` to make sure that it is eventually called. If you do not, your program will leak resources (memory and goroutines) and eventually slow down or crash.

```go
func callBoth(ctx context.Context, urlA, urlB string) {
    // 用传入的 ctx 创建可取消的 Context,  使用 defer cancel() 避免忘了取消而泄露资源
    ctx, cancel := context.WithCancel(ctx)
    defer cancel()

    // 用两个 goroutine 分别请求两个外部服务,  如果遇到了错误用 cancel() 取消另一个服务调用
    var wg sync.WaitGroup
    wg.Add(2)
    go func() {
        defer wg.Done()
        err := callService(ctx, "serviceA", urlA)
        if err != nil {
            cancel()
        }
    }()
    go func() {
        defer wg.Done()
        err := callService(ctx, "serviceB", urlB)
        if err != nil {
            cancel()
        }
    }()
    wg.Wait()
    fmt.Println("done with both")
}

func callService(ctx context.Context, name string, url string) error {
    // 使用 ctx 创建请求,  当 ctx 被取消时,  发出的请求也会被终止
    req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
}
```

### 可取消的任务

```go
func 用_Context_支持取消() {
    startSomeWork := func(ctx context.Context) error {
        done := make(chan struct{})
        go func() {
            defer close(done)
            for i := 1; i <= 5; i++ {
                select {
                // ctx.Done() 返回一个 channel,  当 ctx 被取消时这个 channel 会被关闭
                case <-ctx.Done():
                    fmt.Println("任务被取消了")
                    return
                default:
                    time.Sleep(time.Second) // 总共要 5 秒才能完成任务
                    fmt.Printf("进度: %d\n", i*20)
                }
            }
            fmt.Println("任务完成")
        }()
        <-done
        return ctx.Err()
    }

    // 可以手动 cancel() 提前取消, 注意 defer cancel() 是个好习惯,
    // 既能确保 cancel() 至少被调用一次,  也能确保相关资源被尽早释放 (不必等到 timeout 才释放)
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    err := startSomeWork(ctx) // 一个函数若想支持取消,  则把 context.Context 作为第一个参数
    fmt.Println("error:", err)
}
```

### 封装长期运行的任务

```go
func longRunningThingManager(ctx context.Context, data string) (string, error) {
    // We need to put the data returned from our long-running function into a struct
    type wrapper struct {
        result string
        err    error
    }
    // By buffering the channel, we allow the goroutine to exit,
    // even if the buffered value is never read due to cancellation.
    ch := make(chan wrapper, 1)
    go func() {
        result, err := longRunningThing(ctx, data)
        ch <- wrapper{result, err}
    }()
    select {
    case w := <-ch:
        return w.result, w.err
    case <-ctx.Done():
        return "", ctx.Err() // 被取消了, 返回 Canceled/DeadlineExceeded
    }
}

func longRunningThing(ctx context.Context, data string) (string, error) {
    for i := 0; i < 5; i++ {
        select {
        case <-ctx.Done():
            return "", ctx.Err()
        default:
            time.Sleep(time.Second)
        }
    }
    return "task result", nil
}
```

### 理解 Context Chain

#### ➤ Context Chain 从 context.Background 或 context.TODO 开始

At the top of your asynchronous call-graph, your code probably won’t have been passed a Context. To start the chain, the context package provides you with two functions to create empty instances of Context: context.Background()、context.TODO(). `Background()` simply returns an empty Context. TODO is not meant for use in production, but also returns an empty Context; TODO’s intended purpose is to serve as a placeholder.

#### ➤ 用 context.WithTimeout 创建带超时的 context

The context provides a way to control how long a request runs. The first is `context.WithTimeout`. It takes two parameters, an existing context and time.Duration that specifies the duration until the context automatically cancels. It returns a context that automatically triggers a cancellation after the specified duration as well as a cancellation function that is invoked to cancel the context immediately. The second function is `context.WithDeadline`.

#### ➤ 可以为子任务分配 timeout

When you set a time limit for the overall duration of the request, you might want to subdivide that time. And if you call another service from your service, you might want to limit how long you allow the network call to run, reserving some time for the rest of your processing or for other network calls. You control how long an individual call takes by creating a child context that wraps a parent context using context.WithTimeout or context.WithDeadline.

#### ➤ 如果 parent context 被取消,  那么 child context 会被连带取消

Any timeout that you set on *the child context is bounded by the timeout set on the parent context*; if a parent context times out in two seconds, you can declare that a child context times out in three seconds, but when the parent context times out after two seconds, so will the child.  

### 几个 Context 的方法

```go
func The_Context_Interface() {
    // Background returns an empty Context. It is never canceled, has no values, and has no deadline.
    c := context.Background()                                          // 不带取消功能
    c1, cancel := context.WithCancel(c)                                // 只能手动取消
    c2, cancel := context.WithTimeout(c, time.Second)                  // 在若干秒后取消, 或手动取消
    c3, cancel := context.WithDeadline(c, time.Now().Add(time.Second)) // 在特定时刻取消, 或手动取消

    channel := c.Done()          // 当 c 被取消时这个 channel 会被关闭
    deadline, ok := c.Deadline() // 返回截止时间,  如果无截止时间那么 ok 为 false

    // Err returns a non-nil error value after Done is closed.
    // Err returns Canceled if the context was canceled or DeadlineExceeded if the context's deadline passed.
    err := c.Err()

    // Value returns the value associated with this context for key, or nil if no value is associated with key.
    val := c.Value("key")
}
```

### 用 Context 存储键值对

#### ➤ 一个例子

To check if a value is in a context or any of its parents, use the `Value` method on context.Context. This method takes in a key and returns the value associated with the key. Again, both the key parameter and the value result are declared to be of type `interface{}`. If no value is found for the supplied key, `nil` is returned. Use the `comma, ok` idiom to type assert the returned value to the correct type.

```go
func data_bag() {
    // Context 是不可变对象, 使用 WithValue 添加键值对, 这会返回一个新实例
    // 注意 key 必须支持 == 操作符,  所以不能用 map/slice/function 作为 key
    c := context.WithValue(context.Background(), "userID", "123")

    // 从 Context 取值时要加 type assertion
    ID := c.Value("userID").(string)
    fmt.Println(ID)
}
```

#### ➤ 为什么用自定义类型作为 context key 能减少冲突

```go
func 用接口作为key() {
    type foo int
    type bar int
    m := make(map[interface{}]int)
    m[foo(1)] = 222     // 两个 key 的底层值都是 1
    m[bar(1)] = 333     // 两个 key 的底层值都是 1
    fmt.Printf("%v", m) // map[1:222 1:333]
    
    // 总之,  若用接口类型作为 map key,  只有类型和值都相同时,  两个 key 才相等
}
```

You can see that though the underlying values are the same, the different type information differentiates them within a map. Since the type you define for your package’s keys is unexported, other packages cannot conflict with keys you generate within your package.

### 用 Middleware 设置请求元信息

There is one more use for the context.  
It also provides a way to pass per-request metadata through your program.  

By default, you should prefer to pass data through explicit parameters. However, there are some cases where you cannot pass data explicitly. The most common situation is an HTTP request handler and its associated middleware. As we have seen, all HTTP request handlers have two parameters, one for the request and one for the response. *If you want to make a value available to your handler in middleware, you need to store it in the context*. Some possible situations include extracting a user from a JWT (JSON Web Token) or creating a per-request GUID that is passed through multiple layers of middleware and into your handler and business logic.

#### ➤ 一个例子  
`r.Context()` returns the context.Context associated with the request.  
`r.WithContext(ctx)` returns a new request with the old request’s state combined with the supplied ctx.

```go
// 可能许多包都会往请求上下文塞东西,  为了避免 key 冲突而使用自定义类型
type contextKey string
const contextKeyIsAuthenticated = contextKey("isAuthenticated")

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        sessionValid := true
        if !sessionValid {
            next.ServeHTTP(w, r)
            return
        }
        ctx := context.WithValue(r.Context(), contextKeyIsAuthenticated, true)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

func SomeHandler(w http.ResponseWriter, r *http.Request) {
    isAuthenticated, ok := r.Context().Value(contextKeyIsAuthenticated).(bool)
    if !ok {
        // 不存在相应的 key
    }
    
    if isAuthenticated {
        fmt.Println("用户已通过中间件的校验")
    }
}
```

#### ➤ 推荐在 handler 中把值取出来,  然后传给业务逻辑

In most cases, you want to extract the value from the context in your request handler and pass it in to your business logic explicitly. Go functions have explicit parameters and you shouldn’t use the context as a way to sneak values past the API.

There are some situations where it’s better to keep a value in the context. The tracking GUID that was mentioned earlier is one. This information is meant for management of your application; it is not part of your business state. By leaving a tracking GUID in the context, it passes invisibly through business logic that doesn’t need to know about tracking and is available when your program writes a log message or connects to another server.

## 其他好用的工具

### Using WaitGroups

Sometimes one goroutine needs to wait for multiple goroutines to complete their work. If you are waiting for a single goroutine, you can use the done channel pattern that we saw earlier. But if you are waiting on several goroutines, you need to use a WaitGroup.

As we mentioned earlier, when you have multiple goroutines writing to the same channel, you need to make sure that the channel being written to is only closed once. A `sync.WaitGroup` is perfect for this. Let’s see how it works in a function that processes the values in a channel concurrently, gathers the results into a slice, and returns the slice:  

```go
func main() {
    start := time.Now()
    tasks := generator(1, 2, 3, 4, 5, 6, 7)
    power := func(num int) int {
        time.Sleep(time.Second)
        return num * num
    }
    result := processAndGather(tasks, power, 3)
    fmt.Println(time.Since(start).Seconds(), result)
}

func generator(integers ...int) <-chan int {
    intStream := make(chan int)
    go func() {
        defer close(intStream)
        for _, i := range integers {
            select {
            case intStream <- i:
            }
        }
    }()
    return intStream
}

func processAndGather(in <-chan int, processor func(int) int, num int) []int {
    // 开 num 个 worker 不断处理任务
    out := make(chan int)
    var wg sync.WaitGroup
    wg.Add(num)
    for i := 0; i < num; i++ {
        go func() {
            defer wg.Done()
            for v := range in {
                out <- processor(v)
            }
        }()
    }

    // 等待所有 worker 都退出
    go func() {
        wg.Wait()
        close(out)
    }()

    // 不断接收 worker 的结果,  收集成一个切片
    var result []int
    for v := range out {
        result = append(result, v)
    }
    return result
}
```

In our example, we launch a monitoring goroutine that waits until all of the processing goroutines exit. When they do, the monitoring goroutine calls close on the output channel. The for-range channel loop exits when `out` is closed and the buffer is empty. Finally, the function returns the processed values.

### Using ErrGroup

The Go authors maintain a set of *utilities that supplements the standard library*. Collectively known as the `golang.org/x` packages, they include a type called `ErrGroup` that builds on top of WaitGroup to create a set of goroutines that stop processing when one of them returns an error. Read the `ErrGroup` documentation to learn more.  

```go
func 等待所有goroutine结束并检查是否存在错误() {
    type task struct {
        time  time.Duration
        error error
    }
    start := time.Now()
    group := new(errgroup.Group)
    tasks := []task{
        {1 * time.Second, errors.New("error 1")},
        {2 * time.Second, nil},
        {3 * time.Second, errors.New("error 3")},
    }

    for _, t := range tasks {
        t := t
        group.Go(func() error {
            time.Sleep(t.time)
            return t.error
        })
    }
    // Wait 方法等待所有 goroutine 结束,  并返回首个 non-nil error
    err := group.Wait()
    fmt.Println(time.Since(start).Seconds(), err)
}
```

#### ➤ errgroup.Group 的几个方法

1. `SetLimit` limits the number of active goroutines in this group to at most n. A negative value indicates no limit. Any subsequent call to the `Go` method will block until it can add an active goroutine without exceeding the configured limit.
2. `TryGo` calls the given function in a new goroutine only if the number of active goroutines in the group is currently below the configured limit. The return value reports whether the goroutine was started.

#### ➤ 等待所有子任务完成、并收集结果、处理错误

```go
type Result string
type Search func(ctx context.Context, query string) (Result, error)

func fakeSearch(kind string, duration time.Duration, err error) Search {
    return func(_ context.Context, query string) (Result, error) {
        time.Sleep(duration)
        fmt.Println(kind, "done")
        return Result(fmt.Sprintf("%s result for %q", kind, query)), err
    }
}

func main() {
    searches := []Search{
        fakeSearch("web", 1*time.Second, nil),
        fakeSearch("image", 2*time.Second, nil),
        fakeSearch("video", 3*time.Second, errors.New("video: search error")),
    }
    // 调用 Google 时内部会跑 web、image、video 三个子搜索
    results, err := Google(context.Background(), "golang", searches)
    if err != nil {
        fmt.Println(err)
        return
    }
    for _, result := range results {
        fmt.Println(result)
    }
}

func Google(ctx context.Context, query string, searches []Search) ([]Result, error) {
    // 这个函数演示了: 同时跑多个子任务、等待所有子任务完成、收集结果、处理错误
    group, ctx := errgroup.WithContext(ctx)
    results := make([]Result, len(searches))
    for i, search := range searches {
        i, search := i, search // 避免捕获会变的变量
        group.Go(func() error {
            result, err := search(ctx, query)
            if err == nil {
                results[i] = result // 每个结果都有自己的坑位
            }
            return err
        })
    }
    if err := group.Wait(); err != nil {
        return nil, err // 返回首个 non-nil error
    }
    return results, nil
}
```
