package gbc

import (
	"fmt"
	"gbc/pkg/util"
)

func (cpu *CPU) a16Fetch() uint16 {
	value := cpu.d16Fetch()
	return value
}

func (cpu *CPU) a16FetchJP() uint16 {
	lower := uint16(cpu.FetchMemory8(cpu.Reg.PC + 1)) // M = 1: nn read: memory access for low byte
	cpu.timer(1)
	upper := uint16(cpu.FetchMemory8(cpu.Reg.PC + 2)) // M = 2: nn read: memory access for high byte
	cpu.timer(1)
	value := (upper << 8) | lower
	return value
}

func (cpu *CPU) d8Fetch() byte {
	value := cpu.FetchMemory8(cpu.Reg.PC + 1)
	return value
}

func (cpu *CPU) d16Fetch() uint16 {
	lower, upper := uint16(cpu.FetchMemory8(cpu.Reg.PC+1)), uint16(cpu.FetchMemory8(cpu.Reg.PC+2))
	return (upper << 8) | lower
}

// ------ LD A, *

// LD A,(BC)
func op0x0a(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.A = cpu.FetchMemory8(cpu.Reg.BC())
	cpu.Reg.PC++
}

// LD A,(DE)
func op0x1a(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.A = cpu.FetchMemory8(cpu.Reg.DE())
	cpu.Reg.PC++
}

// LD A,(HL+)
func op0x2a(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.A = cpu.FetchMemory8(cpu.Reg.HL())
	cpu.Reg.setHL(cpu.Reg.HL() + 1)
	cpu.Reg.PC++
}

// LD A,(HL-)
func op0x3a(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.A = cpu.FetchMemory8(cpu.Reg.HL())
	cpu.Reg.setHL(cpu.Reg.HL() - 1)
	cpu.Reg.PC++
}

// LD A,u8
func op0x3e(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.A = cpu.FetchMemory8(cpu.Reg.PC + 1)
	cpu.Reg.PC += 2
}

// LD A, B
func op0x78(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.A = cpu.Reg.B
	cpu.Reg.PC++
}

// LD A, C
func op0x79(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.A = cpu.Reg.C
	cpu.Reg.PC++
}

// LD A, D
func op0x7a(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.A = cpu.Reg.D
	cpu.Reg.PC++
}

// LD A, E
func op0x7b(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.A = cpu.Reg.E
	cpu.Reg.PC++
}

// LD A, H
func op0x7c(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.A = cpu.Reg.H
	cpu.Reg.PC++
}

// LD A, L
func op0x7d(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.A = cpu.Reg.L
	cpu.Reg.PC++
}

// LD A, (HL)
func op0x7e(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.A = cpu.FetchMemory8(cpu.Reg.HL())
	cpu.Reg.PC++
}

// LD A, A
func op0x7f(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.PC++
}

// LD A, (u16)
func op0xfa(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.A = cpu.FetchMemory8(cpu.a16FetchJP())
	cpu.Reg.PC += 3
	cpu.timer(2)
}

// LD A,(FF00+C)
func op0xf2(cpu *CPU, operand1, operand2 int) {
	addr := 0xff00 + uint16(cpu.Reg.C)
	cpu.Reg.A = cpu.fetchIO(addr)
	cpu.Reg.PC++ // mistake?(https://www.pastraiser.com/cpu/gameboy/gameboy_opcodes.html)
}

// ------ LD B, *

// LD B,u8
func op0x06(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.B = cpu.d8Fetch()
	cpu.Reg.PC += 2
}

// LD B,B
func op0x40(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.PC++
}

// LD B,C
func op0x41(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.B = cpu.Reg.C
	cpu.Reg.PC++
}

// LD B,D
func op0x42(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.B = cpu.Reg.D
	cpu.Reg.PC++
}

// LD B,E
func op0x43(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.B = cpu.Reg.E
	cpu.Reg.PC++
}

// LD B,H
func op0x44(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.B = cpu.Reg.H
	cpu.Reg.PC++
}

// LD B,L
func op0x45(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.B = cpu.Reg.L
	cpu.Reg.PC++
}

// LD B,(HL)
func op0x46(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.B = cpu.FetchMemory8(cpu.Reg.HL())
	cpu.Reg.PC++
}

// LD B,A
func op0x47(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.B = cpu.Reg.A
	cpu.Reg.PC++
}

// ------ LD C, *

// LD C,u8
func op0x0e(cpu *CPU, operand1, operand2 int) {
	value := cpu.d8Fetch()
	cpu.Reg.C = value
	cpu.Reg.PC += 2
}

// LD C,B
func op0x48(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.C = cpu.Reg.B
	cpu.Reg.PC++
}

// LD C,C
func op0x49(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.PC++
}

// LD C,D
func op0x4a(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.C = cpu.Reg.D
	cpu.Reg.PC++
}

// LD C,E
func op0x4b(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.C = cpu.Reg.E
	cpu.Reg.PC++
}

// LD C,H
func op0x4c(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.C = cpu.Reg.H
	cpu.Reg.PC++
}

// LD C,L
func op0x4d(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.C = cpu.Reg.L
	cpu.Reg.PC++
}

// LD C,(HL)
func op0x4e(cpu *CPU, operand1, operand2 int) {
	value := cpu.FetchMemory8(cpu.Reg.HL())
	cpu.Reg.C = value
	cpu.Reg.PC++
}

// LD C,A
func op0x4f(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.C = cpu.Reg.A
	cpu.Reg.PC++
}

// ------ LD D, *

// LD D,u8
func op0x16(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.D = cpu.d8Fetch()
	cpu.Reg.PC += 2
}

// LD D,B
func op0x50(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.D = cpu.Reg.B
	cpu.Reg.PC++
}

// LD D,C
func op0x51(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.D = cpu.Reg.C
	cpu.Reg.PC++
}

// LD D,D
func op0x52(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.PC++
}

// LD D,E
func op0x53(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.D = cpu.Reg.E
	cpu.Reg.PC++
}

// LD D,H
func op0x54(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.D = cpu.Reg.H
	cpu.Reg.PC++
}

// LD D,L
func op0x55(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.D = cpu.Reg.L
	cpu.Reg.PC++
}

// LD D,(HL)
func op0x56(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.D = cpu.FetchMemory8(cpu.Reg.HL())
	cpu.Reg.PC++
}

// LD D,A
func op0x57(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.D = cpu.Reg.A
	cpu.Reg.PC++
}

// ------ LD E, *

// LD E,u8
func op0x1e(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.E = cpu.d8Fetch()
	cpu.Reg.PC += 2
}

// LD E,B
func op0x58(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.E = cpu.Reg.B
	cpu.Reg.PC++
}

// LD E,C
func op0x59(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.E = cpu.Reg.C
	cpu.Reg.PC++
}

// LD E,D
func op0x5a(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.E = cpu.Reg.D
	cpu.Reg.PC++
}

// LD E,E
func op0x5b(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.PC++
}

// LD E,H
func op0x5c(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.E = cpu.Reg.H
	cpu.Reg.PC++
}

// LD E,L
func op0x5d(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.E = cpu.Reg.L
	cpu.Reg.PC++
}

// LD E,(HL)
func op0x5e(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.E = cpu.FetchMemory8(cpu.Reg.HL())
	cpu.Reg.PC++
}

// LD E,A
func op0x5f(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.E = cpu.Reg.A
	cpu.Reg.PC++
}

// ------ LD H, *

// LD H,u8
func op0x26(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.H = cpu.d8Fetch()
	cpu.Reg.PC += 2
}

// LD H,B
func op0x60(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.H = cpu.Reg.B
	cpu.Reg.PC++
}

// LD H,C
func op0x61(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.H = cpu.Reg.C
	cpu.Reg.PC++
}

// LD H,D
func op0x62(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.H = cpu.Reg.D
	cpu.Reg.PC++
}

// LD H,E
func op0x63(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.H = cpu.Reg.E
	cpu.Reg.PC++
}

// LD H,H
func op0x64(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.PC++
}

// LD H,L
func op0x65(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.H = cpu.Reg.L
	cpu.Reg.PC++
}

// LD H,(HL)
func op0x66(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.H = cpu.FetchMemory8(cpu.Reg.HL())
	cpu.Reg.PC++
}

// LD H,A
func op0x67(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.H = cpu.Reg.A
	cpu.Reg.PC++
}

// ------ LD L, *

// LD L,u8
func op0x2e(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.L = cpu.d8Fetch()
	cpu.Reg.PC += 2
}

// LD L,B
func op0x68(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.L = cpu.Reg.B
	cpu.Reg.PC++
}

// LD L,C
func op0x69(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.L = cpu.Reg.C
	cpu.Reg.PC++
}

// LD L,D
func op0x6a(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.L = cpu.Reg.D
	cpu.Reg.PC++
}

// LD L,E
func op0x6b(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.L = cpu.Reg.E
	cpu.Reg.PC++
}

// LD L,H
func op0x6c(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.L = cpu.Reg.H
	cpu.Reg.PC++
}

// LD L,L
func op0x6d(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.PC++
}

// LD L,(HL)
func op0x6e(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.L = cpu.FetchMemory8(cpu.Reg.HL())
	cpu.Reg.PC++
}

// LD L,A
func op0x6f(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.L = cpu.Reg.A
	cpu.Reg.PC++
}

// ------ LD (HL), *

// LD (HL),u8
func op0x36(cpu *CPU, operand1, operand2 int) {
	value := cpu.d8Fetch()
	cpu.timer(1)
	cpu.SetMemory8(cpu.Reg.HL(), value)
	cpu.Reg.PC += 2
	cpu.timer(2)
}

// LD (HL),B
func op0x70(cpu *CPU, operand1, operand2 int) {
	cpu.SetMemory8(cpu.Reg.HL(), cpu.Reg.B)
	cpu.Reg.PC++
}

// LD (HL),C
func op0x71(cpu *CPU, operand1, operand2 int) {
	cpu.SetMemory8(cpu.Reg.HL(), cpu.Reg.C)
	cpu.Reg.PC++
}

// LD (HL),D
func op0x72(cpu *CPU, operand1, operand2 int) {
	cpu.SetMemory8(cpu.Reg.HL(), cpu.Reg.D)
	cpu.Reg.PC++
}

// LD (HL),E
func op0x73(cpu *CPU, operand1, operand2 int) {
	cpu.SetMemory8(cpu.Reg.HL(), cpu.Reg.E)
	cpu.Reg.PC++
}

// LD (HL),H
func op0x74(cpu *CPU, operand1, operand2 int) {
	cpu.SetMemory8(cpu.Reg.HL(), cpu.Reg.H)
	cpu.Reg.PC++
}

// LD (HL),L
func op0x75(cpu *CPU, operand1, operand2 int) {
	cpu.SetMemory8(cpu.Reg.HL(), cpu.Reg.L)
	cpu.Reg.PC++
}

