package memory

type memory map[string]string

func (m *memory) Init() {
	m = new(memory)
}

func (m memory) AddUser()                                          {}
func (m memory) Authenticate(username string, passwordHash string) {}
