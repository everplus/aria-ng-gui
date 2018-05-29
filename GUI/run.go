package main

import (
    "os/exec"
)

func main(){
    cmd := exec.Command("./GUI/aria-ng.exe")
    cmd.Output()
}