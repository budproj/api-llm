package model

type SummarizeKeyResultInput struct {
	Objective   Objective `json:"objective"`
	Cycle       Cycle     `json:"cycle"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Goal        float64   `json:"goal"`
	Format      Format    `json:"format"`
	Owner       Owner     `json:"owner"`
	Comments    []Comment `json:"comments"`
	Checklist   []Task    `json:"checklist"`
	CheckIns    []CheckIn `json:"checkIns"`
}

type Objective struct {
	Title string `json:"title"`
}

type Cycle struct {
	Cadence   string `json:"cadence"`
	DateStart string `json:"dateStart"`
	DateEnd   string `json:"dateEnd"`
}

type Format string

const (
	NUMBER     Format = "NUMBER"
	PERCENTAGE Format = "PERCENTAGE"
	COIN_BRL   Format = "COIN_BRL"
	COIN_USD   Format = "COIN_USD"
	COIN_EUR   Format = "COIN_EUR"
	COIN_GBP   Format = "COIN_GBP"
)

type Owner struct {
	Name string `json:"name"`
}

type Comment struct {
	Text      string `json:"text"`
	Author    string `json:"author"`
	CreatedAt string `json:"createdAt"`
}

type Task struct {
	Description string `json:"description"`
	Owner       string `json:"owner"`
	Done        bool   `json:"done"`
}

type CheckIn struct {
	Comment   string  `json:"comment"`
	Author    string  `json:"author"`
	CreatedAt string  `json:"createdAt"`
	Value     float64 `json:"value"`
}