// LD (HL),A
func op0x77(cpu *CPU, operand1, operand2 int) {
	cpu.SetMemory8(cpu.Reg.HL(), cpu.Reg.A)
	cpu.Reg.PC++
}

// ------ others ld

// LD (u16),SP
func op0x08(cpu *CPU, operand1, operand2 int) {
	// Store SP into addresses n16 (LSB) and n16 + 1 (MSB).
	addr := cpu.a16Fetch()
	upper, lower := byte(cpu.Reg.SP>>8), byte(cpu.Reg.SP) // MSB
	cpu.SetMemory8(addr, lower)
	cpu.SetMemory8(addr+1, upper)
	cpu.Reg.PC += 3
	cpu.timer(5)
}

// LD (u16),A
func op0xea(cpu *CPU, operand1, operand2 int) {
	cpu.SetMemory8(cpu.a16FetchJP(), cpu.Reg.A)
	cpu.Reg.PC += 3
	cpu.timer(2)
}

// LD BC,u16
func op0x01(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.setBC(cpu.d16Fetch())
	cpu.Reg.PC += 3
}

// LD DE,u16
func op0x11(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.setDE(cpu.d16Fetch())
	cpu.Reg.PC += 3
}

// LD HL,u16
func op0x21(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.setHL(cpu.d16Fetch())
	cpu.Reg.PC += 3
}

// LD SP,u16
func op0x31(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.SP = cpu.d16Fetch()
	cpu.Reg.PC += 3
}

// LD HL,SP+i8
func op0xf8(cpu *CPU, operand1, operand2 int) {
	delta := int8(cpu.FetchMemory8(cpu.Reg.PC + 1))
	value := int32(cpu.Reg.SP) + int32(delta)
	carryBits := uint32(cpu.Reg.SP) ^ uint32(delta) ^ uint32(value)
	cpu.Reg.setHL(uint16(value))
	cpu.setF(flagZ, false)
	cpu.setF(flagN, false)
	cpu.setF(flagC, util.Bit(carryBits, 8))
	cpu.setF(flagH, util.Bit(carryBits, 4))
	cpu.Reg.PC += 2
}

// LD SP,HL
func op0xf9(cpu *CPU, operand1, operand2 int) {
	cpu.Reg.SP = cpu.Reg.HL()
	cpu.Reg.PC++
}

// LD (FF00+C),A
func op0xe2(cpu *CPU, operand1, operand2 int) {
	addr := 0xff00 + uint16(cpu.Reg.C)
	cpu.SetMemory8(addr, cpu.Reg.A)
	cpu.Reg.PC++ // mistake?(https://www.pastraiser.com/cpu/gameboy/gameboy_opcodes.html)
}

// LD (BC),A
func op0x02(cpu *CPU, operand1, operand2 int) {
	cpu.SetMemory8(cpu.Reg.BC(), cpu.Reg.A)
	cpu.Reg.PC++
}

// LD (DE),A
func op0x12(cpu *CPU, operand1, operand2 int) {
	cpu.SetMemory8(cpu.Reg.DE(), cpu.Reg.A)
	cpu.Reg.PC++
}

// LD (HL+),A
func op0x22(cpu *CPU, operand1, operand2 int) {
	cpu.SetMemory8(cpu.Reg.HL(), cpu.Reg.A)
	cpu.Reg.setHL(cpu.Reg.HL() + 1)
	cpu.Reg.PC++
}

// LD (HL-),A
func op0x32(cpu *CPU, operand1, operand2 int) {
	// (HL)=A, HL=HL-1
	cpu.SetMemory8(cpu.Reg.HL(), cpu.Reg.A)
	cpu.Reg.setHL(cpu.Reg.HL() - 1)
	cpu.Reg.PC++
}

// LD Load
func LD(cpu *CPU, operand1, operand2 int) {
	errMsg := fmt.Sprintf("Error: LD %d %d", operand1, operand2)
	panic(errMsg)
}

// LDH Load High Byte
func LDH(cpu *CPU, operand1, operand2 int) {
	if operand1 == OPERAND_A && operand2 == OPERAND_a8_PAREN { // LD A,($FF00+a8)
		addr := 0xff00 + uint16(cpu.FetchMemory8(cpu.Reg.PC+1))
		cpu.timer(1)
		value := cpu.fetchIO(addr)

		cpu.Reg.A = value
		cpu.Reg.PC += 2
		cpu.timer(2)
	} else if operand1 == OPERAND_a8_PAREN && operand2 == OPERAND_A { // LD ($FF00+a8),A
		addr := 0xff00 + uint16(cpu.FetchMemory8(cpu.Reg.PC+1))
		cpu.timer(1)
		cpu.setIO(addr, cpu.Reg.A)
		cpu.Reg.PC += 2
		cpu.timer(2)
	} else {
		panic(fmt.Errorf("error: LDH %d %d", operand1, operand2))
	}
}

// NOP No operation
func (cpu *CPU) NOP(operand1, operand2 int) {
	cpu.Reg.PC++
}

// INC Increment
func (cpu *CPU) INC(operand1, operand2 int) {
	var value, carryBits byte

	switch operand1 {
	case OPERAND_A:
		value = cpu.Reg.A + 1
		carryBits = cpu.Reg.A ^ 1 ^ value
		cpu.Reg.A = value
	case OPERAND_B:
		value = cpu.Reg.B + 1
		carryBits = cpu.Reg.B ^ 1 ^ value
		cpu.Reg.B = value
	case OPERAND_C:
		value = cpu.Reg.C + 1
		carryBits = cpu.Reg.C ^ 1 ^ value
		cpu.Reg.C = value
	case OPERAND_D:
		value = cpu.Reg.D + 1
		carryBits = cpu.Reg.D ^ 1 ^ value
		cpu.Reg.D = value
	case OPERAND_E:
		value = cpu.Reg.E + 1
		carryBits = cpu.Reg.E ^ 1 ^ value
		cpu.Reg.E = value
	case OPERAND_H:
		value = cpu.Reg.H + 1
		carryBits = cpu.Reg.H ^ 1 ^ value
		cpu.Reg.H = value
	case OPERAND_L:
		value = cpu.Reg.L + 1
		carryBits = cpu.Reg.L ^ 1 ^ value
		cpu.Reg.L = value
	case OPERAND_HL_PAREN:
		value = cpu.FetchMemory8(cpu.Reg.HL()) + 1
		cpu.timer(1)
		carryBits = cpu.FetchMemory8(cpu.Reg.HL()) ^ 1 ^ value
		cpu.SetMemory8(cpu.Reg.HL(), value)
		cpu.timer(2)
	case OPERAND_BC:
		cpu.Reg.setBC(cpu.Reg.BC() + 1)
	case OPERAND_DE:
		cpu.Reg.setDE(cpu.Reg.DE() + 1)
	case OPERAND_HL:
		cpu.Reg.setHL(cpu.Reg.HL() + 1)
	case OPERAND_SP:
		cpu.Reg.SP++
	default:
		panic(fmt.Errorf("error: INC %d %d", operand1, operand2))
	}

	if operand1 != OPERAND_BC && operand1 != OPERAND_DE && operand1 != OPERAND_HL && operand1 != OPERAND_SP {
		cpu.setF(flagZ, value == 0)
		cpu.setF(flagN, false)
		cpu.setF(flagH, util.Bit(carryBits, 4))
	}
	cpu.Reg.PC++
}

// DEC Decrement
func (cpu *CPU) DEC(operand1, operand2 int) {
	var value byte
	var carryBits byte

	switch operand1 {
	case OPERAND_A:
		value = cpu.Reg.A - 1
		carryBits = cpu.Reg.A ^ 1 ^ value
		cpu.Reg.A = value
	case OPERAND_B:
		value = cpu.Reg.B - 1
		carryBits = cpu.Reg.B ^ 1 ^ value
		cpu.Reg.B = value
	case OPERAND_C:
		value = cpu.Reg.C - 1
		carryBits = cpu.Reg.C ^ 1 ^ value
		cpu.Reg.C = value
	case OPERAND_D:
		value = cpu.Reg.D - 1
		carryBits = cpu.Reg.D ^ 1 ^ value
		cpu.Reg.D = value
	case OPERAND_E:
		value = cpu.Reg.E - 1
		carryBits = cpu.Reg.E ^ 1 ^ value
		cpu.Reg.E = value
	case OPERAND_H:
		value = cpu.Reg.H - 1
		carryBits = cpu.Reg.H ^ 1 ^ value
		cpu.Reg.H = value
	case OPERAND_L:
		value = cpu.Reg.L - 1
		carryBits = cpu.Reg.L ^ 1 ^ value
		cpu.Reg.L = value
	case OPERAND_HL_PAREN:
		value = cpu.FetchMemory8(cpu.Reg.HL()) - 1
		cpu.timer(1)
		carryBits = cpu.FetchMemory8(cpu.Reg.HL()) ^ 1 ^ value
		cpu.SetMemory8(cpu.Reg.HL(), value)
		cpu.timer(2)
	case OPERAND_BC:
		cpu.Reg.setBC(cpu.Reg.BC() - 1)
	case OPERAND_DE:
		cpu.Reg.setDE(cpu.Reg.DE() - 1)
	case OPERAND_HL:
		cpu.Reg.setHL(cpu.Reg.HL() - 1)
	case OPERAND_SP:
		cpu.Reg.SP--
	default:
		panic(fmt.Errorf("error: DEC %d %d", operand1, operand2))
	}

	if operand1 != OPERAND_BC && operand1 != OPERAND_DE && operand1 != OPERAND_HL && operand1 != OPERAND_SP {
		cpu.setF(flagZ, value == 0)
		cpu.setF(flagN, true)
		cpu.setF(flagH, util.Bit(carryBits, 4))
	}
	cpu.Reg.PC++
}

// --------- JR ---------

// JR i8
func op0x18(cpu *CPU, operand1, operand2 int) {
	delta := int8(cpu.FetchMemory8(cpu.Reg.PC + 1))
	destination := uint16(int32(cpu.Reg.PC+2) + int32(delta)) // PC+2 because of time after fetch(pc is incremented)
	cpu.Reg.PC = destination
	cpu.timer(3)
}

// JR NZ,i8
func op0x20(cpu *CPU, operand1, operand2 int) {
	if !cpu.f(flagZ) {
		delta := int8(cpu.FetchMemory8(cpu.Reg.PC + 1))
		destination := uint16(int32(cpu.Reg.PC+2) + int32(delta)) // PC+2 because of time after fetch(pc is incremented)
		cpu.Reg.PC = destination
		cpu.timer(3)
	} else {
		cpu.Reg.PC += 2
		cpu.timer(2)
	}
}

