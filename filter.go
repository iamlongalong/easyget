package easyget

var (
	emptyBytes  = []byte("")
	emptyString = ""
)

type IFilter interface {
	Filt(d []byte) []byte
}

type GroupFilter struct {
	filters []IFilter
}

func NewGroupFilter(filters ...IFilter) IFilter {
	return &GroupFilter{
		filters: filters,
	}
}

func (gf *GroupFilter) Filt(d []byte) []byte {

	for _, f := range gf.filters {
		d = f.Filt(d)
	}

	return d
}
