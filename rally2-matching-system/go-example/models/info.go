package models

// Info Basic information describing the algorithm.
//
// swagger:model Info
type Info struct {

	// Name of algorithm
	AlgorithmName string

	// A string enum describing the type of biometric images the algorithm is meant to process
	// Enum: [Face Finger Iris]
	AlgorithmType string

	// Algorithm version identifier
	AlgorithmVersion string

	// Name of the Company which produces the algorithm
	CompanyName string

	// The recommended allocation of CPUs for the deployed docker container.
	RecommendedCpus int64

	// The recommended allocation of memory (MB) for the deployed docker container.
	RecommendedMem int64

	// The email address of an engineer or other technical resource to contact in the event of an error running your service. This field may be left blank if desired.
	TechnicalContactEmail string
}
