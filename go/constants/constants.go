// Package constants package that contains all the enums or constants that can be used in the headers of the proto.
// (For message headers only)
// WARN: do not remove any constant to avoid to break compatibility. Just mark them as deprecated.
//
//go:generate stringer -type=SourceType
package constants

import "fmt"

type SourceType int8

const (
	SourceTypeUnspecified     SourceType = 0 // No source type defined
	SourceTypeDevice          SourceType = 1 // The source of the event was a device (e.g. PMU)
	SourceTypeService         SourceType = 2 // The source of the event was a service (e.g. state estimator)
	SourceTypeExternalService SourceType = 3 // The source of the event was a service external to SynchroGuard platform (e.g. SCADA)
	SourceTypeTestService     SourceType = 4 // The source of the event was a service in test mode.
	_sourceTypeTag            SourceType = iota
)

func NewSourceType(raw int8) (SourceType, error) {
	if raw >= int8(_sourceTypeTag) {
		return SourceTypeUnspecified, fmt.Errorf("invalid source type %d", raw)
	}

	return SourceType(raw), nil
}

func (st SourceType) ToInt8() int8 {
	return int8(st)
}
