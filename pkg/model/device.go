package model

import (
	"errors"
)

type Device struct {
	ID int
	Name string
}

var devices = map[int]Device{
	1: Device{ID: 1, Name: "Device 1"},
	2: Device{ID: 2, Name: "Device 2"},
	3: Device{ID: 3, Name: "Device 3"},
}


func GetDeviceByID(id int) (*Device, error) {
	if id > 3 {
		return nil, errors.New("Device not found")
	}
	
	device := devices[id]
	return &device, nil
}
