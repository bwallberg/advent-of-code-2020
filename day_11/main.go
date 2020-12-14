package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type grid [][]string

const floor string = "."
const seat string = "L"
const occupied string = "#"

func createGrid(lines []string) (grid grid) {
	grid = make([][]string, len(lines))
	for i := range grid {
		for _, p := range lines[i] {
			grid[i] = append(grid[i], string(p))
		}
	}
	return
}

func gridToString(grid grid) string {
	var val string
	for _, row := range grid {
		val += strings.Join(row, "")
	}
	return val
}

func countOccupied(grid grid) int {
	return strings.Count(gridToString(grid), "#")
}

func checkDir(x, y, dirX, dirY, limit int, grid grid) int {
	i := 1
	for i != limit + 1 {
		dy := y + (dirY * i)
		dx := x + (dirX * i)
		if dx < 0 || dy < 0 || dy > len(grid)-1 {
			return 0
		}
		row := grid[dy]
		if dx > len(row)-1 {
			return 0
		}
		if row[dx] == occupied {
			return 1
		}
		if row[dx] == seat {
			return 0
		}
		i++
	}
	return 0
}

func simulate(iGrid grid, limit, acceptance int) grid {
	oGrid := make(grid, len(iGrid))
	for i := range iGrid {
		oGrid[i] = make([]string, len(iGrid[i]))
		copy(oGrid[i], iGrid[i])
	}
	for y := range iGrid {
		for x := range iGrid[y] {
			pos := iGrid[y][x]
			if pos != floor {
				count := 0
				// check adjecent on same row
				count += checkDir(x, y, -1, 0, limit, iGrid)
				count += checkDir(x, y, 1, 0, limit, iGrid)
				// check adjecent on row up/down
				count += checkDir(x, y, 0, -1, limit, iGrid)
				count += checkDir(x, y, 0, 1, limit, iGrid)
				// check adjecent diagonally
				count += checkDir(x, y, 1, 1, limit, iGrid)
				count += checkDir(x, y, -1, 1, limit, iGrid)
				count += checkDir(x, y, 1, -1, limit, iGrid)
				count += checkDir(x, y, -1, -1, limit, iGrid)
				if pos == seat && count == 0 {
					oGrid[y][x] = occupied
				} else if pos == occupied && count >= acceptance {
					oGrid[y][x] = seat
				}
			}
		}
	}
	if gridToString(iGrid) == gridToString(oGrid) {
		return oGrid
	}
	return simulate(oGrid, limit, acceptance)
}

func PartOne(lines []string) int {
	return countOccupied(simulate(createGrid(lines), 1, 4))
}

func PartTwo(lines []string) int {
	return countOccupied(simulate(createGrid(lines), -1, 5))
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(input), "\n")

	fmt.Printf("Part one: %d\n", PartOne(lines))
	fmt.Printf("Part two: %d\n", PartTwo(lines))
}
