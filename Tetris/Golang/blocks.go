package main

// Index 0 = block type
// Index 1 = rotation
// Index 2 = cells
// Index 3 = x or y

type block [][][]int32

type Blocks struct {
	blocks []block
}

// ---------------------------------------------------------------------------------------------------------------------

func (b *Blocks) init() {
	b.lBlockPositions()
	b.jBlockPositions()
	b.iBlockPositions()
	b.oBlockPositions()
	b.sBlockPositions()
	b.tBlockPositions()
	b.zBlockPositions()
}

// ---------------------------------------------------------------------------------------------------------------------

func (b *Blocks) lBlockPositions() {
	var lBlock block
	lBlock = append(lBlock, [][]int32{{0, 2}, {1, 0}, {1, 1}, {1, 2}})
	lBlock = append(lBlock, [][]int32{{0, 1}, {1, 1}, {2, 1}, {2, 2}})
	lBlock = append(lBlock, [][]int32{{1, 0}, {2, 0}, {1, 1}, {1, 2}})
	lBlock = append(lBlock, [][]int32{{0, 0}, {0, 1}, {1, 1}, {2, 1}})

	b.blocks = append(b.blocks, lBlock)
}

// ---------------------------------------------------------------------------------------------------------------------

func (b *Blocks) jBlockPositions() {
	var jBlock block
	jBlock = append(jBlock, [][]int32{{0, 0}, {1, 0}, {1, 1}, {1, 2}})
	jBlock = append(jBlock, [][]int32{{0, 1}, {0, 2}, {1, 1}, {2, 1}})
	jBlock = append(jBlock, [][]int32{{1, 0}, {1, 1}, {1, 2}, {2, 2}})
	jBlock = append(jBlock, [][]int32{{0, 1}, {1, 1}, {2, 0}, {2, 1}})

	b.blocks = append(b.blocks, jBlock)
}

// ---------------------------------------------------------------------------------------------------------------------

func (b *Blocks) iBlockPositions() {
	var iBlock block
	iBlock = append(iBlock, [][]int32{{1, 0}, {1, 1}, {1, 2}, {1, 3}})
	iBlock = append(iBlock, [][]int32{{0, 2}, {1, 2}, {2, 2}, {3, 2}})
	iBlock = append(iBlock, [][]int32{{1, 0}, {1, 1}, {1, 2}, {1, 3}})
	iBlock = append(iBlock, [][]int32{{0, 2}, {1, 2}, {2, 2}, {3, 2}})

	b.blocks = append(b.blocks, iBlock)
}

// ---------------------------------------------------------------------------------------------------------------------

func (b *Blocks) oBlockPositions() {
	var oBlock block
	oBlock = append(oBlock, [][]int32{{0, 0}, {0, 1}, {1, 0}, {1, 1}})
	oBlock = append(oBlock, [][]int32{{0, 0}, {0, 1}, {1, 0}, {1, 1}})
	oBlock = append(oBlock, [][]int32{{0, 0}, {0, 1}, {1, 0}, {1, 1}})
	oBlock = append(oBlock, [][]int32{{0, 0}, {0, 1}, {1, 0}, {1, 1}})

	b.blocks = append(b.blocks, oBlock)
}

// ---------------------------------------------------------------------------------------------------------------------

func (b *Blocks) sBlockPositions() {
	var sBlock block
	sBlock = append(sBlock, [][]int32{{0, 1}, {0, 2}, {1, 0}, {1, 1}})
	sBlock = append(sBlock, [][]int32{{0, 1}, {1, 1}, {1, 2}, {2, 2}})
	sBlock = append(sBlock, [][]int32{{1, 1}, {1, 2}, {2, 0}, {2, 1}})
	sBlock = append(sBlock, [][]int32{{0, 0}, {1, 0}, {1, 1}, {2, 1}})

	b.blocks = append(b.blocks, sBlock)
}

// ---------------------------------------------------------------------------------------------------------------------

func (b *Blocks) tBlockPositions() {
	var tBlock block
	tBlock = append(tBlock, [][]int32{{0, 1}, {1, 0}, {1, 1}, {1, 2}})
	tBlock = append(tBlock, [][]int32{{0, 1}, {1, 1}, {1, 2}, {2, 1}})
	tBlock = append(tBlock, [][]int32{{1, 0}, {1, 1}, {1, 2}, {2, 1}})
	tBlock = append(tBlock, [][]int32{{0, 1}, {1, 0}, {1, 1}, {2, 1}})

	b.blocks = append(b.blocks, tBlock)
}

// ---------------------------------------------------------------------------------------------------------------------

func (b *Blocks) zBlockPositions() {
	var zBlock block
	zBlock = append(zBlock, [][]int32{{0, 0}, {0, 1}, {1, 1}, {1, 2}})
	zBlock = append(zBlock, [][]int32{{0, 2}, {1, 1}, {1, 2}, {2, 1}})
	zBlock = append(zBlock, [][]int32{{1, 0}, {1, 1}, {2, 1}, {2, 2}})
	zBlock = append(zBlock, [][]int32{{0, 1}, {1, 0}, {1, 1}, {2, 0}})

	b.blocks = append(b.blocks, zBlock)
}

// ---------------------------------------------------------------------------------------------------------------------
