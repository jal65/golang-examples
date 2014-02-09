package main

import (
  "strconv";
  "bytes";
  "fmt";
)

func main() {
  var buffer bytes.Buffer;
  for i := 0; i < 32; i++ {
    buffer.WriteString("a");
  }
  s := buffer.String();
  for i := 0; i < 8192; i++ {
    s += strconv.Itoa(i);
  }
  s += buffer.String();
  fmt.Println(s);
}
