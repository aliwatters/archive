# Go Course - Go Lang Coding.
Notes and examples from Udemy course.

[https://www.udemy.com/learn-how-to-code](https://www.udemy.com/learn-how-to-code)

also see: https://github.com/GoesToEleven/GolangTraining

Note on terminology - not sure if all US works like this - but might help with avoiding saying "curly brackets"!

```
{} == BRACES
() == PARENTHESES
[] == BRACKETS
```


## Section 1: Setup.

Installation - already had 1.4 - moved to 1.5.2 (64bit linux)

```
wget https://storage.googleapis.com/golang/go1.5.2.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.5.2.linux-amd64.tar.gz
```


Note on terminology - not sure if all US works like this - but might help with avoiding saying "curly brackets"!

```
{} == BRACES
() == PARENTHESES
[] == BRACKETS
```



## Section 1: Setup.

Installation - already had 1.4 - moved to 1.5.2 (64bit linux)

```
wget https://storage.googleapis.com/golang/go1.5.2.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.5.2.linux-amd64.tar.gz
```

Note: go recommends a workspace folder - mine is in /home/ali/go - made a go-course-01 repo for this.

Note: hit errors with upgrade of Go

```
package runtime: C source files not allowed when not using cgo or SWIG: atomic_amd64x.c defs.c float.c heapdump.c lfstack.c malloc.c mcache.c mcentral.c mem_linux.c mfixalloc.c mgc0.c mheap.c msize.c os_linux.c panic.c parfor.c proc.c runtime.c signal.c signal_amd64x.c signal_unix.c stack.c string.c sys_x86.c vdso_linux_amd64.c
```

Seems there is old c code left over from previous installation:

```
sudo rm -rf /usr/local/go/
wget https://storage.googleapis.com/golang/go1.5.2.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.5.2.linux-amd64.tar.gz
```

Resolved Problem - HELLO WORLD! :)

Additionally - if installing first time - have to set up some go env vars. GOPATH (workspace) GOROOT (/usr/local/bin/go - or whereever) etc - easy done before.

For the course then run: `go get github.com/goestoeleven/golangtraining`

## Section 2

### Packages

Exported / unexported - capitalized is visible outside of the package, lowercase internal. public/private not used.

## Section 4 - scope

Had some problems when I tried to create and use a package called foo. Switched to bar and things worked... odd. I did have foo/foo.go - changed to bar/foo.go - might have been something to do with it.

Func expressions - eg:

inside a function you can assign an anonymous function to a var very similar to JS at this point.

We can also - write functions that return functions - works very similar to javascript (first class functions)

```
func wrapper() func() int {
  x := 0
  return func() int {
    x++
    return x
  }
}

func main() {
  increment := wrapper() // IIFE!
  fmt.Println(increment())
  fmt.Println(increment())
}
```


# Section 5 - fundamentals

## Constants

