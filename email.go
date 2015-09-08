package main
import (
	"log"
	"net/smtp"
	"fmt"
	//"strconv"
)
func main() {
	auth := smtp.PlainAuth(
		"",
		"karthikeyan.net",
		"",
		"smtp.gmail.com",
		)
	fmt.Println(auth)
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"karthikeyan.net@gmail.com",
		[]string{"karthikeyantech@yahoo.co.in"},
		[]byte("This is email body."),
	)
	if err!=nil{
		log.Fatal(err)
		return
	}
}
