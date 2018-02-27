package components

import (
	"fmt"
	"testing"
)

func TestInfluenceBagShuffle(t *testing.T) {
	var bag InfluenceBag

	bag.Influence = append(bag.Influence, Skill)
	bag.Influence = append(bag.Influence, Skill)
	bag.Influence = append(bag.Influence, Skill)
	bag.Influence = append(bag.Influence, Skill)
	bag.Influence = append(bag.Influence, Persuasivness)
	bag.Influence = append(bag.Influence, Persuasivness)
	bag.Influence = append(bag.Influence, Faith)
	bag.Influence = append(bag.Influence, Faith)
	bag.Influence = append(bag.Influence, Faith)

	var preShuffle = fmt.Sprint(bag.Influence)
	bag.shuffle()
	var postShuffle = fmt.Sprint(bag.Influence)
	fmt.Println(preShuffle)
	fmt.Println(postShuffle)

	if preShuffle == postShuffle {
		t.Errorf("shuffle(%q) == %q", preShuffle, postShuffle)
	}

}
