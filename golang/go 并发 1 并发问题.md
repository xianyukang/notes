## Table of Contents
  - [并发程序中的基础问题](#%E5%B9%B6%E5%8F%91%E7%A8%8B%E5%BA%8F%E4%B8%AD%E7%9A%84%E5%9F%BA%E7%A1%80%E9%97%AE%E9%A2%98)
    - [Race Condition](#Race-Condition)
    - [Atomicity](#Atomicity)
    - [加锁解决 Data Race](#%E5%8A%A0%E9%94%81%E8%A7%A3%E5%86%B3-Data-Race)
    - [锁不解决 Race Condition](#%E9%94%81%E4%B8%8D%E8%A7%A3%E5%86%B3-Race-Condition)
    - [加锁为什么不是首选方式?](#%E5%8A%A0%E9%94%81%E4%B8%BA%E4%BB%80%E4%B9%88%E4%B8%8D%E6%98%AF%E9%A6%96%E9%80%89%E6%96%B9%E5%BC%8F)
    - [死锁、活锁、饥饿](#%E6%AD%BB%E9%94%81%E6%B4%BB%E9%94%81%E9%A5%A5%E9%A5%BF)
  - [性能](#%E6%80%A7%E8%83%BD)
    - [并行化的性能](#%E5%B9%B6%E8%A1%8C%E5%8C%96%E7%9A%84%E6%80%A7%E8%83%BD)
    - [何时适合用并发?](#%E4%BD%95%E6%97%B6%E9%80%82%E5%90%88%E7%94%A8%E5%B9%B6%E5%8F%91)



## 并发程序中的基础问题

### Race Condition

➤ race condition 是指什么问题 ?

1. 程序的输出是随机的、不确定的

2. 取决于哪个子过程先执行、哪个子过程跑得快 (也就是取决于运气)

➤ 为什么会出现 race condition ?

A race condition occurs when two or more operations must execute in the correct order,  
but the program has not been written so that this order is guaranteed to be maintained.  
两个操作必须以正确的顺序先后执行,  但程序中忘了编写代码来确保执行顺序

➤ 一个 race condition 例子

```go
func main() {
	var data int

	go func() {
		data++                              // 修改 data
	}()

	if data == 0 {                          // 读取 data
		fmt.Println("data is", data)        // 再次读取 data 
	}
}
```

➤ 这段程序的输出有三种可能:

1. 不打印任何东西,  在第一次读取 data 前把 data 改成了 1
2. 打印了 0,  在第二次读取 data 后才修改 data
3. 打印了 1,  在两次读取 data 之间修改了 data

仅仅只是三行代码,  就能让程序的产生很多的不确定性  
因为根本就没有写任何代码来保证执行顺序,  所以实际输出取决于运气.

➤ 如果运气好 bug 也能运行,  但说不定哪天就挂了

Race conditions are one of the most insidious types of concurrency bugs because they may not show up until years after the code has been placed into production. They are usually precipitated by a change in the environment the code is executing in, or an unprecedented occurrence. In these cases, the code seems to be behaving correctly, but in reality, there’s just a very high chance that the operations will be executed in order. Sooner or later, the program will have an unintended consequence.  

➤ 什么是 data race ? 存在什么问题 ?

像上面的代码, 两个线程同时读写同一块数据,  就叫做 data race,  对读取线程而言、data race 存在如下问题:

1. 观测到的数据不可靠、会失效,  明明检测过 data == 0 才进行打印,  结果打印的时候 data 却是 1
2. 会读到中间状态,  如果一个操作要修改多个数据,  那么可能读到一半改了一半没改这样不一致的中间状态

### Atomicity

```go
type Counter struct {
	mu  sync.Mutex
	ten int
	one int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.one++
	if c.one == 10 {
		c.one = 0
		c.ten++
	}
}

func (c *Counter) Get() (int, int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ten, c.one
}
```

➤ 什么是原子性?

1. 看不到中间状态,  只能看到 `Inc` 执行前/执行后的状态,  `Get` 不可能返回 c.one 为 10 这样的临时/中间状态
2. 一次只能由一个线程进行操作,  例如读的时候不允许其他线程进行读取/修改、`Get/Get`、`Get/Inc`、`Inc/Inc` 之间都是互斥的
3. 要么全都成功、要么全都失败,  如果执行 `Inc` 时发生异常,  相关数据应该回到 `Inc` 之前的状态

➤ 所以怎么判断原子性? 

假设 increment() 函数涉及了 a、b、c 三个变量,  
如果 increment() 函数执行时,  不可能有其他线程同时读取或修改 a、b、c 变量  
那么认为 increment() 函数是原子的.

所以单线程应用中的所有操作都是原子的,  因为总共就一个线程  
多个线程用 i++ 修改各自的 i 变量时,  i++ 也是原子操作,  因为不可能有其他线程读写当前线程独享的 i 变量

### 加锁解决 Data Race

Let’s say we have a data race: two concurrent processes are attempting to access the same area of memory, and the way they are accessing the memory is not atomic. If there is a data race and the output of the program will be completely nondeterministic.

➤ 以独占形式访问共享资源、的一小段代码叫做临界区

In fact, there’s a name for a section of your program that needs exclusive access to a shared resource. This is called a critical section. There are various ways to guard your program’s critical sections, and Go has some better ideas on how to deal with this, but one way to solve this problem is to synchronize access to the memory between your critical sections.

```go
func 加锁保护临界区() {
	var l sync.Mutex
	var data int

	go func() {
		l.Lock() // declare exclusive access
		data++
		l.Unlock()
	}()

	l.Lock()     // declare exclusive access
	if data == 0 {
		fmt.Printf("data is %v \n", data)
	} else {
		fmt.Printf("data is %v \n", data)
	}
	l.Unlock()
}
```

➤ 任何时候访问临界资源都要记得加锁、否则会失去临界区的互斥性/独占性

In this example we’ve created a convention for developers to follow. Anytime developers want to access the `data` variable’s memory, they must first call `Lock`, and when they’re finished they must call `Unlock`. Code between those two statements can then assume it has exclusive access to data; we have successfully synchronized access to the memory. Also note that if developers don’t follow this convention, we have no guarantee of exclusive access!   

### 锁不解决 Race Condition

*You may have noticed that while we have solved our data race, we haven’t actually solved our race condition!* The order of operations in this program is still nondeterministic; we’ve just narrowed the scope of the nondeterminism a bit. In this example, either the goroutine will execute first, or both our if and else blocks will. We still don’t know which will occur first in any given execution of this program. Later, we’ll explore the tools to solve this kind of issue properly.  

### 加锁为什么不是首选方式?

It is true that you can solve some problems by synchronizing access to the memory, but as we just saw, it doesn’t automatically solve data races or logical correctness. Further, it can also create maintenance and performance problems.  

➤ 维护性

Note that earlier we mentioned that we had created a convention for declaring we needed exclusive access to some memory. Conventions are great, but they’re also easy to ignore. By synchronizing access to the memory in this manner, you are counting on all other developers to follow the same convention now and into the future.  

➤ 性能问题

Synchronizing access to the memory in this manner also has performance issue.  
The calls to Lock can make our program slow. 

➤ 复杂性

- What size should my critical sections be?  
- Are my critical sections entered and exited repeatedly? (注意 sync.Mutex 不是可重入锁)

Answering these two questions in the context of your program is an art, and this adds to the difficulty in synchronizing access to the memory.  

### 死锁、活锁、饥饿

- 死锁,  两个线程互相等待对方手上的资源,  并且在取得全部资源前,  不会释放自己手上的资源,  导致无限等待
- [活锁,](https://www.zhihu.com/question/20566246)  虽然线程在运行,  但是在做无用功,  比如不断的失败重试,  但一直不成功,  白白浪费 cpu 资源  
- 饥饿指资源分配不均,  一些线程太贪婪,  导致有的线程很难得到、甚至几乎得不到资源来完成任务  
  Keep in mind that starvation can also apply to CPU, memory, file handles, database connections: any resource that must be shared is a candidate for starvation.  

## 性能

### 并行化的性能

For example, imagine you were writing a program that was largely GUI based: a user is presented with an interface, clicks on some buttons, and stuff happens. This type of program is bounded by one very large sequential portion of the pipeline: human interaction. No matter how many cores you make available to this program, it will always be bounded by how quickly the user can interact with the interface.  

Now consider a different example, calculating digits of pi. Thanks to a class of algorithms called spigot algorithms, *significant gains can be made by making more cores available to your program*, and your new problem becomes how to combine and store the results.  

总之开了多线程也不一定能提升性能,  并行程序的性能受限于有多少代码必须同步运行  
若一个问题的子任务之间没有相互依赖,  则适合用并行化提升性能,  比如请求 10 个网页.

### 何时适合用并发?

Let’s consider an example. Say you are writing a web service that calls three other web services. We send data to two of those services, and then take the results of those two calls and send them to the third, returning the result. The entire process must take less than 50 milliseconds, or an error should be returned. This is a good use of concurrency, because 

1. there are parts of the code that need to *perform I/O that can run without interacting with each other*, 
2. there’s a part where we *combine the results*, and 
3. there’s *a limit on how long* our code needs to run.  
