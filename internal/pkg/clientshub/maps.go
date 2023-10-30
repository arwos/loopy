package clientshub

type Maps struct {
	data map[string]struct{}
}

func NewMaps() *Maps {
	return &Maps{
		data: make(map[string]struct{}, 10),
	}
}

func (v *Maps) Get() []string {
	result := make([]string, 0, len(v.data))
	for value := range v.data {
		result = append(result, value)
	}
	return result
}

func (v *Maps) Set(value ...string) {
	for _, vv := range value {
		v.data[vv] = struct{}{}
	}
}

func (v *Maps) Del(value ...string) {
	for _, vv := range value {
		delete(v.data, vv)
	}
}

func (v *Maps) Has(value string) bool {
	_, ok := v.data[value]
	return ok
}

func (v *Maps) IsEmpty() bool {
	return len(v.data) == 0
}
