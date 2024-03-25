package easyget

type Getter interface {
	Get(key string) (string, bool)
}

type BackGetter interface {
	Getter

	WithBackup(g Getter) BackGetter
}

func WithBackup(g Getter) BackGetter {
	return NewSimpleBackGetter(g)
}

func WithDefault(g Getter, v string) BackGetter {
	return NewSimpleBackGetter(g).WithBackup(NewStaticGetter(v))
}

func NewSimpleBackGetter(gs ...Getter) BackGetter {
	return &SimpleBackGetter{
		getters: gs,
	}
}

type SimpleBackGetter struct {
	getters []Getter
}

func (sbg *SimpleBackGetter) WithBackup(g Getter) BackGetter {
	sbg.getters = append(sbg.getters, g)
	return sbg
}

func (sbg *SimpleBackGetter) Get(key string) (string, bool) {
	for _, g := range sbg.getters {
		if g == nil {
			continue
		}

		v, ok := g.Get(key)
		if ok {
			return v, true
		}
	}

	return "", false
}
