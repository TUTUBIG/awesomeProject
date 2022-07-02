package designPattern

type observer struct {
}

type StockMarket struct {
	observers map[string]*observer
}
