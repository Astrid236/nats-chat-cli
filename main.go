package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	
	"github.com/nats-io/nats.go"

)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run main.go <NATS_SERVER> <CHANNEL> <USERNAME>")
		os.Exit(1)
	}
	
	natsServer := os.Args[1]
	channel := os.Args[2]
	username := os.Args[3]
	
	nc, err := nats.Connect(natsServer)
	if err != nil{
		log.Fatalf("Error conectando al servidor NATS: %v\n", err)
	}
	defer nc.Close()
	
	fmt.Printf("Conectado al servidor NATS: %s\n", natsServer)
	fmt.Printf("Añadido al canal: %s con usuario %s\n", channel, username)
	
	sub, err := nc.Subscribe(channel, func(msg *nats.Msg) {
		fmt.Printf("\r%s\n> ", string(msg.Data))
	})
	if err != nil {
		log.Fatalf("Error al suscribirse al canal: %v\n", err)
	}
	defer sub.Unsubscribe()
	
	input := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	
	for{
		line, err := input.ReadString('\n')
		if err != nil{
			log.Printf("Error leyendo el input: %v\n", err)
			continue
		}
		line = strings.TrimSpace(line)
		if line == ""{
			continue
		}
		
		timestamp := time.Now().Format("15:04:05")
		message := fmt.Sprintf("[%s] %s: %s", timestamp, username, line)
		
		err = nc.Publish(channel, []byte(message))
		if err != nil{
			log.Printf("Error publicando mensaje: %v\n", err)
		}
		fmt.Print("> ")
	
	}

}

