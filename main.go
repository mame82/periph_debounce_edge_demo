// +build linux,arm

package main

import (
	"context"
	"fmt"
	"github.com/mame82/periph_debounce_edge_demo/pgpio"
	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/host"
	"time"
)

func main() {
	/*
	Extending Periph.io gpio PinIn edge detection with context.Context based ExtWaitForEdge method (instead of timeout)
	and debouncing (disabled if debounce duration set to 0).

	Only tested on RPi0W: Aborting internal go routine for edge detection is done by toggling GPIO pull resistor
	between up/down to force real 'WaitForEdge' method to return. This couldn't work on GPIOs without internal pull
	resistor and seems to be the only way to force the aforementioned method to return, in case no edge occurs and timeout
	is set to -1.
	 */

	host.Init()

	// Define GPIO pins
	gpioButton := pgpio.NewP4wnp1PinIO(gpioreg.ByName("GPIO13")) //mechanical button connected to VCC and GPIO13
	gpioLedR := pgpio.NewP4wnp1PinIO(gpioreg.ByName("GPIO23")) //red LED connected to GND and GPIO23
	gpioLedG := pgpio.NewP4wnp1PinIO(gpioreg.ByName("GPIO24")) //green LED connected to GND and GPIO24
	gpioLedB := pgpio.NewP4wnp1PinIO(gpioreg.ByName("GPIO25")) //blue LED connected to GND and GPIO25

	// Quick output test with 3 LEDs or single RGB LED
	delay1 := time.Millisecond * 200
	delay2 := time.Millisecond * 400

	/* Testing output with color LEDs */
	for i := 0; i < 1; i++ {

		gpioLedR.Out(gpio.High)
		time.Sleep(delay1)
		gpioLedG.Out(gpio.High)
		time.Sleep(delay1)
		gpioLedB.Out(gpio.High)

		time.Sleep(delay2)

		gpioLedR.Out(gpio.Low)
		time.Sleep(delay1)
		gpioLedG.Out(gpio.Low)
		time.Sleep(delay1)
		gpioLedB.Out(gpio.Low)
		time.Sleep(delay1)
	}


	/*
	EdgeDetection test with context, no debounce
	 */

	//The following test ends if 1) the timeout is reached or 2) edgeDetection triggers 'edgeCountForCancel' times
	edgeCountForCancel := 5
	timeoutForCancel := time.Second * 20

	counter := 0
	//start edge detection
	gpioButton.In(gpio.PullDown, gpio.FallingEdge)
	ctx,cancel := context.WithTimeout(context.Background(), timeoutForCancel) //Context timeout after 20 seconds and cancels ExtWaitForEdge()
	defer cancel()
	for {
		level,err := gpioButton.ExtWaitForEdge(ctx,0) //<-- extended WaitForEdge method, accepting context and debounce duration
		if err != nil {
			fmt.Println("ExtWaitForEdge error: " + err.Error())
			break
		}
		fmt.Printf("Edge detected, level: %v\n", level)
		counter++

		if counter > edgeCountForCancel {
			cancel() //cancel ExtWaitForEdge, before timeout if more than 5 edges
		}
	}
	//Stop edge detection
	gpioButton.In(gpio.PullNoChange, gpio.NoEdge)


	/*
	Same EdgeDetection test with context, but debounce 100 ms debounce
	 */

	//The following test ends if 1) the timeout is reached or 2) edgeDetection triggers 'edgeCountForCancel' times
	edgeCountForCancel = 10
	timeoutForCancel = time.Second * 40

	counter = 0
	//start edge detection
	gpioButton.In(gpio.PullDown, gpio.FallingEdge)
	ctx,cancel = context.WithTimeout(context.Background(), timeoutForCancel) //Context timeout after 20 seconds and cancels ExtWaitForEdge()
	defer cancel()
	for {
		level,err := gpioButton.ExtWaitForEdge(ctx,100 * time.Millisecond) //<-- extended WaitForEdge method, accepting context and debounce duration
		if err != nil {
			fmt.Println("ExtWaitForEdge error: " + err.Error())
			break
		}
		fmt.Printf("Edge detected, level: %v\n", level)
		counter++

		if counter > edgeCountForCancel {
			cancel() //cancel ExtWaitForEdge, before timeout if more than 5 edges
		}
	}
	//Stop edge detection
	gpioButton.In(gpio.PullNoChange, gpio.NoEdge)




	fmt.Println("Waiting for additional messages for 10 seconds")
	// additional sleep to see some message
	time.Sleep(10*time.Second)
	fmt.Println("Finished")

}
