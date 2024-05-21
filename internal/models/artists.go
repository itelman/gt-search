package models

import (
	"encoding/json"
	"functions/internal/api"
)

type Artist struct {
	ID               int      `json:"id"`
	Image            string   `json:"image"`
	Name             string   `json:"name"`
	Members          []string `json:"members"`
	CreationDate     int      `json:"creationDate"`
	FirstAlbum       string   `json:"firstAlbum"`
	LocationsLink    string   `json:"locations"`
	ConcertDatesLink string   `json:"concertDates"`
	RelationsLink    string   `json:"relations"`

	Locations Locations
	Dates     Dates
	Relations Relations
}

type Locations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`

	LatLngs map[string][]float64
}

type Dates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relations struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func GetAllArtists() ([]*Artist, error) {
	var artists []*Artist
	body, err := api.ParseJson("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &artists); err != nil {
		return nil, err
	}

	return artists, nil
}

func (artist *Artist) GetArtistLocations() (Locations, error) {
	var artistLocations Locations
	body, err := api.ParseJson(artist.LocationsLink)
	if err != nil {
		return Locations{}, err
	}

	err = json.Unmarshal(body, &artistLocations)
	if err != nil {
		return Locations{}, err
	}

	return artistLocations, nil
}

func (artist *Artist) GetArtistDates() (Dates, error) {
	var artistDates Dates
	body, err := api.ParseJson(artist.ConcertDatesLink)
	if err != nil {
		return Dates{}, err
	}

	err = json.Unmarshal(body, &artistDates)
	if err != nil {
		return Dates{}, err
	}

	return artistDates, nil
}

func (artist *Artist) GetArtistRelations() (Relations, error) {
	var artistRelations Relations
	body, err := api.ParseJson(artist.RelationsLink)
	if err != nil {
		return Relations{}, err
	}

	err = json.Unmarshal(body, &artistRelations)
	if err != nil {
		return Relations{}, err
	}

	return artistRelations, nil
}

func (artist *Artist) GetFullInfo() error {
	locations, err := artist.GetArtistLocations()
	if err != nil {
		return err
	}
	artist.Locations = locations

	dates, err := artist.GetArtistDates()
	if err != nil {
		return err
	}
	artist.Dates = dates

	relations, err := artist.GetArtistRelations()
	if err != nil {
		return err
	}
	artist.Relations = relations

	return nil
}
