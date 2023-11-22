package main

import (
	"log"
	"fmt"
	"golang.design/x/hotkey"
	"golang.design/x/hotkey/mainthread"
)

func main() { mainthread.Init(ExceOnKeyDown(hotkey.New([]hotkey.Modifier{}, hotkey.KeyS), Print)) } // Not necessary when use in Fyne, Ebiten or Gio.

func Print() {
	fmt.Println("s")
}

func ExceOnKeyDown(key hotkey, fu func) {
	err := key.Register()
	if err!= nil {
        log.Fatalf("hotkey: failed to register hotkey: %v", err)
        return
    }
	
	<-key.Keydown()
	fu
}