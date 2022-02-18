package auth

import "fmt"

func (s *service) Find() {
	s.repository.Find()
	fmt.Println("service")
}
