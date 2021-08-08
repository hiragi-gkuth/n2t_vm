package vm

import (
	"os"
	"strconv"
)

type ICodeWriter interface {
	SetFileName(hackFilePath string)
	WriteArithmetic(command Command)
	WritePushPop(command Command)
	Close()
}

type CodeWriter struct {
	file          *os.File
	segmentMapper map[string]int
	asmCodes      []string
}

var (
	increaseSP = [2]string{"@SP", "M=M+1"}
	decreaseSP = [2]string{"@SP", "M=M-1"}
	refSP      = [2]string{"@SP", "A=M"}
)

func NewCodeWriter(hackFilePath string) ICodeWriter {
	f, e := os.OpenFile(hackFilePath, os.O_CREATE, 0644)
	if e != nil {
		panic(e.Error())
	}

	// init registers
	asmCodes := []string{}
	// set SP register to 256
	asmCodes = append(asmCodes,
		"@"+strconv.Itoa(O_STACK),
		"D=A",
		"@SP",
		"M=D",
	)

	return &CodeWriter{
		file: f,
		segmentMapper: map[string]int{
			"sp":      0x00,
			"lcl":     0x01,
			"arg":     0x02,
			"this":    0x03,
			"that":    0x04,
			"pointer": 0x03,
			"temp":    0x05,
		},
		asmCodes: asmCodes,
	}
}

func (cw *CodeWriter) SetFileName(hackFilePath string) {
	cw.file.Close()

	f, e := os.OpenFile(hackFilePath, os.O_CREATE, 0644)
	if e != nil {
		panic(e.Error())
	}

	cw.file = f
}

func (cw CodeWriter) WriteArithmetic(command Command) {
	if command.Type != C_ARTITHMETIC {
		panic("Call WriteArithmetic() with not C_ARTITHMETIC")
	}

	switch command.Instruction {
	case I_ADD:
		cw.asmCodes = append(cw.asmCodes, decreaseSP[:]...)

	}
}

func (cw CodeWriter) WritePushPop(command Command) {
	cType := command.Type
	if cType != C_PUSH && cType != C_POP {
		panic("invalid command type for calling WritePushPop")
	}
	seg := *command.Arg1
	loc := *command.Arg2

	// push constant
	if cType == C_PUSH {
		if seg == "constant" {
			// write constant to D register
			cw.asmCodes = append(cw.asmCodes, "@"+strconv.Itoa(loc))
			cw.asmCodes = append(cw.asmCodes, "D=A")
			// ref SP
			cw.asmCodes = append(cw.asmCodes, "@SP")
			cw.asmCodes = append(cw.asmCodes, "A=M")
			// write to stack
			cw.asmCodes = append(cw.asmCodes, "M=D")
			// update SP
			cw.asmCodes = append(cw.asmCodes, "@SP")
			cw.asmCodes = append(cw.asmCodes, "M=M+1")
		}
	}
}

func (cw CodeWriter) Close() {
	cw.file.Close()
}
