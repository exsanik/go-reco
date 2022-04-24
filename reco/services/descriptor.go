package services

import "math"

func DescriptorToBytes(descriptor [128]float32) [512]byte {
	var result [512]byte

	var buffer = result[:0]

	for i := 0; i < 128; i++ {
		var bits uint32 = math.Float32bits(descriptor[i])

		buffer = append(
			buffer,
			byte(bits),
			byte(bits>>8),
			byte(bits>>16),
			byte(bits>>24),
		)
	}

	return result
}

func BytesToDescriptor(bytes [512]byte) [128]float32 {
	var result [128]float32

	var i = 0

	for j := 0; j < 512; j += 4 {
		result[i] = math.Float32frombits(
			uint32(bytes[j]) +
				uint32(bytes[j+1])<<8 +
				uint32(bytes[j+2])<<16 +
				uint32(bytes[j+3])<<24,
		)

		i += 1
	}

	return result
}
