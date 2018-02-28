package main
import (
	"net"
	"time"
	"os"
)
func main() {
	argsWithoutProg := os.Args[1:]
	packet := "Message"
	if len(argsWithoutProg) > 0 { 
		packet = argsWithoutProg[0]
	} 

	for{
		time.Sleep(time.Second)
		ServerAddr,err := net.ResolveUDPAddr("udp","255.255.255.255:40000") 
		if err !=nil{
			continue 
		} 
		Conn, err := net.DialUDP("udp", nil, ServerAddr) 
		if err !=nil{
			continue 
		} 
		for{
			_ , err = Conn.Write([]byte(packet))
			if err != nil{
				break
			} 
			time.Sleep(time.Second) 

		} 
	}
}
