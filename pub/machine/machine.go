package machine

import (
	"errors"
	"net"
	"time"
)

type Status int

const (
	Up Status = iota
	Down
	None
)

func (i Status) String() string {
	switch i {
	case Up:
		return "up"
	case Down:
		return "down"
	case None:
		return "none"
	default:
		return ""
	}
}

type Machine struct {
	Name   string `json:"name"`
	net.IP `json:"ip"`
	UpTime time.Duration `json:"uptime"`
	Status Status        `json:"status"`
}

type PostRequestMachine struct {
	Name   string `json:"name"`
	IP     string `json:"ip"`
	UpTime string `json:"uptime"`
	Status int    `json:"status"`
}

func NewMachine(r PostRequestMachine) (m Machine, err error) {
	m.Name = r.Name
	m.Status = Status(r.Status)
	m.IP = net.ParseIP(r.IP)
	if m.IP == nil {
		return Machine{}, errors.New("ip: failed parse ip")
	}
	m.UpTime, err = time.ParseDuration(r.UpTime)
	if err != nil {
		return Machine{}, err
	}

	return m, err
}

func NewPostRequestMachine(m Machine) (r PostRequestMachine) {
	r.Name = m.Name
	r.IP = m.IP.String()
	r.UpTime = m.UpTime.String()
	r.Status = int(m.Status)

	return r
}
