package src

type SortMode struct {
	mode string
}

func (m *SortMode) ByName() *SortMode {
	return &SortMode{
		mode: "name",
	}
}

func (m *SortMode) ByPath() *SortMode {
	return &SortMode{
		mode: "path",
	}
}

func (m *SortMode) ByCreated() *SortMode {
	return &SortMode{
		mode: "created",
	}
}

func (m *SortMode) ByModified() *SortMode {
	return &SortMode{
		mode: "modified",
	}
}

func (m *SortMode) BySize() *SortMode {
	return &SortMode{
		mode: "size",
	}
}

func (m *SortMode) String() string {
	if m.mode == "" {
		panic("undefined mode")
	}
	return m.mode
}
