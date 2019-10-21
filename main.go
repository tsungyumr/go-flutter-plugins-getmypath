package getmypath

import (
	"errors"
    //  "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "runtime"
    "strings"
	
	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
)

//  Make sure to use the same channel name as was used on the Flutter client side.
const channelName = "samples.flutter.dev/getmypath"

type MyPathPlugin struct{}

var _ flutter.Plugin = &MyPathPlugin{} // compile-time type check

func (p *MyPathPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.HandleFunc("getCurrentPath", HandleGetCurrentPath)
	return nil // no error
}

func HandleGetCurrentPath() (reply interface{}, err error) {
    file, err := exec.LookPath(os.Args[0])
    if err != nil {
        return "", err
    }
    path, err := filepath.Abs(file)
    if err != nil {
        return "", err
    }
    //fmt.Println("path111:", path)
    if runtime.GOOS == "windows" {
        path = strings.Replace(path, "\\", "/", -1)
    }
    //fmt.Println("path222:", path)
    i := strings.LastIndex(path, "/")
    if i < 0 {
        return "", errors.New(`Can't find "/" or "\".`)
    }
    //fmt.Println("path333:", path)
    return string(path[0 : i+1]), nil
}
