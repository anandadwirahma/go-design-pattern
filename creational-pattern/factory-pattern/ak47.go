package main

/*
**notes :
	AK47 as a Concrete Product A (embed gun struct)
*/

type ak47 struct {
	gun
}

func newAk47() iGun {
	return &ak47{
		gun: gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}
