package tempconv

import "testing"

func TestTempConv(t *testing.T) {
	if FToC(Fahrenheit(86)) != Celsius(30) {
		t.Errorf("Expected 86°F to equal 30°C")
	}
	
	if CToF(Celsius(30)) != Fahrenheit(86) {
		t.Errorf("Expected 86°F to equal 30°C")
	}

	if CToK(Celsius(0)) != Kelvin(-273.15) {
		t.Errorf("Expected 0°C to equal -273.15°K, got %v instead", CToK(Celsius(0)))
	}
	
	if KToC(Kelvin(0)) != Celsius(-273.15) {
		t.Errorf("Expected 0°K to equal -273.15°C, got %v instead", KToC(Kelvin(0)))
	}
	
	if k := Kelvin(273.15); KToF(k) != Fahrenheit(32) {
		t.Errorf("Expected 32°F to equal 273.15°K, got %v instead", KToF(k))
	}
}
