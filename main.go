package main
import "testing/fstest"

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
