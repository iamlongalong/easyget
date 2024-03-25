package easyget

type StaticGetter struct {
	val string
}

func NewStaticGetter(val string) *StaticGetter {
	return &StaticGetter{
		val: val,
	}
}

func (sg *StaticGetter) Get(key string) (string, bool) {
	return sg.val, true
}
