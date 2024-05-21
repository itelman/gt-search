package models

import (
	"regexp"
	"strconv"
	"strings"
)

type Result struct {
	Response string
	Keyword  string
	Artist   *Artist
}

func SetResults(query string) ([]Result, error) {
	if IsNonWord(query) {
		return []Result{}, nil
	}

	artists, err := GetAllArtists()
	if err != nil {
		return nil, err
	}

	var list []Result

	for _, artist := range artists {
		resp, key := artist.Search(query)

		if key != "none" {
			result := Result{Response: resp, Keyword: key, Artist: artist}
			list = append(list, result)
		}
	}

	return list, nil
}

func (a *Artist) Search(query string) (string, string) {
	if strings.Contains(a.Name, query) {
		return a.Name, "artist"
	}

	if query != "-" {
		if strings.Contains(a.FirstAlbum, query) {
			return a.FirstAlbum, "first album"
		}
	}

	creationdate := strconv.Itoa(a.CreationDate)
	if strings.Contains(creationdate, query) {
		return creationdate, "creation date"
	}

	for _, member := range a.Members {
		if strings.Contains(member, query) {
			return member, "member"
		}
	}

	if query != "-" {
		artlocs, err := a.GetArtistLocations()
		if err != nil {
			return "", "none"
		}
		for _, location := range artlocs.Locations {
			if strings.Contains(location, query) {
				return location, "location"
			}
		}
	}

	return "", "none"
}

func IsNonWord(query string) bool {
	spaces := regexp.MustCompile(`^[\s\W_]+`)
	regex1 := regexp.MustCompile(`^(-(\w+))`)

	if regex1.MatchString(query) && !spaces.MatchString(query) {
		return false
	}

	return spaces.MatchString(query)
}
