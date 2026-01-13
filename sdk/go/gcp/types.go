package gcp

// Format constants for Artifact Registry
const (
	FormatDocker   = "DOCKER"
	FormatMaven    = "MAVEN"
	FormatNpm      = "NPM"
	FormatPython   = "PYTHON"
	FormatApt      = "APT"
	FormatYum      = "YUM"
	FormatKubeflow = "KUBEFLOW"
	FormatGo       = "GO"
	FormatGeneric  = "GENERIC"
)

// Locations constants
const (
	LocationEU   = "EU"
	LocationUS   = "US"
	LocationASIA = "ASIA"
)

const (
	StorageSSD = "SSD"
	StorageHDD = "HDD"
)

// Regions constants
const (
	// North America
	LocationNorthAmericaNortheast1 = "northamerica-northeast1"
	LocationNorthAmericaNortheast2 = "northamerica-northeast2"
	LocationNorthAmericaSouth1     = "northamerica-south1"
	LocationUsCentral1             = "us-central1"
	LocationUsEast1                = "us-east1"
	LocationUsEast4                = "us-east4"
	LocationUsEast5                = "us-east5"
	LocationUsSouth1               = "us-south1"
	LocationUsWest1                = "us-west1"
	LocationUsWest2                = "us-west2"
	LocationUsWest3                = "us-west3"
	LocationUsWest4                = "us-west4"

	// South America
	LocationSouthAmericaEast1 = "southamerica-east1"
	LocationSouthAmericaWest1 = "southamerica-west1"

	// Europe
	LocationEuropeCentral2   = "europe-central2"
	LocationEuropeNorth1     = "europe-north1"
	LocationEuropeNorth2     = "europe-north2"
	LocationEuropeSouthwest1 = "europe-southwest1"
	LocationEuropeWest1      = "europe-west1"
	LocationEuropeWest2      = "europe-west2"
	LocationEuropeWest3      = "europe-west3"
	LocationEuropeWest4      = "europe-west4"
	LocationEuropeWest6      = "europe-west6"
	LocationEuropeWest8      = "europe-west8"
	LocationEuropeWest9      = "europe-west9"
	LocationEuropeWest10     = "europe-west10"
	LocationEuropeWest12     = "europe-west12"

	// Middle East
	LocationMeCentral1 = "me-central1"
	LocationMeCentral2 = "me-central2"
	LocationMeWest1    = "me-west1"

	// Asia
	LocationAsiaEast1      = "asia-east1"
	LocationAsiaEast2      = "asia-east2"
	LocationAsiaNortheast1 = "asia-northeast1"
	LocationAsiaNortheast2 = "asia-northeast2"
	LocationAsiaNortheast3 = "asia-northeast3"
	LocationAsiaSouth1     = "asia-south1"
	LocationAsiaSouth2     = "asia-south2"
	LocationAsiaSoutheast1 = "asia-southeast1"
	LocationAsiaSoutheast2 = "asia-southeast2"

	// Australia
	LocationAustraliaSoutheast1 = "australia-southeast1"
	LocationAustraliaSoutheast2 = "australia-southeast2"

	// Africa
	LocationAfricaSouth1 = "africa-south1"
)

