package shago

import (
	"crypto/sha256"
	"github.com/1lann/krist-miner/sha2"
)

type generator struct{}

func (g *generator) Sum256Number(data []byte) int64 {
	result := sha256.Sum256(data)
	// Turn last 6 bytes to int64
	return int64(result[31]) + int64(result[30])<<(8*1) +
		int64(result[29])<<(8*2) + int64(result[28])<<(8*3) +
		int64(result[27])<<(8*4) + int64(result[26])<<(8*5)
}

func (g *generator) Sum256NumberCmp(data []byte, work int64) bool {
	result := sha256.Sum256(data)

	value := int64(result[31])
	if value > work {
		return false
	}

	value += int64(result[30]) << (8 * 1)
	if value > work {
		return false
	}

	value += int64(result[29]) << (8 * 2)
	if value > work {
		return false
	}

	value += int64(result[28]) << (8 * 3)
	if value > work {
		return false
	}

	value += int64(result[27]) << (8 * 4)
	if value > work {
		return false
	}

	value += int64(result[26]) << (8 * 5)
	if value > work {
		return false
	}

	return true
}

func init() {
	sha2.RegisterAlgorithm("go", func() sha2.SumNumberAlgorithm {
		return &generator{}
	})
}