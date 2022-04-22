# Concurrency: Channels and Go Routines

- Channels and Go Routines are structures in Go that are used for handling concurrent programming

## Status Checker Project

- We will build a small program that is a Status Checker
- It will allows to send HTTP requests to a list of url
- We want to check if the url are up or down
- We send `Get` requests to different urls to check their status
  - We could do this with a loop and send the request one by one
  - Each requests will execute one after the other

## Serial Approach

```go
func checkLink(link string) {

    // Test the link with a Get call
    _, err := http.Get(link)

    // If error, then we have an issue
    if err != nil {
        fmt.Println(link, "might be down")
        return
    }

    // Else, we are good
    fmt.Println(link, "is up")

}
```

```go
func main() {

    // A list of urls
    urls := []string{
        "https://google.com",
        "https://facebook.com",
        "https://stackoverflow.com",
        "https://golang.org",
        "https://amazon.com",
    }

    // Loop through the links
    for _, url := range urls {
        checkLink(url)
    }

}
```

- There is a distinct delay between each fetching of the urls
- We are sitting for the request to complete before moving
  - This is a *Blocking Pattern*
  - This is a *Sequential* or *Serial*
  - The more we have links, the slower is the performance
  - This is not the best approach

## Parallel Approach

- We should send the requests in parallel independantly
- This is where the concepts of *Go Channels* and *Go Routines* come into play

### Go Routines

- Go Routines can be thought as *Threads*
- When we create a program (compile + execute), we automatically create 1 single Go Routine
  - It is the engine that execute the code
- In our current program, we have a *Blocking Call*: `http.Get()`
  - The code takes some amount of time to execute
  - The whole program is temporarily frozen here until this call is finished
- Instead of making this piece of code blocking, we will make us of Go Routine
- Go Routine implements *Async (Callback) with Threads* in Go
  - Place a `go` keyword in front of it
  - **Only use a `go` keyword in front of a function call**
  - Run the execution inside of a separate Go Routine
  - When hitting the `go` keyword, a brand new Go Routine is created to handle the task separately
  - The control flow is kept by the caller

```go
func main() {

    // A list of urls
    urls := []string{
        "https://google.com",
        "https://facebook.com",
        "https://stackoverflow.com",
        "https://golang.org",
        "https://amazon.com",
    }

    // Loop through the links
    for _, url := range urls {
        // Create a new Go Routine for each call
        go checkLink(url)
    }

}
```

### Theory of Go Routine

- One CPU Core
  - *Go Scheduler*: Runs one Go Routine until it finishes (Blocking)
    - Go Routine
    - Go Routine
    - Go Routine

With the *Go Scheduler*, even though we are creating multiple Go Routine, only one is executed at any given time

- *Go Scheduler* monitor the code that is running inside each Go Routine
  - Even though we have are spanning multiple Go Routine, they are not really executed at the same time
  - The *Go Scheduler* manages which one runs and which one pauses
  - Only 1 Go Routine executed by 1 CPU Core at a time
  - We rely on the *Go Scheduler* to define which one should be executed and when
  - **By default, Go tries to use only 1 CPU core**
    - But this behavior can be easily changed
    - When we have multiple CPU Cores, each one can run 1 single Go Routine at a time
    - When we have multiple CPU Core, we have true concurrency of execution
    - The Scheduler works very quickly behind the scenes
  - Go Philosophy: **Concurrency is not Parallelsim**
    - *Concurrency* is the composition of independently executing processes
      - Dealing with lots of things at once
      - We can have multiple threads executing code
      - If on thread blocks, another one is picked up and worked on
      - We can have concurrency on a single CPU core
    - *Parallelism* is the simultaneous execution of (possibly related) computations: Doing lots of things at once
      - We only get parallelism when we have more than one CPU core

`main` is always executing on the default Go Routine

- Main Routine is created when we launch the program
- Child Routines are created when using the `go` keyword
  - Do not have the same level of *respect* than the main routine
  - This can sometimes create some bugs

### Go Channels

- Executing the program with only `go` keyword will not return anything
  - It might look like the program is not working correctly
- The `main` routine is the single routine that can control when the program exits
  - It creates the children go routines
  - But continues to execute until the end and exit
  - It does not wait on the Children Routines to finish
- We have to use `channels` to make the coordination happen between the main routine and the child routines
  - Used to communicate between different go routines
  - Allows the main routine to be aware when the child go rountines have finalized
  - **The only way that we have to communicate between go routines**
  - Basically an intermediate channel of discussion between the go routines
  - Text messaging between the routines that automatically get sent to all the running routines that have access to that channel
  - Channels are actual values that we can pass around
- **Channels are typed**
  - Similar to all other variables in Go
  - The data that we pass through a channel is typed
- We use channels to:
  - Make the main routine wait until all children routines finish
  - Communicate back to the main routine when a child routine finish
  - Finish the main routine when all other routines finish

```go
// A channel of type string
c := make(chan string)
```

Then, we can pss it around like any other variables

```go
// Loop through the links
for _, url := range urls {
    // Create a new Go Routine for each call
    go checkLink(url, ch)
}
```

And be able to handle it in the defined functions

```go
func checkLink(link string, ch chan string) {...}
```

#### Sending value via a channel

```go
ch <- 500  // Send the value 500 via the channel
```

#### Receiving a value via a channel

- **Main Lesson: Receiving messages from a channel is a blocking call**

```go
num <-ch  // Receiving the value 500 via the channel
```

We can also simply watch the channel and use its value directly

```go
fmt.Println(<-ch)   // Watch the channel and use the value immediately
```

### Fundamentals of Channels

- **Whenever the main routine waits for a message to come through a channel, that is a blocking call**
- The main routine is put to sleep and waits for something to happen

```go
// This is a blocking call
fmt.Println(<-ch)
```

- **Main Lesson: Receiving messages from a channel is a blocking call**
- We need to check the channel multiple times for each urls
- However, for each iterations, the loop will block while waiting for the channel to return a message

```go
for i := 0; i < len(urls); i++ {
    // Receive message from the channel
    // This is a blocking call
    fmt.Println(<-ch)
}
```

- `range` can also be used with channel
  - Wait for the channel to return some value
  - After the channel returns some value, assign it to the variable

```go
// Loop through the links and check the links
for _, url := range urls {
    // Create a new Go Routine for each call
    go checkLink(url, ch)
}

for link := range ch {
    // Then span a new go routine to recheck the link again
    go checkLink(link, ch)
}
```

### Function Literal

- This is like `lambda` or `Anonymous Function`
- `time.sleep` pauses the current Go Routine
- **It is better to put this in a child routine using funcion literal so not to block the main routine**
  - **We do not want to share variable access between multiple go routines**
  - With *function literals*, we have to make sure that we pass the arguments down the line through the function literals
  - This is because Go is a *pass by value* language
  - If we do not pass explicitly like this, Go will only use the memory reference

```go
for l := range ch {
    go func (link string) {
        // time.sleep pauses the current Go Routine
        time.Sleep(time.Second * 2) // Pause of 2 seconds
        // Then span a new go routine to recheck the link again
        checkLink(link, ch)
    }(l)
}
```
