package proc

import (
	"github.com/litsoftware/litmedia/pkg/reflecth"
	"github.com/trustmaster/goflow"
)

func Build(procList ...goflow.Component) *goflow.Graph {
	n := goflow.NewGraph()

	l := len(procList)
	if l > 0 {
		for i, c := range procList {
			name := reflecth.GetName(c)
			_ = n.Add(name, c)

			if i > 0 && i < l {
				_ = n.Connect(reflecth.GetName(procList[i-1]), "Out", name, "In")
			}

			if i == 0 {
				n.MapInPort("In", name, "In")
			}
		}
	}

	return n
}
