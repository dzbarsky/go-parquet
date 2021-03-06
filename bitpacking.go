package main

// Borrowed from https://github.com/fraugster/parquet-go/blob/master/bitbacking32.go

func unpack8int32_1(data []byte) (a [8]int32) {
	_ = data[0]
	a[0] = int32(uint32((data[0]>>0)&1) << 0)
	a[1] = int32(uint32((data[0]>>1)&1) << 0)
	a[2] = int32(uint32((data[0]>>2)&1) << 0)
	a[3] = int32(uint32((data[0]>>3)&1) << 0)
	a[4] = int32(uint32((data[0]>>4)&1) << 0)
	a[5] = int32(uint32((data[0]>>5)&1) << 0)
	a[6] = int32(uint32((data[0]>>6)&1) << 0)
	a[7] = int32(uint32((data[0]>>7)&1) << 0)
	return
}

func pack8int32_1(data [8]int32) []byte {
	return []byte{
		byte(uint32(data[0])<<0 | uint32(data[1])<<1 | uint32(data[2])<<2 | uint32(data[3])<<3 | uint32(data[4])<<4 | uint32(data[5])<<5 | uint32(data[6])<<6 | uint32(data[7])<<7),
	}
}

func unpack8int32_2(data []byte) (a [8]int32) {
	_ = data[1]
	a[0] = int32(uint32((data[0]>>0)&3) << 0)
	a[1] = int32(uint32((data[0]>>2)&3) << 0)
	a[2] = int32(uint32((data[0]>>4)&3) << 0)
	a[3] = int32(uint32((data[0]>>6)&3) << 0)
	a[4] = int32(uint32((data[1]>>0)&3) << 0)
	a[5] = int32(uint32((data[1]>>2)&3) << 0)
	a[6] = int32(uint32((data[1]>>4)&3) << 0)
	a[7] = int32(uint32((data[1]>>6)&3) << 0)
	return
}

func pack8int32_2(data [8]int32) []byte {
	return []byte{
		byte(uint32(data[0])<<0 | uint32(data[1])<<2 | uint32(data[2])<<4 | uint32(data[3])<<6),
		byte(uint32(data[4])<<0 | uint32(data[5])<<2 | uint32(data[6])<<4 | uint32(data[7])<<6),
	}
}

func unpack8int32_3(data []byte) (a [8]int32) {
	_ = data[2]
	a[0] = int32(uint32((data[0]>>0)&7) << 0)
	a[1] = int32(uint32((data[0]>>3)&7) << 0)
	a[2] = int32(uint32((data[0]>>6)&3)<<0 | uint32((data[1]>>0)&1)<<2)
	a[3] = int32(uint32((data[1]>>1)&7) << 0)
	a[4] = int32(uint32((data[1]>>4)&7) << 0)
	a[5] = int32(uint32((data[1]>>7)&1)<<0 | uint32((data[2]>>0)&3)<<1)
	a[6] = int32(uint32((data[2]>>2)&7) << 0)
	a[7] = int32(uint32((data[2]>>5)&7) << 0)
	return
}

func pack8int32_3(data [8]int32) []byte {
	return []byte{
		byte(uint32(data[0])<<0 | uint32(data[1])<<3 | uint32(data[2])<<6),
		byte(uint32(data[2])>>2 | uint32(data[3])<<1 | uint32(data[4])<<4 | uint32(data[5])<<7),
		byte(uint32(data[5])>>1 | uint32(data[6])<<2 | uint32(data[7])<<5),
	}
}

func unpack8int32_4(data []byte) (a [8]int32) {
	_ = data[3]
	a[0] = int32(uint32((data[0]>>0)&15) << 0)
	a[1] = int32(uint32((data[0]>>4)&15) << 0)
	a[2] = int32(uint32((data[1]>>0)&15) << 0)
	a[3] = int32(uint32((data[1]>>4)&15) << 0)
	a[4] = int32(uint32((data[2]>>0)&15) << 0)
	a[5] = int32(uint32((data[2]>>4)&15) << 0)
	a[6] = int32(uint32((data[3]>>0)&15) << 0)
	a[7] = int32(uint32((data[3]>>4)&15) << 0)
	return
}

func pack8int32_4(data [8]int32) []byte {
	return []byte{
		byte(uint32(data[0])<<0 | uint32(data[1])<<4),
		byte(uint32(data[2])<<0 | uint32(data[3])<<4),
		byte(uint32(data[4])<<0 | uint32(data[5])<<4),
		byte(uint32(data[6])<<0 | uint32(data[7])<<4),
	}
}

func unpack8int32_5(data []byte) (a [8]int32) {
	_ = data[4]
	a[0] = int32(uint32((data[0]>>0)&31) << 0)
	a[1] = int32(uint32((data[0]>>5)&7)<<0 | uint32((data[1]>>0)&3)<<3)
	a[2] = int32(uint32((data[1]>>2)&31) << 0)
	a[3] = int32(uint32((data[1]>>7)&1)<<0 | uint32((data[2]>>0)&15)<<1)
	a[4] = int32(uint32((data[2]>>4)&15)<<0 | uint32((data[3]>>0)&1)<<4)
	a[5] = int32(uint32((data[3]>>1)&31) << 0)
	a[6] = int32(uint32((data[3]>>6)&3)<<0 | uint32((data[4]>>0)&7)<<2)
	a[7] = int32(uint32((data[4]>>3)&31) << 0)
	return
}

func pack8int32_5(data [8]int32) []byte {
	return []byte{
		byte(uint32(data[0])<<0 | uint32(data[1])<<5),
		byte(uint32(data[1])>>3 | uint32(data[2])<<2 | uint32(data[3])<<7),
		byte(uint32(data[3])>>1 | uint32(data[4])<<4),
		byte(uint32(data[4])>>4 | uint32(data[5])<<1 | uint32(data[6])<<6),
		byte(uint32(data[6])>>2 | uint32(data[7])<<3),
	}
}

func unpack8int32_6(data []byte) (a [8]int32) {
	_ = data[5]
	a[0] = int32(uint32((data[0]>>0)&63) << 0)
	a[1] = int32(uint32((data[0]>>6)&3)<<0 | uint32((data[1]>>0)&15)<<2)
	a[2] = int32(uint32((data[1]>>4)&15)<<0 | uint32((data[2]>>0)&3)<<4)
	a[3] = int32(uint32((data[2]>>2)&63) << 0)
	a[4] = int32(uint32((data[3]>>0)&63) << 0)
	a[5] = int32(uint32((data[3]>>6)&3)<<0 | uint32((data[4]>>0)&15)<<2)
	a[6] = int32(uint32((data[4]>>4)&15)<<0 | uint32((data[5]>>0)&3)<<4)
	a[7] = int32(uint32((data[5]>>2)&63) << 0)
	return
}

func pack8int32_6(data [8]int32) []byte {
	return []byte{
		byte(uint32(data[0])<<0 | uint32(data[1])<<6),
		byte(uint32(data[1])>>2 | uint32(data[2])<<4),
		byte(uint32(data[2])>>4 | uint32(data[3])<<2),
		byte(uint32(data[4])<<0 | uint32(data[5])<<6),
		byte(uint32(data[5])>>2 | uint32(data[6])<<4),
		byte(uint32(data[6])>>4 | uint32(data[7])<<2),
	}
}

func unpack8int32_7(data []byte) (a [8]int32) {
	_ = data[6]
	a[0] = int32(uint32((data[0]>>0)&127) << 0)
	a[1] = int32(uint32((data[0]>>7)&1)<<0 | uint32((data[1]>>0)&63)<<1)
	a[2] = int32(uint32((data[1]>>6)&3)<<0 | uint32((data[2]>>0)&31)<<2)
	a[3] = int32(uint32((data[2]>>5)&7)<<0 | uint32((data[3]>>0)&15)<<3)
	a[4] = int32(uint32((data[3]>>4)&15)<<0 | uint32((data[4]>>0)&7)<<4)
	a[5] = int32(uint32((data[4]>>3)&31)<<0 | uint32((data[5]>>0)&3)<<5)
	a[6] = int32(uint32((data[5]>>2)&63)<<0 | uint32((data[6]>>0)&1)<<6)
	a[7] = int32(uint32((data[6]>>1)&127) << 0)
	return
}

func pack8int32_7(data [8]int32) []byte {
	return []byte{
		byte(uint32(data[0])<<0 | uint32(data[1])<<7),
		byte(uint32(data[1])>>1 | uint32(data[2])<<6),
		byte(uint32(data[2])>>2 | uint32(data[3])<<5),
		byte(uint32(data[3])>>3 | uint32(data[4])<<4),
		byte(uint32(data[4])>>4 | uint32(data[5])<<3),
		byte(uint32(data[5])>>5 | uint32(data[6])<<2),
		byte(uint32(data[6])>>6 | uint32(data[7])<<1),
	}
}

func unpack8int32_8(data []byte) (a [8]int32) {
	_ = data[7]
	a[0] = int32(uint32((data[0]>>0)&255) << 0)
	a[1] = int32(uint32((data[1]>>0)&255) << 0)
	a[2] = int32(uint32((data[2]>>0)&255) << 0)
	a[3] = int32(uint32((data[3]>>0)&255) << 0)
	a[4] = int32(uint32((data[4]>>0)&255) << 0)
	a[5] = int32(uint32((data[5]>>0)&255) << 0)
	a[6] = int32(uint32((data[6]>>0)&255) << 0)
	a[7] = int32(uint32((data[7]>>0)&255) << 0)
	return
}

