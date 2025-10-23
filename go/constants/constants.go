// Package constants package that contains all the enums or constants that can be used in the headers of the proto.
// (For message headers only)
// WARN: do not remove any constant to avoid to break compatibility. Just mark them as deprecated.
//
//go:generate stringer -type=SourceType
package constants

type SourceType int

const (
	SourceTypeUnspecified     SourceType = 0 // No source type defined
	SourceTypeDevice          SourceType = 1 // The source of the event was a device (e.g. PMU)
	SourceTypeService         SourceType = 2 // The source of the event was a service (e.g. state estimator)
	SourceTypeExternalService SourceType = 3 // The source of the event was a service external to SynchroGuard platform (e.g. SCADA)
	SourceTypeTestService     SourceType = 4 // The source of the event was a service in test mode.
)
