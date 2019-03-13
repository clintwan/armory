package armory

// Slice Slice
type Slice []interface{}

// IndexOf IndexOf
func (arr Slice) IndexOf(ele interface{}) int {
	r := -1
	for idx, val := range arr {
		if ele == val {
			r = idx
		}
	}
	return r
}
