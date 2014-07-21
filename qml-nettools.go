package main

import (
    "fmt"
    "gopkg.in/qml.v0"
    "os"
    "time"
    "net"
    "strconv"
    "github.com/Cubox-/libping"
)

type App struct {
    parent     qml.Object
    Input      qml.Object
    Message    qml.Object
    Btn        qml.Object
    running    bool
}

func (a *App) HandleClick() {
    if a.running != true {
        input := a.Input.String("text")
        if len(input) > 0 {
            go Ping(input, a)
            a.running = true
            a.Btn.Set("text", "Stop")
        }
    } else {
        a.running = false
        a.Btn.Set("text", "Run")
    }
}

func Ping(input string, a *App) {
    chann := make(chan libping.Response)
    go libping.Pinguntil(input, 0, chann, time.Second)
    a.Message.Set("text", "")
    for i := range chann {
        if a.running == false {     // Quite likely very hacky...
            err := libping.Response {
                Error: fmt.Errorf("Stop ping..."),
            }
            chann <- err
            return
        } else if ne, ok := i.Error.(net.Error); ok && ne.Timeout() {
            message := fmt.Sprintf("Request timeout for icmp_seq %d",
                                   strconv.Itoa(i.Seq))
            //fmt.Printf(message)
            a.Message.Call("append", message)
            continue
        } else if i.Error != nil {
            fmt.Println(i.Error)
            a.Message.Set("text", fmt.Sprintf("Error: %s", i.Error))
        } else {
            message := fmt.Sprintf("%d bytes from %s: icmp_seq=%d time=%s",
                                   i.Readsize, i.Destination, i.Seq, i.Delay)
            a.Message.Call("append", message)
            //fmt.Printf(message)
        }
    }
    return
}

func main() {
    if err := run(); err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
    os.Exit(0)
}

func run() error {
    qml.Init(nil)
    engine := qml.NewEngine()

    component, err := engine.LoadFile("main.qml")
    if err != nil {
        return err
    }

    app := App{}

    context := engine.Context()
    context.SetVar("app", &app)

    win := component.CreateWindow(nil)

    app.Message = win.Root().ObjectByName("message")
    app.Input = win.Root().ObjectByName("input")
    app.Btn = win.Root().ObjectByName("btn")

    win.Show()
    win.Wait()

    return nil
}
