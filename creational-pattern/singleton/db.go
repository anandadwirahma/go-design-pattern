package singleton

import (
	"fmt"
	"sync"
)

var (
	instance *db
)

type db struct {
}

func (d db) GetUser(ID string) string {
	return ID
}

var lock = &sync.Mutex{}

var once sync.Once

// GetDBInstance is using singleton
func GetDBInstanceWithLock() *db {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			fmt.Println("Creating single instance now.")
			instance = &db{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}
	return instance
}

func GetDBInstanceWithOnce() *db {
	if instance == nil {
		once.Do(
			func() {
				fmt.Println("Creating single instance now.")
				instance = &db{}
			})
	} else {
		fmt.Println("Single instance already created.")
	}

	return instance
}
