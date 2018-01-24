package space

const EarthYear = 31557600

type Planet = string

var YearMap = map[string]float64 {
	"Mercury": 0.2408467,
	"Venus":  0.61519726,
	"Mars": 1.8808158,
	"Jupiter": 11.862615,
	"Saturn": 29.447498,
	"Uranus": 84.016846,
	"Neptune": 164.79132,
	"Earth": 1.0,
}


func Age(seconds float64, planet string) float64 {
	// Determine the age of someone on the specified planet given the age in seconds on Earth.

	// Calculate the number of Earth years `seconds` converts to
	earth_age := seconds / EarthYear
	// Get the conversion factor that converts Earth years to years on the specified planet
	earth_years_per_year := YearMap[planet]
	// Convert the earth age to planet age.
	planet_years := earth_age / earth_years_per_year
	return planet_years
}
