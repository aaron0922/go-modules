package gomodules

import (
	"fmt"

	"github.com/aaron0922/go-modules/internal/user"
	"github.com/aaron0922/go-modules/pkg/model"
)

func GetUserByID() *model.Device {
	device, err := model.GetDeviceByID(1)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	return device
}

func GetDeviceByID() *user.User {
	user, err := user.GetUserByID(1)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	return user
}
