package ambulance_wl

import (
    "time"

    "slices"
)

func (this *Ambulance) reconcileWaitingList() {
    slices.SortFunc(this.WaitingList, func(left, right WaitingListEntry) int {
        if left.WaitingSince.Before(right.WaitingSince) {
            return -1
        } else if left.WaitingSince.After(right.WaitingSince) {
            return 1
        } else {
            return 0
        }
    })

    // we assume the first entry EstimatedStart is the correct one (computed before previous entry was deleted)
    // but cannot be before current time
    // for sake of simplicity we ignore concepts of opening hours here

    if this.WaitingList[0].EstimatedStart.Before(this.WaitingList[0].WaitingSince) {
        this.WaitingList[0].EstimatedStart = this.WaitingList[0].WaitingSince
    }

    if this.WaitingList[0].EstimatedStart.Before(time.Now()) {
        this.WaitingList[0].EstimatedStart = time.Now()
    }

    nextEntryStart :=
        this.WaitingList[0].EstimatedStart.
            Add(time.Duration(this.WaitingList[0].EstimatedDurationMinutes) * time.Minute)
    for _, entry := range this.WaitingList[1:] {
        if entry.EstimatedStart.Before(nextEntryStart) {
            entry.EstimatedStart = nextEntryStart
        }
        if entry.EstimatedStart.Before(entry.WaitingSince) {
            entry.EstimatedStart = entry.WaitingSince
        }

        nextEntryStart =
            entry.EstimatedStart.
                Add(time.Duration(entry.EstimatedDurationMinutes) * time.Minute)
    }
}
