package main

// A directory can be added like this: fstest.MapFS{"tmp": {Mode: fs.ModeDir}}.
import "testing/fstest"

// Creates a temporary fs in memory.
func main() {
   fs := fstest.MapFS{
      "hello.txt": {
         Data: []byte("hello, world"),
      },
   }
   data, err := fs.ReadFile("hello.txt")
   if err != nil {
      panic(err)
   }
   println(string(data) == "hello, world")
}
