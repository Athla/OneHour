# 01. The Stack Overflow.
- Maybe, just maybe, I forgor to add the end condition to my recursion while trying to implement BFS.

```go
package algorithms

import (
    "github.com/Athla/OneHour/data_structures"
    "log"
)

func Travel(n *data_structures.Node) map[int]*data_structures.Node {
    result := make(map[int]*data_structures.Node)
    if v, ok := result[n.Value]; ok == false {
        result[n.Value] = v
    }
    if n.Neighbours != nil {
        for _, b := range n.Neighbours {
            Travel(b)
        }
    }

    return result
}

func BFS(g *data_structures.Graph) {
    var res map[int]*data_structures.Node
    for _, v := range g.Nodes {
        res = Travel(v)
    }
    log.Println(res)
}
```

- Still, it was fun. And made me notice that coding still fun for me, that's a feeling that I was noticing when doing something for a TUI App that I work in the weekends, but that wasn't happening in my day to day work.
- Learning is fun, breaking the inertia and finding fun is even better.
- Also I gotta fix the recursion call in BFS, due the fact that, for cyclic graphs, it will result in, well, a **Stack Overflow**.
- And one more TUI Project tbd, that's a personal one bc there surely exist one of that, but a terminal list of terms that you've said you would check, but didn't

