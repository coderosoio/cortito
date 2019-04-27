package config

import "strings"

type Service struct {
	Name      string
	Namespace string `default:"io.coderoso.cortito"`
}

func (s *Service) URL() (url string) {
	parts := []string{
		s.Namespace, s.Name,
	}
	url = strings.Join(parts, ".")
	return
}
