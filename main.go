package main

// A directory can be added like this: fstest.MapFS{"tmp": {Mode: fs.ModeDir}}.
import "testing/fstest"

func main() {
   fs := fstest.MapFS{
      "tmp/hello.txt": {
         Data: []byte("hello, world"),
      },
   }
   data, err := fs.ReadFile("tmp/hello.txt")
   if err != nil {
      panic(err)
   }
   println(string(data) == "hello, world")
}
