package TMC

import (
	//"periph.io/x/periph/conn/gpio"
	//"periph.io/x/periph/conn/gpio/gpioreg"
	//"periph.io/x/periph/host"

	"log"

	"github.com/tarm/serial"
)

type TMC_UART struct {
	mtr_id                int32
	ser                   *int32
	r_frame               []int32
	w_frame               []int32
	communication_pause   int
	error_handler_running bool
}

var uart = TMC_UART{0, nil, []int32{0x55, 0, 0, 0}, []int32{0x55, 0, 0, 0, 0, 0, 0, 0}, 0, false}

func Init_uart(serialport string, baudrate int32, mtr_id int32) {

	c := &serial.Config{Name: serialport, Baud: 115200}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	n, err := s.Write([]byte("test"))
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 128)
	n, err = s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%q", buf[:n])

}