// JR Z,i8
func op0x28(cpu *CPU, operand1, operand2 int) {
	if cpu.f(flagZ) {
		delta := int8(cpu.FetchMemory8(cpu.Reg.PC + 1))
		destination := uint16(int32(cpu.Reg.PC+2) + int32(delta)) // PC+2 because of time after fetch(pc is incremented)
		cpu.Reg.PC = destination
		cpu.timer(3)
	} else {
		cpu.Reg.PC += 2
		cpu.timer(2)
	}
}

// JR NC,i8
func op0x30(cpu *CPU, operand1, operand2 int) {
	if !cpu.f(flagC) {
		delta := int8(cpu.FetchMemory8(cpu.Reg.PC + 1))
		destination := uint16(int32(cpu.Reg.PC+2) + int32(delta)) // PC+2 because of time after fetch(pc is incremented)
		cpu.Reg.PC = destination
		cpu.timer(3)
	} else {
		cpu.Reg.PC += 2
		cpu.timer(2)
	}
}

// JR C,i8
func op0x38(cpu *CPU, operand1, operand2 int) {
	if cpu.f(flagC) {
		delta := int8(cpu.FetchMemory8(cpu.Reg.PC + 1))
		destination := uint16(int32(cpu.Reg.PC+2) + int32(delta)) // PC+2 because of time after fetch(pc is incremented)
		cpu.Reg.PC = destination
		cpu.timer(3)
	} else {
		cpu.Reg.PC += 2
		cpu.timer(2)
	}
}

// JR Jump relatively
func JR(cpu *CPU, operand1, operand2 int) {
	result := true
	switch operand1 {
	case OPERAND_r8:
		delta := int8(cpu.FetchMemory8(cpu.Reg.PC + 1))
		destination := uint16(int32(cpu.Reg.PC+2) + int32(delta)) // PC+2 because of time after fetch(pc is incremented)
		cpu.Reg.PC = destination
	case OPERAND_Z:
		if cpu.f(flagZ) {
			delta := int8(cpu.FetchMemory8(cpu.Reg.PC + 1))
			destination := uint16(int32(cpu.Reg.PC+2) + int32(delta)) // PC+2 because of time after fetch(pc is incremented)
			cpu.Reg.PC = destination
		} else {
			cpu.Reg.PC += 2
			result = false
		}
	case OPERAND_C:
		if cpu.f(flagC) {
			delta := int8(cpu.FetchMemory8(cpu.Reg.PC + 1))
			destination := uint16(int32(cpu.Reg.PC+2) + int32(delta)) // PC+2 because of time after fetch(pc is incremented)
			cpu.Reg.PC = destination
		} else {
			cpu.Reg.PC += 2
			result = false
		}
	case OPERAND_NZ:
		if !cpu.f(flagZ) {
			delta := int8(cpu.FetchMemory8(cpu.Reg.PC + 1))
			destination := uint16(int32(cpu.Reg.PC+2) + int32(delta)) // PC+2 because of time after fetch(pc is incremented)
			cpu.Reg.PC = destination
		} else {
			cpu.Reg.PC += 2
			result = false
		}
	case OPERAND_NC:
		if !cpu.f(flagC) {
			delta := int8(cpu.FetchMemory8(cpu.Reg.PC + 1))
			destination := uint16(int32(cpu.Reg.PC+2) + int32(delta)) // PC+2 because of time after fetch(pc is incremented)
			cpu.Reg.PC = destination
		} else {
			cpu.Reg.PC += 2
			result = false
		}
	default:
		panic(fmt.Errorf("error: JR %d %d", operand1, operand2))
	}

	if result {
		cpu.timer(3)
	} else {
		cpu.timer(2)
	}
}

// HALT Halt
func (cpu *CPU) HALT(operand1, operand2 int) {
	cpu.Reg.PC++
	cpu.halt = true

	// ref: https://rednex.github.io/rgbds/gbz80.7.html#HALT
	if !cpu.Reg.IME {
		IE, IF := cpu.RAM[IEIO], cpu.RAM[IFIO]
		pending := IE&IF != 0
		if pending {
			// Some pending
			cpu.halt = false
			PC := cpu.Reg.PC
			cpu.exec()
			cpu.Reg.PC = PC

			// IME turns on due to EI delay.
			cpu.halt = cpu.Reg.IME
		}
	}
}

// STOP stop CPU
func (cpu *CPU) STOP(operand1, operand2 int) {
	if operand1 == OPERAND_0 && operand2 == OPERAND_NONE {
		cpu.Reg.PC += 2
		KEY1 := cpu.FetchMemory8(KEY1IO)
		if util.Bit(KEY1, 0) {
			if util.Bit(KEY1, 7) {
				KEY1 = 0x00
				cpu.boost = 1
			} else {
				KEY1 = 0x80
				cpu.boost = 2
			}
			cpu.SetMemory8(KEY1IO, KEY1)
		}
	} else {
		panic(fmt.Errorf("error: STOP %d %d", operand1, operand2))
	}
}

// XOR xor
func (cpu *CPU) XOR(operand1, operand2 int) {
	var value byte
	switch operand1 {
	case OPERAND_B:
		value = cpu.Reg.A ^ cpu.Reg.B
	case OPERAND_C:
		value = cpu.Reg.A ^ cpu.Reg.C
	case OPERAND_D:
		value = cpu.Reg.A ^ cpu.Reg.D
	case OPERAND_E:
		value = cpu.Reg.A ^ cpu.Reg.E
	case OPERAND_H:
		value = cpu.Reg.A ^ cpu.Reg.H
	case OPERAND_L:
		value = cpu.Reg.A ^ cpu.Reg.L
	case OPERAND_HL_PAREN:
		value = cpu.Reg.A ^ cpu.FetchMemory8(cpu.Reg.HL())
	case OPERAND_A:
		value = 0
	case OPERAND_d8:
		value = cpu.Reg.A ^ cpu.FetchMemory8(cpu.Reg.PC+1)
		cpu.Reg.PC++
	default:
		panic(fmt.Errorf("error: XOR %d %d", operand1, operand2))
	}

	cpu.Reg.A = value
	cpu.setF(flagZ, value == 0)
	cpu.setF(flagN, false)
	cpu.setF(flagH, false)
	cpu.setF(flagC, false)
	cpu.Reg.PC++
}

// JP Jump
func JP(cpu *CPU, operand1, operand2 int) {
	cycle := 1

	switch operand1 {
	case OPERAND_a16:
		destination := cpu.a16FetchJP()
		cycle++
		cpu.Reg.PC = destination
	case OPERAND_HL_PAREN:
		cpu.Reg.PC = cpu.Reg.HL()
	case OPERAND_Z:
		destination := cpu.a16FetchJP()
		if cpu.f(flagZ) {
			cycle++
			cpu.Reg.PC = destination
		} else {
			cpu.Reg.PC += 3
		}
	case OPERAND_C:
		destination := cpu.a16FetchJP()
		if cpu.f(flagC) {
			cycle++
			cpu.Reg.PC = destination
		} else {
			cpu.Reg.PC += 3
		}
	case OPERAND_NZ:
		destination := cpu.a16FetchJP()
		if !cpu.f(flagZ) {
			cycle++
			cpu.Reg.PC = destination
		} else {
			cpu.Reg.PC += 3
		}
	case OPERAND_NC:
		destination := cpu.a16FetchJP()
		if !cpu.f(flagC) {
			cycle++
			cpu.Reg.PC = destination
		} else {
			cpu.Reg.PC += 3
		}
	default:
		panic(fmt.Errorf("error: JP %d %d", operand1, operand2))
	}

	cpu.timer(cycle)
}

// RET Return
func (cpu *CPU) RET(operand1, operand2 int) (result bool) {
	result = true

	switch operand1 {
	case OPERAND_NONE:
		// PC=(SP), SP=SP+2
		cpu.popPC()
	case OPERAND_Z:
		if cpu.f(flagZ) {
			cpu.popPC()
		} else {
			cpu.Reg.PC++
			result = false
		}
	case OPERAND_C:
		if cpu.f(flagC) {
			cpu.popPC()
		} else {
			cpu.Reg.PC++
			result = false
		}
	case OPERAND_NZ:
		if !cpu.f(flagZ) {
			cpu.popPC()
		} else {
			cpu.Reg.PC++
			result = false
		}
	case OPERAND_NC:
		if !cpu.f(flagC) {
			cpu.popPC()
		} else {
			cpu.Reg.PC++
			result = false
		}
	default:
		panic(fmt.Errorf("error: RET %d %d", operand1, operand2))
	}

	return result
}

// RETI Return Interrupt
func (cpu *CPU) RETI(operand1, operand2 int) {
	cpu.popPC()
	cpu.Reg.IME = true
}

// CALL Call subroutine
func CALL(cpu *CPU, operand1, operand2 int) {

	switch operand1 {
	case OPERAND_a16:
		destination := cpu.a16FetchJP()
		cpu.Reg.PC += 3
		cpu.timer(1)
		cpu.pushPCCALL()
		cpu.timer(1)
		cpu.Reg.PC = destination
	case OPERAND_Z:
		if cpu.f(flagZ) {
			destination := cpu.a16FetchJP()
			cpu.Reg.PC += 3
			cpu.timer(1)
			cpu.pushPCCALL()
			cpu.timer(1)
			cpu.Reg.PC = destination
		} else {
			cpu.Reg.PC += 3
			cpu.timer(3)
		}
	case OPERAND_C:
		if cpu.f(flagC) {
			destination := cpu.a16FetchJP()
			cpu.Reg.PC += 3
			cpu.timer(1)
			cpu.pushPCCALL()
			cpu.timer(1)
			cpu.Reg.PC = destination
		} else {
			cpu.Reg.PC += 3
			cpu.timer(3)
		}
	case OPERAND_NZ:
		if !cpu.f(flagZ) {
			destination := cpu.a16FetchJP()
			cpu.Reg.PC += 3
			cpu.timer(1)
			cpu.pushPCCALL()
			cpu.timer(1)
			cpu.Reg.PC = destination
		} else {
			cpu.Reg.PC += 3
			cpu.timer(3)
		}
	case OPERAND_NC:
		if !cpu.f(flagC) {
			destination := cpu.a16FetchJP()
			cpu.Reg.PC += 3
			cpu.timer(1)
			cpu.pushPCCALL()
			cpu.timer(1)
			cpu.Reg.PC = destination
		} else {
			cpu.Reg.PC += 3
			cpu.timer(3)
		}
	default:
		panic(fmt.Errorf("error: CALL %d %d", operand1, operand2))
	}
}

// DI Disable Interrupt
func (cpu *CPU) DI(operand1, operand2 int) {
	cpu.Reg.IME = false
	cpu.Reg.PC++
	if cpu.IMESwitch.Working && cpu.IMESwitch.Value {
		cpu.IMESwitch.Working = false // https://gbdev.gg8.se/wiki/articles/Interrupts 『The effect of EI is delayed by one instruction. This means that EI followed immediately by DI does not allow interrupts between the EI and the DI.』
	}
}

