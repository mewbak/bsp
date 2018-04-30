package plane

import "github.com/galaco/bsp/primitives/common"

type Plane struct {
	Normal common.Vector // normal vector
	Distance float32  // distance from origin
	AxisType int32	  // plane axis identifier
}
