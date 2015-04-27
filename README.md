# webhook
webhook for gitlib


# Installation

`go get github.com/Xuyuanp/webhook`

# Usage

```go
package main

import (
    "log"
    "net/http"
    "os/exec"

    "github.com/Xuyuanp/webhook"
)

func main() {
    wh := webhook.New()

    wh.PushEventHandler = func(event *webhook.PushEvent) {
        if event.Refs != "refs/heads/develop" {
            return
        }
        cmd := exec.Command("sh", "/your/shell/script/location")
        out, err := cmd.Output()
        if err != nil {
            log.Println(err)
            return
        }
        log.Println(string(out))
    }

    log.Fatal(http.ListenAndServe(":1234", wh))
}
```
