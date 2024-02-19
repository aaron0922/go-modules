package gomodules

import (
	"fmt"

	"github.com/aaron0922/go-modules/internal/user"
	"github.com/aaron0922/go-modules/pkg/model"
)

func GetDeviceByID(id int) *model.Device {
	device, err := model.GetDeviceByID(id)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	return device
}

func GetUserByID(id int) *user.User {
	user, err := user.GetUserByID(id)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	return user
}
