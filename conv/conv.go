package conv

/*
  I denne pakken skal alle konverteringfunksjonene
  implementeres. Bruk engelsk.
    FarhenheitToCelsius
    CelsiusToFahrenheit
    KelvinToFarhenheit
    ...
*/

// Konverterer Farhenheit til Celsius
func FarhenheitToCelsius(value float64) float64 {
	// Her skal du implementere funksjonen
	// Du skal ikke formattere float64 i denne funksjonen
	// Gjør formattering i main.go med fmt.Printf eller
	// lignende
	// Legger inn formellen Celsius = (Farhrenheit - 32)*(5/9)
  return (value - 32)*(5.0/9.0)
}

// De andre konverteringsfunksjonene implementere her
// ...
// Konverterer Celsius til Farenheit
func CelsiusToFahrenheit(value float64) float64 {
  return value*(9.0/5.0) + 32
}

// Konverterer Kelvin til Celsius
func KelvinToCelsius(value float64) float64 {
  return value - 273.15
}

// Konverterer Celsius til Kelvin
func CelsiusToKelvin(value float64) float64 {
  return value + 273.15
}

// Konverterer Farenheit til Kelvin
func FarenheitToKelvin(value float64) float64 {
  return (value - 32) * (5.0/9.0) + 273.15
}

// Konverterer Kelvin til Farenheit
func KelvinToFarenheit(value float64) float64 {
  return (value - 273.15)*(9.0/5.0) + 32
}