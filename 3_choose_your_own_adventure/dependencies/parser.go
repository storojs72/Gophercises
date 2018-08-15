package dependencies

type Gopher struct {
	Intro struct {
		Title   string   `json:"title"`
		Story   []string `json:"story"`
		Options []struct {
			Text string `json:"text"`
			Arc  string `json:"arc"`
		} `json:"options"`
	} `json:"intro"`
	NewYork struct {
		Title   string   `json:"title"`
		Story   []string `json:"story"`
		Options []struct {
			Text string `json:"text"`
			Arc  string `json:"arc"`
		} `json:"options"`
	} `json:"new-york"`
	Debate struct {
		Title   string   `json:"title"`
		Story   []string `json:"story"`
		Options []struct {
			Text string `json:"text"`
			Arc  string `json:"arc"`
		} `json:"options"`
	} `json:"debate"`
	SeanKelly struct {
		Title   string   `json:"title"`
		Story   []string `json:"story"`
		Options []struct {
			Text string `json:"text"`
			Arc  string `json:"arc"`
		} `json:"options"`
	} `json:"sean-kelly"`
	MarkBates struct {
		Title   string   `json:"title"`
		Story   []string `json:"story"`
		Options []struct {
			Text string `json:"text"`
			Arc  string `json:"arc"`
		} `json:"options"`
	} `json:"mark-bates"`
	Denver struct {
		Title   string   `json:"title"`
		Story   []string `json:"story"`
		Options []struct {
			Text string `json:"text"`
			Arc  string `json:"arc"`
		} `json:"options"`
	} `json:"denver"`
	Home struct {
		Title   string        `json:"title"`
		Story   []string      `json:"story"`
		Options []interface{} `json:"options"`
	} `json:"home"`
}

