package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Country struct {
	ent.Schema
}

func (Country) Fields() []ent.Field {
	return []ent.Field{
		field.String("code"),
		field.String("name"),
	}
}

type Region struct {
	ent.Schema
}

type State struct {
	ent.Schema
}

func (State) Fields() []ent.Field {
	return []ent.Field{
		field.String("code"),
		field.String("name"),
	}
}

func (State) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cities", City.Type),
		edge.To("airports", Airport.Type),
	}
}

func (State) Mixins() []ent.Mixin {
	return []ent.Mixin{
		XidMixin{},
		TimeMixin{},
	}
}

type City struct {
	ent.Schema
}

func (City) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

func (City) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("state", State.Type).
			Ref("cities").
			Unique(),
		edge.To("airports", Airport.Type),
	}
}

func (City) Mixins() []ent.Mixin {
	return []ent.Mixin{
		XidMixin{},
		TimeMixin{},
	}
}

func (Region) Fields() []ent.Field {
	return []ent.Field{
		field.String("code"),
		field.String("name"),
	}
}

func (Region) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("airports", Airport.Type),
	}
}

func (Region) Mixins() []ent.Mixin {
	return []ent.Mixin{
		XidMixin{},
		TimeMixin{},
	}
}

type Airport struct {
	ent.Schema
}

