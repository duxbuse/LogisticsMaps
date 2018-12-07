package logisticsmaps

/*
Vdrop is a function that calculates the voltage drop for a given lenght of copper cable @ 25 degrees

the formula is

voltage =(lenght * current * 0.017)/area

total lengh of cable run in meters
current in amps
area is cross section of cable in mm^2

input voltage is user defined.

output voltage = input voltage - voltage drop

input: Voltage_input, current, length, area
returns: Voltage_drop, Voltage_output
*/

/*
Vfinal calculates the final output voltage given a specified voltage drop
*/

func Drop(current float64, length float64, area float64) float64 {

	return (current * length * 0.017) / area
}

func Vfinal(Vinput float64, Vdrop float64) float64 {
	return Vinput - Vdrop
}
