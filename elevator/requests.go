package elevator

import (
	"root/config"
	"root/elevio"
)

// HandleButtonPress handles the button press event and updates the elevator state

type Assignments [config.NumFloors][config.NumButtons]bool

func (a Assignments) ReqInDirection(floor int, dir Direction) bool {
	switch dir {
<<<<<<< HEAD
	case Up:
		for f := floor + 1; f < config.NumFloors; f++ {
			for b := 0; b < config.NumButtons; b++ {
				if a[f][b] {
					return true
				}
			}
=======
		case Up:
			for i := floor + 1; i < elevio.NumFloors; i++ {
				if a[i][elevio.BT_HallUp] || a[i][elevio.BT_Cab] {
					return true
				}
			}
			return false
		case Down:
			for i := 0; i < floor; i++ {
				for j := 0; j < 3; j++ {
					if a[i][j] {
						return true
					}
				}
			}
			return false
		default:
			panic("Invalid direction")
		}
}

<<<<<<< HEAD
func EmptyAssigner(floor int, dir Direction, a Assignments, orderDoneC chan<- elevio.ButtonEvent) {
=======

func EmptyAssingner(floor int, dir Direction, a Assingments, orderDoneC chan<- elevio.ButtonEvent) {
	if a[floor][elevio.BT_Cab] {
		orderDoneC <- elevio.ButtonEvent{Floor: floor, Button: elevio.BT_Cab}
	}
	if a[floor][dir] {
		orderDoneC <- elevio.ButtonEvent{Floor: floor, Button: dir.toBT()}
	}
}