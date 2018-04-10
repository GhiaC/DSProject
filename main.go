package main

import (
	"DS/ServicePackage"
	"DS/Agency"
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

var services *lservice.Service
var agencies *Agency.Agency

func main() {
	Init()
	for {
		command := getCommand()
		if len(command) == 3 && command[0] == "add" && command[1] == "agency" {
			agencies.Add(command[2])
		} else if len(command) == 2 && command[0] == "list" && command[1] == "agencies" {
			agencies.List()
			fmt.Println()
		} else if len(command) == 2 && command[0] == "list" && command[1] == "services" {
			services.List()
			fmt.Println()
		} else if len(command) == 4 && command[0] == "list" && command[1] == "services" && command[2] == "from" {
			services.Search(command[3]).Print()
		} else if len(command) == 8 && command[0] == "order" && command[2] == "to" && command[4] == "by" && command[6] == "with" {
			pri, _ := strconv.Atoi(command[7])
			agencies.Search(command[3]).AddOrder(command[5], pri)
		} else if len(command) == 3 && command[0] == "list" && command[1] == "orders" {
			agency := agencies.Search(command[2]).Order
			fmt.Print("LIST :")
			agency.List()
			fmt.Println("Top (dequeue) : ",agency.DeQueue())
		} else if len(command) == 5 && command[0] == "add" && command[1] == "offer" && command[3] == "to" {
			agencies.AddService(command[2], command[4], services)
		} else if len(command) == 4 && command[0] == "delete" && command[2] == "from" {
			tt := agencies.Search(command[3])
			t := tt.RemoveService(command[1])
			if t {
				searchResult := agencies.SearchServiceInAllAgency(command[1])
				if !searchResult {
					services.Remove(services, nil, command[1])
				}
			}
		} else if len(command) == 6 && command[0] == "add" && command[1] == "service" {
			pay, _ := strconv.Atoi(command[5])
			services.Add(lservice.Service{nil, nil, command[2], command[3], command[4], float32(pay)})

		} else if len(command) < 8 && command[0] == "add" && command[1] == "subservice" && command[6] == "to" {
			fmt.Println("True command : add service SubService_Name Car_model Description payment to Service_Name")
		} else if len(command) < 6 && command[0] == "add" && command[1] == "service" {
			fmt.Println("True command : add service Service_Name Car_model Description payment")
		} else if len(command) == 8 && command[0] == "add" && command[1] == "subservice" && command[6] == "to" {
			pay, _ := strconv.Atoi(command[5])
			services.AddSubService(command[2], command[3], command[4], float32(pay), command[7])
		} else {
			fmt.Println("invalid command")
		}
	} //end of while
}
func Init() {
	services = lservice.New()
	agencies = Agency.New()
}

func getCommand() []string {
	fmt.Print("Enter your command: ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	commands := strings.Split(text, "\n")
	text = commands[0]
	text = string(text)
	commands = strings.Split(text, " ")
	return commands
}

//func main() {
//	services = lservice.New()
//	maxheap = Order.New()
//	agencies = Agency.New()
//
//
//	for i := 0; i < 5; i++ {
//		services.Add(lservice.Service{nil, nil, "name" + strconv.Itoa(i), "car1", "des1", 10})
//	}
//	services.AddSubService("SubName3", "SubCarName3", "des", 10, "name3")
//	//services.List()
//	fmt.Println("Agencies")
//	agencies.Add("agency1")
//	agencies.Add("agency2")
//	agencies.Add("agency3")
//	agencies.AddService("name3", "agency2", services)
//	agencies.AddService("name2", "agency2", services)
//	agencies.List()
//
//	maxheap.Add(Order.Order{"c1",1})
//	maxheap.Add(Order.Order{"c1.5",2})
//	maxheap.Add(Order.Order{"c2",13})
//	maxheap.Add(Order.Order{"c3",2})
//	maxheap.Add(Order.Order{"c4",15})
//	maxheap.Add(Order.Order{"c5",25})
//	maxheap.Add(Order.Order{"c6",3})
//	fmt.Println(maxheap.Size())
//	maxheap.List()
//	fmt.Println(maxheap.DeQueue())
//	maxheap.Add(Order.Order{"c7",14})
//	maxheap.Add(Order.Order{"c8",14})
//	fmt.Println(maxheap.Size())
//	maxheap.List()
//	fmt.Println(maxheap.DeQueue())
//	fmt.Println(" ")
//	maxheap.List()
//}
