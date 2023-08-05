package main

import (
	"testing"
)


func TestStore(t *testing.T) {
    concierge := BaggageConcierge(100)

    // Test storage full scenario
    for i := 0; i < 100; i++ {
        bag := Bag{BagType: "carry-on", Size: 1, BagID: i}
        _, err := concierge.Store(bag)
        if err != nil && err.Error() != "Storage is full" {
            t.Errorf("Expected 'Storage is full' error, got: %v", err)
        }
    }

    // Test full for only one bag type scenario
    bag1 := Bag{BagType: "checked", Size: 1, BagID: 201}
    _, err := concierge.Store(bag1)
    if err != nil {
        t.Errorf("Error storing bag: %v", err)
    }

    bag2 := Bag{BagType: "checked", Size: 1, BagID: 202}
    _, err = concierge.Store(bag2)
    if err == nil || err.Error() != "Storage is full for the requested bag type" {
        t.Errorf("Expected 'Storage is full for the requested bag type' error, got: %v", err)
    }

    bag3 := Bag{BagType: "carry-on", Size: 1, BagID: 203}
    _, err = concierge.Store(bag3)
    if err != nil && err.Error() != "Storage is full" {
        t.Errorf("Expected 'Storage is full' error, got: %v", err)
    }
}

