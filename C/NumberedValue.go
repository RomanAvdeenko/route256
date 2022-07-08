package matrix

type NV struct{ N, V int } // Numbered Value
type ByValue []NV

func (v ByValue) Len() int           { return len(v) }
func (v ByValue) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }
func (v ByValue) Less(i, j int) bool { return v[i].V < v[j].V }
