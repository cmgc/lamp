package main

import (
    "fmt"
    "time"
    "log"
)

type Queue []*Point

func (q *Queue) Push(n *Point) {
    *q = append(*q, n)
}

func (q *Queue) Pop() (n *Point) {
    n = (*q)[0]
    *q = (*q)[1:]
    return
}

func (q *Queue) Len() int {
    return len(*q)
}

type Point struct {
    x, y int
}

func isLegal(ar [2]int) bool {
    limit := Point{1000, 1000}
    x := ar[0]
    y := ar[1]
    if x >= 0 && x < limit.x && y >= 0 && y < limit.y {
        return true
    }
    return false
}

func knightSteps(p Point) [][2]int {
    var nextMoves [][2]int
    possible := [][2]int{
        {p.x-2,p.y-1},
        {p.x-2,p.y+1},
        {p.x-1,p.y-2},
        {p.x-1,p.y+2},
        {p.x+1,p.y-2},
        {p.x+1,p.y+2},
        {p.x+2,p.y-1},
        {p.x+2,p.y+1},
    }
    for _, p := range possible {
        if isLegal(p) {
            nextMoves = append(nextMoves, p)
        }
    }
    return nextMoves
}


func bfs(s *Point, f *Point) map[Point]Point {
    parents := make(map[Point]Point)
    queue := new(Queue)
    queue.Push(s)
    found := false
    // fmt.Println("bfs")
    for ; queue.Len() != 0 && !found;  {
        step := queue.Pop()
        if step.x == f.x && step.y == f.y {
            found = true
        }
        for _, nst :=  range knightSteps(*step) {
            temp := Point{nst[0], nst[1]}
            if _, ok := parents[temp]; !ok {
                parents[temp] = *step
                queue.Push(&temp)
            }
        }
    }
    return parents
}

func timeTrack(start time.Time, name string) {
    elapsed := time.Since(start)
    log.Printf("%s took %s", name, elapsed)
}


func main() {
    start := Point{0, 0}
    finish := Point{999, 999}
    // begin := time.Now()
    defer timeTrack(time.Now(), "bfs")
    parents := bfs(&start, &finish)

    // end := time.Since(begin)
    fmt.Println(len(parents))
    // fmt.Println(parents[starst])
    // fmt.Println("time: ", end)
}


