package main

import "testing"

func TestDistanceCalc1 (t *testing.T) {

	testingDistance := func(t testing.TB, got, want int) {
        t.Helper()
        if got != want {
            t.Errorf("got %v, want %v", got, want)
        }
    }


	t.Run(" diagonal correct ", func(t *testing.T) {
       
		got := DistanceCalc(0,0,3,4)
        want := 5

        testingDistance(t, got, want)
    })

	t.Run(" vertical correct ", func(t *testing.T) {
       
		got := DistanceCalc(0,0,0,4)
        want := 4

        testingDistance(t, got, want)
    })

	t.Run(" horizontal correct ", func(t *testing.T) {
       
		got := DistanceCalc(0,4,3,4)
        want := 3

        testingDistance(t, got, want)
    })

	t.Run(" same point correct ", func(t *testing.T) {
       
		got := DistanceCalc(0,0,0,0)
        want := 0

        testingDistance(t, got, want)
    })


}