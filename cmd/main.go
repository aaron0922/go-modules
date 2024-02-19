package main

import (
	"fmt"
	"github.com/aaron0922/go-modules/internal/user"
	"github.com/aaron0922/go-modules/pkg/model"
)

func main() {
	fmt.Println("main start")
	device, err := model.GetDeviceByID(1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("main calling model.GetDeviceByID", device)
	user, err := user.GetUserByID(1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("main calling user.GetUserByID", user)
}