// EI Enable Interrupt
func (cpu *CPU) EI(operand1, operand2 int) {
	// ref: https://github.com/Gekkio/mooneye-gb/blob/master/tests/acceptance/halt_ime0_ei.s#L23
	next := cpu.FetchMemory8(cpu.Reg.PC + 1) // next opcode
	HALT := byte(0x76)
	if next == HALT {
		cpu.Reg.IME = true
		cpu.Reg.PC++
		return
	}

	if !cpu.IMESwitch.Working {
		cpu.IMESwitch = IMESwitch{
			Count:   2,
			Value:   true,
			Working: true,
		}
	}
	cpu.Reg.PC++
}

// CP Compare
func (cpu *CPU) CP(operand1, operand2 int) {
	var value, carryBits byte

	switch operand1 {
	case OPERAND_A:
		value, carryBits = 0, 0
		cpu.setCSub(cpu.Reg.A, cpu.Reg.A)
	case OPERAND_B:
		value = cpu.Reg.A - cpu.Reg.B
		carryBits = cpu.Reg.A ^ cpu.Reg.B ^ value
		cpu.setCSub(cpu.Reg.A, cpu.Reg.B)
	case OPERAND_C:
		value = cpu.Reg.A - cpu.Reg.C
		carryBits = cpu.Reg.A ^ cpu.Reg.C ^ value
		cpu.setCSub(cpu.Reg.A, cpu.Reg.C)
	case OPERAND_D:
		value = cpu.Reg.A - cpu.Reg.D
		carryBits = cpu.Reg.A ^ cpu.Reg.D ^ value
		cpu.setCSub(cpu.Reg.A, cpu.Reg.D)
	case OPERAND_E:
		value = cpu.Reg.A - cpu.Reg.E
		carryBits = cpu.Reg.A ^ cpu.Reg.E ^ value
		cpu.setCSub(cpu.Reg.A, cpu.Reg.E)
	case OPERAND_H:
		value = cpu.Reg.A - cpu.Reg.H
		carryBits = cpu.Reg.A ^ cpu.Reg.H ^ value
		cpu.setCSub(cpu.Reg.A, cpu.Reg.H)
	case OPERAND_L:
		value = cpu.Reg.A - cpu.Reg.L
		carryBits = cpu.Reg.A ^ cpu.Reg.L ^ value
		cpu.setCSub(cpu.Reg.A, cpu.Reg.L)
	case OPERAND_d8:
		value = cpu.Reg.A - cpu.d8Fetch()
		carryBits = cpu.Reg.A ^ cpu.d8Fetch() ^ value
		cpu.setCSub(cpu.Reg.A, cpu.d8Fetch())
		cpu.Reg.PC++
	case OPERAND_HL_PAREN:
		value = cpu.Reg.A - cpu.FetchMemory8(cpu.Reg.HL())
		carryBits = cpu.Reg.A ^ cpu.FetchMemory8(cpu.Reg.HL()) ^ value
		cpu.setCSub(cpu.Reg.A, cpu.FetchMemory8(cpu.Reg.HL()))
	default:
		panic(fmt.Errorf("error: CP %d %d", operand1, operand2))
	}
	cpu.setF(flagZ, value == 0)
	cpu.setF(flagN, true)
	cpu.setF(flagH, util.Bit(carryBits, 4))
	cpu.Reg.PC++
}

// AND And instruction
func (cpu *CPU) AND(operand1, operand2 int) {
	var value byte
	switch operand1 {
	case OPERAND_A:
		value = cpu.Reg.A
	case OPERAND_B:
		value = cpu.Reg.A & cpu.Reg.B
	case OPERAND_C:
		value = cpu.Reg.A & cpu.Reg.C
	case OPERAND_D:
		value = cpu.Reg.A & cpu.Reg.D
	case OPERAND_E:
		value = cpu.Reg.A & cpu.Reg.E
	case OPERAND_H:
		value = cpu.Reg.A & cpu.Reg.H
	case OPERAND_L:
		value = cpu.Reg.A & cpu.Reg.L
	case OPERAND_HL_PAREN:
		value = cpu.Reg.A & cpu.FetchMemory8(cpu.Reg.HL())
	case OPERAND_d8:
		value = cpu.Reg.A & cpu.d8Fetch()
		cpu.Reg.PC++
	default:
		panic(fmt.Errorf("error: AND %d %d", operand1, operand2))
	}

	cpu.Reg.A = value
	cpu.setF(flagZ, value == 0)
	cpu.setF(flagN, false)
	cpu.setF(flagH, true)
	cpu.setF(flagC, false)
	cpu.Reg.PC++
}

// OR or
func (cpu *CPU) OR(operand1, operand2 int) {
	switch operand1 {
	case OPERAND_A:
		cpu.setF(flagZ, cpu.Reg.A == 0)
	case OPERAND_B:
		value := cpu.Reg.A | cpu.Reg.B
		cpu.Reg.A = value
		cpu.setF(flagZ, value == 0)
	case OPERAND_C:
		value := cpu.Reg.A | cpu.Reg.C
		cpu.Reg.A = value
		cpu.setF(flagZ, value == 0)
	case OPERAND_D:
		value := cpu.Reg.A | cpu.Reg.D
		cpu.Reg.A = value
		cpu.setF(flagZ, value == 0)
	case OPERAND_E:
		value := cpu.Reg.A | cpu.Reg.E
		cpu.Reg.A = value
		cpu.setF(flagZ, value == 0)
	case OPERAND_H:
		value := cpu.Reg.A | cpu.Reg.H
		cpu.Reg.A = value
		cpu.setF(flagZ, value == 0)
	case OPERAND_L:
		value := cpu.Reg.A | cpu.Reg.L
		cpu.Reg.A = value
		cpu.setF(flagZ, value == 0)
	case OPERAND_d8:
		value := cpu.Reg.A | cpu.FetchMemory8(cpu.Reg.PC+1)
		cpu.Reg.A = value
		cpu.setF(flagZ, value == 0)
		cpu.Reg.PC++
	case OPERAND_HL_PAREN:
		value := cpu.Reg.A | cpu.FetchMemory8(cpu.Reg.HL())
		cpu.Reg.A = value
		cpu.setF(flagZ, value == 0)
	default:
		panic(fmt.Errorf("error: OR %d %d", operand1, operand2))
	}

	cpu.setF(flagN, false)
	cpu.setF(flagH, false)
	cpu.setF(flagC, false)
	cpu.Reg.PC++
}

// ADD Addition
func (cpu *CPU) ADD(operand1, operand2 int) {
	switch operand1 {
	case OPERAND_A:
		switch operand2 {
		case OPERAND_A:
			value := uint16(cpu.Reg.A) + uint16(cpu.Reg.A)
			carryBits := value
			cpu.Reg.A = byte(value)
			cpu.setF(flagZ, byte(value) == 0)
			cpu.setF(flagN, false)
			cpu.setF(flagH, util.Bit(carryBits, 4))
			cpu.setF(flagC, util.Bit(carryBits, 8))
			cpu.Reg.PC++
		case OPERAND_B:
			value := uint16(cpu.Reg.A) + uint16(cpu.Reg.B)
			carryBits := uint16(cpu.Reg.A) ^ uint16(cpu.Reg.B) ^ value
			cpu.Reg.A = byte(value)
			cpu.setF(flagZ, byte(value) == 0)
			cpu.setF(flagN, false)
			cpu.setF(flagH, util.Bit(carryBits, 4))
			cpu.setF(flagC, util.Bit(carryBits, 8))
			cpu.Reg.PC++
		case OPERAND_C:
			value := uint16(cpu.Reg.A) + uint16(cpu.Reg.C)
			carryBits := uint16(cpu.Reg.A) ^ uint16(cpu.Reg.C) ^ value
			cpu.Reg.A = byte(value)
			cpu.setF(flagZ, byte(value) == 0)
			cpu.setF(flagN, false)
			cpu.setF(flagH, util.Bit(carryBits, 4))
			cpu.setF(flagC, util.Bit(carryBits, 8))
			cpu.Reg.PC++
		case OPERAND_D:
			value := uint16(cpu.Reg.A) + uint16(cpu.Reg.D)
			carryBits := uint16(cpu.Reg.A) ^ uint16(cpu.Reg.D) ^ value
			cpu.Reg.A = byte(value)
			cpu.setF(flagZ, byte(value) == 0)
			cpu.setF(flagN, false)
			cpu.setF(flagH, util.Bit(carryBits, 4))
			cpu.setF(flagC, util.Bit(carryBits, 8))
			cpu.Reg.PC++
		case OPERAND_E:
			value := uint16(cpu.Reg.A) + uint16(cpu.Reg.E)
			carryBits := uint16(cpu.Reg.A) ^ uint16(cpu.Reg.E) ^ value
			cpu.Reg.A = byte(value)
			cpu.setF(flagZ, byte(value) == 0)
			cpu.setF(flagN, false)
			cpu.setF(flagH, util.Bit(carryBits, 4))
			cpu.setF(flagC, util.Bit(carryBits, 8))
			cpu.Reg.PC++
		case OPERAND_H:
			value := uint16(cpu.Reg.A) + uint16(cpu.Reg.H)
			carryBits := uint16(cpu.Reg.A) ^ uint16(cpu.Reg.H) ^ value
			cpu.Reg.A = byte(value)
			cpu.setF(flagZ, byte(value) == 0)
			cpu.setF(flagN, false)
			cpu.setF(flagH, util.Bit(carryBits, 4))
			cpu.setF(flagC, util.Bit(carryBits, 8))
			cpu.Reg.PC++
		case OPERAND_L:
			value := uint16(cpu.Reg.A) + uint16(cpu.Reg.L)
			carryBits := uint16(cpu.Reg.A) ^ uint16(cpu.Reg.L) ^ value
			cpu.Reg.A = byte(value)
			cpu.setF(flagZ, byte(value) == 0)
			cpu.setF(flagN, false)
			cpu.setF(flagH, util.Bit(carryBits, 4))
			cpu.setF(flagC, util.Bit(carryBits, 8))
			cpu.Reg.PC++
		case OPERAND_d8:
			value := uint16(cpu.Reg.A) + uint16(cpu.d8Fetch())
			carryBits := uint16(cpu.Reg.A) ^ uint16(cpu.d8Fetch()) ^ value
			cpu.Reg.A = byte(value)
			cpu.setF(flagZ, byte(value) == 0)
			cpu.setF(flagN, false)
			cpu.setF(flagH, util.Bit(carryBits, 4))
			cpu.setF(flagC, util.Bit(carryBits, 8))
			cpu.Reg.PC += 2
		case OPERAND_HL_PAREN:
			value := uint16(cpu.Reg.A) + uint16(cpu.FetchMemory8(cpu.Reg.HL()))
			carryBits := uint16(cpu.Reg.A) ^ uint16(cpu.FetchMemory8(cpu.Reg.HL())) ^ value
			cpu.Reg.A = byte(value)
			cpu.setF(flagZ, byte(value) == 0)
			cpu.setF(flagN, false)
			cpu.setF(flagH, util.Bit(carryBits, 4))
			cpu.setF(flagC, util.Bit(carryBits, 8))
			cpu.Reg.PC++
		}
	case OPERAND_HL:
		switch operand2 {
		case OPERAND_BC:
			value := uint32(cpu.Reg.HL()) + uint32(cpu.Reg.BC())
			carryBits := uint32(cpu.Reg.HL()) ^ uint32(cpu.Reg.BC()) ^ value
			cpu.Reg.setHL(uint16(value))
			cpu.setF(flagN, false)
			cpu.setF(flagH, util.Bit(carryBits, 12))
			cpu.setF(flagC, util.Bit(carryBits, 16))
			cpu.Reg.PC++
		case OPERAND_DE:
			value := uint32(cpu.Reg.HL()) + uint32(cpu.Reg.DE())
			carryBits := uint32(cpu.Reg.HL()) ^ uint32(cpu.Reg.DE()) ^ value
			cpu.Reg.setHL(uint16(value))
			cpu.setF(flagN, false)
			cpu.setF(flagH, util.Bit(carryBits, 12))
			cpu.setF(flagC, util.Bit(carryBits, 16))
			cpu.Reg.PC++
		case OPERAND_HL:
			value := uint32(cpu.Reg.HL()) + uint32(cpu.Reg.HL())
			carryBits := value
			cpu.Reg.setHL(uint16(value))
			cpu.setF(flagN, false)
			cpu.setF(flagH, util.Bit(carryBits, 12))
			cpu.setF(flagC, util.Bit(carryBits, 16))
			cpu.Reg.PC++
		case OPERAND_SP:
			value := uint32(cpu.Reg.HL()) + uint32(cpu.Reg.SP)
			carryBits := uint32(cpu.Reg.HL()) ^ uint32(cpu.Reg.SP) ^ value
			cpu.Reg.setHL(uint16(value))
			cpu.setF(flagN, false)
			cpu.setF(flagH, util.Bit(carryBits, 12))
			cpu.setF(flagC, util.Bit(carryBits, 16))
			cpu.Reg.PC++
		}
	case OPERAND_SP:
		switch operand2 {
		case OPERAND_r8:
			delta := int8(cpu.FetchMemory8(cpu.Reg.PC + 1))
			value := int32(cpu.Reg.SP) + int32(delta)
			carryBits := uint32(cpu.Reg.SP) ^ uint32(delta) ^ uint32(value)
			cpu.Reg.SP = uint16(value)
			cpu.setF(flagZ, false)
			cpu.setF(flagN, false)
			cpu.setF(flagH, util.Bit(carryBits, 4))
			cpu.setF(flagC, util.Bit(carryBits, 8))
			cpu.Reg.PC += 2
		}
	default:
		panic(fmt.Errorf("error: ADD %d %d", operand1, operand2))
	}
}

