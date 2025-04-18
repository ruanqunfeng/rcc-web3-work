package gobasework

type pair struct{ start, end int }
type MyCalendar []pair

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (c *MyCalendar) Book(start, end int) bool {
	for _, p := range *c {
		if p.start < end && start < p.end {
			return false
		}
	}
	*c = append(*c, pair{start, end})
	return true
}
