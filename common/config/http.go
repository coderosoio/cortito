package config

import "fmt"

// HTTP holds configuration for HTTP.
type HTTP struct {
	Hostname       string      `default:""`
	Port           int         `default:"3000"`
	IsSecure       bool        `yaml:"is_secure" default:"false"`
	AllowedHosts   []string    `yaml:"allowed_hosts"`
	AllowedMethods []string    `yaml:"allowed_methods"`
	AllowedHeaders []string    `yaml:"allowed_headers"`
	SecureHTTP     *SecureHTTP `yaml:"secure_http"`
}

// SecureHTTP holds configuration for HTTPS.
type SecureHTTP struct {
	Port         int    `default:"3443"`
	KeyFilepath  string `yaml:"key_filepath"`
	CertFilepath string `yaml:"cert_filepath"`
}

// Address returns the HTTP address.
func (h *HTTP) Address(withScheme bool) string {
	port := h.Port
	scheme := "http://"
	if h.IsSecure && h.SecureHTTP != nil {
		port = h.SecureHTTP.Port
		scheme = "https://"
	}
	if !withScheme {
		scheme = ""
	}
	return fmt.Sprintf("%s%s:%d", scheme, h.Hostname, port)
}