func pack8int32_8(data [8]int32) []byte {
	return []byte{
		byte(uint32(data[0]) << 0),
		byte(uint32(data[1]) << 0),
		byte(uint32(data[2]) << 0),
		byte(uint32(data[3]) << 0),
		byte(uint32(data[4]) << 0),
		byte(uint32(data[5]) << 0),
		byte(uint32(data[6]) << 0),
		byte(uint32(data[7]) << 0),
	}
}

func unpack8int32_9(data []byte) (a [8]int32) {
	_ = data[8]
	a[0] = int32(uint32((data[0]>>0)&255)<<0 | uint32((data[1]>>0)&1)<<8)
	a[1] = int32(uint32((data[1]>>1)&127)<<0 | uint32((data[2]>>0)&3)<<7)
	a[2] = int32(uint32((data[2]>>2)&63)<<0 | uint32((data[3]>>0)&7)<<6)
	a[3] = int32(uint32((data[3]>>3)&31)<<0 | uint32((data[4]>>0)&15)<<5)
	a[4] = int32(uint32((data[4]>>4)&15)<<0 | uint32((data[5]>>0)&31)<<4)
	a[5] = int32(uint32((data[5]>>5)&7)<<0 | uint32((data[6]>>0)&63)<<3)
	a[6] = int32(uint32((data[6]>>6)&3)<<0 | uint32((data[7]>>0)&127)<<2)
	a[7] = int32(uint32((data[7]>>7)&1)<<0 | uint32((data[8]>>0)&255)<<1)
	return
}

func pack8int32_9(data [8]int32) []byte {
	return []byte{
		byte(uint32(data[0]) << 0),
		byte(uint32(data[0])>>8 | uint32(data[1])<<1),
		byte(uint32(data[1])>>7 | uint32(data[2])<<2),
		byte(uint32(data[2])>>6 | uint32(data[3])<<3),
		byte(uint32(data[3])>>5 | uint32(data[4])<<4),
		byte(uint32(data[4])>>4 | uint32(data[5])<<5),
		byte(uint32(data[5])>>3 | uint32(data[6])<<6),
		byte(uint32(data[6])>>2 | uint32(data[7])<<7),
		byte(uint32(data[7]) >> 1),
	}
}

func unpack8int32_10(data []byte) (a [8]int32) {
	_ = data[9]
	a[0] = int32(uint32((data[0]>>0)&255)<<0 | uint32((data[1]>>0)&3)<<8)
	a[1] = int32(uint32((data[1]>>2)&63)<<0 | uint32((data[2]>>0)&15)<<6)
	a[2] = int32(uint32((data[2]>>4)&15)<<0 | uint32((data[3]>>0)&63)<<4)
	a[3] = int32(uint32((data[3]>>6)&3)<<0 | uint32((data[4]>>0)&255)<<2)
	a[4] = int32(uint32((data[5]>>0)&255)<<0 | uint32((data[6]>>0)&3)<<8)
	a[5] = int32(uint32((data[6]>>2)&63)<<0 | uint32((data[7]>>0)&15)<<6)
	a[6] = int32(uint32((data[7]>>4)&15)<<0 | uint32((data[8]>>0)&63)<<4)
	a[7] = int32(uint32((data[8]>>6)&3)<<0 | uint32((data[9]>>0)&255)<<2)
	return
}

func pack8int32_10(data [8]int32) []byte {
	return []byte{
		byte(uint32(data[0]) << 0),
		byte(uint32(data[0])>>8 | uint32(data[1])<<2),
		byte(uint32(data[1])>>6 | uint32(data[2])<<4),
		byte(uint32(data[2])>>4 | uint32(data[3])<<6),
		byte(uint32(data[3]) >> 2),
		byte(uint32(data[4]) << 0),
		byte(uint32(data[4])>>8 | uint32(data[5])<<2),
		byte(uint32(data[5])>>6 | uint32(data[6])<<4),
		byte(uint32(data[6])>>4 | uint32(data[7])<<6),
		byte(uint32(data[7]) >> 2),
	}
}

func unpack8int32_11(data []byte) (a [8]int32) {
	_ = data[10]
	a[0] = int32(uint32((data[0]>>0)&255)<<0 | uint32((data[1]>>0)&7)<<8)
	a[1] = int32(uint32((data[1]>>3)&31)<<0 | uint32((data[2]>>0)&63)<<5)
	a[2] = int32(uint32((data[2]>>6)&3)<<0 | uint32((data[3]>>0)&255)<<2 | uint32((data[4]>>0)&1)<<10)
	a[3] = int32(uint32((data[4]>>1)&127)<<0 | uint32((data[5]>>0)&15)<<7)
	a[4] = int32(uint32((data[5]>>4)&15)<<0 | uint32((data[6]>>0)&127)<<4)
	a[5] = int32(uint32((data[6]>>7)&1)<<0 | uint32((data[7]>>0)&255)<<1 | uint32((data[8]>>0)&3)<<9)
	a[6] = int32(uint32((data[8]>>2)&63)<<0 | uint32((data[9]>>0)&31)<<6)
	a[7] = int32(uint32((data[9]>>5)&7)<<0 | uint32((data[10]>>0)&255)<<3)
	return
}

func pack8int32_11(data [8]int32) []byte {
	return []byte{
		byte(uint32(data[0]) << 0),
		byte(uint32(data[0])>>8 | uint32(data[1])<<3),
		byte(uint32(data[1])>>5 | uint32(data[2])<<6),
		byte(uint32(data[2]) >> 2),
		byte(uint32(data[2])>>10 | uint32(data[3])<<1),
		byte(uint32(data[3])>>7 | uint32(data[4])<<4),
		byte(uint32(data[4])>>4 | uint32(data[5])<<7),
		byte(uint32(data[5]) >> 1),
		byte(uint32(data[5])>>9 | uint32(data[6])<<2),
		byte(uint32(data[6])>>6 | uint32(data[7])<<5),
		byte(uint32(data[7]) >> 3),
	}
}

func unpack8int32_12(data []byte) (a [8]int32) {
	_ = data[11]
	a[0] = int32(uint32((data[0]>>0)&255)<<0 | uint32((data[1]>>0)&15)<<8)
	a[1] = int32(uint32((data[1]>>4)&15)<<0 | uint32((data[2]>>0)&255)<<4)
	a[2] = int32(uint32((data[3]>>0)&255)<<0 | uint32((data[4]>>0)&15)<<8)
	a[3] = int32(uint32((data[4]>>4)&15)<<0 | uint32((data[5]>>0)&255)<<4)
	a[4] = int32(uint32((data[6]>>0)&255)<<0 | uint32((data[7]>>0)&15)<<8)
	a[5] = int32(uint32((data[7]>>4)&15)<<0 | uint32((data[8]>>0)&255)<<4)
	a[6] = int32(uint32((data[9]>>0)&255)<<0 | uint32((data[10]>>0)&15)<<8)
	a[7] = int32(uint32((data[10]>>4)&15)<<0 | uint32((data[11]>>0)&255)<<4)
	return
}

func pack8int32_12(data [8]int32) []byte {
	return []byte{
		byte(uint32(data[0]) << 0),
		byte(uint32(data[0])>>8 | uint32(data[1])<<4),
		byte(uint32(data[1]) >> 4),
		byte(uint32(data[2]) << 0),
		byte(uint32(data[2])>>8 | uint32(data[3])<<4),
		byte(uint32(data[3]) >> 4),
		byte(uint32(data[4]) << 0),
		byte(uint32(data[4])>>8 | uint32(data[5])<<4),
		byte(uint32(data[5]) >> 4),
		byte(uint32(data[6]) << 0),
		byte(uint32(data[6])>>8 | uint32(data[7])<<4),
		byte(uint32(data[7]) >> 4),
	}
}

func unpack8int32_13(data []byte) (a [8]int32) {
	_ = data[12]
	a[0] = int32(uint32((data[0]>>0)&255)<<0 | uint32((data[1]>>0)&31)<<8)
	a[1] = int32(uint32((data[1]>>5)&7)<<0 | uint32((data[2]>>0)&255)<<3 | uint32((data[3]>>0)&3)<<11)
	a[2] = int32(uint32((data[3]>>2)&63)<<0 | uint32((data[4]>>0)&127)<<6)
	a[3] = int32(uint32((data[4]>>7)&1)<<0 | uint32((data[5]>>0)&255)<<1 | uint32((data[6]>>0)&15)<<9)
	a[4] = int32(uint32((data[6]>>4)&15)<<0 | uint32((data[7]>>0)&255)<<4 | uint32((data[8]>>0)&1)<<12)
	a[5] = int32(uint32((data[8]>>1)&127)<<0 | uint32((data[9]>>0)&63)<<7)
	a[6] = int32(uint32((data[9]>>6)&3)<<0 | uint32((data[10]>>0)&255)<<2 | uint32((data[11]>>0)&7)<<10)
	a[7] = int32(uint32((data[11]>>3)&31)<<0 | uint32((data[12]>>0)&255)<<5)
	return
}

func pack8int32_13(data [8]int32) []byte {
	return []byte{
		byte(uint32(data[0]) << 0),
		byte(uint32(data[0])>>8 | uint32(data[1])<<5),
		byte(uint32(data[1]) >> 3),
		byte(uint32(data[1])>>11 | uint32(data[2])<<2),
		byte(uint32(data[2])>>6 | uint32(data[3])<<7),
		byte(uint32(data[3]) >> 1),
		byte(uint32(data[3])>>9 | uint32(data[4])<<4),
		byte(uint32(data[4]) >> 4),
		byte(uint32(data[4])>>12 | uint32(data[5])<<1),
		byte(uint32(data[5])>>7 | uint32(data[6])<<6),
		byte(uint32(data[6]) >> 2),
		byte(uint32(data[6])>>10 | uint32(data[7])<<3),
		byte(uint32(data[7]) >> 5),
	}
}

