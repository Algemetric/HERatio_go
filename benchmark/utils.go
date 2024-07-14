package benchmark

import (
	scheme "github.com/Algemetric/HERatio/Implementation/Golang"
	"github.com/Algemetric/HERatio/Implementation/Golang/oracle"
	"github.com/Algemetric/HERatio/Implementation/Golang/params"
)

func keychainSetup(pl params.Literal) (*scheme.Keychain, error) {
	// Parameters.
	p, err := params.New(pl)
	if err != nil {
		return nil, err
	}
	// Oracle.
	o := new(oracle.Oracle)
	// Keychain.

	kc, err := scheme.NewKeychain(o, p)
	if err != nil {
		return nil, err
	}

	return kc, nil
}
