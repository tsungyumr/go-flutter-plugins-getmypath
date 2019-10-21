package getmypath

import (
    "fmt"
	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
	"github.com/lxn/win"
	"unsafe"
	"syscall"
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

func HandleGetCurrentPath(arguments interface{}) (reply interface{}, err error) {
    b := make([]uint16, syscall.MAX_PATH)
	var bb win.BROWSEINFO
    rv := win.SHBrowseForFolder(&bb)

	if rv != 0 {
		res := win.SHGetPathFromIDList(rv, (*uint16)(unsafe.Pointer(&b[0])))
		
		if res == true {
			path := syscall.UTF16ToString(b)
			fmt.Println(path)
			return string(path), nil	
		}
	}

    return nil, nil
}
