package lumps

import (
	"bytes"
	"encoding/binary"
	primitives "github.com/galaco/bsp/primitives/worldlight"
	"log"
	"unsafe"
)

// WorldLight is Lump 15: Worldlight
type WorldLight struct {
	LumpGeneric
	data []primitives.WorldLight
}

// Unmarshall Imports this lump from raw byte data
func (lump *WorldLight) Unmarshall(raw []byte, length int32) {
	lump.LumpInfo.SetLength(length)
	if length == 0 {
		return
	}
	lump.data = make([]primitives.WorldLight, length/int32(unsafe.Sizeof(primitives.WorldLight{})))
	err := binary.Read(bytes.NewBuffer(raw), binary.LittleEndian, &lump.data)
	if err != nil {
		log.Fatal(err)
	}
}

// GetData gets internal format structure data
func (lump *WorldLight) GetData() []primitives.WorldLight {
	return lump.data
}

// Marshall dumps this lump back to raw byte data
func (lump *WorldLight) Marshall() ([]byte, error) {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.LittleEndian, lump.data)
	return buf.Bytes(), err
}
