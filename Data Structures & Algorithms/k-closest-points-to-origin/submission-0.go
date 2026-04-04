func kClosest(points [][]int, k int) [][]int {
	ds := []Distance{}

	for _, point := range points {
		ds = append(ds, Distance{
			point: point,
		})
	}

	// sort.Slice(people, func(i, j int) bool {
	// 	return people[i].Age < people[j].Age // Sort by Age, ascending
	// })
}

type Distance struct {
	point []int
	cal int 
}