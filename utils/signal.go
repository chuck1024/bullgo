/**
 * Copyright 2018 godog Author. All Rights Reserved.
 * Author: Chuck1024
 */

package utils

import (
	"github.com/xuyu/logging"
	"os"
	"os/signal"
	"syscall"
)

var (
	Shutdown = make(chan os.Signal)
	Running  = make(chan bool)
	Hup      = make(chan os.Signal)
)

func init(){
	signal.Notify(Shutdown, syscall.SIGINT)
	signal.Notify(Shutdown, syscall.SIGTERM)
	signal.Notify(Hup, syscall.SIGHUP)
}

func Signal() {
	go func() {
		for {
			select {
			case <-Shutdown:
				logging.Info("[Application.signal] receive signal SIGINT or SIGTERM, to stop server...")
				Running <- false
			case <-Hup:
			}
		}
	}()
	logging.Info("[Application.signal] register signal ok")
}
