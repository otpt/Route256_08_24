package main

import (
  "fmt"
  "bufio"
  "os"
  "bytes"
)

func main() {
    var in *bufio.Reader
    var out *bufio.Writer
    in = bufio.NewReader(os.Stdin)
    out = bufio.NewWriter(os.Stdout)
    defer out.Flush()

    var t int
    fmt.Fscan(in, &t)
    // in.ReadString('\n')

    for i := 0; i < t; i++ {
        var n int
        fmt.Fscan(in, &n)
        in.ReadString('\n')

        // fmt.Fprintf(out, "n = %d\n", n)
        var buffer bytes.Buffer
        for j := 0; j < n; j++ {
            text, _ := in.ReadString('\n')
            buffer.WriteString(text)
        }

        fmt.Fprintf(out, "%s\n", buffer.String())
    }
}
