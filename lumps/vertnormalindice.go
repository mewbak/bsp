package lumps

import (
	"bytes"
	"encoding/binary"
	"log"
)

// VertNormalIndice is Lump 31: VertNormalIndice
type VertNormalIndice struct {
	LumpGeneric
	data []uint16
}

// Unmarshall Imports this lump from raw byte data
func (lump *VertNormalIndice) Unmarshall(raw []byte, length int32) {
	lump.LumpInfo.SetLength(length)
	if length == 0 {
		return
	}

	lump.data = make([]uint16, length/int32(2))
	err := binary.Read(bytes.NewBuffer(raw), binary.LittleEndian, &lump.data)
	if err != nil {
		log.Fatal(err)
	}
}

// GetData gets internal format structure data
func (lump *VertNormalIndice) GetData() []uint16 {
	return lump.data
}

// Marshall dumps this lump back to raw byte data
func (lump *VertNormalIndice) Marshall() ([]byte, error) {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.LittleEndian, lump.data)
	return buf.Bytes(), err
}
