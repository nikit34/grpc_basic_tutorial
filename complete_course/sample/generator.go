package sample

import "github.com/nikit34/grpc_basic_tutorial/complete_course/pb"


func Keyboard() *pb.Keyboard {
	keyboard := &pb.Keyboard{
		Layout: randomKeyboardLayout(),
		Backlit: randomBool(),
	}
	return keyboard
}

func NewCPU() *pb.CPU {
	brand := randomCPUBrand()
	name := randomCPUName(brand)
	numberCores := randomInt(2, 8)
	numberThreads := randomInt(numberCores, 12)
	minGhz := randomFloat(2.0, 3.5)
	maxGhz := randomFloat(minGhz, 5.0)

	cpu := &pb.CPU{
		Brand: brand,
		Name: name,
		NumberCores: uint32(numberCores),
		NumberThreads: uint32(numberThreads),
		MinGhz: minGhz,
		MaxGhz: maxGhz,
	}
	return cpu
}