package lazyfs

import (
 	"fmt"
 	"io"
)

//=====================================
type BlockStoreError struct {
	Err string
}

func (e BlockStoreError) Error() string {
	return e.Err
}

// const blockSizeBits uint = 20
// const blockSize uint = (1<<blockSizeBits)

//=====================================
type BlockStore struct {
	source FileSource
	blocks map[uint][]byte

	blockSizeBits uint
	blockSize     int64
}

func OpenBlockStore(source FileSource, bits uint) (*BlockStore, error) {

	fs := BlockStore{
					source: source,
					blocks: make(map[uint][]byte),
					blockSizeBits: bits,
					blockSize: int64(1<<bits),
		}

	return &fs, nil
}

func computeBlocks( off int64, sz int, bits uint ) (startBlock uint, endBlock uint) {
	startBlock = uint(off>>bits)
	endBlock   = uint((off + int64(sz))>>bits + 1)

	return startBlock, endBlock
}


//=====================================
func (fs *BlockStore) ReadAt(p []byte, off int64) (n int, err error) {

	startBlock, endBlock := computeBlocks( off, cap(p), fs.blockSizeBits)

	// Get any blocks we might need
	for i := startBlock; i < endBlock; i++ {
		if _,has := fs.blocks[i]; has == false {
			//fmt.Printf("Need to read block %d\n", i)
			fs.blocks[i] = make([]byte, fs.blockSize)
			n,_ = fs.source.ReadAt( fs.blocks[i], int64(i)*fs.blockSize )

			//fmt.Println( fs.blocks[i])
		}
	}

	endByte := off+int64(cap(p))
	sz,_ := fs.FileSize()
	if endByte > sz {
		endByte = sz
	}

	//fmt.Printf("endByte = %d\n", endByte)

	// Reset length of array
	p = p[:0]

	// Now copy from blocks to p
	for i := startBlock; i < endBlock; i++ {
		st := off - int64(i)*fs.blockSize
		if st < 0 {
			st = 0
		}

		en := endByte - int64( i) *fs. blockSize
		if en > fs.blockSize {
			en = fs.blockSize
		}

		//fmt.Println(i, p, st, en)
		p = append(p, fs.blocks[i][st:en]...)
		//fmt.Println(p)
	}

	return int(endByte-off), nil

	// // Check validity
	// if _, err := fs.HasAt(p, off); err == nil {
	// 	//fmt.Println("Retrieving from store")
	// 	return fs.file.ReadAt(p, off)
	// }
	//
	// //fmt.Println( "Need to update store")
	// n, _ = fs.source.ReadAt(p, off)
	// fs.WriteAt(p[:n], off)
	//
	// return n, nil

	// }
	//
	// n,err =  fs.HasAt( p, off )
	// if err != nil {
	// 	return 0, SparseFileStoreError{"ReadAt: Don't have all of the requested bytes"}
	// }
	//
	// return fs.file.ReadAt( p, off )
}

func (fs *BlockStore) WriteAt(p []byte, off int64) (n int, err error) {
	// n, err = fs.file.WriteAt(p, off)
	//
	// for i := 0; i < n; i++ {
	// 	fs.has[off+int64(i)] = true
	// }
	//
	return 0, nil
}

func (fs *BlockStore) HasAt(p []byte, off int64) (n int, err error) {

	startBlock, endBlock := computeBlocks( off, cap(p), fs.blockSizeBits)

	fmt.Printf("Reading %d bytes at offset %d traverses blocks from %d to %d", cap(p), off, startBlock, endBlock )

	for i := startBlock; i < endBlock; i++ {
		if _,has := fs.blocks[i]; has == false {
			return 0, SparseFileStoreError{"HasAt: Don't have all of the requested blocks"}
		}
	}

	return n, nil
}

func (fs *BlockStore) FileSize() (int64, error) {
	return fs.source.FileSize()
}

func (fs *BlockStore) Reader() io.Reader {
	return fs.source.Reader()
}

func (fs *BlockStore) Path() string {
	return fs.source.Path()
}
