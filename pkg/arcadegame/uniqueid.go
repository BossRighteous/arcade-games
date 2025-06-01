package arcadegame

type UID uint32

var _uniqueId UID = 0

func UniqueID() UID {
	_uniqueId++
	return _uniqueId
}

func ResetUniqueID() {
	_uniqueId = 0
}
