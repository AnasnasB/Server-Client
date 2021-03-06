package main
import "fmt"
import "io/ioutil"
import "net"


func main(){
	conn, err := net.Dial("tcp","127.0.0.1:12444")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	file, err:= ioutil.ReadFile("test.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	conn.Write(file)
	buf := make([]byte, 2000)
	n, err := conn.Read(buf)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(buf[:n]))
}
