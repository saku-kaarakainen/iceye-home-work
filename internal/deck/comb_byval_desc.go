package deck

type ByValueDesc []Combination

func (s ByValueDesc) Len() int {
    return len(s)
}

func (s ByValueDesc) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s ByValueDesc) Less(i, j int) bool {
    return s[i].Value > s[j].Value
}