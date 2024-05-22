package models

import (
	"encoding/json"
	"functions/internal/api"
	"regexp"
	"strconv"
	"strings"
)

type Result struct {
	Response string
	Keyword  string
	Artist   *Artist
}

type Index struct {
	Index []Locations
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
	query = strings.ToLower(query)

	if strings.Contains(strings.ToLower(a.Name), query) {
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
		if strings.Contains(strings.ToLower(member), query) {
			return member, "member"
		}
	}

	if query != "-" {
		artlocs, err := GetLocs(a.ID)
		if err != nil {
			return "", "none"
		}

		for _, location := range artlocs {
			if strings.Contains(location, query) {
				return location, "location"
			}
		}
	}

	return "", "none"
}

func GetLocs(id int) ([]string, error) {
	var index Index

	body, err := api.ParseJsonFile("assets/locations.json")
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &index); err != nil {
		return nil, err
	}

	return index.Index[id-1].Locations, nil
}

func IsNonWord(query string) bool {
	regex := regexp.MustCompile(`(\W|\s){2,}`)

	return regex.MatchString(query) || query == " "
}
