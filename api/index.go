package handler

import (
    "os"
    "github.com/raopx/glance/glance"
    "net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    os.Exit(glance.Main())
}