func unpack8int32_14(data []byte) (a [8]int32) {
	_ = data[13]
	a[0] = int32(uint32((data[0]>>0)&255)<<0 | uint32((data[1]>>0)&63)<<8)
	a[1] = int32(uint32((data[1]>>6)&3)<<0 | uint32((data[2]>>0)&255)<<2 | uint32((data[3]>>0)&15)<<10)
	a[2] = int32(uint32((data[3]>>4)&15)<<0 | uint32((data[4]>>0)&255)<<4 | uint32((data[5]>>0)&3)<<12)
	a[3] = int32(uint32((data[5]>>2)&63)<<0 | uint32((data[6]>>0)&255)<<6)
	a[4] = int32(uint32((data[7]>>0)&255)<<0 | uint32((data[8]>>0)&63)<<8)
	a[5] = int32(uint32((data[8]>>6)&3)<<0 | uint32((data[9]>>0)&255)<<2 | uint32((data[10]>>0)&15)<<10)
	a[6] = int32(uint32((data[10]>>4)&15)<<0 | uint32((data[11]>>0)&255)<<4 | uint32((data[12]>>0)&3)<<12)
	a[7] = int32(uint32((data[12]>>2)&63)<<0 | uint32((data[13]>>0)&255)<<6)
	return
}

func pack8int32_14(data [8]int32) []byte {
	return []byte{
		byte(uint32(data[0]) << 0),
		byte(uint32(data[0])>>8 | uint32(data[1])<<6),
		byte(uint32(data[1]) >> 2),
		byte(uint32(data[1])>>10 | uint32(data[2])<<4),
		byte(uint32(data[2]) >> 4),
		byte(uint32(data[2])>>12 | uint32(data[3])<<2),
		byte(uint32(data[3]) >> 6),
		byte(uint32(data[4]) << 0),
		byte(uint32(data[4])>>8 | uint32(data[5])<<6),
		byte(uint32(data[5]) >> 2),
		byte(uint32(data[5])>>10 | uint32(data[6])<<4),
		byte(uint32(data[6]) >> 4),
		byte(uint32(data[6])>>12 | uint32(data[7])<<2),
		byte(uint32(data[7]) >> 6),
	}
}

func unpack8int32_15(data []byte) (a [8]int32) {
	_ = data[14]
	a[0] = int32(uint32((data[0]>>0)&255)<<0 | uint32((data[1]>>0)&127)<<8)
	a[1] = int32(uint32((data[1]>>7)&1)<<0 | uint32((data[2]>>0)&255)<<1 | uint32((data[3]>>0)&63)<<9)
	a[2] = int32(uint32((data[3]>>6)&3)<<0 | uint32((data[4]>>0)&255)<<2 | uint32((data[5]>>0)&31)<<10)
	a[3] = int32(uint32((data[5]>>5)&7)<<0 | uint32((data[6]>>0)&255)<<3 | uint32((data[7]>>0)&15)<<11)
	a[4] = int32(uint32((data[7]>>4)&15)<<0 | uint32((data[8]>>0)&255)<<4 | uint32((data[9]>>0)&7)<<12)
	a[5] = int32(uint32((data[9]>>3)&31)<<0 | uint32((data[10]>>0)&255)<<5 | uint32((data[11]>>0)&3)<<13)
	a[6] = int32(uint32((data[11]>>2)&63)<<0 | uint32((data[12]>>0)&255)<<6 | uint32((data[13]>>0)&1)<<14)
	a[7] = int32(uint32((data[13]>>1)&127)<<0 | uint32((data[14]>>0)&255)<<7)
	return
}

func pack8int32_15(data [8]int32) []byte {
	return []byte{
		byte(uint32(data[0]) << 0),
		byte(uint32(data[0])>>8 | uint32(data[1])<<7),
		byte(uint32(data[1]) >> 1),
		byte(uint32(data[1])>>9 | uint32(data[2])<<6),
		byte(uint32(data[2]) >> 2),
		byte(uint32(data[2])>>10 | uint32(data[3])<<5),
		byte(uint32(data[3]) >> 3),
		byte(uint32(data[3])>>11 | uint32(data[4])<<4),
		byte(uint32(data[4]) >> 4),
		byte(uint32(data[4])>>12 | uint32(data[5])<<3),
		byte(uint32(data[5]) >> 5),
		byte(uint32(data[5])>>13 | uint32(data[6])<<2),
		byte(uint32(data[6]) >> 6),
		byte(uint32(data[6])>>14 | uint32(data[7])<<1),
		byte(uint32(data[7]) >> 7),
	}
}

func unpack8int32_16(data []byte) (a [8]int32) {
	_ = data[15]
	a[0] = int32(uint32((data[0]>>0)&255)<<0 | uint32((data[1]>>0)&255)<<8)
	a[1] = int32(uint32((data[2]>>0)&255)<<0 | uint32((data[3]>>0)&255)<<8)
	a[2] = int32(uint32((data[4]>>0)&255)<<0 | uint32((data[5]>>0)&255)<<8)
	a[3] = int32(uint32((data[6]>>0)&255)<<0 | uint32((data[7]>>0)&255)<<8)
	a[4] = int32(uint32((data[8]>>0)&255)<<0 | uint32((data[9]>>0)&255)<<8)
	a[5] = int32(uint32((data[10]>>0)&255)<<0 | uint32((data[11]>>0)&255)<<8)
	a[6] = int32(uint32((data[12]>>0)&255)<<0 | uint32((data[13]>>0)&255)<<8)
	a[7] = int32(uint32((data[14]>>0)&255)<<0 | uint32((data[15]>>0)&255)<<8)
	return
}

func pack8int32_16(data [8]int32) []byte {
	return []byte{
		byte(uint32(data[0]) << 0),
		byte(uint32(data[0]) >> 8),
		byte(uint32(data[1]) << 0),
		byte(uint32(data[1]) >> 8),
		byte(uint32(data[2]) << 0),
		byte(uint32(data[2]) >> 8),
		byte(uint32(data[3]) << 0),
		byte(uint32(data[3]) >> 8),
		byte(uint32(data[4]) << 0),
		byte(uint32(data[4]) >> 8),
		byte(uint32(data[5]) << 0),
		byte(uint32(data[5]) >> 8),
		byte(uint32(data[6]) << 0),
		byte(uint32(data[6]) >> 8),
		byte(uint32(data[7]) << 0),
		byte(uint32(data[7]) >> 8),
	}
}

func unpack8int32_17(data []byte) (a [8]int32) {
	_ = data[16]
	a[0] = int32(uint32((data[0]>>0)&255)<<0 | uint32((data[1]>>0)&255)<<8 | uint32((data[2]>>0)&1)<<16)
	a[1] = int32(uint32((data[2]>>1)&127)<<0 | uint32((data[3]>>0)&255)<<7 | uint32((data[4]>>0)&3)<<15)
	a[2] = int32(uint32((data[4]>>2)&63)<<0 | uint32((data[5]>>0)&255)<<6 | uint32((data[6]>>0)&7)<<14)
	a[3] = int32(uint32((data[6]>>3)&31)<<0 | uint32((data[7]>>0)&255)<<5 | uint32((data[8]>>0)&15)<<13)
	a[4] = int32(uint32((data[8]>>4)&15)<<0 | uint32((data[9]>>0)&255)<<4 | uint32((data[10]>>0)&31)<<12)
	a[5] = int32(uint32((data[10]>>5)&7)<<0 | uint32((data[11]>>0)&255)<<3 | uint32((data[12]>>0)&63)<<11)
	a[6] = int32(uint32((data[12]>>6)&3)<<0 | uint32((data[13]>>0)&255)<<2 | uint32((data[14]>>0)&127)<<10)
	a[7] = int32(uint32((data[14]>>7)&1)<<0 | uint32((data[15]>>0)&255)<<1 | uint32((data[16]>>0)&255)<<9)
	return
}

func pack8int32_17(data [8]int32) []byte {
	return []byte{
		byte(uint32(data[0]) << 0),
		byte(uint32(data[0]) >> 8),
		byte(uint32(data[0])>>16 | uint32(data[1])<<1),
		byte(uint32(data[1]) >> 7),
		byte(uint32(data[1])>>15 | uint32(data[2])<<2),
		byte(uint32(data[2]) >> 6),
		byte(uint32(data[2])>>14 | uint32(data[3])<<3),
		byte(uint32(data[3]) >> 5),
		byte(uint32(data[3])>>13 | uint32(data[4])<<4),
		byte(uint32(data[4]) >> 4),
		byte(uint32(data[4])>>12 | uint32(data[5])<<5),
		byte(uint32(data[5]) >> 3),
		byte(uint32(data[5])>>11 | uint32(data[6])<<6),
		byte(uint32(data[6]) >> 2),
		byte(uint32(data[6])>>10 | uint32(data[7])<<7),
		byte(uint32(data[7]) >> 1),
		byte(uint32(data[7]) >> 9),
	}
}

