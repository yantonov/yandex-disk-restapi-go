package src

type ResourceInfoRequestOptions struct {
	Sort_mode    *SortMode
	Limit        *uint32
	Offset       *uint32
	Fields       []string
	Preview_size *PreviewSize
	Preview_crop *bool
}
