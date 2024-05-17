package render

type IGraph interface {
	Prepare()
	Queue()
	Render()
}

type Graph struct {
}

func (g *Graph) Prepare() {
}

func (g *Graph) Queue() {
}

func (g *Graph) Render() {
}
