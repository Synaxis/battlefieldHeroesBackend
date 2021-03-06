package gsum

import (
	"github.com/Synaxis/battlefieldHeroesBackend/backend/network"
	"github.com/Synaxis/battlefieldHeroesBackend/backend/network/codec"
)

const (
	gsumGetSessionID = "GetSessionId"
)

// GameSummary probably stands for Game Summary
type GameSummary struct {
	//
}

func (gsum *GameSummary) answer(client *network.Client, pnum uint32, payload interface{}) {
	client.WriteEncode(&codec.Answer{
		Type:         codec.FeslGameSummary,
		PacketNumber: pnum,
		Payload:      payload,
	})
}