// CPL Complement A Register
func (cpu *CPU) CPL(operand1, operand2 int) {
	cpu.Reg.A = ^cpu.Reg.A
	cpu.setF(flagN, true)
	cpu.setF(flagH, true)
	cpu.Reg.PC++
}

// PREFIXCB is extend instruction
func (cpu *CPU) PREFIXCB(operand1, operand2 int) {
	if operand1 == OPERAND_NONE && operand2 == OPERAND_NONE {
		cpu.Reg.PC++
		cpu.timer(1)
		opcode := prefixCBs[cpu.FetchMemory8(cpu.Reg.PC)]
		instruction, operand1, operand2, cycle := opcode.Ins, opcode.Operand1, opcode.Operand2, opcode.Cycle1

		switch instruction {
		case INS_RLC:
			cpu.RLC(operand1, operand2)
		case INS_RRC:
			cpu.RRC(operand1, operand2)
		case INS_RL:
			cpu.RL(operand1, operand2)
		case INS_RR:
			cpu.RR(operand1, operand2)
		case INS_SLA:
			cpu.SLA(operand1, operand2)
		case INS_SRA:
			cpu.SRA(operand1, operand2)
		case INS_SWAP:
			cpu.SWAP(operand1, operand2)
		case INS_SRL:
			cpu.SRL(operand1, operand2)
		case INS_BIT:
			cpu.BIT(operand1, operand2)
		case INS_RES:
			cpu.RES(operand1, operand2)
		case INS_SET:
			cpu.SET(operand1, operand2)
		default:
			panic(fmt.Errorf("eip: 0x%04x opcode: %v", cpu.Reg.PC, opcode))
		}

		if cycle > 1 {
			cpu.timer(cycle - 1)
		}
	} else {
		panic(fmt.Errorf("error: PREFIXCB %d %d", operand1, operand2))
	}
}

// RLC Rotate n left carry => bit0
func (cpu *CPU) RLC(operand1, operand2 int) {
	var value, bit7 byte
	if operand1 == OPERAND_B && operand2 == OPERAND_NONE {
		value = cpu.Reg.B
		bit7 = value >> 7
		value = (value << 1)
		value = util.SetLSB(value, bit7 != 0)
		cpu.Reg.B = value
	} else if operand1 == OPERAND_C && operand2 == OPERAND_NONE {
		value = cpu.Reg.C
		bit7 = value >> 7
		value = (value << 1)
		value = util.SetLSB(value, bit7 != 0)
		cpu.Reg.C = value
	} else if operand1 == OPERAND_D && operand2 == OPERAND_NONE {
		value = cpu.Reg.D
		bit7 = value >> 7
		value = (value << 1)
		value = util.SetLSB(value, bit7 != 0)
		cpu.Reg.D = value
	} else if operand1 == OPERAND_E && operand2 == OPERAND_NONE {
		value = cpu.Reg.E
		bit7 = value >> 7
		value = (value << 1)
		value = util.SetLSB(value, bit7 != 0)
		cpu.Reg.E = value
	} else if operand1 == OPERAND_H && operand2 == OPERAND_NONE {
		value = cpu.Reg.H
		bit7 = value >> 7
		value = (value << 1)
		value = util.SetLSB(value, bit7 != 0)
		cpu.Reg.H = value
	} else if operand1 == OPERAND_L && operand2 == OPERAND_NONE {
		value = cpu.Reg.L
		bit7 = value >> 7
		value = (value << 1)
		value = util.SetLSB(value, bit7 != 0)
		cpu.Reg.L = value
	} else if operand1 == OPERAND_HL_PAREN && operand2 == OPERAND_NONE {
		value = cpu.FetchMemory8(cpu.Reg.HL())
		cpu.timer(1)
		bit7 = value >> 7
		value = (value << 1)
		value = util.SetLSB(value, bit7 != 0)
		cpu.SetMemory8(cpu.Reg.HL(), value)
		cpu.timer(2)
	} else if operand1 == OPERAND_A && operand2 == OPERAND_NONE {
		value = cpu.Reg.A
		bit7 = value >> 7
		value = (value << 1)
		value = util.SetLSB(value, bit7 != 0)
		cpu.Reg.A = value
	} else {
		panic(fmt.Errorf("error: RLC %d %d", operand1, operand2))
	}

	cpu.setF(flagZ, value == 0)
	cpu.setF(flagN, false)
	cpu.setF(flagH, false)
	cpu.setF(flagC, bit7 != 0)
	cpu.Reg.PC++
}

// RLCA Rotate register A left.
func (cpu *CPU) RLCA(operand1, operand2 int) {
	var value byte
	var bit7 byte
	value = cpu.Reg.A
	bit7 = value >> 7
	value = (value << 1)
	value = util.SetLSB(value, bit7 != 0)
	cpu.Reg.A = value

	cpu.setF(flagZ, false)
	cpu.setF(flagN, false)
	cpu.setF(flagH, false)
	cpu.setF(flagC, bit7 != 0)
	cpu.Reg.PC++
}

// RRC Rotate n right carry => bit7
func (cpu *CPU) RRC(operand1, operand2 int) {
	var value byte
	var bit0 byte
	if operand1 == OPERAND_B && operand2 == OPERAND_NONE {
		value = cpu.Reg.B
		bit0 = value % 2
		value = (value >> 1)
		value = util.SetMSB(value, bit0 != 0)
		cpu.Reg.B = value
	} else if operand1 == OPERAND_C && operand2 == OPERAND_NONE {
		value = cpu.Reg.C
		bit0 = value % 2
		value = (value >> 1)
		value = util.SetMSB(value, bit0 != 0)
		cpu.Reg.C = value
	} else if operand1 == OPERAND_D && operand2 == OPERAND_NONE {
		value = cpu.Reg.D
		bit0 = value % 2
		value = (value >> 1)
		value = util.SetMSB(value, bit0 != 0)
		cpu.Reg.D = value
	} else if operand1 == OPERAND_E && operand2 == OPERAND_NONE {
		value = cpu.Reg.E
		bit0 = value % 2
		value = (value >> 1)
		value = util.SetMSB(value, bit0 != 0)
		cpu.Reg.E = value
	} else if operand1 == OPERAND_H && operand2 == OPERAND_NONE {
		value = cpu.Reg.H
		bit0 = value % 2
		value = (value >> 1)
		value = util.SetMSB(value, bit0 != 0)
		cpu.Reg.H = value
	} else if operand1 == OPERAND_L && operand2 == OPERAND_NONE {
		value = cpu.Reg.L
		bit0 = value % 2
		value = (value >> 1)
		value = util.SetMSB(value, bit0 != 0)
		cpu.Reg.L = value
	} else if operand1 == OPERAND_HL_PAREN && operand2 == OPERAND_NONE {
		value = cpu.FetchMemory8(cpu.Reg.HL())
		cpu.timer(1)
		bit0 = value % 2
		value = (value >> 1)
		value = util.SetMSB(value, bit0 != 0)
		cpu.SetMemory8(cpu.Reg.HL(), value)
		cpu.timer(2)
	} else if operand1 == OPERAND_A && operand2 == OPERAND_NONE {
		value = cpu.Reg.A
		bit0 = value % 2
		value = (value >> 1)
		value = util.SetMSB(value, bit0 != 0)
		cpu.Reg.A = value
	} else {
		panic(fmt.Errorf("error: RRC %d %d", operand1, operand2))
	}
	cpu.setF(flagZ, value == 0)
	cpu.setF(flagN, false)
	cpu.setF(flagH, false)
	cpu.setF(flagC, bit0 != 0)
	cpu.Reg.PC++
}

// RRCA Rotate register A right.
func (cpu *CPU) RRCA(operand1, operand2 int) {
	var value byte
	var lsb bool

	value, lsb = cpu.Reg.A, util.Bit(cpu.Reg.A, 0)
	value = (value >> 1)
	value = util.SetMSB(value, lsb)
	cpu.Reg.A = value

	cpu.setF(flagZ, false)
	cpu.setF(flagN, false)
	cpu.setF(flagH, false)
	cpu.setF(flagC, lsb)
	cpu.Reg.PC++
}