const (
	// Africa
	AfricaSouth1A = "africa-south1-a"
	AfricaSouth1B = "africa-south1-b"
	AfricaSouth1C = "africa-south1-c"

	// Asia
	AsiaEast1A      = "asia-east1-a"
	AsiaEast1B      = "asia-east1-b"
	AsiaEast1C      = "asia-east1-c"
	AsiaEast2A      = "asia-east2-a"
	AsiaEast2B      = "asia-east2-b"
	AsiaEast2C      = "asia-east2-c"
	AsiaNortheast1A = "asia-northeast1-a"
	AsiaNortheast1B = "asia-northeast1-b"
	AsiaNortheast1C = "asia-northeast1-c"
	AsiaNortheast2A = "asia-northeast2-a"
	AsiaNortheast2B = "asia-northeast2-b"
	AsiaNortheast2C = "asia-northeast2-c"
	AsiaNortheast3A = "asia-northeast3-a"
	AsiaNortheast3B = "asia-northeast3-b"
	AsiaNortheast3C = "asia-northeast3-c"
	AsiaSouth1A     = "asia-south1-a"
	AsiaSouth1B     = "asia-south1-b"
	AsiaSouth1C     = "asia-south1-c"
	AsiaSouth2A     = "asia-south2-a"
	AsiaSouth2B     = "asia-south2-b"
	AsiaSouth2C     = "asia-south2-c"
	AsiaSoutheast1A = "asia-southeast1-a"
	AsiaSoutheast1B = "asia-southeast1-b"
	AsiaSoutheast1C = "asia-southeast1-c"
	AsiaSoutheast2A = "asia-southeast2-a"
	AsiaSoutheast2B = "asia-southeast2-b"
	AsiaSoutheast2C = "asia-southeast2-c"

	// Australia
	AustraliaSoutheast1A = "australia-southeast1-a"
	AustraliaSoutheast1B = "australia-southeast1-b"
	AustraliaSoutheast1C = "australia-southeast1-c"
	AustraliaSoutheast2A = "australia-southeast2-a"
	AustraliaSoutheast2B = "australia-southeast2-b"
	AustraliaSoutheast2C = "australia-southeast2-c"

	// Europe
	EuropeCentral2A   = "europe-central2-a"
	EuropeCentral2B   = "europe-central2-b"
	EuropeCentral2C   = "europe-central2-c"
	EuropeNorth1A     = "europe-north1-a"
	EuropeNorth1B     = "europe-north1-b"
	EuropeNorth1C     = "europe-north1-c"
	EuropeNorth2A     = "europe-north2-a"
	EuropeNorth2B     = "europe-north2-b"
	EuropeNorth2C     = "europe-north2-c"
	EuropeSouthwest1A = "europe-southwest1-a"
	EuropeSouthwest1B = "europe-southwest1-b"
	EuropeSouthwest1C = "europe-southwest1-c"
	EuropeWest1B      = "europe-west1-b"
	EuropeWest1C      = "europe-west1-c"
	EuropeWest1D      = "europe-west1-d"
	EuropeWest2A      = "europe-west2-a"
	EuropeWest2B      = "europe-west2-b"
	EuropeWest2C      = "europe-west2-c"
	EuropeWest3A      = "europe-west3-a"
	EuropeWest3B      = "europe-west3-b"
	EuropeWest3C      = "europe-west3-c"
	EuropeWest4A      = "europe-west4-a"
	EuropeWest4B      = "europe-west4-b"
	EuropeWest4C      = "europe-west4-c"
	EuropeWest6A      = "europe-west6-a"
	EuropeWest6B      = "europe-west6-b"
	EuropeWest6C      = "europe-west6-c"
	EuropeWest8A      = "europe-west8-a"
	EuropeWest8B      = "europe-west8-b"
	EuropeWest8C      = "europe-west8-c"
	EuropeWest9A      = "europe-west9-a"
	EuropeWest9B      = "europe-west9-b"
	EuropeWest9C      = "europe-west9-c"
	EuropeWest10A     = "europe-west10-a"
	EuropeWest10B     = "europe-west10-b"
	EuropeWest10C     = "europe-west10-c"
	EuropeWest12A     = "europe-west12-a"
	EuropeWest12B     = "europe-west12-b"
	EuropeWest12C     = "europe-west12-c"

	// Middle East
	MeCentral1A = "me-central1-a"
	MeCentral1B = "me-central1-b"
	MeCentral1C = "me-central1-c"
	MeCentral2A = "me-central2-a"
	MeCentral2B = "me-central2-b"
	MeCentral2C = "me-central2-c"
	MeWest1A    = "me-west1-a"
	MeWest1B    = "me-west1-b"
	MeWest1C    = "me-west1-c"

	// North America
	NorthAmericaNortheast1A = "northamerica-northeast1-a"
	NorthAmericaNortheast1B = "northamerica-northeast1-b"
	NorthAmericaNortheast1C = "northamerica-northeast1-c"
	NorthAmericaNortheast2A = "northamerica-northeast2-a"
	NorthAmericaNortheast2B = "northamerica-northeast2-b"
	NorthAmericaNortheast2C = "northamerica-northeast2-c"
	NorthAmericaSouth1A     = "northamerica-south1-a"
	NorthAmericaSouth1B     = "northamerica-south1-b"
	NorthAmericaSouth1C     = "northamerica-south1-c"
	UsCentral1A             = "us-central1-a"
	UsCentral1B             = "us-central1-b"
	UsCentral1C             = "us-central1-c"
	UsCentral1F             = "us-central1-f"
	UsEast1B                = "us-east1-b"
	UsEast1C                = "us-east1-c"
	UsEast1D                = "us-east1-d"
	UsEast4A                = "us-east4-a"
	UsEast4B                = "us-east4-b"
	UsEast4C                = "us-east4-c"
	UsEast5A                = "us-east5-a"
	UsEast5B                = "us-east5-b"
	UsEast5C                = "us-east5-c"
	UsSouth1A               = "us-south1-a"
	UsSouth1B               = "us-south1-b"
	UsSouth1C               = "us-south1-c"
	UsWest1A                = "us-west1-a"
	UsWest1B                = "us-west1-b"
	UsWest1C                = "us-west1-c"
	UsWest2A                = "us-west2-a"
	UsWest2B                = "us-west2-b"
	UsWest2C                = "us-west2-c"
	UsWest3A                = "us-west3-a"
	UsWest3B                = "us-west3-b"
	UsWest3C                = "us-west3-c"
	UsWest4A                = "us-west4-a"
	UsWest4B                = "us-west4-b"
	UsWest4C                = "us-west4-c"

	// South America
	SouthAmericaEast1A = "southamerica-east1-a"
	SouthAmericaEast1B = "southamerica-east1-b"
	SouthAmericaEast1C = "southamerica-east1-c"
	SouthAmericaWest1A = "southamerica-west1-a"
	SouthAmericaWest1B = "southamerica-west1-b"
	SouthAmericaWest1C = "southamerica-west1-c"
)

