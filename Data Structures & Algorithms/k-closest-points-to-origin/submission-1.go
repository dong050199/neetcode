func kClosest(points [][]int, k int) [][]int {
	ds := []Point{}

	for _, point := range points {
		ds = append(ds, Point{
			point: point,
			val: point[0]*point[0] + point[1]*point[1],
		})
	}

	sort.Slice(ds, func(i, j int) bool {
		return ds[i].val < ds[j].val
	})

	var res [][]int

	for _, d := range ds[:k] {
		res = append(res, d.point)
	}

	return res
}

type Point struct {
	point []int
	val int 
}