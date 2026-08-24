# Go Goroutines, `main`, `return`, `Sleep`, and `WaitGroup` — Connected Notes

## 1. Start with a normal function call

Suppose you have:

```go
func t1() {
    fmt.Println("task 1")
}

func main() {
    t1()
    fmt.Println("main")
}
```

When you write:

```go
t1()
```

this is a **normal synchronous function call**.

The flow is:

```text
main()
  |
  |-- call t1()
  |      |
  |      |-- execute t1
  |      |
  |      |-- t1 reaches return/end
  |
  |-- continue main
  |
  |-- print "main"
  |
  |-- main returns
```

So:

> **A normal function call makes the caller wait until that function returns.**

Technically, every function returns eventually, even if you don't explicitly write `return`.

---

# 2. What does `return` actually mean?

Consider:

```go
func t1() {
    fmt.Println("task 1")
    return
}
```

`return` means:

> **This function is finished; give control back to whoever called it.**

So:

```go
func main() {
    t1()
    t2()
}
```

does:

```text
main
 ↓
t1()
 ↓
t1 executes
 ↓
return from t1
 ↓
back to main
 ↓
t2()
```

So your earlier understanding was correct:

> **For a normal function call, the caller waits until the called function returns before moving to the next statement.**

---

# 3. Now introduce `go`

This changes everything:

```go
go t1()
```

The `go` keyword tells Go:

> **Run this function as a goroutine.**

Now `main` does **not** wait for `t1` to return.

Instead:

```text
main
 |
 |-- go t1()
 |      |
 |      +------> t1 runs independently
 |
 |-- immediately continue
 |
 |-- next statement
```

Compare the two:

```go
t1()
```

means:

```text
START t1
   ↓
WAIT
   ↓
t1 finishes
   ↓
continue
```

while:

```go
go t1()
```

means:

```text
START t1
   ↓
DON'T WAIT
   ↓
continue immediately
          \
           \---- t1 runs concurrently
```

This is the central concept behind your whole question.

---

# 4. Your original program

You had:

```go
func t1() {
    time.Sleep(1000 * time.Millisecond)
    fmt.Println("task 1")
}

func t2() {
    time.Sleep(2000 * time.Millisecond)
    fmt.Println("task 2")
}

func t3() {
    time.Sleep(500 * time.Millisecond)
    fmt.Println("task 3")
}

func t4() {
    time.Sleep(1000 * time.Microsecond)
    fmt.Println("task 4")
}

func main() {
    now := time.Now()

    go t1()
    go t2()
    go t3()
    go t4()

    // ...
}
```

You started **four goroutines**:

```text
main
 ├── t1
 ├── t2
 ├── t3
 └── t4
```

They can execute concurrently.

Their sleep times are:

```text
t1 → 1000 ms
t2 → 2000 ms
t3 → 500 ms
t4 → 1 ms
```

So, conceptually:

```text
0 ms
│
├── t1 starts → sleeps 1000ms
├── t2 starts → sleeps 2000ms
├── t3 starts → sleeps 500ms
└── t4 starts → sleeps 1ms
```

Then:

```text
~1ms      → t4 finishes
~500ms    → t3 finishes
~1000ms   → t1 finishes
~2000ms   → t2 finishes
```

---

# 5. The most important rule: `main` is special

Here's the problem.

Suppose you do:

```go
func main() {
    go t1()
    go t2()
    go t3()
    go t4()
}
```

What happens?

`main` reaches its end almost immediately.

Then:

```text
main finishes
   ↓
program exits
```

The fact that `t1`, `t2`, `t3`, and `t4` are still running **doesn't keep the Go program alive**.

So you might see:

```text
(no task output)
```

or perhaps some task output depending on scheduling.

The important rule is:

> **When the `main` goroutine finishes, the Go program terminates, even if other goroutines haven't finished.**

This is why you saw something like:

```text
14.778µs
```

That tiny duration was basically showing:

> `main` completed almost immediately.

It does **not** mean the goroutines took 14 microseconds to finish.

They didn't finish.

The program ended before they could.

---

# 6. Your temporary fix: `time.Sleep()`

You then added:

```go
time.Sleep(1000 * time.Millisecond)
```

at the end of `main`.

This changes the situation.

Now `main` does:

```text
start t1
start t2
start t3
start t4

        ↓

sleep for 1000ms

        ↓

print time
        ↓
main finishes
        ↓
program exits
```

While `main` is sleeping, the goroutines have time to execute.

So:

```text
~1ms      task 4
~500ms    task 3
~1000ms   task 1
```

But:

```text
~2000ms   task 2
```

hasn't happened yet.

And at ~1000ms:

```text
main wakes up
 ↓
prints elapsed time
 ↓
main exits
 ↓
program terminates
```

Therefore `task 2` may never print.

That's exactly why you observed:

```text
task 4
task 3
task 1
~1 second
```

---

# 7. What `Sleep()` is actually doing

This is where you had a small misconception.

You were thinking:

> "`time.Sleep()` waits for the concurrent functions."

Not exactly.

It does **not** know anything about your goroutines.

When you write:

```go
time.Sleep(3 * time.Second)
```

you're simply saying:

> **"Pause this goroutine for 3 seconds."**

In this case, the goroutine being paused is the `main` goroutine.

So:

```go
time.Sleep(3 * time.Second)
```

means:

```text
main
 ↓
sleep for 3 seconds
 ↓
wake up
 ↓
continue
```

It does **not** mean:

```text
wait until t1 is finished
wait until t2 is finished
wait until t3 is finished
wait until t4 is finished
```

It just happens that while `main` is sleeping, the other goroutines can run.

---

# 8. Why did `Sleep(3000ms)` make everything work?

Your longest task is:

```go
t2 → 2000ms
```

So if you do:

```go
time.Sleep(3000 * time.Millisecond)
```

you keep `main` alive longer than every task.

Conceptually:

```text
0ms
│
├── t1 starts
├── t2 starts
├── t3 starts
└── t4 starts
│
├── 1ms    → t4 finishes
├── 500ms  → t3 finishes
├── 1000ms → t1 finishes
├── 2000ms → t2 finishes
│
└── 3000ms → main wakes
              ↓
           print time
              ↓
           program exits
```

So you get:

```text
task 4
task 3
task 1
task 2
~3 seconds
```

But notice something important:

You **didn't need 3 seconds** for the tasks.

The longest task only takes 2 seconds.

You artificially kept `main` alive for another second.

---

# 9. Concurrent execution is why the total isn't 3.5 seconds

Your tasks take:

```text
t1 = 1000ms
t2 = 2000ms
t3 = 500ms
t4 = 1ms
```

If you ran them normally:

```go
t1()
t2()
t3()
t4()
```

the approximate total would be:

```text
1000 + 2000 + 500 + 1
= 3501ms
```

because they execute one after another.

But with:

```go
go t1()
go t2()
go t3()
go t4()
```

they can overlap.

Conceptually:

```text
             0ms                1000ms             2000ms

t1           [------------------]
t2           [-------------------------------------]
t3           [---------]
t4           [-]
```

Therefore the total work completion time is approximately determined by the **slowest task**:

```text
max(1000, 2000, 500, 1)
= 2000ms
```

That's the benefit of concurrency here.

---

# 10. Another important point: output order

You saw:

```text
task 4
task 3
task 1
task 2
```

That makes sense because their sleep durations are increasing:

```text
t4 → 1ms
t3 → 500ms
t1 → 1000ms
t2 → 2000ms
```

But don't think:

> "Goroutines always execute in this order."

They don't.

The Go scheduler controls when goroutines actually get CPU time.

Concurrency means you should **not rely on goroutine scheduling for ordering** unless you explicitly synchronize it.

The sleeps make this particular order very likely, but they're not a general ordering mechanism.

---

# 11. So what about `return` in a goroutine?

Suppose:

```go
func t1() {
    time.Sleep(1000 * time.Millisecond)
    fmt.Println("task 1")
    return
}
```

If you call:

```go
go t1()
```

the `return` only ends **t1's goroutine**.

It does not tell `main`:

> "Wait for me."

Think:

```text
main
 ├── t1
 ├── t2
 └── t3
```

If `t1` returns:

```text
main
 ├── t1 ✓ finished
 ├── t2
 └── t3
```

The other goroutines continue.

But if `main` returns:

```text
main ✗
 ├── t1
 ├── t2
 └── t3
```

the **whole program terminates**.

So:

> **`return` ends the current function. It does not automatically synchronize goroutines.**

---

# 12. Your temporary solution vs proper solution

### Temporary solution

You used:

```go
time.Sleep(3000 * time.Millisecond)
```

This works approximately like:

```text
"Let's just keep main alive long enough."
```

It's useful for learning/demo purposes.

But it's not a good synchronization strategy because you're guessing how long the work will take.

For example, suppose tomorrow `t2` changes:

```go
time.Sleep(5000 * time.Millisecond)
```

Your:

```go
time.Sleep(3000 * time.Millisecond)
```

is no longer enough.

Or if `t2` takes only 500ms, you're unnecessarily waiting.

---

# 13. The proper solution: `sync.WaitGroup`

Go provides `sync.WaitGroup` for exactly this type of situation.

The idea is:

> **Start the goroutines, then wait until all of them report that they're finished.**

Conceptually:

```text
main
 │
 ├── start t1 ────────────────→ Done()
 ├── start t2 ───────────────────────→ Done()
 ├── start t3 ─────────→ Done()
 ├── start t4 ──→ Done()
 │
 └── Wait()
       │
       │
       ├── waiting...
       ├── waiting...
       ├── waiting...
       │
       └── all Done()
              ↓
           continue
              ↓
          main exits
```

Example:

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func t1(wg *sync.WaitGroup) {
    defer wg.Done()

    time.Sleep(1000 * time.Millisecond)
    fmt.Println("task 1")
}

func t2(wg *sync.WaitGroup) {
    defer wg.Done()

    time.Sleep(2000 * time.Millisecond)
    fmt.Println("task 2")
}

func t3(wg *sync.WaitGroup) {
    defer wg.Done()

    time.Sleep(500 * time.Millisecond)
    fmt.Println("task 3")
}

func t4(wg *sync.WaitGroup) {
    defer wg.Done()

    time.Sleep(1000 * time.Microsecond)
    fmt.Println("task 4")
}

func main() {
    var wg sync.WaitGroup

    wg.Add(4)

    go t1(&wg)
    go t2(&wg)
    go t3(&wg)
    go t4(&wg)

    wg.Wait()

    fmt.Println("all tasks finished")
}
```

Now `main` doesn't guess:

```go
time.Sleep(3 * time.Second)
```

Instead it says:

```go
wg.Wait()
```

which means:

> **"I will wait until all four goroutines call `Done()`."**

---

# 14. The entire concept in one picture

```text
NORMAL FUNCTION CALL
────────────────────

main
 │
 ├── t1()
 │     │
 │     ├── work
 │     └── return
 │
 ├── t2()
 │     │
 │     ├── work
 │     └── return
 │
 └── finish


GOROUTINE
─────────

main
 │
 ├── go t1() ──────────────→ t1 works
 │
 ├── go t2() ───────────────────→ t2 works
 │
 ├── go t3() ─────→ t3 works
 │
 └── main continues
        │
        └── main returns
              ↓
          PROGRAM EXITS


GOROUTINE + SLEEP (YOUR TEMP FIX)
─────────────────────────────────

main
 │
 ├── go t1() ───────────────→ t1
 ├── go t2() ─────────────────────→ t2
 ├── go t3() ───────→ t3
 ├── go t4() ─→ t4
 │
 ├── Sleep(3 sec)
 │       │
 │       ├── t4 finishes
 │       ├── t3 finishes
 │       ├── t1 finishes
 │       └── t2 finishes
 │
 └── main returns
       ↓
    program exits


GOROUTINE + WAITGROUP (PROPER FIX)
───────────────────────────────────

main
 │
 ├── go t1() ───────────────→ Done()
 ├── go t2() ─────────────────────→ Done()
 ├── go t3() ───────→ Done()
 ├── go t4() ─→ Done()
 │
 └── Wait()
       │
       └── waits until ALL are Done()
                    ↓
                continue
                    ↓
                main returns
```

# 15. The mental model to remember

Keep these four rules in your head:

### Rule 1

```go
t1()
```

**Normal call → wait for `t1` to return.**

### Rule 2

```go
go t1()
```

**Goroutine → don't wait for `t1`; continue immediately.**

### Rule 3

```go
time.Sleep(...)
```

**Sleep the current goroutine. It doesn't know or care whether your other goroutines are finished.**

### Rule 4

```go
wg.Wait()
```

**Actually wait for the goroutines you've synchronized with the `WaitGroup`.**

And the most important special case:

> **When `main` returns, the entire Go program ends.**

So your original problem was **not that the goroutines weren't running**.

They were being started.

The problem was:

```text
goroutines are running
       ↓
main doesn't wait
       ↓
main finishes
       ↓
program terminates
       ↓
unfinished goroutines don't get to finish
```

Your `time.Sleep()` was a **temporary way to keep `main` alive**.

`sync.WaitGroup` is the **proper synchronization solution**.