// RL Rotate n rigth through carry bit7 => bit0
func (cpu *CPU) RL(operand1, operand2 int) {
	var value, bit7 byte
	carry := cpu.f(flagC)
	switch operand1 {
	case OPERAND_A:
		value = cpu.Reg.A
		bit7 = value >> 7
		value = (value << 1)
		value = util.SetLSB(value, carry)
		cpu.Reg.A = value
	case OPERAND_B:
		value = cpu.Reg.B
		bit7 = value >> 7
		value = (value << 1)
		value = util.SetLSB(value, carry)
		cpu.Reg.B = value
	case OPERAND_C:
		value = cpu.Reg.C
		bit7 = value >> 7
		value = (value << 1)
		value = util.SetLSB(value, carry)
		cpu.Reg.C = value
	case OPERAND_D:
		value = cpu.Reg.D
		bit7 = value >> 7
		value = (value << 1)
		value = util.SetLSB(value, carry)
		cpu.Reg.D = value
	case OPERAND_E:
		value = cpu.Reg.E
		bit7 = value >> 7
		value = (value << 1)
		value = util.SetLSB(value, carry)
		cpu.Reg.E = value
	case OPERAND_H:
		value = cpu.Reg.H
		bit7 = value >> 7
		value = (value << 1)
		value = util.SetLSB(value, carry)
		cpu.Reg.H = value
	case OPERAND_L:
		value = cpu.Reg.L
		bit7 = value >> 7
		value = (value << 1)
		value = util.SetLSB(value, carry)
		cpu.Reg.L = value
	case OPERAND_HL_PAREN:
		value = cpu.FetchMemory8(cpu.Reg.HL())
		cpu.timer(1)
		bit7 = value >> 7
		value = (value << 1)
		value = util.SetLSB(value, carry)
		cpu.SetMemory8(cpu.Reg.HL(), value)
		cpu.timer(2)
	default:
		panic(fmt.Errorf("error: RL %d %d", operand1, operand2))
	}

	cpu.setF(flagZ, value == 0)
	cpu.setF(flagN, false)
	cpu.setF(flagH, false)
	cpu.setF(flagC, bit7 != 0)
	cpu.Reg.PC++
}

// RLA Rotate register A left through carry.
func (cpu *CPU) RLA(operand1, operand2 int) {
	var value, bit7 byte
	carry := cpu.f(flagC)

	value = cpu.Reg.A
	bit7 = value >> 7
	value = (value << 1)
	value = util.SetLSB(value, carry)
	cpu.Reg.A = value

	cpu.setF(flagZ, false)
	cpu.setF(flagN, false)
	cpu.setF(flagH, false)
	cpu.setF(flagC, bit7 != 0)
	cpu.Reg.PC++
}

// RR Rotate n right through carry bit0 => bit7
func (cpu *CPU) RR(operand1, operand2 int) {
	var value byte
	var lsb bool
	carry := cpu.f(flagC)

	switch operand1 {
	case OPERAND_A:
		value, lsb = cpu.Reg.A, util.Bit(cpu.Reg.A, 0)
		value >>= 1
		value = util.SetMSB(value, carry)
		cpu.Reg.A = value
	case OPERAND_B:
		value, lsb = cpu.Reg.B, util.Bit(cpu.Reg.B, 0)
		value >>= 1
		value = util.SetMSB(value, carry)
		cpu.Reg.B = value
	case OPERAND_C:
		value, lsb = cpu.Reg.C, util.Bit(cpu.Reg.C, 0)
		value >>= 1
		value = util.SetMSB(value, carry)
		cpu.Reg.C = value
	case OPERAND_D:
		value, lsb = cpu.Reg.D, util.Bit(cpu.Reg.D, 0)
		value >>= 1
		value = util.SetMSB(value, carry)
		cpu.Reg.D = value
	case OPERAND_E:
		value, lsb = cpu.Reg.E, util.Bit(cpu.Reg.E, 0)
		value >>= 1
		value = util.SetMSB(value, carry)
		cpu.Reg.E = value
	case OPERAND_H:
		value, lsb = cpu.Reg.H, util.Bit(cpu.Reg.H, 0)
		value >>= 1
		value = util.SetMSB(value, carry)
		cpu.Reg.H = value
	case OPERAND_L:
		value, lsb = cpu.Reg.L, util.Bit(cpu.Reg.L, 0)
		value >>= 1
		value = util.SetMSB(value, carry)
		cpu.Reg.L = value
	case OPERAND_HL_PAREN:
		value = cpu.FetchMemory8(cpu.Reg.HL())
		cpu.timer(1)
		lsb = util.Bit(value, 0)
		value >>= 1
		value = util.SetMSB(value, carry)
		cpu.SetMemory8(cpu.Reg.HL(), value)
		cpu.timer(2)
	default:
		panic(fmt.Errorf("error: RR %d %d", operand1, operand2))
	}

	cpu.setF(flagZ, value == 0)
	cpu.setF(flagN, false)
	cpu.setF(flagH, false)
	cpu.setF(flagC, lsb)
	cpu.Reg.PC++
}

// SLA Shift Left
func (cpu *CPU) SLA(operand1, operand2 int) {
	var value, bit7 byte
	if operand1 == OPERAND_B && operand2 == OPERAND_NONE {
		value = cpu.Reg.B
		bit7 = value >> 7
		value = (value << 1)
		cpu.Reg.B = value
	} else if operand1 == OPERAND_C && operand2 == OPERAND_NONE {
		value = cpu.Reg.C
		bit7 = value >> 7
		value = (value << 1)
		cpu.Reg.C = value
	} else if operand1 == OPERAND_D && operand2 == OPERAND_NONE {
		value = cpu.Reg.D
		bit7 = value >> 7
		value = (value << 1)
		cpu.Reg.D = value
	} else if operand1 == OPERAND_E && operand2 == OPERAND_NONE {
		value = cpu.Reg.E
		bit7 = value >> 7
		value = (value << 1)
		cpu.Reg.E = value
	} else if operand1 == OPERAND_H && operand2 == OPERAND_NONE {
		value = cpu.Reg.H
		bit7 = value >> 7
		value = (value << 1)
		cpu.Reg.H = value
	} else if operand1 == OPERAND_L && operand2 == OPERAND_NONE {
		value = cpu.Reg.L
		bit7 = value >> 7
		value = (value << 1)
		cpu.Reg.L = value
	} else if operand1 == OPERAND_HL_PAREN && operand2 == OPERAND_NONE {
		value = cpu.FetchMemory8(cpu.Reg.HL())
		cpu.timer(1)
		bit7 = value >> 7
		value = (value << 1)
		cpu.SetMemory8(cpu.Reg.HL(), value)
		cpu.timer(2)
	} else if operand1 == OPERAND_A && operand2 == OPERAND_NONE {
		value = cpu.Reg.A
		bit7 = value >> 7
		value = (value << 1)
		cpu.Reg.A = value
	} else {
		panic(fmt.Errorf("error: SLA %d %d", operand1, operand2))
	}

	cpu.setF(flagZ, value == 0)
	cpu.setF(flagN, false)
	cpu.setF(flagH, false)
	cpu.setF(flagC, bit7 != 0)
	cpu.Reg.PC++
}

// SRA Shift Right MSBit dosen't change
func (cpu *CPU) SRA(operand1, operand2 int) {
	var value byte
	var lsb, msb bool
	if operand1 == OPERAND_B && operand2 == OPERAND_NONE {
		value, lsb, msb = cpu.Reg.B, util.Bit(cpu.Reg.B, 0), util.Bit(cpu.Reg.B, 7)
		value = (value >> 1)
		value = util.SetMSB(value, msb)
		cpu.Reg.B = value
	} else if operand1 == OPERAND_C && operand2 == OPERAND_NONE {
		value, lsb, msb = cpu.Reg.C, util.Bit(cpu.Reg.C, 0), util.Bit(cpu.Reg.C, 7)
		value = (value >> 1)
		value = util.SetMSB(value, msb)
		cpu.Reg.C = value
	} else if operand1 == OPERAND_D && operand2 == OPERAND_NONE {
		value, lsb, msb = cpu.Reg.D, util.Bit(cpu.Reg.D, 0), util.Bit(cpu.Reg.D, 7)
		value = (value >> 1)
		value = util.SetMSB(value, msb)
		cpu.Reg.D = value
	} else if operand1 == OPERAND_E && operand2 == OPERAND_NONE {
		value, lsb, msb = cpu.Reg.E, util.Bit(cpu.Reg.E, 0), util.Bit(cpu.Reg.E, 7)
		value = (value >> 1)
		value = util.SetMSB(value, msb)
		cpu.Reg.E = value
	} else if operand1 == OPERAND_H && operand2 == OPERAND_NONE {
		value, lsb, msb = cpu.Reg.H, util.Bit(cpu.Reg.H, 0), util.Bit(cpu.Reg.H, 7)
		value = (value >> 1)
		value = util.SetMSB(value, msb)
		cpu.Reg.H = value
	} else if operand1 == OPERAND_L && operand2 == OPERAND_NONE {
		value, lsb, msb = cpu.Reg.L, util.Bit(cpu.Reg.L, 0), util.Bit(cpu.Reg.L, 7)
		value = (value >> 1)
		value = util.SetMSB(value, msb)
		cpu.Reg.L = value
	} else if operand1 == OPERAND_HL_PAREN && operand2 == OPERAND_NONE {
		value = cpu.FetchMemory8(cpu.Reg.HL())
		cpu.timer(1)
		lsb, msb = util.Bit(value, 0), util.Bit(value, 7)
		value = (value >> 1)
		value = util.SetMSB(value, msb)
		cpu.SetMemory8(cpu.Reg.HL(), value)
		cpu.timer(2)
	} else if operand1 == OPERAND_A && operand2 == OPERAND_NONE {
		value, lsb, msb = cpu.Reg.A, util.Bit(cpu.Reg.A, 0), util.Bit(cpu.Reg.A, 7)
		value = (value >> 1)
		value = util.SetMSB(value, msb)
		cpu.Reg.A = value
	} else {
		panic(fmt.Errorf("error: SRA %d %d", operand1, operand2))
	}

	cpu.setF(flagZ, value == 0)
	cpu.setF(flagN, false)
	cpu.setF(flagH, false)
	cpu.setF(flagC, lsb)
	cpu.Reg.PC++
}