const (
	StandardIsolationPriorityLow    = "PRIORITY_LOW"
	StandardIsolationPriorityMedium = "PRIORITY_MEDIUM"
	StandardIsolationPriorityHigh   = "PRIORITY_HIGH"
)

// Artifact Registry
var ArtifactFormats = struct {
	Docker   string
	Maven    string
	Npm      string
	Python   string
	Apt      string
	Yum      string
	Kubeflow string
	Go       string
	Generic  string
}{
	Docker:   "DOCKER",
	Maven:    "MAVEN",
	Npm:      "NPM",
	Python:   "PYTHON",
	Apt:      "APT",
	Yum:      "YUM",
	Kubeflow: "KUBEFLOW",
	Go:       "GO",
	Generic:  "GENERIC",
}

// Storage
var StorageClasses = struct {
	SSD string
	HDD string
}{
	SSD: "SSD",
	HDD: "HDD",
}

// Regions
var Regions = struct {
	NorthAmerica struct {
		Northeast1 string
		Northeast2 string
		South1     string
		UsCentral1 string
		UsEast1    string
		UsEast4    string
		UsEast5    string
		UsSouth1   string
		UsWest1    string
		UsWest2    string
		UsWest3    string
		UsWest4    string
	}
	SouthAmerica struct {
		East1 string
		West1 string
	}
	Europe struct {
		Central2   string
		North1     string
		North2     string
		Southwest1 string
		West1      string
		West2      string
		West3      string
		West4      string
		West6      string
		West8      string
		West9      string
		West10     string
		West12     string
	}
	MiddleEast struct {
		Central1 string
		Central2 string
		West1    string
	}
	Asia struct {
		East1      string
		East2      string
		Northeast1 string
		Northeast2 string
		Northeast3 string
		South1     string
		South2     string
		Southeast1 string
		Southeast2 string
	}
	Australia struct {
		Southeast1 string
		Southeast2 string
	}
	Africa struct {
		South1 string
	}
}{
	NorthAmerica: struct {
		Northeast1,
		Northeast2,
		South1,
		UsCentral1,
		UsEast1,
		UsEast4,
		UsEast5,
		UsSouth1,
		UsWest1,
		UsWest2,
		UsWest3,
		UsWest4 string
	}{
		Northeast1: "northamerica-northeast1",
		Northeast2: "northamerica-northeast2",
		South1:     "northamerica-south1",
		UsCentral1: "us-central1",
		UsEast1:    "us-east1",
		UsEast4:    "us-east4",
		UsEast5:    "us-east5",
		UsSouth1:   "us-south1",
		UsWest1:    "us-west1",
		UsWest2:    "us-west2",
		UsWest3:    "us-west3",
		UsWest4:    "us-west4",
	},
	SouthAmerica: struct{ East1, West1 string }{
		East1: "southamerica-east1",
		West1: "southamerica-west1",
	},
	Europe: struct {
		Central2,
		North1,
		North2,
		Southwest1,
		West1,
		West2,
		West3,
		West4,
		West6,
		West8,
		West9,
		West10,
		West12 string
	}{
		Central2:   "europe-central2",
		North1:     "europe-north1",
		North2:     "europe-north2",
		Southwest1: "europe-southwest1",
		West1:      "europe-west1",
		West2:      "europe-west2",
		West3:      "europe-west3",
		West4:      "europe-west4",
		West6:      "europe-west6",
		West8:      "europe-west8",
		West9:      "europe-west9",
		West10:     "europe-west10",
		West12:     "europe-west12",
	},
	MiddleEast: struct {
		Central1,
		Central2,
		West1 string
	}{
		Central1: "me-central1",
		Central2: "me-central2",
		West1:    "me-west1",
	},
	Asia: struct {
		East1,
		East2,
		Northeast1,
		Northeast2,
		Northeast3,
		South1,
		South2,
		Southeast1,
		Southeast2 string
	}{
		East1:      "asia-east1",
		East2:      "asia-east2",
		Northeast1: "asia-northeast1",
		Northeast2: "asia-northeast2",
		Northeast3: "asia-northeast3",
		South1:     "asia-south1",
		South2:     "asia-south2",
		Southeast1: "asia-southeast1",
		Southeast2: "asia-southeast2",
	},
	Australia: struct {
		Southeast1,
		Southeast2 string
	}{
		Southeast1: "australia-southeast1",
		Southeast2: "australia-southeast2",
	},
	Africa: struct {
		South1 string
	}{
		South1: "africa-south1",
	},
}

