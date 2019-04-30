package keyvalue

import "strings"

type namespaceMiddleware struct {
	namespace string
	storage   Storage
}

func NamespaceMiddleware(namespace ...string) Middleware {
	return func(storage Storage) Storage {
		return &namespaceMiddleware{
			namespace: strings.Join(namespace, "_"),
			storage:   storage,
		}
	}
}

func (m *namespaceMiddleware) Connect() error {
	return m.storage.Connect()
}

func (m *namespaceMiddleware) Set(key string, value string) error {
	key = m.keyNamespace(key)
	return m.storage.Set(key, value)
}

func (m *namespaceMiddleware) SetIn(key string, field string, value string) error {
	key = m.keyNamespace(key)
	return m.storage.SetIn(key, field, value)
}

func (m *namespaceMiddleware) Get(key string, defaultValue string) (string, error) {
	key = m.keyNamespace(key)
	return m.storage.Get(key, defaultValue)
}

func (m *namespaceMiddleware) GetIn(key string, field string, defaultValue string) (string, error) {
	key = m.keyNamespace(key)
	return m.storage.GetIn(key, field, defaultValue)
}

func (m *namespaceMiddleware) Remove(key string) error {
	key = m.keyNamespace(key)
	return m.storage.Remove(key)
}

func (m *namespaceMiddleware) RemoveIn(key string, field string) error {
	key = m.keyNamespace(key)
	return m.storage.RemoveIn(key, field)
}

func (m *namespaceMiddleware) keyNamespace(key string) string {
	if strings.TrimSpace(m.namespace) == "" {
		return key
	}
	return strings.Join([]string{m.namespace, key}, "_")
}
