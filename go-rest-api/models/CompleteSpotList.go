package models

type CompleteSpotList struct {
    Records []Record `json:"records"`
    Offset  string   `json:"offset"`
}
type Record struct {
    ID          string `json:"id"`
    Fields      Fields `json:"fields"`
    CreatedTime string `json:"createdTime"`
}
type Fields struct {
    SurfBreak               []string `json:"Surf Break"`
    DifficultyLevel         int64    `json:"Difficulty Level"`
    Destination             string   `json:"Destination"`
    Geocode                 string   `json:"Geocode"`
    Influencers             []string `json:"Influencers"`
    MagicSeaweedLink        string   `json:"Magic Seaweed Link"`
    Photos                  []Photo  `json:"Photos"`
    PeakSurfSeasonBegins    string   `json:"Peak Surf Season Begins"`
    DestinationStateCountry string   `json:"Destination State/Country"`
    PeakSurfSeasonEnds      string   `json:"Peak Surf Season Ends"`
    Address                 string   `json:"Address"`
}
type Photo struct {
    ID         string     `json:"id"`
    URL        string     `json:"url"`
    Filename   string     `json:"filename"`
    Size       int64      `json:"size"`
    Type       string     `json:"type"`
    Thumbnails Thumbnails `json:"thumbnails"`
}
type Thumbnails struct {
    Small Full `json:"small"`
    Large Full `json:"large"`
    Full  Full `json:"full"`
}
type Full struct {
    URL    string `json:"url"`
    Width  int64  `json:"width"`
    Height int64  `json:"height"`
}