package draw

import (
	"fmt"

	"github.com/akarki15/dbdot/models"
	"github.com/emicklei/dot"
)

func ERD(tables []models.TableAndColumns, fks []models.ForeignKey) (dot.Graph, error) {
	g := dot.NewGraph(dot.Directed)
	tableToNode := map[string]dot.Node{}
	for _, table := range tables {
		var cols []string
		for _, col := range table.Columns {
			cols = append(cols, fmt.Sprintf("<TR><TD>%s</TD><TD>%s</TD></TR>", col.Name, col.DataType))
		}
		n := g.Node(table.Name).Box()
		n.Attr("shape", "plaintext")
		n.Attr("label", fmt.Sprintf(`<<TABLE BORDER="0" CELLBORDER="1" CELLSPACING="0"><TR><TD colspan="2">%s</TD></TR>%s</TABLE>>`,
			table.Name, cols))
		tableToNode[table.Name] = n
	}

	for _, fk := range fks {
		from, ok := tableToNode[fk.From]
		if !ok {
			// possible because of whitelist
			continue
		}
		to, ok := tableToNode[fk.To]
		if !ok {
			// possible because of whitelist
			continue
		}
		g.Edge(from, to)
	}
	return *g, nil
}
