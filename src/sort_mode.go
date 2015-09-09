package src

import "strings"

type SortMode struct {
	mode string
}

func (m *SortMode) Default() *SortMode {
	return &SortMode{
		mode: "",
	}
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

func (m *SortMode) Reverse() *SortMode {
	if strings.HasPrefix(m.mode, "-") {
		return &SortMode{
			mode: m.mode[1:],
		}
	}
	return &SortMode{
		mode: "-" + m.mode,
	}
}

func (m *SortMode) String() string {
	return m.mode
}

func (m *SortMode) UnmarshalJSON(value []byte) error {
	if value == nil || len(value) == 0 {
		m.mode = ""
		return nil
	}
	m.mode = string(value)
	if strings.HasPrefix(m.mode, "\"") && strings.HasSuffix(m.mode, "\"") {
		m.mode = m.mode[1 : len(m.mode)-1]
	}
	return nil
}
