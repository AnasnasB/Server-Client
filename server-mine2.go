package main
import "encoding/json"
import "fmt"
import "net"
import "sync"

type something struct {
	First struct {
		Id int `json:"id"`
		Name string `json:"name"`
		IsHere bool `json:"isHere"`
		T interface{} `json:"t"`
		Arr []interface{} `json:"arr"`
		Str struct{
			Name string `json:"name"`
			FName string `json: "firstname"`
		}`json:"str"`
	}
}
func main(){
	l, err := net.Listen("tcp","127.0.0.1:12444")
	if err != nil {
		fmt.Println("1 - ",err)
	}
	defer l.Close()
	var wg sync.WaitGroup
	var m sync.Mutex
	
	for {
		conn, err := l.Accept ()
		if err!= nil {
			fmt.Println("2 - ",err)
			continue
		}
		wg.Add(1)
		go Enc(conn, &wg, &m )
	}
}

func Enc(conn net.Conn, wg *sync.WaitGroup, m *sync.Mutex){
	buf := make([]byte, 20000)
	n, err := conn.Read(buf)
	if err != nil {
		conn.Close()
		return
	}
	
	m.Lock()
	
	var s something
	inp := make (map[string]interface{})
	
	err = json.Unmarshal (buf[:n], &inp)
	fmt.Println(inp)
	if err != nil {
		fmt.Println(err)
		return
	}
	first:= inp["first"].(map[string]interface{})
	s.First.Id = int(first["id"].(float64))
	s.First.Name = string(first["name"].(string))
	s.First.IsHere = bool(first["isHere"].(bool))
	s.First.Arr = first["arr"].([]interface{})
	str:= first["str"].(map[string]interface{})
	s.First.Str.Name = string(str["name"].(string))
	s.First.Str.FName = string (str["firstname"].(string))
	
	fmt.Println(s.First)
	data:= []byte("Got it!")
	conn.Write (data)
	m.Unlock()
    wg.Done()
}
