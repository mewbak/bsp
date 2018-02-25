package lumps

import (
	datatypes "github.com/galaco/bsp/lumps/datatypes/plane"
	"encoding/binary"
	"bytes"
	"log"
	"unsafe"
)

/**
	Lump 1: Planes
 */
type Planes struct {
	LumpInfo
	data []datatypes.Plane // MAP_MAX_PLANES = 65536
}

func (lump Planes) FromBytes(raw []byte, length int32) ILump {
	lump.data = make([]datatypes.Plane, length/int32(unsafe.Sizeof(datatypes.Plane{})))
	err := binary.Read(bytes.NewBuffer(raw[:]), binary.LittleEndian, &lump.data)
	if err != nil {
		log.Fatal(err)
	}
	lump.LumpInfo.SetLength(length)

	return lump
}

func (lump Planes) GetData() interface{} {
	return lump.data
}

func (lump Planes) ToBytes() []byte {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, lump.data)
	return buf.Bytes()
}
