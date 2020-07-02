This directory is intentionally empty.

Go runtime is statically linked in the  application & memory is allocated and deallocated automatically;
There are two ways to initialize using make() and new(), they can be used with maps, slices, channels 
new() allocates & results in zero storage to return a memory address but does not initialize the memory; if you try to set a value it will cause an error 
make() allocates & initialize, non zero storage where one can assign the value. 
Mostly one shall use the make() unless a reason to use the new()
example when we do not use make it can cause a panic assignment error, becos we do not initialize the memory beforehand; 
var m map[string]int
m["key"] = 42
fmt.Println(m)
The above code will cause an error becos we have not initialized the memory before assigning, in order to avoid this we shall use the make function
m := make(map[string]int)
m["key"] = 42
fmt.Println(m)

Go runtime has the garbage collector that is responsible for deallocating the memory (GC), objects that are either nil or out of scope


https://golang.org/pkg/runtime/ 