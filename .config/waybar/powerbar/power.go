package main

import (
	"fmt"
	"io"
	"time"

	dbus "github.com/guelfey/go.dbus"
)

const (
	// Unknown      = PowerState(0)
	Charging = PowerState(1)
	// Discharging  = PowerState(2)
	// Empty        = PowerState(3)
	FullyCharged = PowerState(4)
)

type PowerState uint32

func (ps PowerState) String() string {
	switch ps {
	case 1:
		return "Charging"
	case 2:
		return "Discharging"
	case 3:
		return "Empty"
	case 4:
		return "Fully Charged"
	default:
		return "Unknown"
	}
}

type BatteryStatus struct {
	// Capacity is the percentage of how full the battery is
	Capacity float64
	// Usage is the power usage in watts. Negative indicates charging, positive indicates discharging
	Usage float64
	// TimeUntilFull is how long until the battery is full. 0 when fully charged or discharging
	TimeUntilFull time.Duration
	// TimeUntilEmpty is how long until the battery is empty. 0 when charging o
	TimeUntilEmpty time.Duration
	// State is the status of the battery
	State PowerState
	// IsCharging is if the State of the battery == Charging
	IsCharging bool
}

type PowerClient interface {
	io.Closer
	GetBatteryStatus() (BatteryStatus, error)
}

type powerDBusInterface struct {
	Conn   *dbus.Conn
	BatObj *dbus.Object
}

func NewPowerClient() (PowerClient, error) {
	conn, err := dbus.SystemBus()
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to session bus: %w", err)
	}

	batteryObj := conn.Object("org.freedesktop.UPower", "/org/freedesktop/UPower/devices/battery_BAT0")

	return &powerDBusInterface{
		Conn:   conn,
		BatObj: batteryObj,
	}, nil
}

func (pdi *powerDBusInterface) GetBatteryStatus() (BatteryStatus, error) {
	var bs BatteryStatus

	stateProp, err := pdi.BatObj.GetProperty("org.freedesktop.UPower.Device.State")
	if err != nil {
		return bs, fmt.Errorf("could not get online property: %w", err)
	}
	if value, ok := stateProp.Value().(uint32); ok {
		bs.State = PowerState(value)
		bs.IsCharging = PowerState(value) == Charging || PowerState(value) == FullyCharged
	}

	timeToFullProp, err := pdi.BatObj.GetProperty("org.freedesktop.UPower.Device.TimeToFull")
	if err != nil {
		return bs, fmt.Errorf("could not get time to full property: %w", err)
	}
	if value, ok := timeToFullProp.Value().(int64); ok {
		bs.TimeUntilFull = time.Duration(value) * time.Second
	}

	timeToEmptyProp, err := pdi.BatObj.GetProperty("org.freedesktop.UPower.Device.TimeToEmpty")
	if err != nil {
		return bs, fmt.Errorf("could not get time to empty property: %w", err)
	}
	if value, ok := timeToEmptyProp.Value().(int64); ok {
		bs.TimeUntilEmpty = time.Duration(value) * time.Second
	}

	energyRateProp, err := pdi.BatObj.GetProperty("org.freedesktop.UPower.Device.EnergyRate")
	if err != nil {
		return bs, fmt.Errorf("could not get energy rate property: %w", err)
	}
	if value, ok := energyRateProp.Value().(float64); ok {
		bs.Usage = value
	}

	pctProp, err := pdi.BatObj.GetProperty("org.freedesktop.UPower.Device.Percentage")
	if err != nil {
		return bs, fmt.Errorf("could not get percentage property: %w", err)
	}
	if value, ok := pctProp.Value().(float64); ok {
		bs.Capacity = value
	}

	return bs, nil
}

func (pdi *powerDBusInterface) Close() error {
	return pdi.Conn.Close()
}