func unpack8int32_18(data []byte) (a [8]int32) {
	_ = data[17]
	a[0] = int32(uint32((data[0]>>0)&255)<<0 | uint32((data[1]>>0)&255)<<8 | uint32((data[2]>>0)&3)<<16)
	a[1] = int32(uint32((data[2]>>2)&63)<<0 | uint32((data[3]>>0)&255)<<6 | uint32((data[4]>>0)&15)<<14)
	a[2] = int32(uint32((data[4]>>4)&15)<<0 | uint32((data[5]>>0)&255)<<4 | uint32((data[6]>>0)&63)<<12)
	a[3] = int32(uint32((data[6]>>6)&3)<<0 | uint32((data[7]>>0)&255)<<2 | uint32((data[8]>>0)&255)<<10)
	a[4] = int32(uint32((data[9]>>0)&255)<<0 | uint32((data[10]>>0)&255)<<8 | uint32((data[11]>>0)&3)<<16)
	a[5] = int32(uint32((data[11]>>2)&63)<<0 | uint32((data[12]>>0)&255)<<6 | uint32((data[13]>>0)&15)<<14)
	a[6] = int32(uint32((data[13]>>4)&15)<<0 | uint32((data[14]>>0)&255)<<4 | uint32((data[15]>>0)&63)<<12)
	a[7] = int32(uint32((data[15]>>6)&3)<<0 | uint32((data[16]>>0)&255)<<2 | uint32((data[17]>>0)&255)<<10)
	return
}

func pack8int32_18(data [8]int32) []byte {
	return []byte{
		byte(uint32(data[0]) << 0),
		byte(uint32(data[0]) >> 8),
		byte(uint32(data[0])>>16 | uint32(data[1])<<2),
		byte(uint32(data[1]) >> 6),
		byte(uint32(data[1])>>14 | uint32(data[2])<<4),
		byte(uint32(data[2]) >> 4),
		byte(uint32(data[2])>>12 | uint32(data[3])<<6),
		byte(uint32(data[3]) >> 2),
		byte(uint32(data[3]) >> 10),
		byte(uint32(data[4]) << 0),
		byte(uint32(data[4]) >> 8),
		byte(uint32(data[4])>>16 | uint32(data[5])<<2),
		byte(uint32(data[5]) >> 6),
		byte(uint32(data[5])>>14 | uint32(data[6])<<4),
		byte(uint32(data[6]) >> 4),
		byte(uint32(data[6])>>12 | uint32(data[7])<<6),
		byte(uint32(data[7]) >> 2),
		byte(uint32(data[7]) >> 10),
	}
}

func unpack8int32_19(data []byte) (a [8]int32) {
	_ = data[18]
	a[0] = int32(uint32((data[0]>>0)&255)<<0 | uint32((data[1]>>0)&255)<<8 | uint32((data[2]>>0)&7)<<16)
	a[1] = int32(uint32((data[2]>>3)&31)<<0 | uint32((data[3]>>0)&255)<<5 | uint32((data[4]>>0)&63)<<13)
	a[2] = int32(uint32((data[4]>>6)&3)<<0 | uint32((data[5]>>0)&255)<<2 | uint32((data[6]>>0)&255)<<10 | uint32((data[7]>>0)&1)<<18)
	a[3] = int32(uint32((data[7]>>1)&127)<<0 | uint32((data[8]>>0)&255)<<7 | uint32((data[9]>>0)&15)<<15)
	a[4] = int32(uint32((data[9]>>4)&15)<<0 | uint32((data[10]>>0)&255)<<4 | uint32((data[11]>>0)&127)<<12)
	a[5] = int32(uint32((data[11]>>7)&1)<<0 | uint32((data[12]>>0)&255)<<1 | uint32((data[13]>>0)&255)<<9 | uint32((data[14]>>0)&3)<<17)
	a[6] = int32(uint32((data[14]>>2)&63)<<0 | uint32((data[15]>>0)&255)<<6 | uint32((data[16]>>0)&31)<<14)
	a[7] = int32(uint32((data[16]>>5)&7)<<0 | uint32((data[17]>>0)&255)<<3 | uint32((data[18]>>0)&255)<<11)
	return
}

func pack8int32_19(data [8]int32) []byte {
	return []byte{
		byte(uint32(data[0]) << 0),
		byte(uint32(data[0]) >> 8),
		byte(uint32(data[0])>>16 | uint32(data[1])<<3),
		byte(uint32(data[1]) >> 5),
		byte(uint32(data[1])>>13 | uint32(data[2])<<6),
		byte(uint32(data[2]) >> 2),
		byte(uint32(data[2]) >> 10),
		byte(uint32(data[2])>>18 | uint32(data[3])<<1),
		byte(uint32(data[3]) >> 7),
		byte(uint32(data[3])>>15 | uint32(data[4])<<4),
		byte(uint32(data[4]) >> 4),
		byte(uint32(data[4])>>12 | uint32(data[5])<<7),
		byte(uint32(data[5]) >> 1),
		byte(uint32(data[5]) >> 9),
		byte(uint32(data[5])>>17 | uint32(data[6])<<2),
		byte(uint32(data[6]) >> 6),
		byte(uint32(data[6])>>14 | uint32(data[7])<<5),
		byte(uint32(data[7]) >> 3),
		byte(uint32(data[7]) >> 11),
	}
}

func unpack8int32_20(data []byte) (a [8]int32) {
	_ = data[19]
	a[0] = int32(uint32((data[0]>>0)&255)<<0 | uint32((data[1]>>0)&255)<<8 | uint32((data[2]>>0)&15)<<16)
	a[1] = int32(uint32((data[2]>>4)&15)<<0 | uint32((data[3]>>0)&255)<<4 | uint32((data[4]>>0)&255)<<12)
	a[2] = int32(uint32((data[5]>>0)&255)<<0 | uint32((data[6]>>0)&255)<<8 | uint32((data[7]>>0)&15)<<16)
	a[3] = int32(uint32((data[7]>>4)&15)<<0 | uint32((data[8]>>0)&255)<<4 | uint32((data[9]>>0)&255)<<12)
	a[4] = int32(uint32((data[10]>>0)&255)<<0 | uint32((data[11]>>0)&255)<<8 | uint32((data[12]>>0)&15)<<16)
	a[5] = int32(uint32((data[12]>>4)&15)<<0 | uint32((data[13]>>0)&255)<<4 | uint32((data[14]>>0)&255)<<12)
	a[6] = int32(uint32((data[15]>>0)&255)<<0 | uint32((data[16]>>0)&255)<<8 | uint32((data[17]>>0)&15)<<16)
	a[7] = int32(uint32((data[17]>>4)&15)<<0 | uint32((data[18]>>0)&255)<<4 | uint32((data[19]>>0)&255)<<12)
	return
}

func pack8int32_20(data [8]int32) []byte {
	return []byte{
		byte(uint32(data[0]) << 0),
		byte(uint32(data[0]) >> 8),
		byte(uint32(data[0])>>16 | uint32(data[1])<<4),
		byte(uint32(data[1]) >> 4),
		byte(uint32(data[1]) >> 12),
		byte(uint32(data[2]) << 0),
		byte(uint32(data[2]) >> 8),
		byte(uint32(data[2])>>16 | uint32(data[3])<<4),
		byte(uint32(data[3]) >> 4),
		byte(uint32(data[3]) >> 12),
		byte(uint32(data[4]) << 0),
		byte(uint32(data[4]) >> 8),
		byte(uint32(data[4])>>16 | uint32(data[5])<<4),
		byte(uint32(data[5]) >> 4),
		byte(uint32(data[5]) >> 12),
		byte(uint32(data[6]) << 0),
		byte(uint32(data[6]) >> 8),
		byte(uint32(data[6])>>16 | uint32(data[7])<<4),
		byte(uint32(data[7]) >> 4),
		byte(uint32(data[7]) >> 12),
	}
}

func unpack8int32_21(data []byte) (a [8]int32) {
	_ = data[20]
	a[0] = int32(uint32((data[0]>>0)&255)<<0 | uint32((data[1]>>0)&255)<<8 | uint32((data[2]>>0)&31)<<16)
	a[1] = int32(uint32((data[2]>>5)&7)<<0 | uint32((data[3]>>0)&255)<<3 | uint32((data[4]>>0)&255)<<11 | uint32((data[5]>>0)&3)<<19)
	a[2] = int32(uint32((data[5]>>2)&63)<<0 | uint32((data[6]>>0)&255)<<6 | uint32((data[7]>>0)&127)<<14)
	a[3] = int32(uint32((data[7]>>7)&1)<<0 | uint32((data[8]>>0)&255)<<1 | uint32((data[9]>>0)&255)<<9 | uint32((data[10]>>0)&15)<<17)
	a[4] = int32(uint32((data[10]>>4)&15)<<0 | uint32((data[11]>>0)&255)<<4 | uint32((data[12]>>0)&255)<<12 | uint32((data[13]>>0)&1)<<20)
	a[5] = int32(uint32((data[13]>>1)&127)<<0 | uint32((data[14]>>0)&255)<<7 | uint32((data[15]>>0)&63)<<15)
	a[6] = int32(uint32((data[15]>>6)&3)<<0 | uint32((data[16]>>0)&255)<<2 | uint32((data[17]>>0)&255)<<10 | uint32((data[18]>>0)&7)<<18)
	a[7] = int32(uint32((data[18]>>3)&31)<<0 | uint32((data[19]>>0)&255)<<5 | uint32((data[20]>>0)&255)<<13)
	return
}

