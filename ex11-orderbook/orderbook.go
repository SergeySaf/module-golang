package orderbook

type Orderbook struct {
	LastID int
	Bids   []*Order
	Asks   []*Order
	trades []*Trade
}

func New() *Orderbook {
	orderbook := &Orderbook{}
	orderbook.Bids = []*Order{}
	orderbook.Asks = []*Order{}
	orderbook.trades = []*Trade{}

	return orderbook
}

func (orderbook *Orderbook) Match(order *Order) ([]*Trade, *Order) {
	switch order.Side {
	case SideAsk:
		return orderbook.AskOrder(order)
	case SideBid:
		return orderbook.BidOrder(order)
	}
	return nil, nil
}

func (orderbook *Orderbook) AskOrder(order *Order) ([]*Trade, *Order) {
	orderbook.trades = nil
	for i := 0; i < len(orderbook.Bids); i++ {
		currBid := orderbook.Bids[i]
		if order.Price == 0 || order.Price <= currBid.Price {
			trade := &Trade{Bid: currBid, Ask: order, Price: currBid.Price}

			if currBid.Volume > order.Volume {
				trade.Volume = order.Volume
				currBid.Volume -= order.Volume
				order.Volume = 0
			} else {
				trade.Volume = currBid.Volume
				order.Volume -= currBid.Volume
				orderbook.Bids = append(orderbook.Bids[:i], orderbook.Bids[i+1:]...)
				i--
			}
			orderbook.trades = append(orderbook.trades, trade)
			if order.Volume == 0 {
				break
			}
		} else {
			break
		}
	}
	if order.Volume > 1 {
		if order.Price == 0 {
			return orderbook.trades, order
		}
		orderbook.Asks = append(orderbook.Asks, order)
		for i := 0; i < len(orderbook.Asks)-1; i++ {
			if orderbook.Asks[i].Price > orderbook.Asks[i+1].Price {
				orderbook.Asks[i], orderbook.Asks[i+1] = orderbook.Asks[i+1], orderbook.Asks[i]
			} else {
				break
			}
		}
	}
	return orderbook.trades, nil
}

func (orderbook *Orderbook) BidOrder(order *Order) ([]*Trade, *Order) {
	orderbook.trades = nil
	for i := 0; i < len(orderbook.Asks); i++ {
		currAsk := orderbook.Asks[i]
		if order.Price == 0 || order.Price >= currAsk.Price {
			trade := &Trade{Bid: order, Ask: currAsk, Price: currAsk.Price}

			if currAsk.Volume > order.Volume {
				trade.Volume = order.Volume
				currAsk.Volume -= order.Volume
				order.Volume = 0
			} else {
				trade.Volume = currAsk.Volume
				order.Volume -= currAsk.Volume
				orderbook.Asks = append(orderbook.Asks[:i], orderbook.Asks[i+1:]...)
				i--
			}
			orderbook.trades = append(orderbook.trades, trade)
			if order.Volume == 0 {
				break
			}
		} else {
			break
		}
	}
	if order.Volume > 1 {
		if order.Price == 0 {
			return orderbook.trades, order
		}
		orderbook.Bids = append(orderbook.Bids, order)
		for i := len(orderbook.Bids) - 1; i > 0; i-- {
			if orderbook.Bids[i].Price > orderbook.Bids[i-1].Price {
				orderbook.Bids[i], orderbook.Bids[i-1] = orderbook.Bids[i-1], orderbook.Bids[i]
			}
		}
	}
	return orderbook.trades, nil
}
