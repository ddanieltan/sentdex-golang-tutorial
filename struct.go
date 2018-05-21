package main

import "fmt"

const usixteenbitmax float64 = 65535
const kmh_multipler float64 = 1.60934

type car struct {
	gas_pedal      uint16 //min 0 max 65535
	break_pedal    uint16
	steering_wheel int16 //-32k to +32k
	top_speed_kmh  float64
}

func main() {
	a := car{gas_pedal: 22031, break_pedal: 0, steering_wheel: 12561, top_speed_kmh: 225.0}
	fmt.Println(a.gas_pedal)
	fmt.Println(a.kmh())
	fmt.Println(a.mph())
	// a.new_top_speed(500)
	a = newer_top_speed(a,500)
	fmt.Println(a.kmh())
	fmt.Println(a.mph())
}

// value receiver
func (c car) kmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbitmax)
}

func (c car) mph() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbitmax) / kmh_multipler
}

// pointer receiver
func (c *car) new_top_speed(newspeed float64) {
	c.top_speed_kmh = newspeed

}

func newer_top_speed(c car, speed float64) car {
	c.top_speed_kmh=speed
	return c
}