func pack8int32_21(data [8]int32) []byte {
	return []byte{
		byte(uint32(data[0]) << 0),
		byte(uint32(data[0]) >> 8),
		byte(uint32(data[0])>>16 | uint32(data[1])<<5),
		byte(uint32(data[1]) >> 3),
		byte(uint32(data[1]) >> 11),
		byte(uint32(data[1])>>19 | uint32(data[2])<<2),
		byte(uint32(data[2]) >> 6),
		byte(uint32(data[2])>>14 | uint32(data[3])<<7),
		byte(uint32(data[3]) >> 1),
		byte(uint32(data[3]) >> 9),
		byte(uint32(data[3])>>17 | uint32(data[4])<<4),
		byte(uint32(data[4]) >> 4),
		byte(uint32(data[4]) >> 12),
		byte(uint32(data[4])>>20 | uint32(data[5])<<1),
		byte(uint32(data[5]) >> 7),
		byte(uint32(data[5])>>15 | uint32(data[6])<<6),
		byte(uint32(data[6]) >> 2),
		byte(uint32(data[6]) >> 10),
		byte(uint32(data[6])>>18 | uint32(data[7])<<3),
		byte(uint32(data[7]) >> 5),
		byte(uint32(data[7]) >> 13),
	}
}

func unpack8int32_22(data []byte) (a [8]int32) {
	_ = data[21]
	a[0] = int32(uint32((data[0]>>0)&255)<<0 | uint32((data[1]>>0)&255)<<8 | uint32((data[2]>>0)&63)<<16)
	a[1] = int32(uint32((data[2]>>6)&3)<<0 | uint32((data[3]>>0)&255)<<2 | uint32((data[4]>>0)&255)<<10 | uint32((data[5]>>0)&15)<<18)
	a[2] = int32(uint32((data[5]>>4)&15)<<0 | uint32((data[6]>>0)&255)<<4 | uint32((data[7]>>0)&255)<<12 | uint32((data[8]>>0)&3)<<20)
	a[3] = int32(uint32((data[8]>>2)&63)<<0 | uint32((data[9]>>0)&255)<<6 | uint32((data[10]>>0)&255)<<14)
	a[4] = int32(uint32((data[11]>>0)&255)<<0 | uint32((data[12]>>0)&255)<<8 | uint32((data[13]>>0)&63)<<16)
	a[5] = int32(uint32((data[13]>>6)&3)<<0 | uint32((data[14]>>0)&255)<<2 | uint32((data[15]>>0)&255)<<10 | uint32((data[16]>>0)&15)<<18)
	a[6] = int32(uint32((data[16]>>4)&15)<<0 | uint32((data[17]>>0)&255)<<4 | uint32((data[18]>>0)&255)<<12 | uint32((data[19]>>0)&3)<<20)
	a[7] = int32(uint32((data[19]>>2)&63)<<0 | uint32((data[20]>>0)&255)<<6 | uint32((data[21]>>0)&255)<<14)
	return
}

func pack8int32_22(data [8]int32) []byte {
	return []byte{
		byte(uint32(data[0]) << 0),
		byte(uint32(data[0]) >> 8),
		byte(uint32(data[0])>>16 | uint32(data[1])<<6),
		byte(uint32(data[1]) >> 2),
		byte(uint32(data[1]) >> 10),
		byte(uint32(data[1])>>18 | uint32(data[2])<<4),
		byte(uint32(data[2]) >> 4),
		byte(uint32(data[2]) >> 12),
		byte(uint32(data[2])>>20 | uint32(data[3])<<2),
		byte(uint32(data[3]) >> 6),
		byte(uint32(data[3]) >> 14),
		byte(uint32(data[4]) << 0),
		byte(uint32(data[4]) >> 8),
		byte(uint32(data[4])>>16 | uint32(data[5])<<6),
		byte(uint32(data[5]) >> 2),
		byte(uint32(data[5]) >> 10),
		byte(uint32(data[5])>>18 | uint32(data[6])<<4),
		byte(uint32(data[6]) >> 4),
		byte(uint32(data[6]) >> 12),
		byte(uint32(data[6])>>20 | uint32(data[7])<<2),
		byte(uint32(data[7]) >> 6),
		byte(uint32(data[7]) >> 14),
	}
}

func unpack8int32_23(data []byte) (a [8]int32) {
	_ = data[22]
	a[0] = int32(uint32((data[0]>>0)&255)<<0 | uint32((data[1]>>0)&255)<<8 | uint32((data[2]>>0)&127)<<16)
	a[1] = int32(uint32((data[2]>>7)&1)<<0 | uint32((data[3]>>0)&255)<<1 | uint32((data[4]>>0)&255)<<9 | uint32((data[5]>>0)&63)<<17)
	a[2] = int32(uint32((data[5]>>6)&3)<<0 | uint32((data[6]>>0)&255)<<2 | uint32((data[7]>>0)&255)<<10 | uint32((data[8]>>0)&31)<<18)
	a[3] = int32(uint32((data[8]>>5)&7)<<0 | uint32((data[9]>>0)&255)<<3 | uint32((data[10]>>0)&255)<<11 | uint32((data[11]>>0)&15)<<19)
	a[4] = int32(uint32((data[11]>>4)&15)<<0 | uint32((data[12]>>0)&255)<<4 | uint32((data[13]>>0)&255)<<12 | uint32((data[14]>>0)&7)<<20)
	a[5] = int32(uint32((data[14]>>3)&31)<<0 | uint32((data[15]>>0)&255)<<5 | uint32((data[16]>>0)&255)<<13 | uint32((data[17]>>0)&3)<<21)
	a[6] = int32(uint32((data[17]>>2)&63)<<0 | uint32((data[18]>>0)&255)<<6 | uint32((data[19]>>0)&255)<<14 | uint32((data[20]>>0)&1)<<22)
	a[7] = int32(uint32((data[20]>>1)&127)<<0 | uint32((data[21]>>0)&255)<<7 | uint32((data[22]>>0)&255)<<15)
	return
}

func pack8int32_23(data [8]int32) []byte {
	return []byte{
		byte(uint32(data[0]) << 0),
		byte(uint32(data[0]) >> 8),
		byte(uint32(data[0])>>16 | uint32(data[1])<<7),
		byte(uint32(data[1]) >> 1),
		byte(uint32(data[1]) >> 9),
		byte(uint32(data[1])>>17 | uint32(data[2])<<6),
		byte(uint32(data[2]) >> 2),
		byte(uint32(data[2]) >> 10),
		byte(uint32(data[2])>>18 | uint32(data[3])<<5),
		byte(uint32(data[3]) >> 3),
		byte(uint32(data[3]) >> 11),
		byte(uint32(data[3])>>19 | uint32(data[4])<<4),
		byte(uint32(data[4]) >> 4),
		byte(uint32(data[4]) >> 12),
		byte(uint32(data[4])>>20 | uint32(data[5])<<3),
		byte(uint32(data[5]) >> 5),
		byte(uint32(data[5]) >> 13),
		byte(uint32(data[5])>>21 | uint32(data[6])<<2),
		byte(uint32(data[6]) >> 6),
		byte(uint32(data[6]) >> 14),
		byte(uint32(data[6])>>22 | uint32(data[7])<<1),
		byte(uint32(data[7]) >> 7),
		byte(uint32(data[7]) >> 15),
	}
}

func unpack8int32_24(data []byte) (a [8]int32) {
	_ = data[23]
	a[0] = int32(uint32((data[0]>>0)&255)<<0 | uint32((data[1]>>0)&255)<<8 | uint32((data[2]>>0)&255)<<16)
	a[1] = int32(uint32((data[3]>>0)&255)<<0 | uint32((data[4]>>0)&255)<<8 | uint32((data[5]>>0)&255)<<16)
	a[2] = int32(uint32((data[6]>>0)&255)<<0 | uint32((data[7]>>0)&255)<<8 | uint32((data[8]>>0)&255)<<16)
	a[3] = int32(uint32((data[9]>>0)&255)<<0 | uint32((data[10]>>0)&255)<<8 | uint32((data[11]>>0)&255)<<16)
	a[4] = int32(uint32((data[12]>>0)&255)<<0 | uint32((data[13]>>0)&255)<<8 | uint32((data[14]>>0)&255)<<16)
	a[5] = int32(uint32((data[15]>>0)&255)<<0 | uint32((data[16]>>0)&255)<<8 | uint32((data[17]>>0)&255)<<16)
	a[6] = int32(uint32((data[18]>>0)&255)<<0 | uint32((data[19]>>0)&255)<<8 | uint32((data[20]>>0)&255)<<16)
	a[7] = int32(uint32((data[21]>>0)&255)<<0 | uint32((data[22]>>0)&255)<<8 | uint32((data[23]>>0)&255)<<16)
	return
}

func pack8int32_24(data [8]int32) []byte {
	return []byte{
		byte(uint32(data[0]) << 0),
		byte(uint32(data[0]) >> 8),
		byte(uint32(data[0]) >> 16),
		byte(uint32(data[1]) << 0),
		byte(uint32(data[1]) >> 8),
		byte(uint32(data[1]) >> 16),
		byte(uint32(data[2]) << 0),
		byte(uint32(data[2]) >> 8),
		byte(uint32(data[2]) >> 16),
		byte(uint32(data[3]) << 0),
		byte(uint32(data[3]) >> 8),
		byte(uint32(data[3]) >> 16),
		byte(uint32(data[4]) << 0),
		byte(uint32(data[4]) >> 8),
		byte(uint32(data[4]) >> 16),
		byte(uint32(data[5]) << 0),
		byte(uint32(data[5]) >> 8),
		byte(uint32(data[5]) >> 16),
		byte(uint32(data[6]) << 0),
		byte(uint32(data[6]) >> 8),
		byte(uint32(data[6]) >> 16),
		byte(uint32(data[7]) << 0),
		byte(uint32(data[7]) >> 8),
		byte(uint32(data[7]) >> 16),
	}
}

