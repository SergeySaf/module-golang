package main

import (
	"fmt"
	"time"
)

type Shop struct {
	barber *Barber
	room   chan *Client
}

type Barber struct {
	name  string
	sleep chan *Client
}

func (b *Barber) BName() string {
	return b.name
}

type Client struct {
	name string
}

func (client *Client) CName() string {
	return client.name
}

func Create_shop(barber *Barber, num int) *Shop {
	shop := new(Shop)
	shop.barber = barber
	shop.room = make(chan *Client, num)
	return shop
}

func Add_Client(name string) *Client {
	client := new(Client)
	client.name = name
	return client
}

func Add_barber(name string) *Barber {
	barber := new(Barber)
	barber.name = name
	barber.sleep = make(chan *Client)
	return barber
}

func (b *Barber) Rules(shop *Shop) {
	for {
		select {
		case client := <-shop.room:
			fmt.Printf("Barber %s started cuting hair to %s \n", b.BName(), client.CName())
			time.Sleep(time.Second * 6)
			fmt.Printf("Barber %s finished cuting %s \n", b.BName(), client.CName())
		default:
			fmt.Printf("Barber %s asleep... \n", b.BName())
			client := <-b.sleep
			fmt.Printf("Barber %s waked up by %s \n", b.BName(), client.CName())
		}
	}
}

func (client *Client) ClientLogic(shop *Shop) {
	for i := 0; i < 4; i++ {
		select {
		case shop.room <- client:
			fmt.Printf("  %s set on a chair \n", client.CName())
			select {
			case shop.barber.sleep <- client:
				fmt.Printf("  %s wake up a barber \n", client.CName())
			default:
				fmt.Printf("  %s, barber is busy\n", client.CName())
			}
			return
		default:
			fmt.Printf("There are no free chair, %s f*ck off!!!\n", client.CName())
			time.Sleep(time.Second * 3)
		}
		time.Sleep(time.Second * 1)
	}
}

func main() {
	barberName := "Nick"
	barber := Add_barber(barberName)
	shop := Create_shop(barber, 3)
	fmt.Printf("Welcome to %s's Barbershop!\n", barberName)

	go barber.Rules(shop)
	time.Sleep(time.Second * 1)
	k := 0
	clients := []string{"Dima", "Vicky", "Vanyok", "Horhe", "Vasya", "Sereha"}
	for _, n := range clients {
		client := Add_Client(n)
		go client.ClientLogic(shop)
		t := time.Duration((6 - k) / 2)
		time.Sleep(time.Second * t)
		k++
	}
	fmt.Scanln()
}
