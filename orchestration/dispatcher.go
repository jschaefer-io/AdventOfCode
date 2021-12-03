package orchestration

import (
    "fmt"
    "sort"
    "strings"
    "sync"
    "time"
)

type WorkLoad map[string]string

type Dispatcher struct {
    solvers map[string]Solver
}

func (d *Dispatcher) AddSolver(id string, solver Solver) {
    d.solvers[id] = solver
}

func (d *Dispatcher) Process(results chan<- Result, workload WorkLoad) {
    var all sync.WaitGroup

    for key, value := range workload {
        all.Add(1)
        solver, ok := d.solvers[key]
        if !ok {
            panic(fmt.Sprintf("%s has no solver", key))
        }
        go solver.handle(key, value, results)
    }
}

func (d *Dispatcher) Start(workload WorkLoad) {
    limit := len(workload)
    if limit == 0 {
        fmt.Println("No Workload provided")
        return
    }
    resList := make([]Result, 0)

    results := make(chan Result)
    d.Process(results, workload)

    for res := range results {
        resList = append(resList, res)
        currentCount := len(resList)
        var sb strings.Builder
        sb.WriteString(fmt.Sprintf("%d of %d -- ", currentCount, limit))
        sb.WriteString(fmt.Sprintf("Finished: %s", res.Id))
        if res.Err != nil {
            sb.WriteString(" (with error)")
        }
        fmt.Println(sb.String())
        if currentCount == limit {
            close(results)
        }
    }

    sort.Slice(resList, func(i, j int) bool {
        return strings.Compare(resList[i].Id, resList[j].Id) == -1
    })

    var sumDurations time.Duration = 0
    for _, res := range resList {
        sumDurations += res.Duration
    }

    fmt.Println("\nSolutions:")
    for _, res := range resList {
        var sb strings.Builder
        sb.WriteString(fmt.Sprintf("Day %s (%s):", res.Id, res.Duration))
        if res.Err != nil {
            sb.WriteString(fmt.Sprintf("\nError: %s", res.Err))
        } else {
            for i, answer := range res.Results {
                sb.WriteString(fmt.Sprintf("\n%d) %s", i+1, answer))
            }
        }
        sb.WriteString("\n")
        fmt.Println(sb.String())
    }
}

var MainDispatcher Dispatcher

func init() {
    MainDispatcher = Dispatcher{
        solvers: make(map[string]Solver),
    }
}
