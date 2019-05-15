package trans

import (
	"../../sd"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
)

const defaultBufSize = 4096

type FileInfo struct {
	Name    string
	Size    int
	Message []byte
	op      uint32
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

//增量传输
func RsyncToServer(filename string, user sd.UserProfile) {
	/*
		cmd := exec.Command("/usr/bin/expect", "/Users/sequin_yf/go/src/CloudHybrid/mainprocedure/src/trans/rsync.sh", "/Users/sequin_yf/go/src/CloudHybrid/monitor_dir/fff", "sequin", "ting199787")
		cmd.Stderr = os.Stdout
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	*/

	command := `/Users/sequin_yf/go/src/CloudHybrid/mainprocedure/src/trans/rsync.sh`
	fmt.Println(user.Passwd, user.Name)
	cmd := exec.Command("/usr/bin/expect", command, filename, user.Name, user.Passwd)
	output, err := cmd.Output()
	if err != nil {
		log.Printf("Execute Shell:%s failed with error:%s\n", command, err)
		return
	}

	log.Printf("Execute Shell:%s finished with output:\n%s\n", command, string(output))

}

//全量传输
func UploadToServer(conn net.Conn, filepath string, op uint32) bool {

	var file *os.File

	err := checkFileIsExist(filepath)
	if err != true {
		log.Fatal("file not exist")
	}
	file, er := os.Open(filepath)
	if er != nil {
		log.Println(filepath + "cant open")
	}

	var temp_size int = 0
	temp_file := make([]byte, defaultBufSize)

	for {
		n, _ := file.Read(temp_file)
		if 0 == n {
			break
		}
		temp_size += n
	}

	fi := &FileInfo{
		filepath,
		temp_size,
		temp_file,
		op,
	}

	b, er := json.Marshal(fi)
	if er != nil {
		log.Fatal("marshal")
	}
	fmt.Println(string(b))
	conn.Write(b)
	return true
}
