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
	StandardIsolationPriorityLow = "PRIORITY_LOW"
)
