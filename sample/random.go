package sample

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
	v1 "github.com/sazid/learngrpc/api/v1"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomKeyboardLayout() v1.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return v1.Keyboard_QWERTY
	case 2:
		return v1.Keyboard_QWERTZ
	default:
		return v1.Keyboard_AZERTY
	}
}

func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD", "Qualcomm", "MediaTek")
}

func randomCPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFromSet(
			"Core i3",
			"Core i5",
			"Core i7",
			"Core i9",
		)
	} else if brand == "AMD" {
		return randomStringFromSet(
			"Ryzen 5 1600",
			"Ryzen 5 2600",
			"Ryzen 5 3600",
			"Ryzen 5 5600",
		)
	} else if brand == "Qualcomm" {
		return randomStringFromSet(
			"SD855",
			"SD855+",
			"SD865",
			"SD888",
			"SD888+",
			"SD 8 Gen1",
		)
	} else if brand == "MediaTek" {
		return randomStringFromSet(
			"Dimensity 1000",
			"Dimensity 1200",
		)
	}

	return "unknown"
}

func randomGPUBrand() string {
	return randomStringFromSet("Nvidia", "AMD")
}

func randomGPUName(brand string) string {
	if brand == "Nvidia" {
		return randomStringFromSet(
			"RTX 2060",
			"RTX 2070",
			"RTX 3060",
			"RTX 3070",
		)
	} else if brand == "AMD" {
		return randomStringFromSet(
			"RX 580",
			"RX 590",
			"RX 5700-XT",
			"RX Vega-56",
		)
	}

	return "unknown"
}

func randomScreenResolution() *v1.Screen_Resolution {
	width := randomInt(1080, 4320)
	height := width * 16 / 9
	return &v1.Screen_Resolution{
		Width:  uint32(width),
		Height: uint32(height),
	}
}

func randomScreenPanel() v1.Screen_Panel {
	switch rand.Intn(2) {
	case 1:
		return v1.Screen_IPS
	default:
		return v1.Screen_OLED
	}
}

func randomLaptopBrand() string {
	return randomStringFromSet("Acer", "Asus", "HP")
}

func randomNumberItems[T any](n int, f func() T) []T {
	res := make([]T, n)
	for i := 0; i < n; i++ {
		res[i] = f()
	}
	return res
}

func randomId() string {
	return uuid.New().String()
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}
