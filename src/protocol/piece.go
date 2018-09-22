// Implementation of data piece
package protocol

import (
	"encoding/binary"
	"unsafe"
	"fmt"
)


type Piece struct {
	length uint64
	position uint64
	data *[]byte
}


func Encode(p *Piece)  []byte {
	size := unsafe.Sizeof(p.length) + unsafe.Sizeof(p.position) + uintptr(len(*p.data))
	buffer := make([]byte, size)
	written := binary.PutUvarint(buffer, p.length)
	written += binary.PutUvarint(buffer[written:], p.position)
	for i := written; i < len(*p.data); i++ {
		buffer[i] = (*p.data)[i - written]
	}
	return buffer
}

/*func Decode([]byte bytes) Piece {

}*/


func main() {
	var piece []byte = []byte("Hello, World!")
	//const word_length uint64 = len(piece)
	//var bytes []byte = make([]byte, len(piece))
	var p Piece = Piece{uint64(len(piece)), 1, &piece}
	fmt.Println(Encode(&p))
}