// Zones
var Zones = struct {
	NorthAmerica struct {
		Northeast1A, Northeast1B, Northeast1C              string
		Northeast2A, Northeast2B, Northeast2C              string
		South1A, South1B, South1C                          string
		UsCentral1A, UsCentral1B, UsCentral1C, UsCentral1F string
		UsEast1B, UsEast1C, UsEast1D                       string // Note: East1 não tem zona A
		UsEast4A, UsEast4B, UsEast4C                       string
		UsEast5A, UsEast5B, UsEast5C                       string
		UsSouth1A, UsSouth1B, UsSouth1C                    string
		UsWest1A, UsWest1B, UsWest1C                       string
		UsWest2A, UsWest2B, UsWest2C                       string
		UsWest3A, UsWest3B, UsWest3C                       string
		UsWest4A, UsWest4B, UsWest4C                       string
	}
	SouthAmerica struct {
		East1A, East1B, East1C string
		West1A, West1B, West1C string
	}
	Europe struct {
		Central2A, Central2B, Central2C       string
		North1A, North1B, North1C             string
		North2A, North2B, North2C             string
		Southwest1A, Southwest1B, Southwest1C string
		West1B, West1C, West1D                string
		West2A, West2B, West2C                string
		West3A, West3B, West3C                string
		West4A, West4B, West4C                string
		West6A, West6B, West6C                string
		West8A, West8B, West8C                string
		West9A, West9B, West9C                string
		West10A, West10B, West10C             string
		West12A, West12B, West12C             string
	}
	Asia struct {
		East1A, East1B, East1C                string
		East2A, East2B, East2C                string
		Northeast1A, Northeast1B, Northeast1C string
		Northeast2A, Northeast2B, Northeast2C string
		Northeast3A, Northeast3B, Northeast3C string
		South1A, South1B, South1C             string
		South2A, South2B, South2C             string
		Southeast1A, Southeast1B, Southeast1C string
		Southeast2A, Southeast2B, Southeast2C string
	}
	Australia struct {
		Southeast1A, Southeast1B, Southeast1C string
		Southeast2A, Southeast2B, Southeast2C string
	}
	MiddleEast struct {
		Central1A, Central1B, Central1C string
		Central2A, Central2B, Central2C string
		West1A, West1B, West1C          string
	}
	Africa struct {
		South1A, South1B, South1C string
	}
}{
	NorthAmerica: struct {
		Northeast1A, Northeast1B, Northeast1C              string
		Northeast2A, Northeast2B, Northeast2C              string
		South1A, South1B, South1C                          string
		UsCentral1A, UsCentral1B, UsCentral1C, UsCentral1F string
		UsEast1B, UsEast1C, UsEast1D                       string
		UsEast4A, UsEast4B, UsEast4C                       string
		UsEast5A, UsEast5B, UsEast5C                       string
		UsSouth1A, UsSouth1B, UsSouth1C                    string
		UsWest1A, UsWest1B, UsWest1C                       string
		UsWest2A, UsWest2B, UsWest2C                       string
		UsWest3A, UsWest3B, UsWest3C                       string
		UsWest4A, UsWest4B, UsWest4C                       string
	}{
		// Canada
		Northeast1A: "northamerica-northeast1-a",
		Northeast1B: "northamerica-northeast1-b",
		Northeast1C: "northamerica-northeast1-c",
		Northeast2A: "northamerica-northeast2-a",
		Northeast2B: "northamerica-northeast2-b",
		Northeast2C: "northamerica-northeast2-c",
		South1A:     "northamerica-south1-a",
		South1B:     "northamerica-south1-b",
		South1C:     "northamerica-south1-c",
		// US Central
		UsCentral1A: "us-central1-a",
		UsCentral1B: "us-central1-b",
		UsCentral1C: "us-central1-c",
		UsCentral1F: "us-central1-f",
		// US East
		UsEast1B: "us-east1-b",
		UsEast1C: "us-east1-c",
		UsEast1D: "us-east1-d",
		UsEast4A: "us-east4-a",
		UsEast4B: "us-east4-b",
		UsEast4C: "us-east4-c",
		UsEast5A: "us-east5-a",
		UsEast5B: "us-east5-b",
		UsEast5C: "us-east5-c",
		// US South
		UsSouth1A: "us-south1-a",
		UsSouth1B: "us-south1-b",
		UsSouth1C: "us-south1-c",
		// US West
		UsWest1A: "us-west1-a",
		UsWest1B: "us-west1-b",
		UsWest1C: "us-west1-c",
		UsWest2A: "us-west2-a",
		UsWest2B: "us-west2-b",
		UsWest2C: "us-west2-c",
		UsWest3A: "us-west3-a",
		UsWest3B: "us-west3-b",
		UsWest3C: "us-west3-c",
		UsWest4A: "us-west4-a",
		UsWest4B: "us-west4-b",
		UsWest4C: "us-west4-c",
	},
	SouthAmerica: struct {
		East1A, East1B, East1C string
		West1A, West1B, West1C string
	}{
		East1A: "southamerica-east1-a",
		East1B: "southamerica-east1-b",
		East1C: "southamerica-east1-c",
		West1A: "southamerica-west1-a",
		West1B: "southamerica-west1-b",
		West1C: "southamerica-west1-c",
	},
	Europe: struct {
		Central2A, Central2B, Central2C       string
		North1A, North1B, North1C             string
		North2A, North2B, North2C             string
		Southwest1A, Southwest1B, Southwest1C string
		West1B, West1C, West1D                string
		West2A, West2B, West2C                string
		West3A, West3B, West3C                string
		West4A, West4B, West4C                string
		West6A, West6B, West6C                string
		West8A, West8B, West8C                string
		West9A, West9B, West9C                string
		West10A, West10B, West10C             string
		West12A, West12B, West12C             string
	}{
		Central2A:   "europe-central2-a",
		Central2B:   "europe-central2-b",
		Central2C:   "europe-central2-c",
		North1A:     "europe-north1-a",
		North1B:     "europe-north1-b",
		North1C:     "europe-north1-c",
		North2A:     "europe-north2-a",
		North2B:     "europe-north2-b",
		North2C:     "europe-north2-c",
		Southwest1A: "europe-southwest1-a",
		Southwest1B: "europe-southwest1-b",
		Southwest1C: "europe-southwest1-c",
		West1B:      "europe-west1-b",
		West1C:      "europe-west1-c",
		West1D:      "europe-west1-d",
		West2A:      "europe-west2-a",
		West2B:      "europe-west2-b",
		West2C:      "europe-west2-c",
		West3A:      "europe-west3-a",
		West3B:      "europe-west3-b",
		West3C:      "europe-west3-c",
		West4A:      "europe-west4-a",
		West4B:      "europe-west4-b",
		West4C:      "europe-west4-c",
		West6A:      "europe-west6-a",
		West6B:      "europe-west6-b",
		West6C:      "europe-west6-c",
		West8A:      "europe-west8-a",
		West8B:      "europe-west8-b",
		West8C:      "europe-west8-c",
		West9A:      "europe-west9-a",
		West9B:      "europe-west9-b",
		West9C:      "europe-west9-c",
		West10A:     "europe-west10-a",
		West10B:     "europe-west10-b",
		West10C:     "europe-west10-c",
		West12A:     "europe-west12-a",
		West12B:     "europe-west12-b",
		West12C:     "europe-west12-c",
	},
	Asia: struct {
		East1A, East1B, East1C                string
		East2A, East2B, East2C                string
		Northeast1A, Northeast1B, Northeast1C string
		Northeast2A, Northeast2B, Northeast2C string
		Northeast3A, Northeast3B, Northeast3C string
		South1A, South1B, South1C             string
		South2A, South2B, South2C             string
		Southeast1A, Southeast1B, Southeast1C string
		Southeast2A, Southeast2B, Southeast2C string
	}{
		East1A:      "asia-east1-a",
		East1B:      "asia-east1-b",
		East1C:      "asia-east1-c",
		East2A:      "asia-east2-a",
		East2B:      "asia-east2-b",
		East2C:      "asia-east2-c",
		Northeast1A: "asia-northeast1-a",
		Northeast1B: "asia-northeast1-b",
		Northeast1C: "asia-northeast1-c",
		Northeast2A: "asia-northeast2-a",
		Northeast2B: "asia-northeast2-b",
		Northeast2C: "asia-northeast2-c",
		Northeast3A: "asia-northeast3-a",
		Northeast3B: "asia-northeast3-b",
		Northeast3C: "asia-northeast3-c",
		South1A:     "asia-south1-a",
		South1B:     "asia-south1-b",
		South1C:     "asia-south1-c",
		South2A:     "asia-south2-a",
		South2B:     "asia-south2-b",
		South2C:     "asia-south2-c",
		Southeast1A: "asia-southeast1-a",
		Southeast1B: "asia-southeast1-b",
		Southeast1C: "asia-southeast1-c",
		Southeast2A: "asia-southeast2-a",
		Southeast2B: "asia-southeast2-b",
		Southeast2C: "asia-southeast2-c",
	},
	Australia: struct {
		Southeast1A, Southeast1B, Southeast1C string
		Southeast2A, Southeast2B, Southeast2C string
	}{
		Southeast1A: "australia-southeast1-a",
		Southeast1B: "australia-southeast1-b",
		Southeast1C: "australia-southeast1-c",
		Southeast2A: "australia-southeast2-a",
		Southeast2B: "australia-southeast2-b",
		Southeast2C: "australia-southeast2-c",
	},
	MiddleEast: struct {
		Central1A, Central1B, Central1C string
		Central2A, Central2B, Central2C string
		West1A, West1B, West1C          string
	}{
		Central1A: "me-central1-a",
		Central1B: "me-central1-b",
		Central1C: "me-central1-c",
		Central2A: "me-central2-a",
		Central2B: "me-central2-b",
		Central2C: "me-central2-c",
		West1A:    "me-west1-a",
		West1B:    "me-west1-b",
		West1C:    "me-west1-c",
	},
	Africa: struct {
		South1A, South1B, South1C string
	}{
		South1A: "africa-south1-a",
		South1B: "africa-south1-b",
		South1C: "africa-south1-c",
	},
}
