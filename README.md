# Periph.io GPIO with debounced edge detection and context.Context based ExWaitForEdge

See this tweet by @marcaruel for ref: https://twitter.com/marcaruel/status/1060714157505437696

The code is extracted from a still private project (P4wnP1 successor) and thus isn't
cleaned or polished. It is not meant for production use, but for sharing ideas.

Code builds for ARM(v6) Linux and is meant to be run on a Pi0 (maybe others work).
It couldn't work with GPIOs without internal pull resistors, as they are used to abort 
WaitForEdge

## Demo setup:
- push button connected with 5V and GPIO13
- red LED connected with GND and GPIO23 (don't forget proper resistor)
- green LED connected with GND and GPIO24 (don't forget proper resistor)
- blue LED connected with GND and GPIO25 (don't forget proper resistor)

## Addtional notes
1) Edges detected between `In()` and `ExtWaitForEdge` aren't preserved (there's a boolean in `startEdgeDetection` to change this behavior, but it isn't tested as it doesn't allign with my use case)
2) Instead of using `In()` to define the debounce duration, i chose to aplly it on every call to `ExtWaitForEdge`. This allows adaptive scaling of the value between wait calls (and conflicts with preserving detected edges as highlighted above)
3) The `ExtWaitForEdge` method isn't meant to be called concurrently in this implementation (no NotifyAll like approach)
4) Wait for edge return the GPIO level, as read when the Edge interrupt occurs (this would be needed, if debouncing and event preserving are used together with gpio.BothEdges. For rising/falling edge high/low is returned without reading GPII state)
