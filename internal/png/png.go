package png

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"pngme/internal/chunk"
	"pngme/pkg/utils"
)

var (
	standartHeader = [8]byte{137, 80, 78, 71, 13, 10, 26, 10}
)

type Png struct {
	chunks []chunk.Chunk
}

func FromChunks(chunks []chunk.Chunk) *Png {
	png := new(Png)
	png.chunks = make([]chunk.Chunk, 0, len(chunks))
	png.chunks = append(png.chunks, chunks...)
	return png
}

func FromFile(f *os.File) (*Png, error) {
	png := new(Png)
	png.chunks = make([]chunk.Chunk, 0)
	r := bufio.NewReader(f)
	head, err := r.Peek(8)
	if err != nil {
		return nil, fmt.Errorf("png from file: %v", err)
	}
	if !bytes.Equal(head, standartHeader[:]) {
		return nil, fmt.Errorf("invalid png header: %v", head)
	}

	chunks := make([]chunk.Chunk, 0)
	for {
		size, err := r.Peek(4)
		if err != nil {
			return nil, err
		}

		ln := binary.BigEndian.Uint32(size)
		if err != nil {
			return nil, err
		}

		tp, err := r.Peek(4)
		if err != nil {
			return nil, err
		}

		if utils.B2S(tp) == chunk.IEND {
			break
		}

		data, err := r.Peek(int(ln))
		if err != nil {
			return nil, err
		}

		chunk, err := chunk.New(tp, data)
		if err != nil {
			return nil, err
		}

		// skip crc
		_, err = r.Discard(4)
		if err != nil {
			return nil, err
		}

		chunks = append(chunks, *chunk)
	}

	return png, nil
}

func (p *Png) Header() []byte {
	return standartHeader[:]
}

func (p *Png) Marshal() ([]byte, error) {
	buf := new(bytes.Buffer)

	if _, err := buf.Write(p.Header()); err != nil {
		return nil, err
	}

	for i := range p.chunks {
		rawChnk, err := p.chunks[i].Marshal()
		if err != nil {
			return nil, err
		}
		if err = binary.Write(buf, binary.BigEndian, rawChnk); err != nil {
			return nil, err
		}
	}

	end, err := chunk.MakeIEND().Marshal()
	if err != nil {
		return nil, err
	}
	if _, err := buf.Write(end); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (p *Png) ChunkByType(chnkT string) *chunk.Chunk {
	for _, chunk := range p.chunks {
		if chunk.Type().String() == chnkT {
			return &chunk
		}
	}
	return nil
}

func (p *Png) AppendChunk(chnk *chunk.Chunk) {
	p.chunks = append(p.chunks, *chnk)
}

func (p *Png) RemoveChunk(chnkT string) *chunk.Chunk {
	for i, chunk := range p.chunks {
		if chunk.Type().String() == chnkT {
			p.chunks = append(p.chunks[:i], p.chunks[i+1:]...)
			return &chunk
		}
	}
	return nil
}
