package main

/*
**notes :
	Musket as a Concrete Product B (embed gun struct)
*/

type musket struct {
	gun
}

func newMusket() iGun {
	return &musket{
		gun: gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}
