package main

type Collector struct {
	distinct map[Match]struct{}
}

func NewCollector() *Collector {
	return &Collector{
		distinct: map[Match]struct{}{},
	}
}

func (c *Collector) Push(matches ...Match) {
	for _, s := range matches {
		if _, ok := c.distinct[s]; !ok {
			c.distinct[s] = struct{}{}
		}
	}
}

func (c *Collector) Result() []Match {
	res := make([]Match, 0, len(c.distinct))
	for k := range c.distinct {
		res = append(res, k)
	}
	return res
}
