## jo shorthand syntax hai usko func ke bahar declare ni kr skte ho but u can declare full syntax

##Yes. make allocates and initializes the internal data structures needed by certain Go types.

- For maps:

- m := make(map[string]string)

- Go creates an empty map and returns a reference to it. After that, you can safely insert values:

- m["key1"] = "value1"

##Yes.
```

In Go, _ is called the blank identifier. You can use it when you want to ignore a value.

Example 1: Ignore a return value
func getData() (string, int) {
    return "hello", 42
}

func main() {
    msg, _ := getData()
    fmt.Println(msg)
}
```

##range
```
For strings:

for i, c := range s
i = starting byte index of the rune
c = rune (Unicode code point)
range automatically decodes UTF-8 characters
English letters are usually 1 byte each
Many non-English characters (Hindi, Chinese, emojis) occupy multiple bytes, so indices may jump (0, 3, 6, ...)

This is one of the reasons Go uses UTF-8 strings by default and provides rune to work correctly with
```
## closure
```
Why isn't `count` destroyed when `counter()` returns?

Normally, local variables are associated with a function's stack frame. However, Go performs escape analysis.

The compiler sees that `count` is used by a function that outlives `counter()`.

Because of that, `count` cannot safely remain only on the stack.

The compiler moves `count` to the heap (or otherwise ensures it remains valid).

The returned closure holds a reference to that variable.

So after counter() returns, the stack frame may be gone, but the captured variable still exists because something (the closure) still references it.



```

# why closures are in heap
```
go step by step
1) first 
increment:=counter()
here increment is getting the inner function which is count+=1 if we are storing count in stack means it gets destroyed already and count is not declared yet
and when we run it it will give error
so one thing we can do is if we keep count variable inside the func it will not give error or putting the count variable in heap

```