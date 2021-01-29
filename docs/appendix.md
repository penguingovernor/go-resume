\newpage

# Appendix

```go
// Skills represents a skill that you might have.
// Keywords are usually a subset of that skill.
type Skills struct {
	Level    string   `json:"level,omitempty" yaml:"level,omitempty"`
	Keywords []string `json:"keywords" yaml:"keywords"`
	Name     string   `json:"name" yaml:"name"`
}
```

: Definition for skills in Go {#lst:skills.go}

```yaml
- name: Programming Languages
  keywords:
    - Go
    - Python
    - Java
    - C
    - C++
- name: Spoken Languages
  keywords:
    - English (Native)
    - Spanish (Native)
    - French (Conversational)
```

: Example yaml file {#lst:skills.yaml}

```json
{
  "selectedTemplate": 1,
  "headings": {
    "work": "Work Experience",
    "education": "Education",
    "projects": "Projects",
    "awards": "",
    "skills": "Skills"
  },
  "basics": {
    "name": "Jorge Henriquez",
    "email": "contact@jorgehenriquez.dev",
    "phone": "(661) 243-7834",
    "location": {
      "address": "Bakersfield, CA"
    },
    "website": "https://jorgehenriquez.dev"
  },
  "education": [
    {
      "institution": "University of California, Santa Cruz",
      "location": "Santa Cruz, CA",
      "area": "Computer Engineering with Honors",
      "studyType": "BS",
      "startDate": "September 2016",
      "endDate": "December 2020",
      "gpa": "3.35"
    }
  ],
  "work": [],
  "skills": [
    {
      "level": "",
      "keywords": ["C/C++", "Go", "Verilog", "Chisel3"],
      "name": "Programming Languages"
    },
    {
      "keywords": ["React", "Babel"],
      "name": "Frameworks and Tools"
    }
  ],
  "projects": [],
  "awards": [
    {
      "title": "",
      "date": "",
      "awarder": "",
      "summary": ""
    }
  ],
  "sections": [
    "templates",
    "profile",
    "education",
    "work",
    "skills",
    "projects",
    "awards"
  ]
}
```

: Sample `resume.json` {#lst:resume.json}
