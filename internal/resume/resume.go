package resume

// Resume represents the latex resume template.
type Resume struct {
	SelectedTemplate int         `json:"selectedTemplate"`
	Headings         Headings    `json:"headings"`
	Basics           Basics      `json:"basics"`
	Education        []Education `json:"education"`
	Work             []Work      `json:"work"`
	Skills           []Skills    `json:"skills"`
	Projects         []Projects  `json:"projects"`
	Awards           []Awards    `json:"awards"`
	Sections         []string    `json:"sections"`
}

// Headings define what headings to include.
// An empty heading field denotes omission.
type Headings struct {
	Work      string `json:"work"`
	Education string `json:"education"`
	Projects  string `json:"projects"`
	Awards    string `json:"awards"`
	Skills    string `json:"skills"`
}

// Location represents a city,state combination.
type Location struct {
	Address string `json:"address"`
}

// Basics represents basic user information.
type Basics struct {
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Phone    string   `json:"phone"`
	Location Location `json:"location"`
	Website  string   `json:"website"`
}

// Education represents basic education field.
type Education struct {
	Institution string `json:"institution"`
	Location    string `json:"location"`
	Area        string `json:"area"`
	StudyType   string `json:"studyType"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	Gpa         string `json:"gpa"`
}

// Work highlights
type Work struct {
	Position   string   `json:"position"`
	Website    string   `json:"website,omitempty"`
	EndDate    string   `json:"endDate"`
	Highlights []string `json:"highlights"`
	Company    string   `json:"company"`
	Location   string   `json:"location"`
	StartDate  string   `json:"startDate"`
}

// Skills represents a skill that you might have.
// Keywords are usually a subset of that skill.
type Skills struct {
	Level    string   `json:"level,omitempty" yaml:"level,omitempty"`
	Keywords []string `json:"keywords" yaml:"keywords"`
	Name     string   `json:"name" yaml:"name"`
}

// Projects are projects.
type Projects struct {
	Name        string   `json:"name"`
	Keywords    []string `json:"keywords"`
	Description string   `json:"description"`
	URL         string   `json:"url"`
}

// Awards are well, awards.
type Awards struct {
	Title   string `json:"title"`
	Date    string `json:"date"`
	Awarder string `json:"awarder"`
	Summary string `json:"summary"`
}
