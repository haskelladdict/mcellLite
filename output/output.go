// Package output contains reaction and visualization data related functionality
//
// Copyright 2015 Markus Dittrich
// Licensed under BSD license, see LICENSE file for details
package output

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"os"
	"path"

	"github.com/haskelladdict/mcellLite/mol"
)

// WriteCB writes the current molecule position to a CellBlender compatible
// input format
// NOTE: This function does a lot of type casting mostly due to the way
// CellBlender format was designed. E.g., while positions are internally stored
// as 64bit doubles, CellBlender format writes them as 32bit floats. Eventually,
// the CellBlender format should probably be overhauled or completele redesigned.
func WriteCB(molMap mol.MolMap, outPath, fileName string, iter int64) error {
	// if output dir does not exist we create it
	if err := os.MkdirAll(outPath, 0700); err != nil {
		return fmt.Errorf("in writeCB: %s", err)
	}

	p := fmt.Sprintf("%s.cellbin.%04d.dat", fileName, iter)
	p = path.Join(outPath, p)
	file, err := os.Create(p)
	if err != nil {
		return fmt.Errorf("in writeCB: %s", err)
	}
	defer file.Close()

	// write version info
	var version uint32 = 1
	if err := writeUint32(file, version); err != nil {
		return err
	}

	// write molecules
	for sp, mols := range molMap {
		// write species info
		if err := writeUint8(file, uint8(len(sp))); err != nil {
			return err
		}
		if n, err := file.Write([]byte(sp)); err != nil || len(sp) != n {
			return err
		}

		// write mol type (0 = volume molecule)
		if err := writeUint8(file, 0); err != nil {
			return err
		}

		// write number of molecules and then molecule info
		if err := writeUint32(file, 3*uint32(len(mols))); err != nil {
			return err
		}
		for _, mol := range mols {
			if err := writeUint32(file, math.Float32bits(float32(mol.R.X))); err != nil {
				return err
			}
			if err := writeUint32(file, math.Float32bits(float32(mol.R.Y))); err != nil {
				return err
			}
			if err := writeUint32(file, math.Float32bits(float32(mol.R.Z))); err != nil {
				return err
			}
		}
	}
	return nil
}

// writeUint8 write an uint8 to the provided Writer in little endian format
func writeUint8(w io.Writer, i uint8) error {
	buf := []byte{i}
	if n, err := w.Write(buf); err != nil || n != 1 {
		return fmt.Errorf("in WriteUint8: %s", err)
	}
	return nil
}

// writeUint16 write an uint16 to the provided Writer in little endian format
func writeUint16(w io.Writer, i uint16) error {
	buf := make([]byte, 2)
	binary.LittleEndian.PutUint16(buf, i)
	if n, err := w.Write(buf); err != nil || n != 2 {
		return fmt.Errorf("in WriteUint16: %s", err)
	}
	return nil
}

// writeUint32 write an uint32 to the provided Writer in little endian format
func writeUint32(w io.Writer, i uint32) error {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, i)
	if n, err := w.Write(buf); err != nil || n != 4 {
		return fmt.Errorf("in WriteUint32: %s", err)
	}
	return nil
}

// writeUint64 write an uint64 to the provided Writer in little endian format
func writeUint64(w io.Writer, i uint64) error {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, i)
	if n, err := w.Write(buf); err != nil || n != 8 {
		return fmt.Errorf("in WriteUint64: %s", err)
	}
	return nil
}
