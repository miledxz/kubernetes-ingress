package annotations

import (
	"errors"

	"github.com/haproxytech/client-native/v2/models"

	"github.com/haproxytech/kubernetes-ingress/controller/utils"
)

type DefaultOption struct {
	name     string
	defaults *models.Defaults
}

func NewDefaultOption(n string, d *models.Defaults) *DefaultOption {
	return &DefaultOption{
		name:     n,
		defaults: d,
	}
}

func (a *DefaultOption) GetName() string {
	return a.name
}

func (a *DefaultOption) Process(input string) error {
	if input == "" {
		switch a.name {
		case "http-server-close", "http-keep-alive":
			a.defaults.HTTPConnectionMode = ""
		case "dontlognull":
			a.defaults.Dontlognull = ""
		case "logasap":
			a.defaults.Logasap = ""
		default:
			return errors.New("unknown param")
		}
		return nil
	}

	enabled, err := utils.GetBoolValue(input, a.name)
	if err != nil {
		return err
	}
	if enabled {
		switch a.name {
		case "http-server-close":
			a.defaults.HTTPConnectionMode = "http-server-close"
		case "http-keep-alive":
			a.defaults.HTTPConnectionMode = "http-keep-alive"
		case "dontlognull":
			a.defaults.Dontlognull = "enabled"
		case "logasap":
			a.defaults.Logasap = "enabled"
		default:
			return errors.New("unknown param")
		}
	} else {
		switch a.name {
		case "http-server-close", "http-keep-alive":
			a.defaults.HTTPConnectionMode = ""
		case "dontlognull":
			a.defaults.Dontlognull = "disabled"
		case "logasap":
			a.defaults.Logasap = "disabled"
		default:
			return errors.New("unknown param")
		}
	}
	return nil
}
