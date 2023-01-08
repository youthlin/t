package t

import "golang.org/x/text/language"

func Match(supported []string, userAccept []string) (tag language.Tag, index int, c language.Confidence) {
	return MatchTag(Tags(supported), Tags(userAccept))
}

func MatchTag(supportedTags []language.Tag, userAcceptTags []language.Tag) (tag language.Tag, index int, c language.Confidence) {
	matcher := language.NewMatcher(supportedTags)
	return matcher.Match(userAcceptTags...)
}
