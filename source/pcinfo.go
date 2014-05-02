package main
import (
         "fmt"
         "os/exec"
         "log"
)

func main() {
    out, err := exec.Command("sw_vers").Output()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s/n", out)
}
