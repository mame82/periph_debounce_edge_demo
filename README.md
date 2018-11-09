# Periph.io GPIO with debounced edge detection and context.Context based ExWaitForEdge

See this tweet by @marcaruel for ref: https://twitter.com/marcaruel/status/1060714157505437696

The code is extracted from a still private project (P4wnP1 successor) and thus isn't
cleaned or polished. It is not meant for production use, but for sharing ideas.

Code builds for ARM(v6) Linux and is meant to be run on a Pi0 (maybe others work).
It couldn't work with GPIOs without internal pull resistors, as they are used to abort 
WaitForEdge


Demo setup:
- push button connected with 5V and GPIO13
- red LED connected with GND and GPIO23 (don't forget proper resistor)
- green LED connected with GND and GPIO24 (don't forget proper resistor)
- blue LED connected with GND and GPIO25 (don't forget proper resistor)
