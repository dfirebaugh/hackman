package store

import (
	"hackman/internal/model"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

var (
	MemberCache     []model.Member
	CurrentList     []model.Member
	CurrentMember   model.Member
	CurrentUsername string
	CurrentUserPass string
)

func GetMembers() {
	CurrentList = MemberCache
}

func Search(term string) []model.Member {
	return slowSearch(term)
}

func slowSearch(term string) []model.Member {
	members := []model.Member{}
	if len(MemberCache) == 0 {
		GetMembers()
	}

	for _, m := range MemberCache {
		if isMatch(term, m.Name) { // checkmatch
			members = append(members, m)
		}
	}
	CurrentList = members
	return members
}

func isMatch(term string, name string) bool {
	return fuzzy.Match(term, name)
}
