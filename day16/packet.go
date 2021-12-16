package main

import (
	"math"
)

type packet struct {
	version    int
	type_id    int
	literal    int
	subPackets []packet
}

func parsePacket(bits bitList) (packet, bitList) {
	type_id := bits[3:6]
	if type_id.intValue() == 4 {
		return parseLiteralPacket(bits)
	} else {
		return parseOperatorPacket(bits)
	}
}

func newLiteralPacket(version bitList, type_id, literal bitList) packet {
	return packet{version.intValue(), type_id.intValue(), literal.intValue(), []packet{}}
}

func newOperatorPacket(version bitList, type_id bitList, subPackets []packet) packet {
	return packet{version.intValue(), type_id.intValue(), 0, subPackets}
}

func (p *packet) versionSum() int {
	versionSum := int((*p).version)

	for _, sub := range (*p).subPackets {
		versionSum += sub.versionSum()
	}

	return versionSum
}

func (p *packet) value() int {
	var result int

	if p.type_id == 0 {
		for _, sub := range p.subPackets {
			result += sub.value()
		}
	} else if p.type_id == 1 {
		result = 1
		for _, sub := range p.subPackets {
			result *= sub.value()
		}
	} else if p.type_id == 2 {
		result = math.MaxInt
		for _, sub := range p.subPackets {
			val := sub.value()
			if result > val {
				result = val
			}
		}
	} else if p.type_id == 3 {
		result = 0
		for _, sub := range p.subPackets {
			val := sub.value()
			if result < val {
				result = val
			}
		}
	} else if p.type_id == 4 {
		result = int(p.literal)
	} else if p.type_id == 5 {
		if p.subPackets[0].value() > p.subPackets[1].value() {
			result = 1
		} else {
			result = 0
		}
	} else if p.type_id == 6 {
		if p.subPackets[0].value() < p.subPackets[1].value() {
			result = 1
		} else {
			result = 0
		}
	} else if p.type_id == 7 {
		if p.subPackets[0].value() == p.subPackets[1].value() {
			result = 1
		} else {
			result = 0
		}
	}

	return int(result)
}

func parseLiteralPacket(bits bitList) (packet, bitList) {
	version, rest := bits[:3], bits[3:]
	bits = rest

	type_id, rest := bits[:3], bits[3:]
	bits = rest

	literal := bitList{}
	endLiteral := false
	for !endLiteral {
		prefix, literalPart, rest := bits[0], bits[1:5], bits[5:]

		bits = rest
		literal = append(literal, literalPart...)
		endLiteral = !prefix
	}

	return newLiteralPacket(version, type_id, literal), bits
}

func parseOperatorPacket(bits bitList) (packet, bitList) {
	packets := []packet{}
	version, rest := bits[:3], bits[3:]
	bits = rest

	type_id, rest := bits[:3], bits[3:]
	bits = rest

	length_type, rest := bits[0], bits[1:]
	bits = rest

	if length_type {
		length, rest := bits[:11], bits[11:]
		bits = rest

		intLength := length.intValue()
		for i := 0; i < intLength; i++ {
			packet, rest := parsePacket(bits)
			bits = rest
			packets = append(packets, packet)
		}
	} else {
		length, rest := bits[:15], bits[15:]
		bits = rest

		intLength := length.intValue()
		subPackets, rest := bits[:intLength], bits[intLength:]
		bits = rest

		for len(subPackets) > 1 {
			packet, rest := parsePacket(subPackets)
			packets = append(packets, packet)
			subPackets = rest
		}

	}

	return newOperatorPacket(version, type_id, packets), bits
}
