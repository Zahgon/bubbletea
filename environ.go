package tea

import uv "github.com/charmbracelet/ultraviolet"

type EnvMsg uv.Environ

func (msg EnvMsg) Getenv(key string) (v string) { _ = "STUB: not implemented"; return "" }

func (msg EnvMsg) LookupEnv(key string) (s string, v bool) {
	_ = "STUB: not implemented"
	return "", false
}
