package main

type Algorithm []bool

func (al *Algorithm) GetBit(bits []bool) bool {
	key := 0

	if bits[0] {
		key += 256
	}
	if bits[1] {
		key += 128
	}
	if bits[2] {
		key += 64
	}
	if bits[3] {
		key += 32
	}
	if bits[4] {
		key += 16
	}
	if bits[5] {
		key += 8
	}
	if bits[6] {
		key += 4
	}
	if bits[7] {
		key += 2
	}
	if bits[8] {
		key += 1
	}

	return (*al)[key]
}

func (al *Algorithm) Blips() bool {
	return (*al)[0] && !(*al)[len(*al)-1]
}
