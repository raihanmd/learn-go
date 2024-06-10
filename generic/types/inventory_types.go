package types

type Inventory[T any] []T

func (i *Inventory[T]) Add(item T) {
	*i = append(*i, item)
}

func (i *Inventory[T]) Remove(item T) {
	*i = append(*i, item)
}

func (i *Inventory[T]) Print() {
	for _, item := range *i {
		println(item)
	}
}
