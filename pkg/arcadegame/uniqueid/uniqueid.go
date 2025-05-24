package uniqueid

type UID int

var _uniqueId UID = -1

func UniqueID() UID {
	_uniqueId++
	return _uniqueId
}
