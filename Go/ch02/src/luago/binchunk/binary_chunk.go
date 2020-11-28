package binchunk

//定义binaryChunk结构体
type binaryChunk struct {
	header  //头部
	sizeUpvalues byte //主函数upvalue数量
	mainFunc     *Prototype //主函数原型
}

//定义header
type header struct {
	signature       [4]byte //签名，用于快速识别文件格式
	version         byte    //记录二进制chunk文件对应的Lua版本号
	format          byte    //格式号
	luacData        [6]byte //进一步校验
	cintSize        byte    //记录cint类型所占字节
	sizetSize       byte    //记录size_t类型所占字节
	instructionSize byte    //记录Lua虚拟机指令所占字节
	luaIntegerSize  byte    //记录整数所占字节
	luaNumberSize   byte    //记录浮点数所占字节
	luacInt         int64   //记录大小端方式
	luacNum         float64 //记录浮点数格式
}

const (
	LUA_SIGNATURE    = "\x1bLua"
	LUAC_VERSION     = 0x53
	LUAC_FORMAT      = 0
	LUAC_DATA        = "\x19\x93\r\n\x1a\n"
	CINT_SIZE        = 4
	CSZIET_SIZE      = 8
	INSTRUCTION_SIZE = 4
	LUA_INTEGER_SIZE = 8
	LUA_NUMBER_SIZE  = 8
	LUAC_INT         = 0x5678
	LUAC_NUM         = 370.5
)

//定义Prototype
type Prototype struct{
	Source string //源文件名，记录二进制chunk是由哪个源文件编译的，仅在主函数原型里生效
	LineDefined uint32
	LastLineDefined uint32
	NumParams byte
	IsVararg byte
	MaxStackSize byte
	Code []uint32
	Constants []interface{}
	Upvalues []Upvalue
	Protos []*Prototype
	LineInfo []uint32
	LocVars []LocVar
	UpvalueNames []string
}

const(
	TAG_NIL =0x00
	TAG_BOOLEAN = 0x01
	TAG_NUMBER = 0x03
	TAG_INTEGER = 0x13
	TAG_SHORT_STR =0x04
	TAG_LONG_STR =0x14
)

type Upvalue struct {
	Instack byte
	Idx byte
}

type LocVar struct {
	VarName string
	StartPC uint32
	EndPC uint32
}

func Undump(data []byte) *Prototype{
	reader:=&reader{data}
	reader.checkHeader()
	reader.readByte()
	return reader.readProto("")

}
