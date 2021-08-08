package vm

import (
	"log"
	"os"
	"strconv"
	"strings"
)

var cTypeMapper = map[Instruction]CommandType{
	I_ADD:      C_ARTITHMETIC,
	I_SUB:      C_ARTITHMETIC,
	I_NEG:      C_ARTITHMETIC,
	I_EQ:       C_ARTITHMETIC,
	I_GT:       C_ARTITHMETIC,
	I_LT:       C_ARTITHMETIC,
	I_AND:      C_ARTITHMETIC,
	I_OR:       C_ARTITHMETIC,
	I_NOT:      C_ARTITHMETIC,
	I_PUSH:     C_PUSH,
	I_POP:      C_POP,
	I_LABEL:    C_LABEL,
	I_IF_GOTO:  C_IF,
	I_FUNCTION: C_FUNCTION,
	I_RETURN:   C_RETURN,
	I_CALL:     C_CALL,
}

var cLenMapper = map[Instruction]uint8{
	"add":      1,
	"sub":      1,
	"neg":      1,
	"eq":       1,
	"gt":       1,
	"lt":       1,
	"and":      1,
	"or":       1,
	"not":      1,
	"push":     3,
	"pop":      3,
	"label":    2,
	"if-goto":  1,
	"function": 3,
	"return":   1,
	"call":     2,
}

type Command struct {
	Instruction Instruction
	Type        CommandType
	Arg1        *string
	Arg2        *int
}

type IParser interface {
	HasMoreCommands() bool
	Advance()
	CommandType() CommandType
	Arg1() *string
	Arg2() *int
}

type Parser struct {
	commands []*Command
	seeker   int
}

func NewParser(vmFilePath string) IParser {
	raw, e := os.ReadFile(vmFilePath)
	if e != nil {
		panic(e.Error())
	}

	return &Parser{
		commands: vmRawCodeToCommands(raw),
		seeker:   0,
	}
}

func (p Parser) HasMoreCommands() bool {
	return len(p.commands) > p.seeker
}

func (p *Parser) Advance() {
	p.seeker++
}

func (p Parser) CommandType() CommandType {
	return p.commands[p.seeker].Type
}

func (p Parser) Arg1() *string {
	return p.commands[p.seeker].Arg1
}
func (p Parser) Arg2() *int {
	return p.commands[p.seeker].Arg2
}

func vmRawCodeToCommands(rawCode []byte) (commands []*Command) {
	rawCodeArray := strings.Split(string(rawCode), "\n")
	commands = make([]*Command, 0, len(rawCodeArray))

	// skipping comments and empty line
	// parse command string to Command structure
	for _, line := range rawCodeArray {
		line = strings.TrimSpace(line)

		// skip comment
		commentPos := strings.Index(line, "//")
		if commentPos != -1 {
			line = line[:commentPos]
		}

		// skip empty line
		if strings.Compare(line, "") == 0 {
			continue
		}

		instruction := Instruction(strings.Split(line, " ")[0])
		// panic when we get unknown instruction
		if _, ok := cTypeMapper[instruction]; !ok {
			log.Fatal("Unknown instruction: " + instruction)
		}

		commandLen := cLenMapper[instruction]
		command := Command{
			Instruction: instruction,
			Type:        cTypeMapper[instruction],
			Arg1:        nil,
			Arg2:        nil,
		}

		// adds arguments
		if commandLen == 2 {
			command.Arg1 = &strings.Split(line, " ")[1]
		}
		if commandLen == 3 {
			arg2, e := strconv.Atoi(strings.Split(line, " ")[2])
			if e != nil {
				log.Fatalf("%s :: Invalid argument2 for this instruction", e.Error())
			}
			command.Arg2 = &arg2
		}

		// append command
		commands = append(commands, &command)
	}
	return
}
