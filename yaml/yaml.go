// Package yaml adds support for conversion of manually entered data stories to Go structs
// and creation of edges based on keys extracted from struct fields.
package yaml

import (
	"fmt"
	"log"

	"github.com/ghodss/yaml"
)

type Entity struct {
	KnownAs            string `json:"knownAs"` // Affects YAML field names too.
	Ownership          string `json:"ownership"`
	OrganizationalForm string `json:"organizationalForm"`
	Description        string `json:"description"`
	Comment            string `json:"comment"`
	Governmental       string `json:"governmental"`
	Domain             string `json:"domain"`
	Name               string `json:"name"`
	EndUser            bool   `json:"endUser"`
	ServiceProvider    bool   `json:"serviceProvider"`
	/*	BUSINESS_RELATIONSHIP_WITH_Entity []string //EDGE.BUSINESS_RELATIONSHIP_WITH_Entity
		ADOPTED_Initiative                []string //EDGE.ADOPTED_Initiative
		OPERATES_Project                  []string //EDGE.OPERATES_Project
		PROVIDES_Service                  []string //WHERE: #REF_TO_SERVICE
		LOCATED_IN_Place                  []string //WHERE: #REF_TO_PLACE
	*/
}

type Service struct {
	KnownAs                     string  `json:"knownAs"`
	PowerDeliveryMW             float32 `json:"powerDeliveryMW"`
	PolarPowerCapacityMWhYear   float32 `json:"solarPowerCapacityMWhYear"`
	BiomassCarbonImpact         string  `json:"biomassCarbonImpact"`
	HydroCaptureType            string  `json:"hydroCaptureType"`
	PowerStorageType            string  `json:"powerStorageType"`
	GridStabilization           string  `json:"gridStabilization"`
	PowerGenerationMW           float32 `json:"powerGenerationMW"`
	GeoThermalPowerCapacityMW   float32 `json:"geoThermalPowerCapacityMW"`
	SustainabiltyFocused        bool    `json:"sustainabiltyFocused"`
	ElectricUtilityService      bool    `json:"electricUtilityService"`
	SolarCaptureTechnology      string  `json:"solarCaptureTechnology"`
	GeoThermalCaptureType       string  `json:"geoThermalCaptureType"`
	Domain                      string  `json:"domain"`
	SolarPlantType              string  `json:"solarPlantType"`
	WindPowerCapacityMWhYear    float32 `json:"windPowerCapacityMWhYear"`
	WindCaptureTechnology       string  `json:"windCaptureTechnology"`
	EconomicModel               string  `json:"economicModel"`
	GeographicalScale           string  `json:"geographicalScale"`
	ElectricTransmissionService bool    `json:"electricTransmissionService"`
	Comment                     string  `json:"comment"`
	Description                 string  `json:"description"`
	NuclearPowerTechnology      string  `json:"nuclearPowerTechnology"`
	PowerStorageMW              float32 `json:"powerStorageMW"`
	WindPowerCapacityMW         float32 `json:"windPowerCapacityMW"`
	NuclearPowerCapacityMW      float32 `json:"nuclearPowerCapacityMW"`
	SolarPowerCapacityMW        float32 `json:"solarPowerCapacityMW"`
	WindPlantType               string  `json:"windPlantType"`
	BiomassPowerMWhYear         float32 `json:"biomassPowerMWhYear"`
	BiomassCaptureType          string  `json:"biomassCaptureType"`
	TargetUsers                 string  `json:"targetUsers"`
	HydroPowerCapacityMW        float32 `json:"hydroPowerCapacityMW"`
	ElectricWholesalingService  bool    `json:"electricWholesalingService"`
	/* PROVIDED_TO_Entity          []string //WHERE: #REF_TO_ENTITY
	 */
}

type Project struct {
	KnownAs             string `json:"knownAs"`
	OperationalDate     string `json:"operationalDate"`
	UsesNaturalResource string `json:"usesNaturalResource"`
	UsesBuiltResource   string `json:"usesBuiltResource"`
	Description         string `json:"description"`
	Comment             string `json:"comment"`
	Domain              string `json:"domain"`
	ProjectState        string `json:"projectState"`
	/*	LOCATED_IN_Place    []string //WHERE: #REF_TO_PLACE
		ENABLES_Service     []string //WHERE: #REF_TO_SERVICE
	*/
}

type Initiative struct {
	KnownAs                         string `json:"knownAs"`
	Comment                         string `json:"comment"`
	Domain                          string `json:"domain"`
	Purpose                         string `json:"purpose"`
	DecarbonizationPlan             bool   `json:"decarbonizationPlan"`
	MandatedDecarbonizationTargets  string `json:"mandatedDecarbonizationTargets"`
	VoluntaryDecarbonizationTargets string `json:"voluntaryDecarbonizationTargets"`
	Description                     string `json:"description"`
	/*	COMPLIES_WITH_Regulation        []string //EDGE.COMPLIES_WITH_Regulation]
		SUPPORTS_Project                string   //WHERE: #REF_TO_PROJECT
		APPLIES_TO_Place                string   //WHERE: #REF_TO_PLACE
		SUPPORTS_Service                string   //WHERE: #REF_TO_SERVICE
	*/
}

