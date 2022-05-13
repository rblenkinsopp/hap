package accessory

import (
	"github.com/brutella/hap/service"
)

type SecuritySystem struct {
	*A

	SecuritySystem *service.SecuritySystem
}

// NewSecuritySystem returns a security system accessory.
func NewSecuritySystem(info Info) *SecuritySystem {
	a := New(info, TypeSecuritySystem)

	security := service.NewSecuritySystem()
	a.Ss = append(a.Ss, security.S)

	return &SecuritySystem{
		A:              a,
		SecuritySystem: security,
	}
}
