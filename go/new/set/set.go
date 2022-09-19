package set

type Interface interface {
	Add(v interface{}) bool
	Remove(v interface{}) bool
	IsElementOf(v interface{}) bool
	Size() int
}

type Emptier interface {
	Empty()
}

type HashSet struct {
	items map[interface{}]status
}
type status bool

const statusExists status = true

func New() *HashSet {
	return &HashSet{items: make(map[interface{}]status)}
}

func (set *HashSet) Add(item interface{}) bool {
	if _, exists := set.items[item]; exists {
		return false
	}
	set.items[item] = statusExists
	return true
}

func (set *HashSet) Remove(item interface{}) bool {
	if _, exists := set.items[item]; !exists {
		return false
	}
	delete(set.items, item)
	return true
}
func (set *HashSet) Size() int {
	return len(set.items)
}

func (set *HashSet) IsElementOf(item interface{}) bool {
	if _, exists := set.items[item]; exists {
		return false
	}
	return true
}
func (set *HashSet) Values() []interface{} {
	values := make([]interface{}, len(set.items))
	for item := range set.items{
		values = append(values, item)
	}
	return values
}
