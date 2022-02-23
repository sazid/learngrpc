package sample

import (
	"math/rand"

	"github.com/golang/protobuf/ptypes"
	v1 "github.com/sazid/learngrpc/api/v1"
)

// NewKeyboard returns a new sample keyboard
func NewKeyboard() *v1.Keyboard {
	keyboard := &v1.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}
	return keyboard
}

func NewCPU() *v1.CPU {
	brand := randomCPUBrand()
	name := randomCPUName(brand)

	numberCores := randomInt(2, 8)
	numberThreads := randomInt(numberCores, 12)

	minGhz := randomFloat64(2.0, 3.5)
	maxGhz := randomFloat64(minGhz, 5.0)

	cpu := &v1.CPU{
		Brand:         brand,
		Name:          name,
		NumberCores:   uint32(numberCores),
		NumberThreads: uint32(numberThreads),
		MinGhz:        minGhz,
		MaxGhz:        maxGhz,
	}
	return cpu
}

func NewGPU() *v1.GPU {
	brand := randomGPUBrand()
	name := randomGPUName(brand)

	minGhz := randomFloat64(1.0, 1.5)
	maxGhz := randomFloat64(minGhz, 2.0)

	memory := &v1.Memory{
		Value: uint64(randomInt(2, 6)),
		Unit:  v1.Memory_GIGABYTE,
	}

	gpu := &v1.GPU{
		Brand:  brand,
		Name:   name,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: memory,
	}
	return gpu
}

func NewRAM() *v1.Memory {
	ram := &v1.Memory{
		Value: uint64(randomInt(4, 32)),
		Unit:  v1.Memory_GIGABYTE,
	}
	return ram
}

func newSSD() *v1.Storage {
	ssd := &v1.Storage{
		Driver: v1.Storage_SSD,
		Memory: &v1.Memory{
			Value: uint64(randomInt(64, 1024)),
			Unit:  v1.Memory_GIGABYTE,
		},
	}
	return ssd
}

func newHDD() *v1.Storage {
	hdd := &v1.Storage{
		Driver: v1.Storage_HDD,
		Memory: &v1.Memory{
			Value: uint64(randomInt(1, 6)),
			Unit:  v1.Memory_TERABYTE,
		},
	}
	return hdd
}

func NewStorage() *v1.Storage {
	if rand.Intn(2) == 1 {
		return newHDD()
	}
	return newSSD()
}

func NewScreen() *v1.Screen {
	screen := &v1.Screen{
		SizeInch:   randomFloat32(10, 18),
		Resolution: randomScreenResolution(),
		Panel:      randomScreenPanel(),
		Multitouch: randomBool(),
	}
	return screen
}

func NewLaptop() *v1.Laptop {
	laptop := &v1.Laptop{
		Id:       randomId(),
		Brand:    randomLaptopBrand(),
		Name:     "Custom Laptop",
		Cpu:      NewCPU(),
		Ram:      NewRAM(),
		Gpus:     randomNumberItems(2, NewGPU),
		Screen:   NewScreen(),
		Storages: randomNumberItems(2, NewStorage),
		Weight: &v1.Laptop_WeightKg{
			WeightKg: randomFloat64(1, 3),
		},
		PriceUsd:    randomFloat64(500, 1500),
		ReleaseYear: uint32(randomInt(2020, 2024)),
		UpdatedAt:   ptypes.TimestampNow(),
	}
	return laptop
}
