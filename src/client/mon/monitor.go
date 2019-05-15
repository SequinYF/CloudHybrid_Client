package m

import (
	"../../sd"
	"../trans"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
)

//客户端监控程序
func Client_monitor(user sd.UserProfile) {

	fmt.Println("Monitor start...")

	//建立监控队列
	watch, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	defer watch.Close()

	//阻塞
	done := make(chan bool)

	go func() {
		for {
			select {
			case ev := <-watch.Events:
				{
					if ev.Op&fsnotify.Create == fsnotify.Create {
						log.Println("create file", ev.Name)
						trans.RsyncToServer(ev.Name, user)
					}
					if ev.Op&fsnotify.Write == fsnotify.Write {
						log.Println("write file", ev.Name)
					}
					if ev.Op&fsnotify.Remove == fsnotify.Remove {
						log.Println("remove file", ev.Name)
						trans.RsyncToServer(ev.Name, user)
					}
					if ev.Op&fsnotify.Rename == fsnotify.Rename {
						log.Println("rename file", ev.Name)
					}
					if ev.Op&fsnotify.Chmod == fsnotify.Chmod {
						log.Println("change chomd", ev.Name)
					}
				}
			case err := <-watch.Errors:
				{
					log.Println("errror", err)
					return
				}
			}
		}
	}()

	err = watch.Add(sd.FilePath)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
