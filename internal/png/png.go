package png

type header [8]byte

var standartHeader header = [...]byte{137, 80, 78, 71, 13, 10, 26, 10}

func StandartHeader() header {
	return standartHeader
}

type png struct {
	
}
