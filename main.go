package main

import (
	"fmt"
	"machine"
	"math"
	"time"
)

func voltageToDB(vIn, vOut float64) float64 {
	if vIn <= 0 || vOut <= 0 {
		return 0
	}
	return 20 * math.Log10(vOut/vIn)
}

func main() {
	machine.InitADC()

	referenceVoltage := 3.3

	adc0 := machine.ADC{Pin: machine.ADC0} // ADC1 = GP26 = GPIO31
	adc0.Configure(machine.ADCConfig{})
	adc1 := machine.ADC{Pin: machine.ADC1} // ADC1 = GP27 = GPIO32
	adc1.Configure(machine.ADCConfig{})

	for {
		time.Sleep(500 * time.Millisecond)
		println()

		rawVIn := adc0.Get() // Raw, 0-65535
		rawVOut := adc1.Get()

		vIn := (float64(rawVIn) / 65535.0) * referenceVoltage // convert raw value
		vOut := (float64(rawVOut) / 65535.0) * referenceVoltage

		fmt.Printf("Vin   (V): %.8f\n", vIn)
		fmt.Printf("Vout  (V): %.8f\n", vOut)
		fmt.Printf("Gain  (dB): %.2f\n\n", voltageToDB(vIn, vOut))
	}
}
