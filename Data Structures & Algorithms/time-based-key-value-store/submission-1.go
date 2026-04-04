type TimeMap struct {
	Map map[string][]MapValue
}

type MapValue struct {
	Value string
	TimeStamp int
}

func Constructor() TimeMap {
	m := make(map[string][]MapValue)
	return TimeMap{
		Map: m,
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.Map[key] = append(this.Map[key],MapValue{
		Value: value,
		TimeStamp: timestamp,
	})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	s := this.Map[key]
	l , r := 0, len(s) - 1
	for l < r {
		m := l + (r - l) / 2
		if s[m].TimeStamp == timestamp {
			return s[m].Value
		}
		if s[m].TimeStamp > timestamp {
			r = m
			continue
		}
		l = m + 1
	}
	return s[l].Value
}
// [1,2,3,4,5,6]
