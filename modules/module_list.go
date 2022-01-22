package modules

type ModuleList struct {
	modules []Module
}

func NewModuleList() *ModuleList {
	return &ModuleList{}
}

func (m *ModuleList) AddAllAvailable() *ModuleList {
	m.AddModule(NewSimpleWiresModule())
	m.AddModule(NewBigButtonModule())
	m.AddModule(NewPasswordModule())

	return m
}

func (m *ModuleList) AddModule(module Module) *ModuleList {
	m.modules = append(m.modules, module)
	return m
}

func (m *ModuleList) GetNames() []string {
	result := make([]string, len(m.modules))
	for i, module := range m.modules {
		result[i] = module.Name()
	}
	return result
}

func (m *ModuleList) GetByName(moduleName string) Module {
	for _, module := range m.modules {
		if module.Name() != moduleName {
			continue
		}
		return module
	}
	return nil
}