iota - increments by 1 each time used (neat)
untyped constants - allows 42 to be used with any numeric value (regardless of it's type!)
but there are also typed constants which can be used as static typed

## Pointers

show the address with & so

```
a := 42
fmt.Println("a's address is", &a)
```


note: fmt.Scan - reads into a memory address, eg.

```
var note string
fmt.Scan(&note)
```

create pointers with * and dereference pointers with *

```
a := 43
var b *int = &a
fmt.Println("b has address", b, "and value", *b)
```

# Section 6 - control flow

for loops - does more than in most languages!

```
// standard
for init; cond; post { }

// like a while
for condition { }

// endless (break to get out)
for { }

// also for will iterate over an array
for key, value := range oldMap {
  newMap[key] = value
}

// value can be ignored with
for key := range m { ... }

// key can be ignored with the blank identifier _
for _, value := range m { ... }

// strings can be treated as array like!
for pos, char := range "some utf-8 string" { ... }

// and you can multiple init & post
// reverse a in place
for i, j := 0, len(a) - 1; i < j; i = i + 1, j-1 {
  a[i], a[j] = a[j], a[i]
}

```


## Rune

a rune is character from utf-8 - it is also int32 -- maps to the utf-8 character set (4 bytes - 4*8 = 32)

```
x := 'a'
// x - is now the rune value of a - a int32 value 92
```


## String Literals

```
a := `Preserve  whitespace newlines etc and "quotes" don't need
				escaped`
b := "Single line strings"
```

## Switch statements

```
switch "value" {
	case "value":
		fmt.Println("Value found")
	case "b":
		fmt.Println("b found")
		fallthrough
	case "c":
		fmt.Println("c found") // b will also print this
	case "d", "e":
		fmt.Println("d or e found")
	default:
		fmt.Println("the default case")
}

switch {
case x < y: f1()
case x < z: f2()
case x == 4: f3()
}


switch i := x.(type) {
case nil:
	printString("x is nil")                // type of i is type of x (interface{})
case int:
	printInt(i)                            // type of i is int
case float64:
	printFloat64(i)                        // type of i is float64
case func(int) float64:
	printFunction(i)                       // type of i is func(int) float64
case bool, string:
	printString("type is bool or string")  // type of i is type of x (interface{})
default:
	printString("don't know the type")     // type of i is type of x (interface{})
}
```

Impressive - again go is really powerful. 

## more on if

```
b := true

if food := "chocolate"; b {
	fmt.Println(food) // initialized and scoped in this block only - keep scope tight
}
```

## Functions
 - see examles

## Data Stucts

### Array
 - in golang, numbered sequence of a single type that does not change in size.

### Slice
 - has an underlying array that go dynamically resizes if capacity is exceeded
 - the initialization of a slice can specify these (useful if we have some idea of order of magnitude)

```
mySlice := make([]T, size, capacity)
```

 - as elements are appended to a slice - capacity doubles in the low numbers and then moves up blocks of bytes - eg. 1536, 2048..

```
[] length: 0 capacity: 10
---------------------
10 :	 length: 11 capacity: 20
20 :	 length: 21 capacity: 40
40 :	 length: 41 capacity: 80
80 :	 length: 81 capacity: 160
160 :	 length: 161 capacity: 320
320 :	 length: 321 capacity: 672
672 :	 length: 673 capacity: 1536
1536 :	 length: 1537 capacity: 2048
2048 :	 length: 2049 capacity: 2560
2560 :	 length: 2561 capacity: 3552
3552 :	 length: 3553 capacity: 5120
5120 :	 length: 5121 capacity: 7168
7168 :	 length: 7169 capacity: 9216
9216 :	 length: 9217 capacity: 12288
12288 :	 length: 12289 capacity: 15360
15360 :	 length: 15361 capacity: 19456
19456 :	 length: 19457 capacity: 24576
24576 :	 length: 24577 capacity: 30720
30720 :	 length: 30721 capacity: 38912
38912 :	 length: 38913 capacity: 49152
49152 :	 length: 49153 capacity: 61440
61440 :	 length: 61441 capacity: 76800
76800 :	 length: 76801 capacity: 96256
96256 :	 length: 96257 capacity: 120832
```

 - there's an underlying algo which is adaptive - which is very cool
 - note: append must be used after size is exceeded.
 - concat is via append and a variadic (spread)

```
  mySlice = append(mySlice, otherSlice...)
```

 - deleting from a slice - seems a little naff

```
  mySlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
  mySlice = append(mySlice[:4], mySlice[5:]...)
  // [0 1 2 3 5 6 7 8 9]
```

#### Creating Slices

A few options

 - ```var slice []string``` - empty must use append
 - ```slice := []string{}``` -  empty must use append
 - ```slice := []string{"a", "b", "c"}``` - first 3 slots filled can be accessed eg. ```slice[1]```
 - ```slice := make([]string, 3, 10)``` - as above - size is accessible, capacity is efficient if we know the likely size of the slice
 - ```slice := make([]string, 10)``` - size and capacity are 10


Note: make seems to be the idiomatic slice creation method.

## Map

note: maps == dicts - ```val := m["key"]``` will get a zero value if key not present

idiom for ensuring that a val exists vs zero
"comma ok" idiom
```
  m := make(map[string]int)

  if val, ok := m["key"]; ok {
    // do something with val
  }
```

 - problems with ```var m map[string]string``` -- no way to add stuff in!


## Hash 
 - underlying structure in maps
 - thing -> hash algo -> bucket

## OOP

- no classes - create a type
- no instantiating - create a variable of type

- methods are added as so: (note this or self is not idiomatic)

```
type person struct {
	first string
	last  string
	age   int
}

func (p person) fullName() string {
	return p.first + " " + p.last
}
```


## Interfaces

- has them - simply implemented and implicitly - interface is a type

```
type Square struct {
	side float64
}

func (z Square) area() float64 {
	return z.side * z.side
}

type Shape interface {
	area() float64
}

func info(z Shape) {
	// Square or Circle or anything that has an area func which returns float64 can be z
}
```

Square implements the Shape interface - so if there are functions that require a Shape, Square can use them.
Square above maybe refered to as a concrete type. Interface types never implement the functionality directly.

### Empty interfaces

```interface{}``` - everything conforms to this interface so it acts as a wildcard for types


## Recievers

Simply: 

```
Recievers		Values (will accept)
------------------------------------
(t T)			T and *T
(t *T)			*T
```

## Casting and Assertion

```
(int)val  // cast
val.(int) // assert
```

## Concurrency

### Wait Groups

From the sync package;

```
var wg.sync.WaitGroup

func main() {
	wg.Add(2)
	go foo("foo:")
	go foo("bar:")
	wg.Wait()
}

func foo(msg string) {
	for i := 0; i < 100; i++ {
		fmt.Println(msg, i)
	}
	wg.Done()
}
```

Note that the scope of foo must have access to wg; wg.Done() looks like it decrements the count from Add() and allows Wait() to finish.

### Concurrency vs Parrallelism

concurrency - independently executing processes - interleaving
parallelism - simultaneous execution of computations - forking

### Race conditions

Two or more processes modifying simultaneously updating a var - leads to unexpected results. Even ++ can do this - see example 23 - 04

Run with -race flag to see

```
go run -race main.go
```
### Mutex

Mutual exclusion object

### Channels

make is used: - make create slices, maps and channels only!

```
channel := make(chan T, buffer)

eg.

c := make(chan int) // simple int channel
```

close - closes the channel to NEW items

range will drain a closed channel (as well as reading a open channel)

semaphore pattern - one channel or more - and signal to a done channel (bool) that the other channels are completed.

NOTE: 
race conditions - multiple processes (go routines) affecting the same var
deadlocks - a sender has no reciever on a channel and is blocking - which means that the receiver never gets set up!

### Deadlock

example:

```
func main() {
	c := make(chan int)
	c <- 1 // put something on the channel - but - blocks until something is ready to recieve it
	fmt.Println(<-c) // recieve something from chan c - but doesn't runas blocked
}
```

### Concurrency Additional Notes:
A nice article on concurrency: Dancing with Go’s Mutexes 
Memory Allocation
A great Stack Overflow post on memory allocation http://stackoverflow.com/questions/34197248/how-can-i-store-reference-to-the-result-of-an-operation-in-go
A great article on memory allocation http://golangtutorials.blogspot.ae/2011/06/memory-variables-in-memory-and-pointers.html


## Next Steps

https://goo.gl/k5VKHd -- Building Web Apps with Go

