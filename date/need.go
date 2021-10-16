package date

import "time"

type Need struct {
	location   Location
	dateFormat format
}

type format struct {
	dateType Format
	isSet    bool
}

func (n *Need) Location(local Location) {
	n.location = local
}

func (n *Need) Format(dateFormat Format) {
	n.dateFormat = format{
		dateType: dateFormat,
		isSet:    true,
	}
}

func (n *Need) GetLocation() *time.Location {
	if n.location == CST {
		return time.FixedZone("CST", 8*3600)
	}

	if n.location == UTC {
		return time.UTC
	}

	if n.location == Local {
		return time.Local
	}

	if n.location == ShangHai {

		location, err := time.LoadLocation(ShangHai.Value())
		if err == nil {
			return location
		}

		n.location = CST
		return n.GetLocation()
	}

	return time.Local
}
