package day11

import (
	"bufio"
	"fmt"
	"os"
)

type place struct {
	isSeat   bool
	occupied bool
}

func (p *place) String() string {
	if p.occupied {
		return "#"
	}
	if p.isSeat {
		return "L"
	}
	return "."
}

func Run(input string) {
	f, _ := os.Open(input)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	room := parseData(scanner)

	fmt.Printf("Challenge 1: %v occupied places\n", challenge1(room))
	fmt.Printf("Challenge 2: %v occupied places\n", challenge2(room))

}

func parseData(input *bufio.Scanner) (room [][]place) {
	for input.Scan() {
		line := input.Text()
		var roomRow []place
		for _, char := range line {
			place := place{
				occupied: false,
			}
			switch char {
			case 'L':
				place.isSeat = true
			case '.':
				place.isSeat = false
			case '#':
				place.isSeat = true
				place.occupied = true

			}
			roomRow = append(roomRow, place)
		}
		room = append(room, roomRow)
	}
	return room
}

func challenge1(room [][]place) (occupiedCount int) {
	_, _, count := _challenge1(room)
	return count
}

func _challenge1(room [][]place) (newRoom [][]place, changed bool, occupiedCount int) {
	changed = false
	newRoom = make([][]place, len(room))
	for i := range newRoom {
		newRoom[i] = make([]place, len(room[0]))
	}
	for x := 0; x < len(room); x++ {
		for y := 0; y < len(room[0]); y++ {
			if room[x][y].isSeat {
				var iMin, jMin, iMax, jMax int
				iMin = x - 1
				if iMin < 0 {
					iMin = 0
				}
				jMin = y - 1
				if jMin < 0 {
					jMin = 0
				}
				iMax = x + 1
				if iMax == len(room) {
					iMax--
				}
				jMax = y + 1
				if jMax == len(room[0]) {
					jMax--
				}
				var occupiedCounter int
				for i := iMin; i <= iMax; i++ {
					for j := jMin; j <= jMax; j++ {
						if room[i][j].isSeat {
							if room[i][j].occupied {
								occupiedCounter++
							}
						}
					}
				}
				if (occupiedCounter - 1) < 0 {
					place := place{
						isSeat:   true,
						occupied: true,
					}
					newRoom[x][y] = place
					changed = true
				} else if (occupiedCounter-1) >= 4 && room[x][y].occupied {
					place := place{
						isSeat:   true,
						occupied: false,
					}
					newRoom[x][y] = place
					changed = true
				} else {
					newRoom[x][y] = room[x][y]
				}
			}
		}
	}
	if changed {
		return _challenge1(newRoom)
	}
	return room, false, countOccupied(room)
}

func countOccupied(room [][]place) (occupiedCounter int) {
	for x := 0; x < len(room); x++ {
		for y := 0; y < len(room[0]); y++ {
			if room[x][y].isSeat && room[x][y].occupied {
				occupiedCounter++
			}
		}
	}
	return
}

func challenge2(room [][]place) (occupiedCount int) {
	_, _, count := _challenge2(room)
	return count
}

func _challenge2(room [][]place) (newRoom [][]place, changed bool, occupiedCount int) {
	changed = false
	newRoom = make([][]place, len(room))
	for i := range newRoom {
		newRoom[i] = make([]place, len(room[0]))
	}
	for y := 0; y < len(room); y++ {
		for x := 0; x < len(room[0]); x++ {
			if room[y][x].isSeat {
				var occupiedCounter int
				for i := x + 1; i < len(room[0]); i++ {
					if room[y][i].isSeat {
						if room[y][i].occupied {
							occupiedCounter++
						}
						break
					}
				}
				for i := x + 1; i < len(room[0]); i++ {
					if y+(i-x) < len(room) {
						if room[y+(i-x)][i].isSeat {
							if room[y+(i-x)][i].occupied {
								occupiedCounter++
							}
							break
						}
					} else {
						break
					}
				}
				for i := x + 1; i < len(room[0]); i++ {
					if y-(i-x) >= 0 {
						if room[y-(i-x)][i].isSeat {
							if room[y-(i-x)][i].occupied {
								occupiedCounter++
							}
							break
						}
					} else {
						break
					}
				}
				for i := x - 1; i >= 0; i-- {
					if room[y][i].isSeat {
						if room[y][i].occupied {
							occupiedCounter++
						}
						break
					}
				}
				for i := x - 1; i >= 0; i-- {
					if y+(x-i) < len(room) {
						if room[y+(x-i)][i].isSeat {
							if room[y+(x-i)][i].occupied {
								occupiedCounter++
							}
							break
						}
					} else {
						break
					}
				}
				for i := x - 1; i >= 0; i-- {
					if y-(x-i) >= 0 {
						if room[y-(x-i)][i].isSeat {
							if room[y-(x-i)][i].occupied {
								occupiedCounter++
							}
							break
						}
					} else {
						break
					}
				}
				for i := y + 1; i < len(room); i++ {
					if room[i][x].isSeat {
						if room[i][x].occupied {
							occupiedCounter++
						}
						break
					}
				}
				for i := y - 1; i >= 0; i-- {
					if room[i][x].isSeat {
						if room[i][x].occupied {
							occupiedCounter++
						}
						break
					}
				}
				if occupiedCounter == 0 && !room[y][x].occupied {
					place := place{
						isSeat:   true,
						occupied: true,
					}
					newRoom[y][x] = place
					changed = true
				} else if occupiedCounter >= 5 && room[y][x].occupied {
					place := place{
						isSeat:   true,
						occupied: false,
					}
					newRoom[y][x] = place
					changed = true
				} else {
					newRoom[y][x] = room[y][x]
				}
			}
		}
	}
	if changed {
		return _challenge2(newRoom)
	}
	return room, false, countOccupied(room)
}

func printRoom(room [][]place) {
	for x := 0; x < len(room); x++ {
		for y := 0; y < len(room[0]); y++ {
			fmt.Print(room[x][y].String())
		}
		fmt.Print("\n")
	}
	fmt.Print("\n\n")
}