func unpack8int32_25(data []byte) (a [8]int32) {
	_ = data[24]
	a[0] = int32(uint32((data[0]>>0)&255)<<0 | uint32((data[1]>>0)&255)<<8 | uint32((data[2]>>0)&255)<<16 | uint32((data[3]>>0)&1)<<24)
	a[1] = int32(uint32((data[3]>>1)&127)<<0 | uint32((data[4]>>0)&255)<<7 | uint32((data[5]>>0)&255)<<15 | uint32((data[6]>>0)&3)<<23)
	a[2] = int32(uint32((data[6]>>2)&63)<<0 | uint32((data[7]>>0)&255)<<6 | uint32((data[8]>>0)&255)<<14 | uint32((data[9]>>0)&7)<<22)
	a[3] = int32(uint32((data[9]>>3)&31)<<0 | uint32((data[10]>>0)&255)<<5 | uint32((data[11]>>0)&255)<<13 | uint32((data[12]>>0)&15)<<21)
	a[4] = int32(uint32((data[12]>>4)&15)<<0 | uint32((data[13]>>0)&255)<<4 | uint32((data[14]>>0)&255)<<12 | uint32((data[15]>>0)&31)<<20)
	a[5] = int32(uint32((data[15]>>5)&7)<<0 | uint32((data[16]>>0)&255)<<3 | uint32((data[17]>>0)&255)<<11 | uint32((data[18]>>0)&63)<<19)
	a[6] = int32(uint32((data[18]>>6)&3)<<0 | uint32((data[19]>>0)&255)<<2 | uint32((data[20]>>0)&255)<<10 | uint32((data[21]>>0)&127)<<18)
	a[7] = int32(uint32((data[21]>>7)&1)<<0 | uint32((data[22]>>0)&255)<<1 | uint32((data[23]>>0)&255)<<9 | uint32((data[24]>>0)&255)<<17)
	return
}

func pack8int32_25(data [8]int32) []byte {
	return []byte{
		byte(uint32(data[0]) << 0),
		byte(uint32(data[0]) >> 8),
		byte(uint32(data[0]) >> 16),
		byte(uint32(data[0])>>24 | uint32(data[1])<<1),
		byte(uint32(data[1]) >> 7),
		byte(uint32(data[1]) >> 15),
		byte(uint32(data[1])>>23 | uint32(data[2])<<2),
		byte(uint32(data[2]) >> 6),
		byte(uint32(data[2]) >> 14),
		byte(uint32(data[2])>>22 | uint32(data[3])<<3),
		byte(uint32(data[3]) >> 5),
		byte(uint32(data[3]) >> 13),
		byte(uint32(data[3])>>21 | uint32(data[4])<<4),
		byte(uint32(data[4]) >> 4),
		byte(uint32(data[4]) >> 12),
		byte(uint32(data[4])>>20 | uint32(data[5])<<5),
		byte(uint32(data[5]) >> 3),
		byte(uint32(data[5]) >> 11),
		byte(uint32(data[5])>>19 | uint32(data[6])<<6),
		byte(uint32(data[6]) >> 2),
		byte(uint32(data[6]) >> 10),
		byte(uint32(data[6])>>18 | uint32(data[7])<<7),
		byte(uint32(data[7]) >> 1),
		byte(uint32(data[7]) >> 9),
		byte(uint32(data[7]) >> 17),
	}
}

func unpack8int32_26(data []byte) (a [8]int32) {
	_ = data[25]
	a[0] = int32(uint32((data[0]>>0)&255)<<0 | uint32((data[1]>>0)&255)<<8 | uint32((data[2]>>0)&255)<<16 | uint32((data[3]>>0)&3)<<24)
	a[1] = int32(uint32((data[3]>>2)&63)<<0 | uint32((data[4]>>0)&255)<<6 | uint32((data[5]>>0)&255)<<14 | uint32((data[6]>>0)&15)<<22)
	a[2] = int32(uint32((data[6]>>4)&15)<<0 | uint32((data[7]>>0)&255)<<4 | uint32((data[8]>>0)&255)<<12 | uint32((data[9]>>0)&63)<<20)
	a[3] = int32(uint32((data[9]>>6)&3)<<0 | uint32((data[10]>>0)&255)<<2 | uint32((data[11]>>0)&255)<<10 | uint32((data[12]>>0)&255)<<18)
	a[4] = int32(uint32((data[13]>>0)&255)<<0 | uint32((data[14]>>0)&255)<<8 | uint32((data[15]>>0)&255)<<16 | uint32((data[16]>>0)&3)<<24)
	a[5] = int32(uint32((data[16]>>2)&63)<<0 | uint32((data[17]>>0)&255)<<6 | uint32((data[18]>>0)&255)<<14 | uint32((data[19]>>0)&15)<<22)
	a[6] = int32(uint32((data[19]>>4)&15)<<0 | uint32((data[20]>>0)&255)<<4 | uint32((data[21]>>0)&255)<<12 | uint32((data[22]>>0)&63)<<20)
	a[7] = int32(uint32((data[22]>>6)&3)<<0 | uint32((data[23]>>0)&255)<<2 | uint32((data[24]>>0)&255)<<10 | uint32((data[25]>>0)&255)<<18)
	return
}

func pack8int32_26(data [8]int32) []byte {
	return []byte{
		byte(uint32(data[0]) << 0),
		byte(uint32(data[0]) >> 8),
		byte(uint32(data[0]) >> 16),
		byte(uint32(data[0])>>24 | uint32(data[1])<<2),
		byte(uint32(data[1]) >> 6),
		byte(uint32(data[1]) >> 14),
		byte(uint32(data[1])>>22 | uint32(data[2])<<4),
		byte(uint32(data[2]) >> 4),
		byte(uint32(data[2]) >> 12),
		byte(uint32(data[2])>>20 | uint32(data[3])<<6),
		byte(uint32(data[3]) >> 2),
		byte(uint32(data[3]) >> 10),
		byte(uint32(data[3]) >> 18),
		byte(uint32(data[4]) << 0),
		byte(uint32(data[4]) >> 8),
		byte(uint32(data[4]) >> 16),
		byte(uint32(data[4])>>24 | uint32(data[5])<<2),
		byte(uint32(data[5]) >> 6),
		byte(uint32(data[5]) >> 14),
		byte(uint32(data[5])>>22 | uint32(data[6])<<4),
		byte(uint32(data[6]) >> 4),
		byte(uint32(data[6]) >> 12),
		byte(uint32(data[6])>>20 | uint32(data[7])<<6),
		byte(uint32(data[7]) >> 2),
		byte(uint32(data[7]) >> 10),
		byte(uint32(data[7]) >> 18),
	}
}

func unpack8int32_27(data []byte) (a [8]int32) {
	_ = data[26]
	a[0] = int32(uint32((data[0]>>0)&255)<<0 | uint32((data[1]>>0)&255)<<8 | uint32((data[2]>>0)&255)<<16 | uint32((data[3]>>0)&7)<<24)
	a[1] = int32(uint32((data[3]>>3)&31)<<0 | uint32((data[4]>>0)&255)<<5 | uint32((data[5]>>0)&255)<<13 | uint32((data[6]>>0)&63)<<21)
	a[2] = int32(uint32((data[6]>>6)&3)<<0 | uint32((data[7]>>0)&255)<<2 | uint32((data[8]>>0)&255)<<10 | uint32((data[9]>>0)&255)<<18 | uint32((data[10]>>0)&1)<<26)
	a[3] = int32(uint32((data[10]>>1)&127)<<0 | uint32((data[11]>>0)&255)<<7 | uint32((data[12]>>0)&255)<<15 | uint32((data[13]>>0)&15)<<23)
	a[4] = int32(uint32((data[13]>>4)&15)<<0 | uint32((data[14]>>0)&255)<<4 | uint32((data[15]>>0)&255)<<12 | uint32((data[16]>>0)&127)<<20)
	a[5] = int32(uint32((data[16]>>7)&1)<<0 | uint32((data[17]>>0)&255)<<1 | uint32((data[18]>>0)&255)<<9 | uint32((data[19]>>0)&255)<<17 | uint32((data[20]>>0)&3)<<25)
	a[6] = int32(uint32((data[20]>>2)&63)<<0 | uint32((data[21]>>0)&255)<<6 | uint32((data[22]>>0)&255)<<14 | uint32((data[23]>>0)&31)<<22)
	a[7] = int32(uint32((data[23]>>5)&7)<<0 | uint32((data[24]>>0)&255)<<3 | uint32((data[25]>>0)&255)<<11 | uint32((data[26]>>0)&255)<<19)
	return
}

