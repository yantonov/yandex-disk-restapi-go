package src

import "fmt"

type PreviewSize struct {
	size string
}

func (s *PreviewSize) PredefinedSizeS() *PreviewSize {
	return &PreviewSize{
		size: "S",
	}
}

func (s *PreviewSize) PredefinedSizeM() *PreviewSize {
	return &PreviewSize{
		size: "M",
	}
}

func (s *PreviewSize) PredefinedSizeL() *PreviewSize {
	return &PreviewSize{
		size: "L",
	}
}

func (s *PreviewSize) PredefinedSizeXL() *PreviewSize {
	return &PreviewSize{
		size: "XL",
	}
}

func (s *PreviewSize) PredefinedSizeXXL() *PreviewSize {
	return &PreviewSize{
		size: "XXL",
	}
}

func (s *PreviewSize) PredefinedSizeXXXL() *PreviewSize {
	return &PreviewSize{
		size: "XXXL",
	}
}

func (s *PreviewSize) ExactWidth(width uint32) *PreviewSize {
	return &PreviewSize{
		size: fmt.Sprintf("%dx", width),
	}
}

func (s *PreviewSize) ExactHeight(height uint32) *PreviewSize {
	return &PreviewSize{
		size: fmt.Sprintf("x%d", height),
	}
}

func (s *PreviewSize) ExactSize(width uint32, height uint32) *PreviewSize {
	return &PreviewSize{
		size: fmt.Sprintf("%dx%d", width, height),
	}
}

func (s *PreviewSize) String() string {
	return s.size
}
