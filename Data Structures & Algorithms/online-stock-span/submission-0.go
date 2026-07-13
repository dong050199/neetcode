type Node struct {
	price int
	span  int
}

type StockSpanner struct {
	spanner []Node
}

func Constructor() StockSpanner {
	return StockSpanner{
		spanner: []Node{},
	}
}

func (this *StockSpanner) Next(price int) int {
	span := 1
	for len(this.spanner) > 0 && price >= this.spanner[len(this.spanner)-1].price {
		span += this.spanner[len(this.spanner)-1].span
		this.spanner = this.spanner[:len(this.spanner)-1]
	}

	this.spanner = append(this.spanner, Node{
		price: price,
		span:  span,
	})

	return span
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */