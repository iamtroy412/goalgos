package main

import "fmt"

type Point struct {
    x int
    y int
}

var dirs = [4][2]int{
    {-1, 0},
    {1, 0},
    {0, -1},
    {0, 1},
}

func walk(maze [][]string, wall string, curr, end Point, seen [][]bool, path *[]Point) bool {
    // off map
    if curr.x < 0 || curr.x >= len(maze[0]) || curr.y < 0 || curr.y >= len(maze) {
        return false
    }

    // on wall
    if maze[curr.y][curr.x] == wall {
        return false
    }

    // at the end
    if curr.x == end.x && curr.y == end.y {
        *path = append(*path, end)
        return true
    }

    // been here before
    if seen[curr.y][curr.x] {
        return false
    }

    // pre
    seen[curr.y][curr.x] = true
    *path = append(*path, curr)
    
    // recurse
    for i := 0; i < len(dirs); i++ {
        x := dirs[i][0]
        y := dirs[i][1]
        if walk(maze, wall, Point{x: curr.x + x, y: curr.y + y}, end, seen, path) {
            return true
        }
    }

    // post
    *path = (*path)[:len(*path) - 1]
    return false
}

func solve(maze [][]string, wall string, start, end Point) []Point {
    seen := make([][]bool, len(maze))
    for i := range seen {
        seen[i] = make([]bool, len(maze[0]))
    }
    var path []Point

    walk(maze, wall, start, end, seen, &path)
    return path
}

func main() {
    fmt.Println("running")
    maze := [][]string{
        {"#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "S", "#"},
        {"#", " ", " ", " ", " ", " ", " ", " ", " ", "#", " ", "#"},
        {"#", " ", " ", " ", " ", " ", " ", " ", " ", "#", " ", "#"},
        {"#", " ", "#", "#", "#", "#", "#", "#", "#", "#", " ", "#"},
        {"#", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "#"},
        {"#", "E", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#"},
    }
    wall := "#"
    start := Point{
        x: 10,
        y: 0,
    }
    end := Point{
        x: 1,
        y: 5,
    }
    fmt.Println(solve(maze, wall, start, end))
}
