package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/logrusorgru/aurora"
)

type Pair struct {
	x  int
	y  int
	id int
}

type Cell struct {
	val       int
	is_center bool
}

type Area struct {
	area   int
	is_inf bool
}

func (cell *Cell) is_some() bool {
	return cell.val >= 0
}

func (cell *Cell) unwrap() int {
	return cell.val
}

func (cell *Cell) invalidate() {
	cell.val = -1
}

func (cell Cell) String() string {
	if cell.is_some() && cell.is_center {
		return fmt.Sprintf("%d", aurora.Bold(aurora.Green(cell.val)))
	} else {
		if cell.is_some() {
			return fmt.Sprintf("%d", cell.val)
		}
		return "."
	}
}

func parseInput() []Pair {

	var coords []Pair
	scanner := bufio.NewScanner(os.Stdin)

	id := 0

	for scanner.Scan() {
		tmp_coords := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(strings.TrimSpace(tmp_coords[0]))
		y, _ := strconv.Atoi(strings.TrimSpace(tmp_coords[1]))
		coords = append(coords, Pair{x: x, y: y, id: id})
		id++
	}

	return coords
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func dist(a *Pair, b *Pair) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func find_closest(a *Pair, coords []Pair) Cell {

	// there may be more than one closest point
	closest := []Pair{coords[0]}

	for idx := range coords {

		p := coords[idx]

		last_point := closest[len(closest)-1]

		if last_point != p {

			dist_to_last := dist(a, &last_point)
			dist_to_a := dist(a, &p)

			if dist_to_last == dist_to_a {
				// we have another point the same distance away from "a"
				closest = append(closest, p)

			} else if dist_to_last > dist_to_a {
				// we've found a new closest point, forget older ones
				closest = []Pair{p}
			}
		}
	}

	if len(closest) > 1 {
		return Cell{val: -1, is_center: false}
	}

	return Cell{val: closest[0].id, is_center: false}

}

func prepareField(coords []Pair) [][]Cell {

	max_x, max_y := 0, 0

	// find max x and max y coordinates
	for _, p := range coords {

		if p.x > max_x {
			max_x = p.x
		}
		if p.y > max_y {
			max_y = p.y
		}

	}

	// prepare field for Voronoi areas
	var field = make([][]Cell, max_y+1)
	for i := range field {
		field[i] = make([]Cell, max_x+1)
	}

	return field

}

func traverse(y int, x int, id int, field [][]Cell) int {
	// function returns -1 if Voronoi region area is infinite
	// and a positive integer which represents an area otherwise

	// if something goes beyond the field this implies it will
	// be continue infinitely

	if x-1 < 0 {
		return -1
	}
	if y-1 < 0 {
		return -1
	}
	if y+1 > len(field) {
		return -1
	}
	if x+1 > len(field[y]) {
		return -1
	}

	cur_value := field[y][x].unwrap()

	if cur_value != id {
		return 0
	}

	field[y][x].invalidate()

	left := traverse(y, x-1, id, field)
	right := traverse(y, x+1, id, field)
	top := traverse(y-1, x, id, field)
	bottom := traverse(y+1, x, id, field)

	if bottom == -1 || top == -1 || left == -1 || right == -1 {
		return -1
	}

	return 1 + left + right + top + bottom

}

func scanAreas(field [][]Cell) int {
	max_area := 0

	// traverse field and designate Voronoi area
	for y := range field {
		for x := range field[y] {
			if field[y][x].is_some() {
				area := traverse(y, x, field[y][x].unwrap(), field)
				if area > max_area {
					max_area = area
				}
			}
		}
	}
	return max_area
}

func findRegion(pts []Pair, max_dist int) int {
	total_area := 0

	for y := 0; y < max_dist; y++ {
		for x := 0; x < max_dist; x++ {

			net_dist := 0
			cur_point := Pair{x: x, y: y}

			for _, p := range pts {
				net_dist += dist(&p, &cur_point)
				if net_dist >= max_dist {
					break
				}
			}

			if net_dist < max_dist {
				total_area += 1
			}
		}
	}

	return total_area
}

func visualize(field [][]Cell) {
	for y := range field {
		for x := range field[y] {
			fmt.Print(field[y][x])
		}
		fmt.Println()
	}
}

func main() {

	coords := parseInput()
	field := prepareField(coords)

	// fill in Voronoi areas
	for y := range field {
		for x := range field[y] {
			field[y][x] = find_closest(&Pair{x: x, y: y}, coords)
		}
	}

	// set up points
	for idx := range coords {
		p := coords[idx]
		field[p.y][p.x] = Cell{val: p.id, is_center: true}
	}

	//visualize(field)
	fmt.Printf("Max area is: %d\n", scanAreas(field))
	fmt.Printf("Region size: %d", findRegion(coords, 10000))
}
