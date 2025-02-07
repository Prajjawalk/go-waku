package cliutils

import (
	"strings"

	"github.com/multiformats/go-multiaddr"
)

type MultiaddrSlice struct {
	Values *[]multiaddr.Multiaddr
}

func (k *MultiaddrSlice) Set(value string) error {
	ma, err := multiaddr.NewMultiaddr(value)
	if err != nil {
		return err
	}
	*k.Values = append(*k.Values, ma)
	return nil
}

func (v *MultiaddrSlice) String() string {
	if v.Values == nil {
		return ""
	}

	var output []string
	for _, v := range *v.Values {
		output = append(output, v.String())
	}

	return strings.Join(output, ", ")
}

type MultiaddrValue struct {
	Value **multiaddr.Multiaddr
}

func (v *MultiaddrValue) Set(value string) error {
	ma, err := multiaddr.NewMultiaddr(value)
	if err != nil {
		return err
	}

	*v.Value = new(multiaddr.Multiaddr)
	**v.Value = ma

	return nil
}

func (v *MultiaddrValue) String() string {
	if v.Value == nil || *v.Value == nil {
		return ""
	}
	return (**v.Value).String()
}
