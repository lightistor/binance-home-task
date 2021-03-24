package main

type symbols []*SymbolData

func (s symbols) Len() int      { return len(s) }
func (s symbols) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// ByVolume implements sort.Interface
type ByVolume struct{ symbols }

func (s ByVolume) Less(i, j int) bool { return s.symbols[i].Volume.GreaterThan(s.symbols[j].Volume) }

// ByTradeCount implements sort.Interface
type ByTradeCount struct{ symbols }

func (s ByTradeCount) Less(i, j int) bool {
	return s.symbols[i].TradeCount > s.symbols[j].TradeCount
}
