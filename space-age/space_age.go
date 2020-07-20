package space

type Planet string

const EarthSecondsAYear = 31557600

//seconds / (PlanetPeriod * 31557600 )
func Age(seconds float64, planet Planet) float64 {
	if seconds != 0 && planet != "" {
		switch planet {
		case "Mercury":
			return seconds / (0.2408467 * EarthSecondsAYear)
		case "Venus":
			return seconds / (0.61519726 * EarthSecondsAYear)
		case "Earth":
			return seconds / EarthSecondsAYear
		case "Mars":
			return seconds / (1.8808158 * EarthSecondsAYear)
		case "Jupiter":
			return seconds / (11.862615 * EarthSecondsAYear)
		case "Saturn":
			return seconds / (29.447498 * EarthSecondsAYear)
		case "Uranus":
			return seconds / (84.016846 * EarthSecondsAYear)
		case "Neptune":
			return seconds / (164.79132 * EarthSecondsAYear)
		}
	}
	return 0
}
