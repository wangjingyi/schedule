package main

import "fmt"

//////////////////////////////////////////////////////////////////
// This program try to schedule optometrists to do the eye test
// the time slots are from 9am to 9pm and a unit is 30 minutes
// so people can schedule 9:00, 9:30, 10, 10:30 ... 9:00pm
// assume each optometrist have to use 3 units (90 minutes)
// to finish his work and can start taking the next one
// For a given time slot, say 9:30am to see if that slot
// is available

const numSlots = 24

func inRange(i int) bool {
  return i >= 0 && i < numSlots
}

var index = map[string]int{
  "9:00": 0,
  "9:30": 1,
  "10:00": 2,
  "10:30": 3,
  "11:00": 4,
  "11:30": 5,
  "12:00": 6,
  "12:30": 7,
  "13:00": 8,
  "13:30": 9,
  "14:00": 10,
  "14:30": 11,
  "15:00": 12,
  "15:30": 13,
  "16:00": 14,
  "16:30": 15,
  "17:00": 16,
  "17:30": 17,
  "18:00": 18,
  "18:30": 19,
  "19:00": 20,
  "19:30": 21,
  "20:00": 22,
  "20:30": 23,
}

type Optometrist struct {
   name string
   slots [numSlots]byte /* 0 means available, 1 means occupy */
}

func (op *Optometrist) clear() {
  for i := 0; i < numSlots; i++ {
    op.slots[i] >>= 1
  }
}

func (op *Optometrist) assign(t string) {
  beg := index[t]
  for i := beg - 2; inRange(i) && i <= beg + 2; i++ {
    op.slots[i] = 1
  }
}

func (op *Optometrist) available(t string) bool {
  return op.slots[index[t]] == 0
}

type Schedule []*Optometrist

// return false means time slot is not available
// so assignment was failed
func (s Schedule) assign(t string) bool {
   for _, op := range s {
     if op.available(t) {
       op.assign(t)
       return true
     }
   }
   return false
}

func (s Schedule) available(t string) bool {
  for _, op := range s {
    if op.available(t) {
      return true
    }
  }
  return false
}

func (s Schedule) clear() {
  for _, op := range s {
    op.clear()
  }
}

func main() {
  var schedule = Schedule {
                  &Optometrist{ name:  "A" },
                  &Optometrist{ name:  "B" },
                  &Optometrist{ name:  "C" },
                }
  fmt.Printf("%v\n", schedule.assign("11:30"))
  fmt.Printf("%v\n", schedule.assign("10:00"))
  fmt.Printf("%v\n", schedule.assign("10:00"))
  fmt.Printf("%v\n", schedule.assign("11:00"))
  fmt.Printf("%v\n", schedule.assign("10:30"))
}