// SWAP Swap n[5:8] and n[0:4]
func (cpu *CPU) SWAP(operand1, operand2 int) {
	var value byte

	switch operand1 {
	case OPERAND_B:
		B := cpu.Reg.B
		B03 := B & 0x0f
		B47 := B >> 4
		value = (B03 << 4) | B47
		cpu.Reg.B = value
	case OPERAND_C:
		C := cpu.Reg.C
		C03 := C & 0x0f
		C47 := C >> 4
		value = (C03 << 4) | C47
		cpu.Reg.C = value
	case OPERAND_D:
		D := cpu.Reg.D
		D03 := D & 0x0f
		D47 := D >> 4
		value = (D03 << 4) | D47
		cpu.Reg.D = value
	case OPERAND_E:
		E := cpu.Reg.E
		E03 := E & 0x0f
		E47 := E >> 4
		value = (E03 << 4) | E47
		cpu.Reg.E = value
	case OPERAND_H:
		H := cpu.Reg.H
		H03 := H & 0x0f
		H47 := H >> 4
		value = (H03 << 4) | H47
		cpu.Reg.H = value
	case OPERAND_L:
		L := cpu.Reg.L
		L03 := L & 0x0f
		L47 := L >> 4
		value = (L03 << 4) | L47
		cpu.Reg.L = value
	case OPERAND_HL_PAREN:
		data := cpu.FetchMemory8(cpu.Reg.HL())
		cpu.timer(1)
		data03 := data & 0x0f
		data47 := data >> 4
		value = (data03 << 4) | data47
		cpu.SetMemory8(cpu.Reg.HL(), value)
		cpu.timer(2)
	case OPERAND_A:
		A := cpu.Reg.A
		A03 := A & 0x0f
		A47 := A >> 4
		value = (A03 << 4) | A47
		cpu.Reg.A = value
	default:
		panic(fmt.Errorf("error: SWAP %d %d", operand1, operand2))
	}

	cpu.setF(flagZ, value == 0)
	cpu.setF(flagN, false)
	cpu.setF(flagH, false)
	cpu.setF(flagC, false)
	cpu.Reg.PC++
}

// SRL Shift Right MSBit = 0
func (cpu *CPU) SRL(operand1, operand2 int) {
	var value byte
	var bit0 byte

	switch operand1 {
	case OPERAND_B:
		value = cpu.Reg.B
		bit0 = value % 2
		value = (value >> 1)
		cpu.Reg.B = value
	case OPERAND_C:
		value = cpu.Reg.C
		bit0 = value % 2
		value = (value >> 1)
		cpu.Reg.C = value
	case OPERAND_D:
		value = cpu.Reg.D
		bit0 = value % 2
		value = (value >> 1)
		cpu.Reg.D = value
	case OPERAND_E:
		value = cpu.Reg.E
		bit0 = value % 2
		value = (value >> 1)
		cpu.Reg.E = value
	case OPERAND_H:
		value = cpu.Reg.H
		bit0 = value % 2
		value = (value >> 1)
		cpu.Reg.H = value
	case OPERAND_L:
		value = cpu.Reg.L
		bit0 = value % 2
		value = (value >> 1)
		cpu.Reg.L = value
	case OPERAND_HL_PAREN:
		value = cpu.FetchMemory8(cpu.Reg.HL())
		cpu.timer(1)
		bit0 = value % 2
		value = (value >> 1)
		cpu.SetMemory8(cpu.Reg.HL(), value)
		cpu.timer(2)
	case OPERAND_A:
		value = cpu.Reg.A
		bit0 = value % 2
		value = (value >> 1)
		cpu.Reg.A = value
	default:
		panic(fmt.Errorf("error: SRL %d %d", operand1, operand2))
	}

	cpu.setF(flagZ, value == 0)
	cpu.setF(flagN, false)
	cpu.setF(flagH, false)
	cpu.setF(flagC, bit0 == 1)
	cpu.Reg.PC++
}

// BIT Test bit n
func (cpu *CPU) BIT(operand1, operand2 int) {
	var value bool
	targetBit := operand1 - OPERAND_0
	switch operand2 {
	case OPERAND_B:
		value = util.Bit(cpu.Reg.B, targetBit)
	case OPERAND_C:
		value = util.Bit(cpu.Reg.C, targetBit)
	case OPERAND_D:
		value = util.Bit(cpu.Reg.D, targetBit)
	case OPERAND_E:
		value = util.Bit(cpu.Reg.E, targetBit)
	case OPERAND_H:
		value = util.Bit(cpu.Reg.H, targetBit)
	case OPERAND_L:
		value = util.Bit(cpu.Reg.L, targetBit)
	case OPERAND_HL_PAREN:
		value = util.Bit(cpu.FetchMemory8(cpu.Reg.HL()), targetBit)
	case OPERAND_A:
		value = util.Bit(cpu.Reg.A, targetBit)
	}

	cpu.setF(flagZ, !value)
	cpu.setF(flagN, false)
	cpu.setF(flagH, true)
	cpu.Reg.PC++
}

// RES Clear bit n
func (cpu *CPU) RES(operand1, operand2 int) {
	targetBit := operand1 - OPERAND_0
	switch operand2 {
	case OPERAND_B:
		mask := ^(byte(1) << targetBit)
		cpu.Reg.B &= mask
	case OPERAND_C:
		mask := ^(byte(1) << targetBit)
		cpu.Reg.C &= mask
	case OPERAND_D:
		mask := ^(byte(1) << targetBit)
		cpu.Reg.D &= mask
	case OPERAND_E:
		mask := ^(byte(1) << targetBit)
		cpu.Reg.E &= mask
	case OPERAND_H:
		mask := ^(byte(1) << targetBit)
		cpu.Reg.H &= mask
	case OPERAND_L:
		mask := ^(byte(1) << targetBit)
		cpu.Reg.L &= mask
	case OPERAND_HL_PAREN:
		mask := ^(byte(1) << targetBit)
		value := cpu.FetchMemory8(cpu.Reg.HL()) & mask
		cpu.timer(1)
		cpu.SetMemory8(cpu.Reg.HL(), value)
		cpu.timer(2)
	case OPERAND_A:
		mask := ^(byte(1) << targetBit)
		cpu.Reg.A &= mask
	}
	cpu.Reg.PC++
}

// SET Clear bit n
func (cpu *CPU) SET(operand1, operand2 int) {
	targetBit := operand1 - OPERAND_0
	switch operand2 {
	case OPERAND_B:
		mask := byte(1) << targetBit
		cpu.Reg.B |= mask
	case OPERAND_C:
		mask := byte(1) << targetBit
		cpu.Reg.C |= mask
	case OPERAND_D:
		mask := byte(1) << targetBit
		cpu.Reg.D |= mask
	case OPERAND_E:
		mask := byte(1) << targetBit
		cpu.Reg.E |= mask
	case OPERAND_H:
		mask := byte(1) << targetBit
		cpu.Reg.H |= mask
	case OPERAND_L:
		mask := byte(1) << targetBit
		cpu.Reg.L |= mask
	case OPERAND_HL_PAREN:
		mask := byte(1) << targetBit
		value := cpu.FetchMemory8(cpu.Reg.HL()) | mask
		cpu.timer(1)
		cpu.SetMemory8(cpu.Reg.HL(), value)
		cpu.timer(2)
	case OPERAND_A:
		mask := byte(1) << targetBit
		cpu.Reg.A |= mask
	}
	cpu.Reg.PC++
}

// PUSH value
func (cpu *CPU) PUSH(operand1, operand2 int) {
	cpu.timer(1)
	switch operand1 {
	case OPERAND_BC:
		cpu.pushBC()
	case OPERAND_DE:
		cpu.pushDE()
	case OPERAND_HL:
		cpu.pushHL()
	case OPERAND_AF:
		cpu.pushAF()
	default:
		panic(fmt.Errorf("error: PUSH %d %d", operand1, operand2))
	}
	cpu.Reg.PC++
	cpu.timer(2)
}

// POP value
func (cpu *CPU) POP(operand1, operand2 int) {
	switch operand1 {
	case OPERAND_BC:
		cpu.popBC()
	case OPERAND_DE:
		cpu.popDE()
	case OPERAND_HL:
		cpu.popHL()
	case OPERAND_AF:
		cpu.popAF()
	default:
		panic(fmt.Errorf("error: POP %d %d", operand1, operand2))
	}
	cpu.Reg.PC++
	cpu.timer(2)
}

// SUB subtract
func (cpu *CPU) SUB(operand1, operand2 int) {
	switch operand1 {
	case OPERAND_A:
		cpu.setCSub(cpu.Reg.A, cpu.Reg.A)
		cpu.Reg.A = 0
		cpu.setF(flagZ, true)
		cpu.setF(flagN, true)
		cpu.setF(flagH, false)
		cpu.Reg.PC++
	case OPERAND_B:
		value := cpu.Reg.A - cpu.Reg.B
		carryBits := cpu.Reg.A ^ cpu.Reg.B ^ value
		cpu.setCSub(cpu.Reg.A, cpu.Reg.B)
		cpu.Reg.A = value
		cpu.setF(flagZ, value == 0)
		cpu.setF(flagN, true)
		cpu.setF(flagH, util.Bit(carryBits, 4))
		cpu.Reg.PC++
	case OPERAND_C:
		value := cpu.Reg.A - cpu.Reg.C
		carryBits := cpu.Reg.A ^ cpu.Reg.C ^ value
		cpu.setCSub(cpu.Reg.A, cpu.Reg.C)
		cpu.Reg.A = value
		cpu.setF(flagZ, value == 0)
		cpu.setF(flagN, true)
		cpu.setF(flagH, util.Bit(carryBits, 4))
		cpu.Reg.PC++
	case OPERAND_D:
		value := cpu.Reg.A - cpu.Reg.D
		carryBits := cpu.Reg.A ^ cpu.Reg.D ^ value
		cpu.setCSub(cpu.Reg.A, cpu.Reg.D)
		cpu.Reg.A = value
		cpu.setF(flagZ, value == 0)
		cpu.setF(flagN, true)
		cpu.setF(flagH, util.Bit(carryBits, 4))
		cpu.Reg.PC++
	case OPERAND_E:
		value := cpu.Reg.A - cpu.Reg.E
		carryBits := cpu.Reg.A ^ cpu.Reg.E ^ value
		cpu.setCSub(cpu.Reg.A, cpu.Reg.E)
		cpu.Reg.A = value
		cpu.setF(flagZ, value == 0)
		cpu.setF(flagN, true)
		cpu.setF(flagH, util.Bit(carryBits, 4))
		cpu.Reg.PC++
	case OPERAND_H:
		value := cpu.Reg.A - cpu.Reg.H
		carryBits := cpu.Reg.A ^ cpu.Reg.H ^ value
		cpu.setCSub(cpu.Reg.A, cpu.Reg.H)
		cpu.Reg.A = value
		cpu.setF(flagZ, value == 0)
		cpu.setF(flagN, true)
		cpu.setF(flagH, util.Bit(carryBits, 4))
		cpu.Reg.PC++
	case OPERAND_L:
		value := cpu.Reg.A - cpu.Reg.L
		carryBits := cpu.Reg.A ^ cpu.Reg.L ^ value
		cpu.setCSub(cpu.Reg.A, cpu.Reg.L)
		cpu.Reg.A = value
		cpu.setF(flagZ, value == 0)
		cpu.setF(flagN, true)
		cpu.setF(flagH, util.Bit(carryBits, 4))
		cpu.Reg.PC++
	case OPERAND_d8:
		value := cpu.Reg.A - cpu.d8Fetch()
		carryBits := cpu.Reg.A ^ cpu.d8Fetch() ^ value
		cpu.setCSub(cpu.Reg.A, cpu.d8Fetch())
		cpu.Reg.A = value
		cpu.setF(flagZ, value == 0)
		cpu.setF(flagN, true)
		cpu.setF(flagH, util.Bit(carryBits, 4))
		cpu.Reg.PC += 2
	case OPERAND_HL_PAREN:
		value := cpu.Reg.A - cpu.FetchMemory8(cpu.Reg.HL())
		carryBits := cpu.Reg.A ^ cpu.FetchMemory8(cpu.Reg.HL()) ^ value
		cpu.setCSub(cpu.Reg.A, cpu.FetchMemory8(cpu.Reg.HL()))
		cpu.Reg.A = value
		cpu.setF(flagZ, value == 0)
		cpu.setF(flagN, true)
		cpu.setF(flagH, util.Bit(carryBits, 4))
		cpu.Reg.PC++
	default:
		panic(fmt.Errorf("error: SUB %d %d", operand1, operand2))
	}
}

