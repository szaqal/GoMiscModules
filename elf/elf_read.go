package elf

import (
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

// ReadResult Product of ELF header read.
type ReadResult struct {
}

// ParseELFHeader parsess ELF header
func ParseELFHeader(fileName string) (*ReadResult, error) {

	headerData, err := ReadELFHeader(fileName)
	if err != nil {
		return nil, err
	}

	fmt.Printf("\n e_ident: %s", hex.EncodeToString(headerData[0:18]))
	fmt.Printf("\n e_ident[EI_CLASS]: %s", hex.EncodeToString(headerData[4:5]))   //1 or 2 to signify 32- or 64-bit format, respectively.
	fmt.Printf("\n e_ident[EI_VERSION]: %s", hex.EncodeToString(headerData[5:6])) // Always 1
	fmt.Printf("\n e_ident[EI_OSABI]: %s", hex.EncodeToString(headerData[7:8]))   //It is often set to 0 regardless of the target platform.
	fmt.Printf("\n e_machine: %s", hex.EncodeToString(headerData[19:20]))

	return &ReadResult{}, nil
}

// ReadELFHeader reads ELF header
func ReadELFHeader(fileName string) ([]byte, error) {

	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	// ELF header 40 bytes
	buf := make([]byte, 40)
	for {
		bytesRead, err := file.Read(buf)

		if bytesRead > 0 {
			valHex := hex.EncodeToString(buf[:bytesRead])
			fmt.Printf("ELF header: %s", valHex)
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("read %d bytes: %v", bytesRead, err)
			break
		}
		break
	}

	return buf, err
}
