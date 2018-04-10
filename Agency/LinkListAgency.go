package Agency

import (
	"DS/ServicePackage"
	"fmt"
	"DS/Order"
	"DS/Arraylist"
)

type Agency struct {
	Name     string
	Services *arraylist.List
	Order    *Order.TreeOrder
	DLink    *Agency
}

func New() *Agency {
	return &Agency{"first", arraylist.New(), Order.New(), nil}
}
func NewAgency(name string) *Agency {
	return &Agency{name, arraylist.New(), Order.New(), nil}
}

func (first *Agency) Add(agencyName string) {
	for first.DLink != nil {
		first = first.DLink
	}
	if first.DLink == nil {
		first.DLink = NewAgency(agencyName)
	}
}

func (first *Agency)AddOrder(name string,pri int){
	first.Order.Add(Order.Order{name,pri})
}
func (first *Agency) SearchServiceInAllAgency(str string) (bool) {
	for first.DLink != nil {
		if first.DLink.SearchInServices(str) {
			return true
		}
		first = first.DLink
	}
	return false
}
func (first *Agency) SearchInServices(str string) (bool) {
	for i := 0; i < first.Services.Size(); i++ {
		service, _ := first.Services.Get(i)
		serviceTemp := service.(*lservice.Service)
		if serviceTemp.Name == str {
			return true
		}
	}
	return false
}
func (first *Agency) RemoveService(str string) (bool) {
	for i := 0; i < first.Services.Size(); i++ {
		service, _ := first.Services.Get(i)
		serviceTemp := service.(*lservice.Service)
		if serviceTemp.Name == str {
			first.Services.Remove(i)
			return true
		}
	}
	return false
}
func (first *Agency) AddService(nameService, agencyName string, services *lservice.Service) {
	servicePtr := services.Search(nameService)
	agency := first.Search(agencyName)
	if agency != nil && servicePtr != nil {
		agency.Services.Add(servicePtr)
	}
}

func (first *Agency) Search(agencyName string) *Agency {
	first = first.DLink
	for first != nil {
		if agencyName == first.Name {
			return first
		}
		first = first.DLink
	}
	return nil
}
func (first *Agency) List() {
	for first.DLink != nil {
		first = first.DLink
		first.Print()
	}
}
func (agency *Agency) Print() {
	fmt.Print("[Name: ", agency.Name)
	if agency.Services.Size() > 0 {
		fmt.Print(",Services : ")
	}
	for i := 0; i < agency.Services.Size(); i++ {
		service1 ,_ := agency.Services.Get(i)
		ser := service1.(*lservice.Service)
		ser.Print()
	}
	fmt.Print("] ")
}