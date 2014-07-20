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

const (
)

type App struct {
	parent     qml.Object
	Input      qml.Object
	Message    qml.Object
	pinging    bool
}

func (a *App) HandleClick() {
	input := a.Input.String("text")
	if len(input) > 0 {
		go ping(input, a)
	}
}

func ping(input string, a *App){
	a.pinging = true
    chann := make(chan libping.Response, 100)
    go libping.Pinguntil(input, 0, chann, time.Second)
    for i := range chann {
        if ne, ok := i.Error.(net.Error); ok && ne.Timeout() {
        	message := fmt.Sprintf("Request timeout for icmp_seq %d\n",
        	                       strconv.Itoa(i.Seq))
            //fmt.Printf(message)
            a.Message.Set("text", message)
            continue
        } else if i.Error != nil {
            fmt.Println(i.Error)
            a.Message.Set("text", fmt.Sprintf("Error: %s", i.Error))
        } else {
        	message := fmt.Sprintf("%d bytes from %s: icmp_seq=%d time=%s\n",
        	                       i.Readsize, i.Destination, i.Seq, i.Delay)
        	a.Message.Set("text", message)
            //fmt.Printf(message)
        }
    }
    a.pinging = false
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

	app := App{
	}

	context := engine.Context()
	context.SetVar("app", &app)

	win := component.CreateWindow(nil)

	app.Message = win.Root().ObjectByName("message")
	app.Input = win.Root().ObjectByName("input")

	win.Show()
	win.Wait()

	return nil
}
