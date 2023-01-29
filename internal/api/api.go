package api

import (
	"fmt"
	"os"
	"pngme/internal/chunk"
	"pngme/internal/png"
	"pngme/pkg/utils"
)

func EncodeMessage(dst string, head string, msg string) error {
	f, err := os.Open(dst)
	if err != nil {
		return fmt.Errorf("could not open file %s: %v\n", dst, err)
	}
	defer f.Close()

	fpng, err := png.FromFile(f)
	if err != nil {
		return fmt.Errorf("could not read file %s: %v\n", dst, err)
	}

	chnk, err := chunk.New(utils.S2B(head), utils.S2B(msg))
	if err != nil {
		return fmt.Errorf("could not create chunk: %v\n", err)
	}

	fpng.AppendChunk(chnk)

	data, err := fpng.Marshal()
	if err != nil {
		return fmt.Errorf("could not marshal png: %v\n", err)
	}

	if err := os.WriteFile(dst, data, 0777); err != nil {
		return fmt.Errorf("could to write file %s: %v\n", dst, err)
	}
	return nil
}

func DecodeMessage(src string, head string) error {
	f, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("could not open file %s: %v\n", src, err)
	}
	defer f.Close()

	fpng, err := png.FromFile(f)
	if err != nil {
		return fmt.Errorf("could not read file %s: %v\n", src, err)
	}

	chnks := fpng.ChunksByType(head)

	fmt.Printf("-----------------------------Messages %s------------------------------\n", head)
	for i, chunk := range chnks {
		chunk.Print(i)
	}

	return nil
}

func PrintPngFile(src string) error {
	f, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("could not open file %s: %v\n", src, err)
	}
	defer f.Close()

	fpng, err := png.FromFile(f)
	if err != nil {
		return fmt.Errorf("could not read file %s: %v\n", src, err)
	}
	fmt.Printf("-----------------------------File %s------------------------------\n", src)
	for i, chunk := range fpng.Chunks() {
		chunk.Print(i)
	}
	return nil
}

func RemoveMessage(src string, head string) error {
	f, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("could not open file %s: %v\n", src, err)
	}
	defer f.Close()

	fpng, err := png.FromFile(f)
	if err != nil {
		return fmt.Errorf("could not read file %s: %v\n", src, err)
	}

	chnk := fpng.RemoveChunk(head)
	if chnk == nil {
		fmt.Printf(`Сhunk with "%s" header was not found\n`, head)
		return nil
	}
	fmt.Printf("Chunk successfully deleted\n")

	data, err := fpng.Marshal()
	if err != nil {
		return fmt.Errorf("could not marshal png: %v\n", err)
	}

	if err := os.WriteFile(src, data, 0777); err != nil {
		return fmt.Errorf("could to write file %s: %v\n", src, err)
	}

	return nil
}
