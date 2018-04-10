package lservice

import (
	"fmt"
)

type Service struct {
	Next                        *Service
	SubService                  *Service
	Name, CarModel, Description string
	Payment                     float32
}

func New() *Service {
	return &Service{}
}
func (first *Service) Remove(servicePtr, previous *Service, str string) (bool){
	if servicePtr == nil {
		return false
	}
	if servicePtr.Name == str {
		previous.Next = servicePtr.Next
		return true
	}
	result := first.Remove(servicePtr.SubService,servicePtr, str)
	if result {
		return result
	}
	result = first.Remove(servicePtr.Next,servicePtr, str)
	if result {
		return result
	}
	return false
}
func NewService(Name, CarModel, Description string, payment float32) *Service {
	return &Service{Name: Name, CarModel: CarModel, Description: Description, Payment: payment, Next: nil, SubService: nil}
}

func (first *Service) List() {
	list(first.Next)
}
func list(servicePtr *Service) {
	fmt.Print("{")
	for servicePtr != nil {
		servicePtr.Print()
		servicePtr = servicePtr.Next
	}
	fmt.Print("}")
}

func (service *Service) Print() {
	fmt.Print(" [Name: ", service.Name, ",", "car model: ", service.CarModel, ",Payment: ", service.Payment)
	if (service.SubService != nil) {
		fmt.Print(",Child <")
		service.SubService.Print()
		fmt.Print("> ")
	}
	fmt.Print("] ")
}
func (first *Service) Add(service Service) {
	for first.Next != nil {
		first = first.Next
	}
	first.Next = &service
}

func (first *Service) AddSubService(subServiceName, carModel, descriptionCustomer string, pay float32, serviceName string) {
	resultSearch := first.SearchService(first, serviceName)
	if resultSearch != nil {
		newService := NewService(subServiceName, carModel, descriptionCustomer, pay)
		cur := resultSearch.SubService
		if cur == nil {
			resultSearch.SubService = newService
		} else {
			for cur.Next != nil {
				cur = cur.Next
			} // check
			cur.Next = newService
		}
	} else {
		fmt.Println("not find service")
	}
}
func (first *Service) Search(serviceName string) *Service {
	resultSearch := first.SearchService(first, serviceName)
	return resultSearch
}

func (first *Service) SearchService(servicePtr *Service, serviceName string) *Service {
	if servicePtr == nil {
		return nil
	}
	if servicePtr.Name == serviceName {
		return servicePtr
	}
	result := first.SearchService(servicePtr.SubService, serviceName)
	if result != nil {
		return result
	}
	result = first.SearchService(servicePtr.Next, serviceName)
	if result != nil {
		return result
	}
	return nil
}
