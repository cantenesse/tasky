package main

import (
    "fmt"
    "encoding/json"
    "io/ioutil"
    "os"
    "github.com/docker/docker/api/types"
    "github.com/docker/docker/client"
)

type Task struct {
    Name          string `json:"name"`
    Email         string `json:"email"`
    Owner         string `json:"owner"`
    Schedule      string `json:"schedule"`
    EmailSubject  string `json:"subject"`
    Command       string `json:"command"`

}

func main() {
    t := readTask()
    fmt.Println(t.Name)
}

func readTask() Task {
    raw, err := ioutil.ReadFile("/tmp/task.json")
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    var task Task
    json.Unmarshal(raw, &task)

    return task

}