func (Airport) Fields() []ent.Field {
	return []ent.Field{
		field.String("site_id"),
		field.Enum("facility_type").
			NamedValues(
				"Airport", "AIRPORT",
				"Balloonport", "BALLOONPORT",
				"SeaplaneBase", "SEAPLANE_BASE",
				"Gliderport", "GLIDERPORT",
				"Heliport", "HELIPORT",
				"Ultralight", "ULTRALIGHT",
			),
		field.String("airport_id").
			Comment("FAA Airport Identifier, 3-4 characters (e.g. ORD, AL03)"),
		field.String("ado_code").
			Comment("FAA District Office Code (e.g. DCA, TEC, MEM)"),
		field.String("name").
			Comment("Common name for airport (e.g. Gifford Fld)"),
		field.Enum("ownership_type").
			NamedValues(
				"Public", "PU",
				"Private", "PR",
				"MilitaryArmy", "MR",
				"MilitaryAirForce", "MA",
				"MilitaryNavy", "MN",
				"MilitaryCoastGuard", "CG",
			),
		field.Enum("usage").
			NamedValues(
				"Public", "PU",
				"Private", "PR",
			),
		field.Float32("latitude"),
		field.Float32("longitude"),
		field.Enum("location_survey_method").
			NamedValues(
				"Estimated", "E",
				"Surveyed", "S",
			),
		field.Float32("elevation"),
		field.Enum("elevation_survey_method").
			NamedValues(
				"Estimated", "E",
				"Surveyed", "S",
			),
		field.Int8("magnetic_variance"),
		field.Int8("magnetic_variance_year").
			Comment("Magnetic variation epoch year"),
		field.Int8("traffic_pattern_altitude"),
		field.String("chart_name"),
		field.Uint8("distance_to_city"),
		field.Enum("direction_to_city").
			NamedValues(
				"North", "N",
				"NorthNortheast", "NNE",
				"Northeast", "NE",
				"EastNortheast", "ENE",
				"East", "E",
				"EastSouthest", "ESE",
				"Southeast", "SE",
				"SouthSoutheast", "SSE",
				"South", "S",
				"SouthSouthwest", "SSW",
				"Southwest", "SW",
				"WestSouthwest", "WSW",
				"West", "W",
				"WestNorthwest", "WNW",
				"Northwest", "NW",
				"NorthNothwest", "NNW",
			),
		field.Uint8("acreage"),
		field.String("artcc_id"),
		field.String("artcc_name"),
		field.String("computer_id"),
		field.Bool("fss_on_airport").
			Comment("Tie-In FSS Physically Located on Facility"),
		field.String("fss_id"), //TODO: breakout to edge
		field.String("fss_name"),
		field.String("phone_number"),
		field.String("toll_free_number"), // TODO: data quality?
		field.String("alternate_fss_id"), // TODO: breakout to edge
		field.String("alternate_fss_name"),
		field.String("alternate_toll_free_number"),
		field.String("notam_id"),
		field.Bool("notam_flag"). //TODO: better name?
						Comment("Availability of NOTAM 'D' Service at Airport"),
		field.Time("activation_date"), //TODO: only have YYYY/MM, better type?
		field.Enum("status").
			NamedValues(
				"Operational", "O",
				"ClosedIndefinitely", "CI",
				"ClosedPermanently", "CP",
			),
		field.Enum("far_139_class").
			Values(
				"I",
				"II",
				"III",
				"IV",
			).
			Comment("FAR 139 Airport Class Code"), //TODO: byte or uint8?
		field.Enum("arff_index").
			Values(
				"A",
				"B",
				"C",
				"D",
				"E",
				"L",
			).
			Comment("Airport ARFF Certification Type Code"),
		field.Time("arff_certification_date").
			Comment("Airport ARFF Certificiation Date"),
		field.Enum("far_139_carrier_service_code").
			NamedValues(
				"Scheduled", "S",
				"Unscheduled", "U",
			),
		field.String("asp_code").
			Comment("Combination of 1 to 7 codes that indiciate type of federal agreements existing at the airport\n" +
				"N - National Plan of Integrated Airport Systems\n" +
				"B - Installation of Navigational Facilities on Privately Owned Airports under F&E Program\n" +
				"G - Grant Agreements under FAAP/ADAP/AIP\n" +
				"H - Complaiance with Accessibility to the Handicapped\n" +
				"P - Surplus Property Agreement Under Public Law 289\n" +
				"R - Surplus Property Agreement Under REgulation 16-WAA\n" +
				"S - Conveyance Under Section 16, Federal Airport Act of 1946 or Section 23, Airport and Airway Development Act of 1970\n" +
				"V - Advance Planning Agreement Under FAAP\n" +
				"X - Obligations Assumed By Transfer\n" +
				"Y - Assurances Pursuant to Title VI, Civil Rights Act  of 1964\n" +
				"Z - Conveyance Under Section 303(C), Federal Aviation Act of 1958\n" +
				"1 - Grant Agreement Has Expired; however, agreement remains in effect for this facility as long as it is public use\n" +
				"2 - Section 303(C) Authority from FAA Act of 1958 has expired; however, agreement remains in effect for this facility as long as itt is public use\n" +
				"3 - AP-4 Agreement under DLAND or DCLA has expired\n" +
				"NONE - No Grant Agreement Exists\n" +
				"BLANK - No Grant Agreement Exists"), //TODO: make this sane...
		field.Enum("airspace_analysis_determination").
			NamedValues(
				"Conditional", "CONDITIONAL",
				"NotAnalyzed", "NOT_ANALYZED",
				"NoObjection", "NO_OBJECTION",
				"Objectionable", "OBJECTIONABLE",
			).
			Comment("Airport Airspace Analysis Determination"),
		field.Bool("customs_airport_of_entry"),
		field.Bool("customs_landing_rights").
			Comment("A landing rights airport is any airport, other than an international airport or user fee airport, at which flights from a foreign area are given permissions by Customs to land."),
		field.Bool("joint_use").
			Comment("Facility has Military/Civil Join Use Agreement that allows Civil Operations at a Military Airport."),
		field.Bool("military_landing_rights").
			Comment("mil_lndg_flag; Airport has entered into an Agreement that Grants Landing Rights to the Military"),
		field.Enum("inspection_method").
			NamedValues(
				"Federal", "F",
				"State", "S",
				"Contractor", "C",
				"PublicUseMailoutProgram", "1",
				"PrivateUseMailoutProgram", "2",
			).
			Comment("Airport Inspection Method"),
		field.Enum("inspector_code").
			NamedValues(
				"FaaAirportFieldPersonnel", "F",
				"StateAeronauticalPersonnel", "S",
				"PrivateContractPersonnel", "C",
				"Owner", "N",
			).
			Comment("Agency/Group Performing Physical Inspection"),
		field.Time("last_inspection").
			Comment("Last Physical Inspection Date"),
		field.Time("last_info_response").
			Comment("Last Date Information Request was completed by Facility Owner or Manager"),
		field.String("fuel_types").
			Comment("Fuel Types available for public use at the Airport.\n" +
				"Comma-separated list of values (or blank)\n" +
				"100,100LL,A,A+,A++,A++10,A1,A1+,J5,J8,J8+10,J,MOGAS,UL91,UL94,UL100"),
		field.Enum("airframe_repair_service").
			NamedValues(
				"Major", "MAJOR",
				"Minor", "MINOR",
				"None", "NONE",
			),
		field.Enum("power_plant_repair_service").
			NamedValues(
				"Major", "MAJOR",
				"Minor", "MINOR",
				"None", "NONE",
			),
		field.Enum("bottled_oxygen_type").
			NamedValues(
				"High", "HIGH",
				"Low", "LOW",
				"HighLow", "HIGH_LOW",
				"None", "NONE",
			),
		field.Enum("bulk_oxygen_type").
			NamedValues(
				"High", "HIGH",
				"Low", "LOW",
				"HighLow", "HIGH_LOW",
				"None", "NONE",
			),
		field.Enum("lighting_schedule").
			NamedValues(
				"SeeRemark", "SEE_RMK",
				"SunsetSunrise", "SUNSET_SUNRISE",
			).
			Comment("Airport Lighting Schedule; blank, SS-SR (sunset to sunrise), or 'SEE RMK' (details are in facility remark data)"),
		field.Enum("beacon_lighting_schedule").
			NamedValues(
				"SeeRemark", "SEE_RMK",
				"SunsetSunrise", "SUNSET_SUNRISE",
			).
			Comment("Airport Beacon Lighting Schedule; blank, SS-SR (sunset to sunrise), or 'SEE RMK' (details are in facility remark data)"),
		field.Enum("tower_type").
			NamedValues(
				"AirTrafficControlTower", "ATCT",
				"NonAirTrafficControlTower", "NON_ATCT",
				"AirTrafficControlTowerWithApproachControl", "ATCT_AC",
				"AirForceAirTrafficControlTowerWithRadarApproachControl", "ATCT_RAPCON",
				"NavyAirTrafficControlWithRadarApproachControl", "ATCT_RATCF",
				"AirTrafficControlTowerWithTerminalRadarApproachControl", "ATCT_TRACON",
			),
		field.Enum("segmented_circle_marker").
			NamedValues(
				"Yes", "Y",
				"No", "N",
				"YesLighted", "Y_L",
			),
		field.Enum("beacon_lens_color").
			NamedValues(
				"WhiteGreen", "WG", // Lighted Land Airport
				"WhiteYellow", "WY", // Lighted Seaplane Base
				"WhiteGreenYellow", "WGY", // Heliport
				"SplitWhiteGreen", "SWG", // Lighted Military Airport
				"White", "W", // Unlighted Land Airport
				"Yellow", "Y", // Unlighted Seaplane Base
				"Green", "G", // Lighted Land Airport
				"None", "N",
			),
		field.Bool("landing_fee").
			Comment("Landing Fee charged to Non-Commercial Users of Airport"),
		field.Bool("medical_use").
			Comment("Landing Facility is used for Medical Purposes"),
		field.Uint("based_single_engine").
			Comment("Number of single engine general aviation aircraft based at airport"),
		field.Uint("based_multi_engine_aircraft"),
		field.Uint("based_jet_engine_aircraft"),
		field.Uint("based_helicopters"),
		field.Uint("based_gliders"),
		field.Uint("based_military_aircraft"),
		field.Uint("based_ultralight_aircraft"),
		field.Uint("commercial_operations").
			Comment("Commercial Services Count"),
		field.Uint("commuter_operations"),
		field.Uint("air_taxi_operations"),
		field.Uint("general_aviation_local_operations").
			Comment("General Aviation local operations (within traffic pattern or within 20-mile radius of airport"),
		field.Uint("general_aviation_itinerant_operations").
			Comment("General Aviation Itinerant Operations (GA ops excluding commuter, air taxi, and non-local)"),
		field.Uint("military_operations"),
		field.Time("annual_operations_date").
			Comment("12-month ending date on which annual operations data in above six fields is based"),
		field.String("airport_position_source"),
		field.Time("airport_position_source_date"),
		field.String("airport_elevation_source"),
		field.Time("airport_elevation_source_date"),
		field.Bool("contract_fuel").
			Comment("Contract Fuel Available; useless data, only 3 airports have this field set"),
		field.Bool("transient_buoy_storage").
			Comment("Buoy Transient Storage Facilities; per the FAA w/o further context"),
		field.Bool("transient_hangar_storage"),
		field.Bool("transient_tie_down_storage"),
		field.String("other_services").
			Comment("Comma-delimited list of other airport services; " +
				"AFRT (Air Freight), " +
				"AGRI (Crop Dusting), " +
				"AMB (Air Ambulance), " +
				"AVNCS (Avionics), " +
				"BCHGR (Beaching Gear), " +
				"CARGO (Cargo Handling), " +
				"CHTR (Charter), " +
				"GLD (Glider), " +
				"INSTR (Pilot Instruction), " +
				"PAJA (Parachute Jump Activity), " +
				"RNTL (Aircraft Rental), " +
				"SALES (Aircraft Sales), " +
				"SURV (Annual Surveying), " +
				"TOW (Glider Towing Services)"),
		field.Enum("wind_indicator_flag").
			NamedValues(
				"No", "NO_WIND_INDICATOR",
				"Unlighted", "UNLIGHTED",
				"Lighted", "LIGHTED",
			),
		field.String("icao_id"),
		field.Bool("minimum_operational_network").
			Comment("A MON airport will ensure that a pilot will always be within 100nm of an airport with an instrument appr not dependent on GPS"),
		field.Bool("customs_user_fee").
			Comment("US CUSTOMS USER FEE ARPT; per FAA, without further context"),
		field.Int("altitude_correction_temperature").
			Comment("CTA, Cold Temperature Airport; Altitude Correction Required At or Below this Temperature Given in Celsius"),
	}
}

func (Airport) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("state", State.Type).
			Ref("airports").
			Unique(),
		edge.From("city", City.Type).
			Ref("airports").
			Unique(),
		edge.From("region", Region.Type).
			Ref("airports").
			Unique(),
	}
}

func (Airport) Mixins() []ent.Mixin {
	return []ent.Mixin{
		XidMixin{},
		TimeMixin{},
	}
}

func (Airport) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.MultiOrder(),
		entgql.QueryField(),
	}
}
