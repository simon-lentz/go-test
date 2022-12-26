package types

type EDGES struct {
	BUSINESS_RELATIONSHIP_WITH_Entity struct {
		businessRelationshipType string
		WHERE                    string //"#REF_TO_ENTITY", will use KnownAs as common value
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
