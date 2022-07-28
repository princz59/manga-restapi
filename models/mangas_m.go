package models

type MangaDirectories struct {
	ID          	int    `json:"id"`
	Name        	string `json:"name"`
	Status      	string `json:"status"`
	Genres	    	[]string `json:"genres"`
	URL_PATH		string `json:"url_path"`
	COVER_URL_PATH 	string `json:"cover_url_path"`
}