package builder

type Ship struct {
	ArmorType        string
	RadarType        string
	SonarType        string
	ArtilleryCaliber int
	AirDefenceType   int
}

type Builder interface {
	BuildArmor(armor string)
	BuildRadar(radar string)
	BuildSonar(sonar string)
	BuildArtillery(caliber int)
	BuildAirDefence(caliber int)
}

type ShipBuilder struct {
	ship Ship
}

func (sb *ShipBuilder) BuildArmor(armor string) {
	sb.ship.ArmorType = armor
}

func (sb *ShipBuilder) BuildAirDefence(caliber int) {
	sb.ship.AirDefenceType = caliber
}

func (sb *ShipBuilder) BuildRadar(radar string) {
	sb.ship.RadarType = radar
}

func (sb *ShipBuilder) BuildSonar(sonar string) {
	sb.ship.RadarType = sonar
}

func (sb *ShipBuilder) BuildArtillery(caliber int) {
	sb.ship.ArtilleryCaliber = caliber
}

func (sb *ShipBuilder) GetShip() Ship {
	return sb.ship
}

type NavalPort struct {
	ship ShipBuilder
}

func (np *NavalPort) Build(armor, radar, sonar string, artillery, airguns int) {
	np.ship.BuildArmor(armor)
	np.ship.BuildRadar(radar)
	np.ship.BuildSonar(sonar)
	np.ship.BuildArtillery(artillery)
	np.ship.BuildAirDefence(airguns)
}

func mainBuilder() {
	destroyer := &ShipBuilder{}

	navyPort := &NavalPort{ship: *destroyer}
}