func pack8int32_27(data [8]int32) []byte {
	return []byte{
		byte(uint32(data[0]) << 0),
		byte(uint32(data[0]) >> 8),
		byte(uint32(data[0]) >> 16),
		byte(uint32(data[0])>>24 | uint32(data[1])<<3),
		byte(uint32(data[1]) >> 5),
		byte(uint32(data[1]) >> 13),
		byte(uint32(data[1])>>21 | uint32(data[2])<<6),
		byte(uint32(data[2]) >> 2),
		byte(uint32(data[2]) >> 10),
		byte(uint32(data[2]) >> 18),
		byte(uint32(data[2])>>26 | uint32(data[3])<<1),
		byte(uint32(data[3]) >> 7),
		byte(uint32(data[3]) >> 15),
		byte(uint32(data[3])>>23 | uint32(data[4])<<4),
		byte(uint32(data[4]) >> 4),
		byte(uint32(data[4]) >> 12),
		byte(uint32(data[4])>>20 | uint32(data[5])<<7),
		byte(uint32(data[5]) >> 1),
		byte(uint32(data[5]) >> 9),
		byte(uint32(data[5]) >> 17),
		byte(uint32(data[5])>>25 | uint32(data[6])<<2),
		byte(uint32(data[6]) >> 6),
		byte(uint32(data[6]) >> 14),
		byte(uint32(data[6])>>22 | uint32(data[7])<<5),
		byte(uint32(data[7]) >> 3),
		byte(uint32(data[7]) >> 11),
		byte(uint32(data[7]) >> 19),
	}
}

func unpack8int32_28(data []byte) (a [8]int32) {
	_ = data[27]
	a[0] = int32(uint32((data[0]>>0)&255)<<0 | uint32((data[1]>>0)&255)<<8 | uint32((data[2]>>0)&255)<<16 | uint32((data[3]>>0)&15)<<24)
	a[1] = int32(uint32((data[3]>>4)&15)<<0 | uint32((data[4]>>0)&255)<<4 | uint32((data[5]>>0)&255)<<12 | uint32((data[6]>>0)&255)<<20)
	a[2] = int32(uint32((data[7]>>0)&255)<<0 | uint32((data[8]>>0)&255)<<8 | uint32((data[9]>>0)&255)<<16 | uint32((data[10]>>0)&15)<<24)
	a[3] = int32(uint32((data[10]>>4)&15)<<0 | uint32((data[11]>>0)&255)<<4 | uint32((data[12]>>0)&255)<<12 | uint32((data[13]>>0)&255)<<20)
	a[4] = int32(uint32((data[14]>>0)&255)<<0 | uint32((data[15]>>0)&255)<<8 | uint32((data[16]>>0)&255)<<16 | uint32((data[17]>>0)&15)<<24)
	a[5] = int32(uint32((data[17]>>4)&15)<<0 | uint32((data[18]>>0)&255)<<4 | uint32((data[19]>>0)&255)<<12 | uint32((data[20]>>0)&255)<<20)
	a[6] = int32(uint32((data[21]>>0)&255)<<0 | uint32((data[22]>>0)&255)<<8 | uint32((data[23]>>0)&255)<<16 | uint32((data[24]>>0)&15)<<24)
	a[7] = int32(uint32((data[24]>>4)&15)<<0 | uint32((data[25]>>0)&255)<<4 | uint32((data[26]>>0)&255)<<12 | uint32((data[27]>>0)&255)<<20)
	return
}

func pack8int32_28(data [8]int32) []byte {
	return []byte{
		byte(uint32(data[0]) << 0),
		byte(uint32(data[0]) >> 8),
		byte(uint32(data[0]) >> 16),
		byte(uint32(data[0])>>24 | uint32(data[1])<<4),
		byte(uint32(data[1]) >> 4),
		byte(uint32(data[1]) >> 12),
		byte(uint32(data[1]) >> 20),
		byte(uint32(data[2]) << 0),
		byte(uint32(data[2]) >> 8),
		byte(uint32(data[2]) >> 16),
		byte(uint32(data[2])>>24 | uint32(data[3])<<4),
		byte(uint32(data[3]) >> 4),
		byte(uint32(data[3]) >> 12),
		byte(uint32(data[3]) >> 20),
		byte(uint32(data[4]) << 0),
		byte(uint32(data[4]) >> 8),
		byte(uint32(data[4]) >> 16),
		byte(uint32(data[4])>>24 | uint32(data[5])<<4),
		byte(uint32(data[5]) >> 4),
		byte(uint32(data[5]) >> 12),
		byte(uint32(data[5]) >> 20),
		byte(uint32(data[6]) << 0),
		byte(uint32(data[6]) >> 8),
		byte(uint32(data[6]) >> 16),
		byte(uint32(data[6])>>24 | uint32(data[7])<<4),
		byte(uint32(data[7]) >> 4),
		byte(uint32(data[7]) >> 12),
		byte(uint32(data[7]) >> 20),
	}
}

func unpack8int32_29(data []byte) (a [8]int32) {
	_ = data[28]
	a[0] = int32(uint32((data[0]>>0)&255)<<0 | uint32((data[1]>>0)&255)<<8 | uint32((data[2]>>0)&255)<<16 | uint32((data[3]>>0)&31)<<24)
	a[1] = int32(uint32((data[3]>>5)&7)<<0 | uint32((data[4]>>0)&255)<<3 | uint32((data[5]>>0)&255)<<11 | uint32((data[6]>>0)&255)<<19 | uint32((data[7]>>0)&3)<<27)
	a[2] = int32(uint32((data[7]>>2)&63)<<0 | uint32((data[8]>>0)&255)<<6 | uint32((data[9]>>0)&255)<<14 | uint32((data[10]>>0)&127)<<22)
	a[3] = int32(uint32((data[10]>>7)&1)<<0 | uint32((data[11]>>0)&255)<<1 | uint32((data[12]>>0)&255)<<9 | uint32((data[13]>>0)&255)<<17 | uint32((data[14]>>0)&15)<<25)
	a[4] = int32(uint32((data[14]>>4)&15)<<0 | uint32((data[15]>>0)&255)<<4 | uint32((data[16]>>0)&255)<<12 | uint32((data[17]>>0)&255)<<20 | uint32((data[18]>>0)&1)<<28)
	a[5] = int32(uint32((data[18]>>1)&127)<<0 | uint32((data[19]>>0)&255)<<7 | uint32((data[20]>>0)&255)<<15 | uint32((data[21]>>0)&63)<<23)
	a[6] = int32(uint32((data[21]>>6)&3)<<0 | uint32((data[22]>>0)&255)<<2 | uint32((data[23]>>0)&255)<<10 | uint32((data[24]>>0)&255)<<18 | uint32((data[25]>>0)&7)<<26)
	a[7] = int32(uint32((data[25]>>3)&31)<<0 | uint32((data[26]>>0)&255)<<5 | uint32((data[27]>>0)&255)<<13 | uint32((data[28]>>0)&255)<<21)
	return
}

func pack8int32_29(data [8]int32) []byte {
	return []byte{
		byte(uint32(data[0]) << 0),
		byte(uint32(data[0]) >> 8),
		byte(uint32(data[0]) >> 16),
		byte(uint32(data[0])>>24 | uint32(data[1])<<5),
		byte(uint32(data[1]) >> 3),
		byte(uint32(data[1]) >> 11),
		byte(uint32(data[1]) >> 19),
		byte(uint32(data[1])>>27 | uint32(data[2])<<2),
		byte(uint32(data[2]) >> 6),
		byte(uint32(data[2]) >> 14),
		byte(uint32(data[2])>>22 | uint32(data[3])<<7),
		byte(uint32(data[3]) >> 1),
		byte(uint32(data[3]) >> 9),
		byte(uint32(data[3]) >> 17),
		byte(uint32(data[3])>>25 | uint32(data[4])<<4),
		byte(uint32(data[4]) >> 4),
		byte(uint32(data[4]) >> 12),
		byte(uint32(data[4]) >> 20),
		byte(uint32(data[4])>>28 | uint32(data[5])<<1),
		byte(uint32(data[5]) >> 7),
		byte(uint32(data[5]) >> 15),
		byte(uint32(data[5])>>23 | uint32(data[6])<<6),
		byte(uint32(data[6]) >> 2),
		byte(uint32(data[6]) >> 10),
		byte(uint32(data[6]) >> 18),
		byte(uint32(data[6])>>26 | uint32(data[7])<<3),
		byte(uint32(data[7]) >> 5),
		byte(uint32(data[7]) >> 13),
		byte(uint32(data[7]) >> 21),
	}
}

func unpack8int32_30(data []byte) (a [8]int32) {
	_ = data[29]
	a[0] = int32(uint32((data[0]>>0)&255)<<0 | uint32((data[1]>>0)&255)<<8 | uint32((data[2]>>0)&255)<<16 | uint32((data[3]>>0)&63)<<24)
	a[1] = int32(uint32((data[3]>>6)&3)<<0 | uint32((data[4]>>0)&255)<<2 | uint32((data[5]>>0)&255)<<10 | uint32((data[6]>>0)&255)<<18 | uint32((data[7]>>0)&15)<<26)
	a[2] = int32(uint32((data[7]>>4)&15)<<0 | uint32((data[8]>>0)&255)<<4 | uint32((data[9]>>0)&255)<<12 | uint32((data[10]>>0)&255)<<20 | uint32((data[11]>>0)&3)<<28)
	a[3] = int32(uint32((data[11]>>2)&63)<<0 | uint32((data[12]>>0)&255)<<6 | uint32((data[13]>>0)&255)<<14 | uint32((data[14]>>0)&255)<<22)
	a[4] = int32(uint32((data[15]>>0)&255)<<0 | uint32((data[16]>>0)&255)<<8 | uint32((data[17]>>0)&255)<<16 | uint32((data[18]>>0)&63)<<24)
	a[5] = int32(uint32((data[18]>>6)&3)<<0 | uint32((data[19]>>0)&255)<<2 | uint32((data[20]>>0)&255)<<10 | uint32((data[21]>>0)&255)<<18 | uint32((data[22]>>0)&15)<<26)
	a[6] = int32(uint32((data[22]>>4)&15)<<0 | uint32((data[23]>>0)&255)<<4 | uint32((data[24]>>0)&255)<<12 | uint32((data[25]>>0)&255)<<20 | uint32((data[26]>>0)&3)<<28)
	a[7] = int32(uint32((data[26]>>2)&63)<<0 | uint32((data[27]>>0)&255)<<6 | uint32((data[28]>>0)&255)<<14 | uint32((data[29]>>0)&255)<<22)
	return
}

