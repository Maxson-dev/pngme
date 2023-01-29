package png

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
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

	head := make([]byte, 8)
	_, err := io.ReadFull(f, head)
	if err != nil {
		return nil, fmt.Errorf("png from file: %v\n", err)
	}

	if !bytes.Equal(head, standartHeader[:]) {
		return nil, fmt.Errorf("invalid png header: %v\n", head)
	}

	chunks := make([]chunk.Chunk, 0)
	for {
		buf := make([]byte, 4)
		_, err := io.ReadFull(f, buf)
		if err != nil {
			return nil, err
		}
		ln := binary.BigEndian.Uint32(buf)

		_, err = io.ReadFull(f, buf)
		if err != nil {
			return nil, err
		}

		if utils.B2S(buf) == chunk.IEND {
			break
		}

		data := make([]byte, ln)
		_, err = io.ReadFull(f, data)
		if err != nil {
			return nil, err
		}

		chunk, err := chunk.New(buf, data)
		if err != nil {
			return nil, err
		}

		// skip crc
		io.ReadFull(f, buf)
		if err != nil {
			return nil, err
		}

		chunks = append(chunks, *chunk)
	}

	png.chunks = chunks

	return png, nil
}

func (p *Png) Chunks() []chunk.Chunk {
	return p.chunks
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

func (p *Png) ChunksByType(chnkT string) []chunk.Chunk {
	res := make([]chunk.Chunk, 0)
	for _, chunk := range p.chunks {
		if chunk.Type().String() == chnkT {
			res = append(res, chunk)
		}
	}
	return res
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
