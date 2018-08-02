package orderbook

type Orderbook struct {
	book_bid []*Order
	book_ask []*Order
	n        int
	m        int
}

func New() *Orderbook {
	return &Orderbook{}
}
func New_trade() *Trade {
	return &Trade{}
}

func (or *Orderbook) Add_bid(order *Order) {
	if or.n < 1 {
		or.book_bid = append(or.book_bid, order)
		or.n++
	} else if order.Price > or.book_bid[or.n-1].Price {
		or.book_bid = append(or.book_bid, order)
		or.n++
	} else {
		for i := or.n - 1; i >= 0; i-- {
			if i == 0 {
				or.book_bid = append([]*Order{order}, or.book_bid...)
				or.n++
			} else if order.Price > or.book_bid[i].Price {
				newB := or.book_bid[:i+1]
				newB = append(newB, order)
				or.book_bid = append(newB, or.book_bid[i+1:]...)
				or.n++
				break
			}
		}
	}
}

func (or *Orderbook) Add_ask(order *Order) {
	if or.m < 1 {
		or.book_ask = append(or.book_ask, order)
		or.m++
	} else if order.Price < or.book_ask[or.m-1].Price {
		or.book_ask = append(or.book_ask, order)
		or.m++
	} else {
		for i := or.m - 1; i >= 0; i-- {
			if i == 0 {
				or.book_ask = append([]*Order{order}, or.book_ask...)
				or.m++
			} else if order.Price < or.book_ask[i].Price {
				newB := or.book_ask[:i+1]
				newB = append(newB, order)
				or.book_ask = append(newB, or.book_ask[i+1:]...)
				or.m++
				break
			}
		}
	}
}

func (or *Orderbook) Del_ask() {
	or.book_ask = or.book_ask[:or.m-1]
	or.m--
}
func (or *Orderbook) Del_bid() {
	or.book_bid = or.book_bid[:or.n-1]
	or.n--
}

func (or *Orderbook) Search_ask(order *Order) ([]*Trade, *Order) {
	var trade []*Trade
	for i := len(or.book_ask) - 1; i >= 0; i-- {
		if or.book_ask[i].Price <= order.Price {
			tr := New_trade()
			tr.Bid = order
			tr.Ask = or.book_ask[i]
			tr.Price = or.book_ask[i].Price
			if or.book_ask[i].Volume < order.Volume {
				tr.Volume = or.book_ask[i].Volume
				trade = append(trade, tr)
				order.Volume -= tr.Volume
				or.Del_ask()
			} else if or.book_ask[i].Volume >= order.Volume {
				tr.Volume = order.Volume
				trade = append(trade, tr)
				if or.book_ask[i].Volume == order.Volume {
					or.Del_ask()
				} else {
					or.book_ask[i].Volume -= order.Volume
				}
				order.Volume = 0
				break
			}
		}
	}
	if order.Volume > 0 {
		or.Add_bid(order)
		return trade, order
	}
	return trade, nil
}

func (or *Orderbook) Search_bid(order *Order) ([]*Trade, *Order) {
	var trade []*Trade
	for i := len(or.book_bid) - 1; i >= 0; i-- {
		if or.book_bid[i].Price >= order.Price {
			tr := New_trade()
			tr.Ask = order
			tr.Bid = or.book_bid[i]
			tr.Price = or.book_bid[i].Price
			if or.book_bid[i].Volume < order.Volume {
				tr.Volume = or.book_bid[i].Volume
				trade = append(trade, tr)
				order.Volume -= tr.Volume
				or.Del_bid()
			} else if or.book_bid[i].Volume >= order.Volume {
				tr.Volume = order.Volume
				trade = append(trade, tr)
				if or.book_bid[i].Volume == order.Volume {
					or.Del_bid()
				} else {
					or.book_bid[i].Volume -= order.Volume
				}
				order.Volume = 0
				break
			}
		}
	}
	if order.Volume > 0 {
		or.Add_ask(order)
		return trade, order
	}
	return trade, nil
}

func (or *Orderbook) Market_ask(order *Order) ([]*Trade, *Order) {
	var trade []*Trade
	for i := len(or.book_ask) - 1; i >= 0; i-- {
		tr := New_trade()
		tr.Bid = order
		tr.Ask = or.book_ask[i]
		tr.Price = or.book_ask[i].Price
		if or.book_ask[i].Volume < order.Volume {
			tr.Volume = or.book_ask[i].Volume
			trade = append(trade, tr)
			order.Volume -= tr.Volume
			or.Del_ask()
		} else if or.book_ask[i].Volume >= order.Volume {
			tr.Volume = order.Volume
			trade = append(trade, tr)
			if or.book_ask[i].Volume == order.Volume {
				or.Del_ask()
			} else {
				or.book_ask[i].Volume -= order.Volume
			}
			order.Volume = 0
			break
		}
	}
	if order.Volume == 0 {
		return trade, nil
	} else {
		return trade, order
	}
}

func (or *Orderbook) Market_bid(order *Order) ([]*Trade, *Order) {
	var trade []*Trade
	for i := len(or.book_bid) - 1; i >= 0; i-- {
		tr := New_trade()
		tr.Bid = order
		tr.Ask = or.book_bid[i]
		tr.Price = or.book_bid[i].Price
		if or.book_bid[i].Volume < order.Volume {
			tr.Volume = or.book_bid[i].Volume
			trade = append(trade, tr)
			order.Volume -= tr.Volume
			or.Del_bid()
		} else if or.book_bid[i].Volume >= order.Volume {
			tr.Volume = order.Volume
			trade = append(trade, tr)
			if or.book_bid[i].Volume == order.Volume {
				or.Del_bid()
			} else {
				or.book_bid[i].Volume -= order.Volume
			}
			order.Volume = 0
			break
		}
	}
	if order.Volume == 0 {
		return trade, nil
	} else {
		return trade, order
	}
}

func (orderbook *Orderbook) Match(order *Order) ([]*Trade, *Order) {
	if order.Kind == 2 {
		if order.Side == 1 {
			trade, _ := orderbook.Search_ask(order)
			if len(trade) > 0 {
				return trade, nil
			} else {
				return nil, nil
			}
		} else if order.Side == 2 {
			trade, _ := orderbook.Search_bid(order)
			if len(trade) > 0 {
				return trade, nil
			} else {
				return nil, nil
			}
		}
	} else if order.Kind == 1 {
		if order.Side == 1 {
			trade, ord := orderbook.Market_ask(order)
			if len(trade) > 0 {
				return trade, ord
			} else {
				return nil, ord
			}
		} else if order.Side == 2 {
			trade, ord := orderbook.Market_bid(order)
			if len(trade) > 0 {
				return trade, ord
			} else {
				return nil, ord
			}
		}
	}
	return nil, nil
}
