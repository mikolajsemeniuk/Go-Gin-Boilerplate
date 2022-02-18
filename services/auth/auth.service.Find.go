package auth

import "fmt"

func (service *service) Find() {
	service.repository.Find()
	fmt.Println("service")
}
