package main

type Manager struct {
	showcase map[string]Product
}

func NewManager() *Manager {
	return &Manager{make(map[string]Product)}
}

func (self *Manager) register(name string, proto Product) {
	self.showcase[name] = proto
}

func (self *Manager) create(name string) Product {
	p := self.showcase[name]
	return p.createClone()
}
