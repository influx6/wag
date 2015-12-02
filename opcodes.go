package wag

const (
	opNop                 = 0x00
	opBlock               = 0x01
	opLoop                = 0x02
	opIf                  = 0x03
	opIfElse              = 0x04
	opSelect              = 0x05
	opBr                  = 0x06
	opBrIf                = 0x07
	opTableswitch         = 0x08
	opI8_Const            = 0x09
	opI32_Const           = 0x0a
	opI64_Const           = 0x0b
	opF64_Const           = 0x0c
	opF32_Const           = 0x0d
	opGetLocal            = 0x0e
	opSetLocal            = 0x0f
	opGetGlobal           = 0x10
	opSetGlobal           = 0x11
	opCall                = 0x12
	opCallIndirect        = 0x13
	opReturn              = 0x14
	opUnreachable         = 0x15
	opI32_LoadMem8S       = 0x20
	opI32_LoadMem8U       = 0x21
	opI32_LoadMem16S      = 0x22
	opI32_LoadMem16U      = 0x23
	opI64_LoadMem8S       = 0x24
	opI64_LoadMem8U       = 0x25
	opI64_LoadMem16S      = 0x26
	opI64_LoadMem16U      = 0x27
	opI64_LoadMem32S      = 0x28
	opI64_LoadMem32U      = 0x29
	opI32_LoadMem         = 0x2a
	opI64_LoadMem         = 0x2b
	opF32_LoadMem         = 0x2c
	opF64_LoadMem         = 0x2d
	opI32_StoreMem8       = 0x2e
	opI32_StoreMem16      = 0x2f
	opI64_StoreMem8       = 0x30
	opI64_StoreMem16      = 0x31
	opI64_StoreMem32      = 0x32
	opI32_StoreMem        = 0x33
	opI64_StoreMem        = 0x34
	opF32_StoreMem        = 0x35
	opF64_StoreMem        = 0x36
	opResizeMemory_I32    = 0x39
	opResizeMemory_I64    = 0x3a
	opI32_Add             = 0x40
	opI32_Sub             = 0x41
	opI32_Mul             = 0x42
	opI32_SDiv            = 0x43
	opI32_UDiv            = 0x44
	opI32_SRem            = 0x45
	opI32_URem            = 0x46
	opI32_AND             = 0x47
	opI32_OR              = 0x48
	opI32_XOR             = 0x49
	opI32_SHL             = 0x4a
	opI32_SHR             = 0x4b
	opI32_SAR             = 0x4c
	opI32_EQ              = 0x4d
	opI32_NE              = 0x4e
	opI32_SLT             = 0x4f
	opI32_SLE             = 0x50
	opI32_ULT             = 0x51
	opI32_ULE             = 0x52
	opI32_SGT             = 0x53
	opI32_SGE             = 0x54
	opI32_UGT             = 0x55
	opI32_UGE             = 0x56
	opI32_CLZ             = 0x57
	opI32_CTZ             = 0x58
	opI32_PopCnt          = 0x59
	opI32_NOT             = 0x5a
	opI64_Add             = 0x5b
	opI64_Sub             = 0x5c
	opI64_Mul             = 0x5d
	opI64_SDiv            = 0x5e
	opI64_UDiv            = 0x5f
	opI64_SRem            = 0x60
	opI64_URem            = 0x61
	opI64_AND             = 0x62
	opI64_OR              = 0x63
	opI64_XOR             = 0x64
	opI64_SHL             = 0x65
	opI64_SHR             = 0x66
	opI64_SAR             = 0x67
	opI64_EQ              = 0x68
	opI64_NE              = 0x69
	opI64_SLT             = 0x6a
	opI64_SLE             = 0x6b
	opI64_ULT             = 0x6c
	opI64_ULE             = 0x6d
	opI64_SGT             = 0x6e
	opI64_SGE             = 0x6f
	opI64_UGT             = 0x70
	opI64_UGE             = 0x71
	opI64_CLZ             = 0x72
	opI64_CTZ             = 0x73
	opI64_PopCnt          = 0x74
	opF32_Add             = 0x75
	opF32_Sub             = 0x76
	opF32_Mul             = 0x77
	opF32_Div             = 0x78
	opF32_Min             = 0x79
	opF32_Max             = 0x7a
	opF32_Abs             = 0x7b
	opF32_Neg             = 0x7c
	opF32_CopySign        = 0x7d
	opF32_Ceil            = 0x7e
	opF32_Floor           = 0x7f
	opF32_Trunc           = 0x80
	opF32_Nearest         = 0x81
	opF32_Sqrt            = 0x82
	opF32_EQ              = 0x83
	opF32_NE              = 0x84
	opF32_LT              = 0x85
	opF32_LE              = 0x86
	opF32_GT              = 0x87
	opF32_GE              = 0x88
	opF64_Add             = 0x89
	opF64_Sub             = 0x8a
	opF64_Mul             = 0x8b
	opF64_Div             = 0x8c
	opF64_Min             = 0x8d
	opF64_Max             = 0x8e
	opF64_Abs             = 0x8f
	opF64_Neg             = 0x90
	opF64_CopySign        = 0x91
	opF64_Ceil            = 0x92
	opF64_Floor           = 0x93
	opF64_Trunc           = 0x94
	opF64_Nearest         = 0x95
	opF64_Sqrt            = 0x96
	opF64_EQ              = 0x97
	opF64_NE              = 0x98
	opF64_LT              = 0x99
	opF64_LE              = 0x9a
	opF64_GT              = 0x9b
	opF64_GE              = 0x9c
	opI32_SConvert_F32    = 0x9d
	opI32_SConvert_F64    = 0x9e
	opI32_UConvert_F32    = 0x9f
	opI32_UConvert_F64    = 0xa0
	opI32_Convert_I64     = 0xa1
	opI64_SConvert_F32    = 0xa2
	opI64_SConvert_F64    = 0xa3
	opI64_UConvert_F32    = 0xa4
	opI64_UConvert_F64    = 0xa5
	opI64_SConvert_I32    = 0xa6
	opI64_UConvert_I32    = 0xa7
	opF32_SConvert_I32    = 0xa8
	opF32_UConvert_I32    = 0xa9
	opF32_SConvert_I64    = 0xaa
	opF32_UConvert_I64    = 0xab
	opF32_Convert_F64     = 0xac
	opF32_Reinterpret_I32 = 0xad
	opF64_SConvert_I32    = 0xae
	opF64_UConvert_I32    = 0xaf
	opF64_SConvert_I64    = 0xb0
	opF64_UConvert_I64    = 0xb1
	opF64_Convert_F32     = 0xb2
	opF64_Reinterpret_I64 = 0xb3
	opI32_Reinterpret_F32 = 0xb4
	opI64_Reinterpret_F64 = 0xb5
)
