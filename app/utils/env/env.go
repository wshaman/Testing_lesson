package env

import "os"

func EnvOrDef(name, def string) (val string) {
	if val = os.Getenv(name); val != "" {
		return val
	}
	return def
}