type Regulation struct {
	KnownAs           string `json:"knownAs"`
	RelatedRegulation string `json:"relatedRegulation"`
	Description       string `json:"description"`
	Compliance        string `json:"compliance"`
	Comment           string `json:"comment"`
	Domain            string `json:"domain"`
	StatutoryTextLink string `json:"statutoryTextLink"`
	Purpose           string `json:"purpose"`
	/*
	   APPLIES_TO_Place   string //WHERE: #REF_TO_PLACE
	   APPLIES_TO_Entity  string //WHERE: #REF_TO_ENTITY
	   APPLIES_TO_Service string //WHERE: #REF_TO_SERVICE
	   APPLIES_TO_Project string //WHERE: #REF_TO_PROJECT
	*/
}

type Place struct {
	KnownAs             string    `json:"knownAs"`
	Description         string    `json:"description"`
	State               string    `json:"state"`
	City                string    `json:"city"`
	InternationalRegion string    `json:"internationalRegion"`
	MultistateRegion    string    `json:"multistateRegion"`
	GridCoordinates     []float32 `json:"gridCoordinates"`
	Comment             string    `json:"comment"`
	Country             string    `json:"country"`
	County              string    `json:"county"`
}

type GovernmentAuthority struct {
	KnownAs          string `json:"knownAs"`
	RegulatoryAgency bool   `json:"regulatoryAgency"`
	Description      string `json:"description"`
	Comment          string `json:"comment"`
	Level            string `json:"level"`
	Domain           string `json:"domain"`
	/*	ADOPTED_Regulation []string //EDGE.ADOPTED_Regulation
		REGULATES_Project  string   //WHERE: #REF_TO_PROJECT
		REGULATES_Service  string   //WHERE: #REF_TO_SERVICE
		REGULATES_Entity   string   //WHERE: #REF_TO_ENTITY
		LOCATED_IN_Place   string   //WHERE: #REF_TO_PLACE
	*/
}

type EDGE struct {
	BUSINESS_RELATIONSHIP_WITH_Entity struct {
		businessRelationshipType string
		WHERE                    string //"#REF_TO_ENTITY"
	}
	ADOPTED_Initiative struct {
		adoptionDate string
		WHERE        string //"#REF_TO_INITIATIVE"
	}
	PROVIDES_Service struct {
		WHERE string //"#REF_TO_SERVICE"
	}
	OPERATES_Project struct {
		role  string
		WHERE string //"#REF_TO_PROJECT"
	}
	LOCATED_IN_Place struct {
		WHERE string //"#REF_TO_PLACE"
	}
	PROVIDED_TO_Entity struct {
		WHERE string //"#REF_TO_ENTITY"
	}
	ENABLES_Service struct {
		WHERE string //"#REF_TO_SERVICE"
	}
	COMPLIES_WITH_Regulation struct {
		compliance string
		WHERE      string //"#REF_TO_REGULATION"
	}
	SUPPORTS_Project struct {
		WHERE string //"#REF_TO_PROJECT"
	}
	APPLIES_TO_Place struct {
		WHERE string //"#REF_TO_PLACE"
	}
	SUPPORTS_Service struct {
		WHERE string //"#REF_TO_SERVICE"
	}
	APPLIES_TO_Entity struct {
		WHERE string //"#REF_TO_ENTITY"
	}
	APPLIES_TO_Service struct {
		WHERE string //"#REF_TO_SERVICE"
	}
	APPLIES_TO_Project struct {
		WHERE string //"#REF_TO_PROJECT"
	}
	ADOPTED_Regulation struct {
		adoptionDate string
		WHERE        string //"#REF_TO_REGULATION"
	}
	REGULATES_Project struct {
		WHERE string //"#REF_TO_PROJECT"
	}
	REGULATES_Service struct {
		WHERE string //"#REF_TO_SERVICE"
	}
	REGULATES_Entity struct {
		WHERE string //"#REF_TO_ENTITY"
	}
}

func Convert() string {
	// Marshal an Entity struct to YAML.
	DTE := Entity{"DTE", "private member", "for-profit corporate", "", "", "State", "Electric Power", "Delaware Total Power", false, true}
	y, err := yaml.Marshal(DTE)
	if err != nil {
		log.Print(err.Error())
	}
	fmt.Println(string(y))

	// Unmarshal the YAML back into an Entity struct.
	var DelawareTotalPower Entity
	err = yaml.Unmarshal(y, &DelawareTotalPower)
	if err != nil {
		log.Print(err.Error())
	}
	fmt.Println(DelawareTotalPower)
	return "test complete"
}

/*
#REF_TO_ENTITY: {
	knownAs: *"" | string
}
#REF_TO_INITIATIVE: {
	knownAs: *"" | string
}
#REF_TO_PLACE: {
	knownAs: *"" | string
}
#REF_TO_PROJECT: {
	knownAs: *"" | string
}
#REF_TO_REGULATION: {
	knownAs: *"" | string
}
#REF_TO_SERVICE: {
	knownAs: *"" | string
}
*/
