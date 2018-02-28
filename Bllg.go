package main

import (
        "bytes"
        "fmt"
        "io"
        "os"
        "os/exec"
)

// ='ls -tlr'

var c1 *exec.Cmd
var c2 *exec.Cmd

func executeCmds() {

        r, w := io.Pipe()
        c1.Stdout = w
        c2.Stdin = r

        var b2 bytes.Buffer
        c2.Stdout = &b2

        c1.Start()
        c2.Start()
        c1.Wait()
        w.Close()
        c2.Wait()
        io.Copy(os.Stdout, &b2)
        return
}

func execute_command1(args2 string) {
        c1 = exec.Command("ls", "-l")
        //Stdout, _ := c1.Output()
        c2 = exec.Command("grep", "-i", args2)
        executeCmds()
        return
}

func execute_command2(args2 string) {
        c1 = exec.Command("ls", "-lt")
        //Stdout, _ := c1.Output()
        c2 = exec.Command("grep", "-i", args2)
        executeCmds()
        return
}

func execute_command3(args2 string) {
        c1 = exec.Command("ls", "-ltr")
        //Stdout, _ := c1.Output()
        c2 = exec.Command("grep", "-i", args2)
        executeCmds()
        return
}

func display_help() {
        fmt.Println("Program llg.bin")
        fmt.Println("this message maybe displayed because of an error in the inputs")
        fmt.Println("llg.bin or if alias just llg - plus search string llg test - ")
        fmt.Println("the search string is not case sensitive")
        fmt.Println("or llg.bin search_string t or if alias just llg search_string t  order by time")
        fmt.Println(" ie llg searchstring or llg searchstring t or llg searchstring tl")
        fmt.Println(" ")
        return
}

func main() {
        args := os.Args[1:]
        // args is the first input
        if len(args) == 0 {
                display_help()
                return
        }
        if len(args) == 1 {
                execute_command1(args[0])
        } else if len(args) == 2 {
                if args[1] == "t" {
                        execute_command2(args[0])
                } else if args[1] == "tl" {
                        execute_command3(args[0])
                } else {
                        display_help()
                }

        } else {
                display_help()
        }
        return
}
