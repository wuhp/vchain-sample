package main

import (
    "fmt"
    "github.com/wuhp/vchain"
)

func main() {
    path := "/tmp/simple.vlog"
    fmt.Printf("Please see result in file `%s`\n", path)

    vchain.SetOutput(path)

    r1 := vchain.NewRequest("", "S1", "C1")
    vchain.Log(r1, "r1 `%s`,  before calling r2", r1.Uuid)

    r2 := vchain.NewRequest(r1.Uuid, "S2", "C1")
    vchain.Log(r2, "r2 `%s`, before calling r3", r2.Uuid)

    r3 := vchain.NewRequest(r2.Uuid, "S3", "C1")
    vchain.Log(r3, "r3 `%s`", r3.Uuid)
    r3.EndWithCommit()

    vchain.Log(r2, "r2 `%s`, after calling r3", r2.Uuid)
    r2.EndWithCommit()

    vchain.Log(r1, "r1 `%s`,  after calling r2", r1.Uuid)
    vchain.Log(r1, "r1 `%s`,  before calling r4", r1.Uuid)


    r4 := vchain.NewRequest(r1.Uuid, "S4", "C1")
    vchain.Log(r4, "r4 `%s`", r4.Uuid)
    r4.EndWithCommit()

    vchain.Log(r1, "r1 `%s`,  after calling r4", r1.Uuid)
    r1.EndWithCommit()
}
