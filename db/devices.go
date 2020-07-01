package db

type ExistingDevice struct {
	ID         string
	NumberInDB string
}

var Devices = [...]ExistingDevice {
	{"01963d76-d80b-4fb7-b759-a7276e6bbdb8", "90000001"},
}