package main
import "fmt"
func main() {	

	mychannel := make(chan string,4)

	go func(){
		for i:=0;i<4;i++{
			mychannel <- "Hello from channel message "+fmt.Sprint(i+1)
			fmt.Printf("Sent message %d to channel\n",i+1)
		}
		close(mychannel)
	}()

	for msg := range mychannel{
		fmt.Println("Received message from channel:",msg)
	}
}
