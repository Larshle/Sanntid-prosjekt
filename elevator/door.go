package elevator

import (
	"root/config"
	"root/elevio"
	"time"
)

const (
	DoorOpenDuration = 2 * time.Second
)
>>>>>>> parent of 34a4414 (Merge pull request #2 from Larshle/UpdatingNEWassignments)

type DoorState int

const (
	Open DoorState = iota
	Closed
	Obstructed
	InCountDown
	Obstructed
)

func Door(doorClosedC chan<- bool, doorOpenC <-chan bool) {

	elevio.SetDoorOpenLamp(false)
	obstructionC := make(chan bool)
	go elevio.PollObstructionSwitch(obstructionC)

	obstruction := false
	timeCounter := time.NewTimer(time.Hour)
	var ds DoorState = Closed

	for {
		select {
		case obstruction = <-obstructionC:
			if !obstruction && ds == Obstructed {
				elevio.SetDoorOpenLamp(false)
				doorClosedC <- true
				ds = Closed
			}

		case <-doorOpenC:
			switch ds {
			case Closed:
				elevio.SetDoorOpenLamp(true)
				timeCounter = time.NewTimer(DoorOpenDuration)
				ds = InCountDown
			case InCountDown:
				timeCounter = time.NewTimer(DoorOpenDuration)
				
			case Obstructed:
				timeCounter = time.NewTimer(DoorOpenDuration)
				ds = InCountDown

			default:
				panic("Door state not implemented")
			}
		case <-timeCounter.C:
			if ds != InCountDown {
				panic("Door state not implemented")
			}
			if obstruction {
				ds = Obstructed
			} else {
				elevio.SetDoorOpenLamp(false)
				doorClosedC <- true
				ds = Closed
			}
		}
	}
}