// RRA Rotate register A right through carry.
func (cpu *CPU) RRA(operand1, operand2 int) {
	carry := cpu.f(flagC)
	regA := cpu.Reg.A
	cpu.setF(flagC, util.Bit(regA, 0))
	if carry {
		regA = (1 << 7) | (regA >> 1)
	} else {
		regA = (0 << 7) | (regA >> 1)
	}
	cpu.Reg.A = regA
	cpu.setF(flagZ, false)
	cpu.setF(flagN, false)
	cpu.setF(flagH, false)
	cpu.Reg.PC++
}

// ADC Add the value n8 plus the carry flag to A
func (cpu *CPU) ADC(operand1, operand2 int) {
	var carry, value, value4 byte
	var value16 uint16
	if cpu.f(flagC) {
		carry = 1
	} else {
		carry = 0
	}

	switch operand1 {
	case OPERAND_A:
		switch operand2 {
		case OPERAND_A:
			value = cpu.Reg.A + carry + cpu.Reg.A
			value4 = (cpu.Reg.A & 0b1111) + carry + (cpu.Reg.A & 0b1111)
			value16 = uint16(cpu.Reg.A) + uint16(cpu.Reg.A) + uint16(carry)
		case OPERAND_B:
			value = cpu.Reg.B + carry + cpu.Reg.A
			value4 = (cpu.Reg.B & 0b1111) + carry + (cpu.Reg.A & 0b1111)
			value16 = uint16(cpu.Reg.B) + uint16(carry) + uint16(cpu.Reg.A)
		case OPERAND_C:
			value = cpu.Reg.C + carry + cpu.Reg.A
			value4 = (cpu.Reg.C & 0b1111) + carry + (cpu.Reg.A & 0b1111)
			value16 = uint16(cpu.Reg.C) + uint16(carry) + uint16(cpu.Reg.A)
		case OPERAND_D:
			value = cpu.Reg.D + carry + cpu.Reg.A
			value4 = (cpu.Reg.D & 0b1111) + carry + (cpu.Reg.A & 0b1111)
			value16 = uint16(cpu.Reg.D) + uint16(carry) + uint16(cpu.Reg.A)
		case OPERAND_E:
			value = cpu.Reg.E + carry + cpu.Reg.A
			value4 = (cpu.Reg.E & 0b1111) + carry + (cpu.Reg.A & 0b1111)
			value16 = uint16(cpu.Reg.E) + uint16(carry) + uint16(cpu.Reg.A)
		case OPERAND_H:
			value = cpu.Reg.H + carry + cpu.Reg.A
			value4 = (cpu.Reg.H & 0b1111) + carry + (cpu.Reg.A & 0b1111)
			value16 = uint16(cpu.Reg.H) + uint16(carry) + uint16(cpu.Reg.A)
		case OPERAND_L:
			value = cpu.Reg.L + carry + cpu.Reg.A
			value4 = (cpu.Reg.L & 0b1111) + carry + (cpu.Reg.A & 0b1111)
			value16 = uint16(cpu.Reg.L) + uint16(carry) + uint16(cpu.Reg.A)
		case OPERAND_HL_PAREN:
			data := cpu.FetchMemory8(cpu.Reg.HL())
			value = data + carry + cpu.Reg.A
			value4 = (data & 0x0f) + carry + (cpu.Reg.A & 0b1111)
			value16 = uint16(data) + uint16(cpu.Reg.A) + uint16(carry)
		case OPERAND_d8:
			data := cpu.d8Fetch()
			value = data + carry + cpu.Reg.A
			value4 = (data & 0x0f) + carry + (cpu.Reg.A & 0b1111)
			value16 = uint16(data) + uint16(cpu.Reg.A) + uint16(carry)
			cpu.Reg.PC++
		}
	default:
		panic(fmt.Errorf("error: ADC %d %d", operand1, operand2))
	}
	cpu.Reg.A = value
	cpu.setF(flagZ, value == 0)
	cpu.setF(flagN, false)
	cpu.setF(flagH, util.Bit(value4, 4))
	cpu.setF(flagC, util.Bit(value16, 8))
	cpu.Reg.PC++
}

// SBC Subtract the value n8 and the carry flag from A
func (cpu *CPU) SBC(operand1, operand2 int) {
	var carry, value, value4 byte
	var value16 uint16
	if cpu.f(flagC) {
		carry = 1
	} else {
		carry = 0
	}

	switch operand1 {
	case OPERAND_A:
		switch operand2 {
		case OPERAND_A:
			value = cpu.Reg.A - (cpu.Reg.A + carry)
			value4 = (cpu.Reg.A & 0b1111) - ((cpu.Reg.A & 0b1111) + carry)
			value16 = uint16(cpu.Reg.A) - (uint16(cpu.Reg.A) + uint16(carry))
		case OPERAND_B:
			value = cpu.Reg.A - (cpu.Reg.B + carry)
			value4 = (cpu.Reg.A & 0b1111) - ((cpu.Reg.B & 0b1111) + carry)
			value16 = uint16(cpu.Reg.A) - (uint16(cpu.Reg.B) + uint16(carry))
		case OPERAND_C:
			value = cpu.Reg.A - (cpu.Reg.C + carry)
			value4 = (cpu.Reg.A & 0b1111) - ((cpu.Reg.C & 0b1111) + carry)
			value16 = uint16(cpu.Reg.A) - (uint16(cpu.Reg.C) + uint16(carry))
		case OPERAND_D:
			value = cpu.Reg.A - (cpu.Reg.D + carry)
			value4 = (cpu.Reg.A & 0b1111) - ((cpu.Reg.D & 0b1111) + carry)
			value16 = uint16(cpu.Reg.A) - (uint16(cpu.Reg.D) + uint16(carry))
		case OPERAND_E:
			value = cpu.Reg.A - (cpu.Reg.E + carry)
			value4 = (cpu.Reg.A & 0b1111) - ((cpu.Reg.E & 0b1111) + carry)
			value16 = uint16(cpu.Reg.A) - (uint16(cpu.Reg.E) + uint16(carry))
		case OPERAND_H:
			value = cpu.Reg.A - (cpu.Reg.H + carry)
			value4 = (cpu.Reg.A & 0b1111) - ((cpu.Reg.H & 0b1111) + carry)
			value16 = uint16(cpu.Reg.A) - (uint16(cpu.Reg.H) + uint16(carry))
		case OPERAND_L:
			value = cpu.Reg.A - (cpu.Reg.L + carry)
			value4 = (cpu.Reg.A & 0b1111) - ((cpu.Reg.L & 0b1111) + carry)
			value16 = uint16(cpu.Reg.A) - (uint16(cpu.Reg.L) + uint16(carry))
		case OPERAND_HL_PAREN:
			data := cpu.FetchMemory8(cpu.Reg.HL())
			value = cpu.Reg.A - (data + carry)
			value4 = (cpu.Reg.A & 0b1111) - ((data & 0x0f) + carry)
			value16 = uint16(cpu.Reg.A) - (uint16(data) + uint16(carry))
		case OPERAND_d8:
			data := cpu.d8Fetch()
			value = cpu.Reg.A - (data + carry)
			value4 = (cpu.Reg.A & 0b1111) - ((data & 0x0f) + carry)
			value16 = uint16(cpu.Reg.A) - (uint16(data) + uint16(carry))
			cpu.Reg.PC++
		}
	default:
		panic(fmt.Errorf("error: SBC %d %d", operand1, operand2))
	}
	cpu.Reg.A = value
	cpu.setF(flagZ, value == 0)
	cpu.setF(flagN, true)
	cpu.setF(flagH, util.Bit(value4, 4))
	cpu.setF(flagC, util.Bit(value16, 8))
	cpu.Reg.PC++
}

// DAA Decimal adjust
func (cpu *CPU) DAA(operand1, operand2 int) {
	A := uint8(cpu.Reg.A)
	// ref: https://forums.nesdev.com/viewtopic.php?f=20&t=15944
	if !cpu.f(flagN) {
		if cpu.f(flagC) || A > 0x99 {
			A += 0x60
			cpu.setF(flagC, true)
		}
		if cpu.f(flagH) || (A&0x0f) > 0x09 {
			A += 0x06
		}
	} else {
		if cpu.f(flagC) {
			A -= 0x60
		}
		if cpu.f(flagH) {
			A -= 0x06
		}
	}

	cpu.Reg.A = A
	cpu.setF(flagZ, A == 0)
	cpu.setF(flagH, false)
	cpu.Reg.PC++
}

// RST Push present address and jump to vector address
func (cpu *CPU) RST(operand1, operand2 int) {
	vector := map[int]uint16{OPERAND_00H: 0x00, OPERAND_08H: 0x08, OPERAND_10H: 0x10, OPERAND_18H: 0x18, OPERAND_20H: 0x20, OPERAND_28H: 0x28, OPERAND_30H: 0x30, OPERAND_38H: 0x38}[operand1]
	destination := uint16(vector)
	cpu.Reg.PC++
	cpu.pushPC()
	cpu.Reg.PC = destination
}

// SCF Set Carry Flag
func (cpu *CPU) SCF(operand1, operand2 int) {
	cpu.setF(flagN, false)
	cpu.setF(flagH, false)
	cpu.setF(flagC, true)
	cpu.Reg.PC++
}

// CCF Complement Carry Flag
func (cpu *CPU) CCF(operand1, operand2 int) {
	cpu.setF(flagN, false)
	cpu.setF(flagH, false)
	cpu.setF(flagC, !cpu.f(flagC))
	cpu.Reg.PC++
}