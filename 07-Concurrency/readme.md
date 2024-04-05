# Concurrency: Go Routines and Channels

---

- [Status Checker Project](#status-checker-project)
- [Naive / Blocking Serial Approach](#naive--blocking-serial-approach)
- [Parallel / Concurrent Approach](#parallel--concurrent-approach)
  - [Go Routines](#go-routines)
  - [Theory of Go Routine](#theory-of-go-routine)
  - [Go Channels](#go-channels)
    - [Quick Comparison: `async/await`](#quick-comparison-asyncawait)
    - [Sending Message Via A Channel](#sending-message-via-a-channel)
    - [Receiving/Awaiting Message From A Channel](#receivingawaiting-message-from-a-channel)
  - [Fundamental Concepts of Channels](#fundamental-concepts-of-channels)
  - [Repeating Channel Calls](#repeating-channel-calls)
  - [Function Literal](#function-literal)
    - [Using Function Literal With `go`](#using-function-literal-with-go)

---

- Channels and Go Routines are structures in Go that are used for handling concurrent programming

## Status Checker Project

- We will build a small program that is a Website Status Checker
- It will allow to send HTTP requests to a list of urls
- We want to check if the urls are up or down
- We send `Get` requests to different urls to check their status
  - We could do this with a loop and send the request one by one
  - Each requests will execute one after the other
  - **This approach is *Blocking***

## Naive / Blocking Serial Approach

```go
// Check if a URL is reachable or not.
func checkUrl(url string) {
    // Test the url with a Get call
    _, err := http.Get(url)
    // If error, then we have an issue
    if err != nil {
        fmt.Println(url, "might be down")
        return
    }

    // Else, we are good
    fmt.Println(url, "is up")
}

// Main entry of the application.
func main() {
    // A list of urls
    urls := []string{
        "https://google.com",
        "https://facebook.com",
        "https://stackoverflow.com",
        "https://golang.org",
        "https://amazon.com",
    }

    // Loop through the urls
    for _, url := range urls {
        checkUrl(url)
    }
}
```

- There is a distinct delay between each fetching of the urls
- **We are waiting for the request to complete before moving to the next one**
  - **This is a *Blocking Pattern***
  - This is *Sequential* or *Serial*
  - The more we have urls, the slower is the performance
  - This is not the best approach

## Parallel / Concurrent Approach

- We should send the requests in parallel independantly
- This is where the concepts of *Go Routines* and *Go Channels* come into play

### Go Routines

- Go Routines can be thought as *Threads*
  - They are *Green Threads*
- **When we create an executable program (compile + execute), we automatically create 1 single Go Routine**
  - It is the engine that execute the code
- In our current program, we have a *Blocking Call* in `main()`: `http.Get()`
  - The code takes some amount of time to execute
  - **The whole program is temporarily frozen here until this call is finished**
- Instead of making this piece of code *blocking*, we will make us of Go Routine
- **Go Routine implements *Async (Callback) with Threads* in Go**
  - Place a `go` keyword in front of the function call
  - **Only use a `go` keyword in front of a function call**
  - This will run the execution inside of a separate Go Routine
  - **When hitting the `go` keyword, a brand new Go Routine is created to handle the task separately**
  - The control flow is kept by the caller (`main`)

```go
// Main entry of the application.
func main() {
    // A list of urls
    urls := []string{
        "https://google.com",
        "https://facebook.com",
        "https://stackoverflow.com",
        "https://golang.org",
        "https://amazon.com",
    }

    // Loop through the urls
    for _, url := range urls {
        // Create a new Go Routine for each call
        go checkUrl(url)
    }
}
```

### Theory of Go Routine

- ***Go Scheduler***
  - Runs on One CPU Core: By default, Go will only use 1 CPU
  - Ony runs a single Go Routine at a time until it finishes or blocks
  - Essentially makes blocking calls
  - With the *Go Scheduler*, even though we are creating multiple Go Routines, only one is executed at any given time
  - We rely on the *Go Scheduler* to determine which Go Routine should be executed
- ***Go Scheduler* monitor the code that is running inside each Go Routine**
  - Even though we have them spanning multiple Go Routine, they are not really executed at the same time
  - The *Go Scheduler* manages which one runs and which one pauses
  - Only 1 Go Routine executed by 1 CPU Core at a time
  - We rely on the *Go Scheduler* to define which one should be executed and when
- **By default, Go uses only 1 CPU core**
- **But this behavior can be changed**
  - *When we have multiple CPU Cores, each core can run 1 single Go Routine at a time*
  - The *Go Scheduler* can then assign one Go Routine per core
  - **When we have multiple CPU Cores, we have true concurrency of execution**
  - The Scheduler works very quickly behind the scenes
- **Go Philosophy: Concurrency is not Parallelsim**
  - ***Concurrency* is the composition of independently executing processes**
    - Dealing with lots of things at once
    - We can have multiple Go Routine executing code
    - If one Go Routine blocks, another one is picked up and worked on
    - ***We can have concurrency even on a single CPU core***
  - ***Parallelism* is the simultaneous execution of (possibly related) computations**
    - Doing lots of things at once
    - We only get parallelism when we have more than one CPU core
    - It is possible that each core also handle concurrent executing processes
    - So multiple cores execution, and each core can handle multiple concurent tasks
- ***`main()` is always executing on the default Go Routine***
  - *Main Routine is created when we launch the program*
  - *Child Routines are created when using the `go` keyword*
    - Do not have the same level of *respect* than the main routine
    - This can sometimes create some bugs

```go
// Main entry of the application.
func main() {
    // A list of urls
    urls := []string{
        "https://google.com",
        "https://facebook.com",
        "https://stackoverflow.com",
        "https://golang.org",
        "https://amazon.com",
    }

    // Loop through the urls
    for _, url := range urls {
        // Create a new Go Routine for each call
        go checkUrl(url)
    }
}
```

HOWEVER...

- **Executing the program with only `go` keyword will not return anything**
  - It might look like the program is not working correctly
- The `main()` routine is the single routine that can control when the program exits
  - It creates the children go routines
  - But continues to execute until the end and exit
  - **`main()` does not wait on the Children Routines to finish**

### Go Channels

- **We have to use `channels` to make the coordination happen between the main routine and the child routines**
  - *Used to communicate, coodinate, and orchestrate exxecutions between different go routines*
  - Allows the main routine to be aware when the child go rountines have finished
  - **The only way that we have to allow communication between go routines**
  - Basically an intermediate channel of discussion between the go routines
  - *Text messaging* between the routines that automatically get sent to all the running routines that have access and listening on that channel
  - ***Channels are actual values that we can pass around***
- **Channels are typed**
  - Similar to all other variables in Go
  - *The data that we pass through a channel must match its type*
- We use channels to:
  - Make the main routine wait until all children routines finish
  - Communicate back to the main routine when a child routine finish
  - Finish the main routine when all other routines finish

```go
// Creating a channel of type string
ch := make(chan string)
```

- Then, we can pass it around like any other variables
- We treat it like any other values

```go
// Loop through the urls
for _, url := range urls {
    // Create a new Go Routine for each call
    // Pass the channel to the new routine
    go checkUrl(url, ch)
}
```

- And be able to handle it in the defined functions

```go
func checkUrl(url string, ch chan string) {
    ... // Use channel here
}
```

- Think of channel as a *tube* of communication
  - One end *Sends* the message
  - One end *Receeives* the message

#### Quick Comparison: `async/await`

- `go` is similar to `async`
- Receiving value from `channel` is similar to `await`

#### Sending Message Via A Channel

```go
// Send the value 500 via the channel
ch <- 500
```

#### Receiving/Awaiting Message From A Channel

- **NOTE: Receiving messages from a channel is a blocking call**

```go
// Receiving the value 500 via the channel
num <-ch
```

- We can also simply watch the channel and use its value directly when it is available

```go
// Watch the channel and use the value immediately when available
fmt.Println(<-ch)
```

- So we can do the following to send message via the channel

```go
// Check if a URL is reachable or not.
func checkUrl(url string, ch chan string) {
    // Test the url with a Get call
    _, err := http.Get(url)
    var message string

    // Error Handling
    if err != nil {
        message = fmt.Sprintf("%s might be down", url)
        // Send message via the channel
        ch <- message
        return
    }

    // Else, we are good
    message = fmt.Sprintf("%s is up", url)

    // Send message via the channel
    ch <- message
}
```

- And the following to receive data from the channel

```go
// Main entry of the application.
func main() {
    // A list of urls
    urls := []string{
        "https://google.com",
        "https://facebook.com",
        "https://stackoverflow.com",
        "https://golang.org",
        "https://amazon.com",
    }

    // Channel for communicating with go routines
    ch := make(chan string)

    // Loop through the urls
    for _, url := range urls {
        // Create a new Go Routine for each call
        // Pass the channel to the new routine
        go checkUrl(url, ch)
    }

    // Receive the messages from the channel
    // NOTE: This code is blocking
    fmt.Println(<-ch)
}
```

- **However, this approach will yield only one result each time**

### Fundamental Concepts of Channels

- **Whenever the main routine waits for a message to come through a channel, that is a blocking call**
  - The main routine is put to sleep and waits for something to happen before continuing
  - Once something happen, it continues execution

```go
// This is a blocking call
fmt.Println(<-ch)
```

- **Receiving messages from a channel is always a blocking call**
- Once `main` receive the message, it finishes its execution
  - So we would need to check the channel multiple times for each urls
  - We can use a `for` loop for multiple iterations of listening from the channel
  - However, for each iterations, the loop will block while waiting for the channel to return a message

```go
for i := 0; i < len(urls); i++ {
    // Receive message from the channel
    // This is a blocking call
    fmt.Println(<-ch)
}
```

### Repeating Channel Calls

- We could setup an infinite calls if we keep on passing the channel
- We could keep on sending the `url` as the message via the channel
- Then keep on calling `checkUrl` back with the returned message

```go
// Check if a URL is reachable or not.
func checkUrl(url string, ch chan string) {
    // Test the url with a Get call
    _, err := http.Get(url)
    // Error Handling
    if err != nil {
        fmt.Println(url, "might be down")
        // Send the url via the channel
        ch <- url
        return
    }

    // Else, we are good
    fmt.Println(url, "is up")

    // Send the url via the channel
    ch <- url
}
```

```go
func main() {
    // A slice of urls
    urls := []string{
        "https://google.com",
        "https://facebook.com",
        "https://stackoverflow.com",
        "https://go.dev",
        "https://amazon.com",
    }

    // Channel for communicating with go routines
    ch := make(chan string)

    // Loop through the urls
    for _, url := range urls {
        // Create a new Go Routine for each call
        // Pass the channel to the new routine
        go checkUrl(url, ch)
    }

    // Keep on checking the url in an infinite loop
    for {
        // Receive message from the channel
        // Span a new go routine to recheck the url again
        // This is a blocking call
        go checkUrl(<-ch, ch)
    }
}
```

- *Alternative Syntax*: `range` can also be used with channels
  - Wait for the channel to return some value
  - After the channel returns some value, assign it to the variable
- **`for l := range ch`: This is completely equivalent to the above code**
  - `range ch` will wait for the channel to return some value
  - Once `ch` return a value, it assigns it to `l`

```go
// Loop through the urls
for _, url := range urls {
    // Create a new Go Routine for each call
    // Pass the channel to the new routine
    go checkUrl(url, ch)
}

// Alternative Syntax
// Keep on checking the url in an infinite loop
for l := range ch {
    // Receive message from the channel
    // Span a new go routine to recheck the url again
    // This is a blocking call
    go checkUrl(l, ch)
}
```

- We could add a slight pause between each call of `checkUrl()`

### Function Literal

- This is like `lambda` or `Anonymous Function`

```go
func(p str) {
    // p is param
    fmt.Println(p)
}(param)
```

#### Using Function Literal With `go`

- `time.sleep` pauses the current Go Routine
- *But if we use it directly in the `main` routine, it will block the `main` routine instead*
- **It is better to put this in a child routine using function literal so not to block the main routine**
  - **We do not want to share variable access between multiple go routines**
- With *function literals*, we have to make sure that we pass the arguments down the line through the function literals
  - This is because Go is a *pass by value* language
  - If we do not pass explicitly like this by value, Go will only use the memory reference from the outer scope
  - **Never attempt to reference the same variable inside of 2 separate `go` routines: Always pass by value as an argument**

```go
// We should add a slight pause between each new call
// Keep on checking the url in an infinite loop
for l := range ch {
    // Receive message from the channel
    // This is a blocking call
    // Use a function literal so to not block the main routine
    go func(lnk string) {
        // NOTE: time.Sleep pauses the current Go Routine, which is the function literal
        // Pause of 2 seconds
        time.Sleep(time.Second * 2)

        // Recheck the url again
        checkUrl(lnk, ch)
    }(l)
}
```
