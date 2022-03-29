//* Copyright (c) 2020, Alex Lewontin
//* All rights reserved.
//*
//* Redistribution and use in source and binary forms, with or without
//* modification, are permitted provided that the following conditions are met:
//*
//* - Redistributions of source code must retain the above copyright notice, this
//* list of conditions and the following disclaimer.
//* - Redistributions in binary form must reproduce the above copyright notice,
//* this list of conditions and the following disclaimer in the documentation
//* and/or other materials provided with the distribution.
//*
//* THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
//* ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
//* WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
//* DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
//* FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
//* DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
//* SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
//* CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
//* OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
//* OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package riverboat

import (
	. "github.com/alexclewontin/riverboat/eval"
	"github.com/google/uuid"
)

type player struct {
	Username   string
	ID         string
	Position   uint
	Ready      bool
	In         bool
	Called     bool
	Left       bool
	TotalBuyIn uint
	Stack      uint
	Bet        uint
	TotalBet   uint
	Cards      [2]Card
}

func (p *player) allIn() bool {
	return p.In && (p.Stack == 0)
}

func (p *player) initialize() {
	*p = player{}

	p.ID = uuid.New().String()
	p.Ready = false
	p.In = false
	p.Called = false

}

//putInChips is simply a helper function that transfers the amounts between fields
func (p *player) putInChips(amt uint) {
	if p.Stack > amt {
		p.Bet += amt
		p.TotalBet += amt
		p.Stack -= amt
	} else {
		p.Bet += p.Stack
		p.TotalBet += p.Stack
		p.Stack = 0
	}
}

func (p *player) returnChips(amt uint) {
	if p.TotalBet > amt {
		p.TotalBet -= amt
		p.Stack += amt
	} else {
		p.Stack += p.TotalBet
		p.TotalBet = 0
	}
}
