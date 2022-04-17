package sample

import "github.com/nikit34/grpc_basic_tutorial/complete_course/pb"


func Keyboard() *pb.Keyboard {
	keyboard := &pb.Keyboard{
		Layout: randomKeyboardLayout(),
		Backlit: randomBool(),
	}
	return keyboard
}