func pack8int32_30(data [8]int32) []byte {
	return []byte{
		byte(uint32(data[0]) << 0),
		byte(uint32(data[0]) >> 8),
		byte(uint32(data[0]) >> 16),
		byte(uint32(data[0])>>24 | uint32(data[1])<<6),
		byte(uint32(data[1]) >> 2),
		byte(uint32(data[1]) >> 10),
		byte(uint32(data[1]) >> 18),
		byte(uint32(data[1])>>26 | uint32(data[2])<<4),
		byte(uint32(data[2]) >> 4),
		byte(uint32(data[2]) >> 12),
		byte(uint32(data[2]) >> 20),
		byte(uint32(data[2])>>28 | uint32(data[3])<<2),
		byte(uint32(data[3]) >> 6),
		byte(uint32(data[3]) >> 14),
		byte(uint32(data[3]) >> 22),
		byte(uint32(data[4]) << 0),
		byte(uint32(data[4]) >> 8),
		byte(uint32(data[4]) >> 16),
		byte(uint32(data[4])>>24 | uint32(data[5])<<6),
		byte(uint32(data[5]) >> 2),
		byte(uint32(data[5]) >> 10),
		byte(uint32(data[5]) >> 18),
		byte(uint32(data[5])>>26 | uint32(data[6])<<4),
		byte(uint32(data[6]) >> 4),
		byte(uint32(data[6]) >> 12),
		byte(uint32(data[6]) >> 20),
		byte(uint32(data[6])>>28 | uint32(data[7])<<2),
		byte(uint32(data[7]) >> 6),
		byte(uint32(data[7]) >> 14),
		byte(uint32(data[7]) >> 22),
	}
}

func unpack8int32_31(data []byte) (a [8]int32) {
	_ = data[30]
	a[0] = int32(uint32((data[0]>>0)&255)<<0 | uint32((data[1]>>0)&255)<<8 | uint32((data[2]>>0)&255)<<16 | uint32((data[3]>>0)&127)<<24)
	a[1] = int32(uint32((data[3]>>7)&1)<<0 | uint32((data[4]>>0)&255)<<1 | uint32((data[5]>>0)&255)<<9 | uint32((data[6]>>0)&255)<<17 | uint32((data[7]>>0)&63)<<25)
	a[2] = int32(uint32((data[7]>>6)&3)<<0 | uint32((data[8]>>0)&255)<<2 | uint32((data[9]>>0)&255)<<10 | uint32((data[10]>>0)&255)<<18 | uint32((data[11]>>0)&31)<<26)
	a[3] = int32(uint32((data[11]>>5)&7)<<0 | uint32((data[12]>>0)&255)<<3 | uint32((data[13]>>0)&255)<<11 | uint32((data[14]>>0)&255)<<19 | uint32((data[15]>>0)&15)<<27)
	a[4] = int32(uint32((data[15]>>4)&15)<<0 | uint32((data[16]>>0)&255)<<4 | uint32((data[17]>>0)&255)<<12 | uint32((data[18]>>0)&255)<<20 | uint32((data[19]>>0)&7)<<28)
	a[5] = int32(uint32((data[19]>>3)&31)<<0 | uint32((data[20]>>0)&255)<<5 | uint32((data[21]>>0)&255)<<13 | uint32((data[22]>>0)&255)<<21 | uint32((data[23]>>0)&3)<<29)
	a[6] = int32(uint32((data[23]>>2)&63)<<0 | uint32((data[24]>>0)&255)<<6 | uint32((data[25]>>0)&255)<<14 | uint32((data[26]>>0)&255)<<22 | uint32((data[27]>>0)&1)<<30)
	a[7] = int32(uint32((data[27]>>1)&127)<<0 | uint32((data[28]>>0)&255)<<7 | uint32((data[29]>>0)&255)<<15 | uint32((data[30]>>0)&255)<<23)
	return
}

func pack8int32_31(data [8]int32) []byte {
	return []byte{
		byte(uint32(data[0]) << 0),
		byte(uint32(data[0]) >> 8),
		byte(uint32(data[0]) >> 16),
		byte(uint32(data[0])>>24 | uint32(data[1])<<7),
		byte(uint32(data[1]) >> 1),
		byte(uint32(data[1]) >> 9),
		byte(uint32(data[1]) >> 17),
		byte(uint32(data[1])>>25 | uint32(data[2])<<6),
		byte(uint32(data[2]) >> 2),
		byte(uint32(data[2]) >> 10),
		byte(uint32(data[2]) >> 18),
		byte(uint32(data[2])>>26 | uint32(data[3])<<5),
		byte(uint32(data[3]) >> 3),
		byte(uint32(data[3]) >> 11),
		byte(uint32(data[3]) >> 19),
		byte(uint32(data[3])>>27 | uint32(data[4])<<4),
		byte(uint32(data[4]) >> 4),
		byte(uint32(data[4]) >> 12),
		byte(uint32(data[4]) >> 20),
		byte(uint32(data[4])>>28 | uint32(data[5])<<3),
		byte(uint32(data[5]) >> 5),
		byte(uint32(data[5]) >> 13),
		byte(uint32(data[5]) >> 21),
		byte(uint32(data[5])>>29 | uint32(data[6])<<2),
		byte(uint32(data[6]) >> 6),
		byte(uint32(data[6]) >> 14),
		byte(uint32(data[6]) >> 22),
		byte(uint32(data[6])>>30 | uint32(data[7])<<1),
		byte(uint32(data[7]) >> 7),
		byte(uint32(data[7]) >> 15),
		byte(uint32(data[7]) >> 23),
	}
}

func unpack8int32_32(data []byte) (a [8]int32) {
	_ = data[31]
	a[0] = int32(uint32((data[0]>>0)&255)<<0 | uint32((data[1]>>0)&255)<<8 | uint32((data[2]>>0)&255)<<16 | uint32((data[3]>>0)&255)<<24)
	a[1] = int32(uint32((data[4]>>0)&255)<<0 | uint32((data[5]>>0)&255)<<8 | uint32((data[6]>>0)&255)<<16 | uint32((data[7]>>0)&255)<<24)
	a[2] = int32(uint32((data[8]>>0)&255)<<0 | uint32((data[9]>>0)&255)<<8 | uint32((data[10]>>0)&255)<<16 | uint32((data[11]>>0)&255)<<24)
	a[3] = int32(uint32((data[12]>>0)&255)<<0 | uint32((data[13]>>0)&255)<<8 | uint32((data[14]>>0)&255)<<16 | uint32((data[15]>>0)&255)<<24)
	a[4] = int32(uint32((data[16]>>0)&255)<<0 | uint32((data[17]>>0)&255)<<8 | uint32((data[18]>>0)&255)<<16 | uint32((data[19]>>0)&255)<<24)
	a[5] = int32(uint32((data[20]>>0)&255)<<0 | uint32((data[21]>>0)&255)<<8 | uint32((data[22]>>0)&255)<<16 | uint32((data[23]>>0)&255)<<24)
	a[6] = int32(uint32((data[24]>>0)&255)<<0 | uint32((data[25]>>0)&255)<<8 | uint32((data[26]>>0)&255)<<16 | uint32((data[27]>>0)&255)<<24)
	a[7] = int32(uint32((data[28]>>0)&255)<<0 | uint32((data[29]>>0)&255)<<8 | uint32((data[30]>>0)&255)<<16 | uint32((data[31]>>0)&255)<<24)
	return
}

func pack8int32_32(data [8]int32) []byte {
	return []byte{
		byte(uint32(data[0]) << 0),
		byte(uint32(data[0]) >> 8),
		byte(uint32(data[0]) >> 16),
		byte(uint32(data[0]) >> 24),
		byte(uint32(data[1]) << 0),
		byte(uint32(data[1]) >> 8),
		byte(uint32(data[1]) >> 16),
		byte(uint32(data[1]) >> 24),
		byte(uint32(data[2]) << 0),
		byte(uint32(data[2]) >> 8),
		byte(uint32(data[2]) >> 16),
		byte(uint32(data[2]) >> 24),
		byte(uint32(data[3]) << 0),
		byte(uint32(data[3]) >> 8),
		byte(uint32(data[3]) >> 16),
		byte(uint32(data[3]) >> 24),
		byte(uint32(data[4]) << 0),
		byte(uint32(data[4]) >> 8),
		byte(uint32(data[4]) >> 16),
		byte(uint32(data[4]) >> 24),
		byte(uint32(data[5]) << 0),
		byte(uint32(data[5]) >> 8),
		byte(uint32(data[5]) >> 16),
		byte(uint32(data[5]) >> 24),
		byte(uint32(data[6]) << 0),
		byte(uint32(data[6]) >> 8),
		byte(uint32(data[6]) >> 16),
		byte(uint32(data[6]) >> 24),
		byte(uint32(data[7]) << 0),
		byte(uint32(data[7]) >> 8),
		byte(uint32(data[7]) >> 16),
		byte(uint32(data[7]) >> 24),
	}
}
