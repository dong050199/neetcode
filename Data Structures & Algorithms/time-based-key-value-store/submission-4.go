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
	// get the value from this.Map
	if _, exist := this.Map[key]; !exist {
		return ""
	}

	values := this.Map[key]
	// find the value from values
	l , r := 0, len(values) - 1
	for l < r {
		m := l + (r - l) / 2
		if values[m].TimeStamp == timestamp {
			return values[m].Value
		}
		if values[m].TimeStamp < timestamp {
			l = m + 1
		} else {
			r = m
		}
	}

	return values[r].Value
}
