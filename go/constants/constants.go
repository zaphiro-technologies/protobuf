// Package constants package that contains all the enums or constants that can be used in the headers of the proto.
// (For message headers only)
// WARN: do not remove/rename any constant to avoid to break compatibility. Just mark them as deprecated.
// Also all the string must be in PascalCase.
package constants

import "fmt"

type SourceType string

//nolint:lll
const (
	SourceTypeUnspecified     SourceType = "Unspecified"     // No source type defined
	SourceTypeDevice          SourceType = "Device"          // The source of the event was a device (e.g. PMU)
	SourceTypeService         SourceType = "Service"         // The source of the event was a service (e.g. state estimator)
	SourceTypeExternalService SourceType = "ExternalService" // The source of the event was a service external to SynchroGuard platform (e.g. SCADA)
	SourceTypeTestService     SourceType = "TestService"     // The source of the event was a service in test mode.
)

func (st SourceType) String() string {
	return string(st)
}

func NewSourceType(s string) (SourceType, error) {
	switch SourceType(s) {
	case SourceTypeUnspecified:
		return SourceTypeUnspecified, nil
	case SourceTypeDevice:
		return SourceTypeDevice, nil
	case SourceTypeService:
		return SourceTypeService, nil
	case SourceTypeExternalService:
		return SourceTypeExternalService, nil
	case SourceTypeTestService:
		return SourceTypeTestService, nil
	default:
		return SourceTypeUnspecified, fmt.Errorf("unexpected SourceType %q", s)
	}
}
