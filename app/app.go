package app

import (
	"fmt"
	"log"
	"log/syslog"
        "time"
)

// The wrapper of your app
func yourApp(s server) {

        
        check_folders_exist();

	// This is just some sample code to do something
	time.Sleep(1 * time.Second)
	s.winlog.Info(1, "Still running")

	time.Sleep(2 * time.Second)
	s.winlog.Info(1, "And running")

	time.Sleep(3 * time.Second)
	s.winlog.Info(1, "But the service will keep running")

	// Notice that if this exits, the service continues to run
	// You can launch web servers, etc.
}

func check_k_exist()
{
   if exists("/c/k"){
      return;
      }
   os.MkdirAll("/c/k,       os.ModePerm);
   os.MkdirAll("/c/k/lock", os.ModePerm);
   os.MkdirAll("/c/k/etc",  os.ModePerm);
   os.MkdirAll("/c/k/bin",  os.ModePerm);
   os.MkdirAll("/c/k/log",  os.ModePerm);
   log(1,"Creating base folders");
   create_lock("k");
}


func create_lock(thename) {
   path = "/c/k/lock/" + thename + ".lock";
   newFile, err := os.Create("test.txt")
   if err != nil {
      log(err,"Failed to create lock: " + thename);
      return;
   }
   newFile.Close()
}
 
func exists(filePath string) (exists bool) {
  exists = true

  if _, err := os.Stat(filePath); os.IsNotExist(err) {
    exists = false
  }
  

  return
}

func log(status,message)
{
 s.winlog.Info(status, message);

}
