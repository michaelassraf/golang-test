package main

import "fmt"

type Permission struct {
	Name    string
	Actions []string
}
type Role struct {
	Name        string
	Permissions []Permission
}

func main() {
	// ...
	adminRole := Role{
		Name: "admin",
		Permissions: []Permission{
			{Name: "user", Actions: []string{"read", "create", "update", "delete"}},
			{Name: "product", Actions: []string{"read", "create", "update", "delete"}},
		},
	}
	userRole := Role{
		Name: "user",
		Permissions: []Permission{
			{Name: "user", Actions: []string{"read"}},
			{Name: "product", Actions: []string{"read"}},
		},
	}
	fmt.Println(adminRole)
	fmt.Println(userRole)
}
