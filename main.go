package main

import (
    "errors"
    "fmt"
)

type Bag struct {
    BagType string
    Size    int
    BagID   int
}

type Ticket struct {
    BagID int
}

type Bin struct {
    Bags []Bag
    Capacity int
}

type BaggageConciergeImpl struct {
    Bins []Bin
    MaxBinCapacity int
}

func BaggageConcierge(maxBinCapacity int) *BaggageConciergeImpl {
    return &BaggageConciergeImpl{
        Bins:           []Bin{{}},
        MaxBinCapacity: maxBinCapacity,
    }
}

func (bc *BaggageConciergeImpl) Store(bag Bag) (Ticket, error) {
   for i := range bc.Bins {
        if (bag.BagType == "carry-on" && len(bc.Bins[i].Bags) < bc.Bins[i].Capacity && bag.Size <= 2) ||
           (bag.BagType == "checked" && len(bc.Bins[i].Bags) == 0 && bag.Size == 1) {
            bc.Bins[i].Bags = append(bc.Bins[i].Bags, bag)
            return Ticket{BagID: bag.BagID}, nil
        } else if bag.BagType == "carry-on" && len(bc.Bins[i].Bags) >= bc.Bins[i].Capacity {
            return Ticket{}, errors.New("Storage is full")
        } else if bag.BagType == "checked" && len(bc.Bins[i].Bags) >= bc.Bins[i].Capacity {
            return Ticket{}, errors.New("Storage is full for the requested bag type")
        }
    }

    return Ticket{}, errors.New("Storage is full")

}

func (bc *BaggageConciergeImpl) Retrieve(ticket Ticket) (Bag, error) {
    for i := range bc.Bins {
        for j, storedBag := range bc.Bins[i].Bags {
            if storedBag.BagID == ticket.BagID {
                bc.Bins[i].Bags = append(bc.Bins[i].Bags[:j], bc.Bins[i].Bags[j+1:]...)
                return storedBag, nil
            }
        }
    }

    return Bag{}, errors.New("bag not found")
}



func main() {
    concierge := BaggageConcierge(100)

    bag1 := Bag{BagType: "carry-on", Size: 2, BagID: 1}
    bag2 := Bag{BagType: "checked", Size: 1, BagID: 2}
    bag3 := Bag{BagType: "carry-on", Size: 2, BagID: 3}
    bag4 := Bag{BagType: "checked", Size: 1, BagID: 4}

    ticket1, err := concierge.Store(bag1)
    if err != nil {
        fmt.Println("Error storing bag:", err)
        return
    }
    fmt.Println("Bag 1 stored with ticket:", ticket1)

    ticket2, err := concierge.Store(bag2)
    if err != nil {
        fmt.Println("Error storing bag:", err)
        return
    }
    fmt.Println("Bag 2 stored with ticket:", ticket2)

    ticket3, err := concierge.Store(bag3)
    if err != nil {
        fmt.Println("Error storing bag:", err)
        return
    }
    fmt.Println("Bag 3 stored with ticket:", ticket3)

    ticket4, err := concierge.Store(bag4)
    if err != nil {
        fmt.Println("Error storing bag:", err)
        return
    }
    fmt.Println("Bag 4 stored with ticket:", ticket4)

    retrievedBag, err := concierge.Retrieve(ticket2)
    if err != nil {
        fmt.Println("Error retrieving bag:", err)
        return
    }
    fmt.Println("Retrieved Bag:", retrievedBag)

    nonExistentTicket := Ticket{BagID: 12345}
    _, err = concierge.Retrieve(nonExistentTicket)
    if err != nil {
        fmt.Println("Error retrieving bag:", err)
    }
}


