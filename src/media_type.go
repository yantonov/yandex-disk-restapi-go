package src

type MediaType struct {
	media_type string
}

func (m *MediaType) Audio() *MediaType {
	return &MediaType{
		media_type: "audio",
	}
}

func (m *MediaType) Backup() *MediaType {
	return &MediaType{
		media_type: "backup",
	}
}

func (m *MediaType) Book() *MediaType {
	return &MediaType{
		media_type: "book",
	}
}

func (m *MediaType) Compressed() *MediaType {
	return &MediaType{
		media_type: "compressed",
	}
}

func (m *MediaType) Data() *MediaType {
	return &MediaType{
		media_type: "data",
	}
}

func (m *MediaType) Development() *MediaType {
	return &MediaType{
		media_type: "development",
	}
}

func (m *MediaType) Diskimage() *MediaType {
	return &MediaType{
		media_type: "diskimage",
	}
}

func (m *MediaType) Document() *MediaType {
	return &MediaType{
		media_type: "document",
	}
}

func (m *MediaType) Encoded() *MediaType {
	return &MediaType{
		media_type: "encoded",
	}
}

func (m *MediaType) Executable() *MediaType {
	return &MediaType{
		media_type: "executable",
	}
}

func (m *MediaType) Flash() *MediaType {
	return &MediaType{
		media_type: "flash",
	}
}

func (m *MediaType) Font() *MediaType {
	return &MediaType{
		media_type: "font",
	}
}

func (m *MediaType) Image() *MediaType {
	return &MediaType{
		media_type: "image",
	}
}

func (m *MediaType) Settings() *MediaType {
	return &MediaType{
		media_type: "settings",
	}
}

func (m *MediaType) Spreadsheet() *MediaType {
	return &MediaType{
		media_type: "spreadsheet",
	}
}

func (m *MediaType) Text() *MediaType {
	return &MediaType{
		media_type: "text",
	}
}

func (m *MediaType) Unknown() *MediaType {
	return &MediaType{
		media_type: "unknown",
	}
}

func (m *MediaType) Video() *MediaType {
	return &MediaType{
		media_type: "video",
	}
}

func (m *MediaType) Web() *MediaType {
	return &MediaType{
		media_type: "web",
	}
}

func (m *MediaType) String() string {
	return m.media_type
}
