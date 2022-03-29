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
	"testing"
)

func TestIntegration_Scenarios(t *testing.T) {

	t.Run("Scenario 1", func(t *testing.T) {
		g := NewGame()

		pn_a := g.AddPlayer()
		g.AddPlayer()
		g.AddPlayer()

		err := Deal(g, pn_a, 0)

		if err != ErrIllegalAction {
			t.Error("Test failed - Deal must return ErrIllegalAction as 0 players are marked ready.")
		}
	})

	t.Run("Scenario 2", func(t *testing.T) {
		g := NewGame()

		pn_a := g.AddPlayer()
		g.AddPlayer()
		g.AddPlayer()

		err := ToggleReady(g, pn_a, 0)

		if err != ErrIllegalAction {
			t.Error("Test failed - ToggleReady must return ErrIllegalAction as player 0 has not bought in.")
		}
	})

	t.Run("Scenario 3", func(t *testing.T) {
		var err error
		g := NewGame()

		pn_a := g.AddPlayer()
		pn_b := g.AddPlayer()
		pn_c := g.AddPlayer()

		err = BuyIn(g, pn_a, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = BuyIn(g, pn_b, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = BuyIn(g, pn_c, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = ToggleReady(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		err = ToggleReady(g, pn_b, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		err = ToggleReady(g, pn_c, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		err = Deal(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - error dealing: %s", err)
		}
	})

	t.Run("Scenario 4", func(t *testing.T) {
		var err error
		g := NewGame()

		pn_a := g.AddPlayer()
		pn_b := g.AddPlayer()
		pn_c := g.AddPlayer()

		err = BuyIn(g, pn_a, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = BuyIn(g, pn_b, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = BuyIn(g, pn_c, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = ToggleReady(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		err = ToggleReady(g, pn_b, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		err = ToggleReady(g, pn_c, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		err = Deal(g, pn_b, 0)

		if err != ErrIllegalAction {
			t.Errorf("Test failed - must return ErrIllegalAction as pn_b is not the dealer")
		}
	})

	t.Run("Scenario 5", func(t *testing.T) {
		var err error
		g := NewGame()

		pn_a := g.AddPlayer()
		pn_b := g.AddPlayer()
		pn_c := g.AddPlayer()

		err = BuyIn(g, pn_a, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = BuyIn(g, pn_b, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = BuyIn(g, pn_c, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = ToggleReady(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		err = ToggleReady(g, pn_b, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		err = ToggleReady(g, pn_c, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		err = Deal(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - error dealing: %s", err)
		}

		err = Bet(g, pn_a, 25)

		if err != nil {

			t.Errorf("Test failed - error betting: %s", err)
		}

		if g.players[pn_a].Bet != 25 {
			t.Errorf("Betting mechanic not working.")
		}
	})

	t.Run("Scenario 6 simple", func(t *testing.T) {
		var err error
		g := NewGame()

		pn_a := g.AddPlayer()
		pn_b := g.AddPlayer()
		pn_c := g.AddPlayer()

		err = BuyIn(g, pn_a, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = BuyIn(g, pn_b, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = BuyIn(g, pn_c, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = ToggleReady(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		err = ToggleReady(g, pn_b, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		err = ToggleReady(g, pn_c, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		// Preflop

		err = Deal(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - error dealing: %s", err)
		}

		err = Bet(g, pn_a, 25)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_b, 15)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_c, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		// Flop

		err = Bet(g, pn_b, 25)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_c, 25)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_a, 25)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		// Turn

		err = Bet(g, pn_b, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_c, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		//River
		err = Bet(g, pn_b, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_c, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

	})

	t.Run("Scenario 7 fold", func(t *testing.T) {
		var err error
		g := NewGame()

		pn_a := g.AddPlayer()
		pn_b := g.AddPlayer()
		pn_c := g.AddPlayer()

		err = BuyIn(g, pn_a, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = BuyIn(g, pn_b, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = BuyIn(g, pn_c, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = ToggleReady(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		err = ToggleReady(g, pn_b, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		err = ToggleReady(g, pn_c, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		// Preflop

		err = Deal(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - error dealing: %s", err)
		}

		err = Bet(g, pn_a, 25)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_b, 15)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_c, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		// Flop

		err = Bet(g, pn_b, 25)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Fold(g, pn_c, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_a, 25)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		// Turn
		err = Bet(g, pn_b, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		//River

		err = Bet(g, pn_b, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

	})

	t.Run("Scenario 8 reraise", func(t *testing.T) {
		var err error
		g := NewGame()

		pn_a := g.AddPlayer()
		pn_b := g.AddPlayer()
		pn_c := g.AddPlayer()

		err = BuyIn(g, pn_a, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = BuyIn(g, pn_b, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = BuyIn(g, pn_c, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = ToggleReady(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		err = ToggleReady(g, pn_b, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		err = ToggleReady(g, pn_c, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		// Preflop

		err = Deal(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - error dealing: %s", err)
		}

		err = Bet(g, pn_a, 25)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_b, 15)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_c, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		// Flop

		err = Bet(g, pn_b, 25)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Fold(g, pn_c, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_a, 50)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_b, 25)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		// Turn

		err = Bet(g, pn_b, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		//River

		err = Bet(g, pn_b, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

	})
	t.Run("Scenario 9 set username", func(t *testing.T) {
		g := NewGame()

		pn_a := g.AddPlayer()
		g.AddPlayer()
		g.AddPlayer()

		username := "test"

		err := setUsername(g, pn_a, username)

		if err != nil {
			t.Error("Test failed - Could not set username.")
		}
		view := g.GeneratePlayerView(pn_a)

		if view.Players[0].Username != username {
			t.Error("Test failed - Could not set username.")
		}
	})

	t.Run("Scenario 10 add players at positions", func(t *testing.T) {
		g := NewGame()

		pn_a, err := g.AddPlayerPosition(0)
		if err != nil {
			t.Error("Test failed - failed to add player a 0th position")
		}

		_, err = g.AddPlayerPosition(1)
		if err != nil {
			t.Error("Test failed - failed to add player a 1th position")
		}

		_, err = g.AddPlayerPosition(3)
		if err != ErrOutOfBounds {
			t.Error("Test failed - AddPlayerPosition must return ErrOutOfBounds as when adding a player at position 3 to a game with only 2 players.")
		}

		g.AddPlayer()
		view := g.GenerateOmniView()
		if len(view.Players) != 3 {
			t.Error("Test failed - there should be 3 players in the game")
		}

		username := "test"
		setUsername(g, pn_a, username)

		g.AddPlayerPosition(0)
		view = g.GenerateOmniView()
		if len(view.Players) != 4 {
			t.Error("Test failed - there should be 4 players in the game")
		}
		if view.Players[1].Username != username {
			t.Error("Test failed - pn_a should be pushed to position 1 after insertion of player at position 0")
		}

	})

	t.Run("Scenario 11 set player positions", func(t *testing.T) {
		g := NewGame()

		pn_a := g.AddPlayer()
		setUsername(g, pn_a, "pn_a")
		err := setPosition(g, pn_a, 2)
		if err != nil {
			t.Error("Test failed - setting valid position should not cause error")
		}

		pn_b := g.AddPlayer()
		setUsername(g, pn_b, "pn_b")
		err = setPosition(g, pn_b, 5)
		if err != nil {
			t.Error("Test failed - setting valid position should not cause error")
		}

		pn_c := g.AddPlayer()
		setUsername(g, pn_c, "pn_c")
		err = setPosition(g, pn_c, 9)
		if err != nil {
			t.Error("Test failed - setting valid position should not cause error")
		}

		pn_d := g.AddPlayer()
		setUsername(g, pn_d, "pn_d")
		err = setPosition(g, pn_d, 1)
		if err != nil {
			t.Error("Test failed - setting valid position should not cause error")
		}
		if g.players[0].Username != "pn_d" {
			t.Error("Test failed - player order does not match position")
		}

		pn_e := g.AddPlayer()
		err = setPosition(g, pn_e, 9)
		if err != ErrInvalidPosition {
			t.Error("Test failed - setting repeated position should raise error")
		}
	})
}
