package clay

// Sortable lets Clay systems order systems and plugins.
type Sortable interface {
	Order() int
}
