# Go Resume

**_Go Resume_** is a resume tailoring tool with super powers 🚀

## Building 🛠

### Dependencies

- [Go](https://golang.org)
- [NodeJS](https://nodejs.org)
- [Latex](https://www.latex-project.org/)

### Installation Steps

1. Clone the repo with `git clone`

```sh
git clone --recursive https://github.com/penguingovernor/go-resume.git
```

2. Build the commands with `go build`

```sh
# Builds the resume-start binary
go build ./cmd/resume-start
# Builds the resume-tailor binary
go build ./cmd/resume-tailor
```

## Usage 📚

1. Start the resumake GUI with `resume-start`

```sh
# This will build and the resumake GUI on localhost:3000
./resume-start
```

2. Build your base resume on the GUI and downloads your `resume.json`.

3. Optionally, stop `resume-start`.

4. If `resume-start` was killed, it can now be re-ran with `./resume-start -skip-build -no-client`, this will start instantly.

5. Make a file called `skills.yaml` that is based on a job description.

Example:

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

7. Build your resume with resume tailor

```sh
# This assumes that your files are called resume.json and skills.yaml
./resume-tailor > resume.pdf
```

On subsequent runs do steps 4,5, and 7.

### Options

```sh
./resume-start -help
Usage of ./resume-start:
  -log string
        the file where resume-start logs to (default "./.resume-start.log")
  -no-client
        disable running the client
  -resumake-dir string
        the directory where resumake.io resides (default "./resumake.io")
  -skip-build
        skip building resumake.io dependencies
```

```sh
./resume-tailor -help
Usage: ./resume-tailor [OPTIONS] [resume.json skills.yaml] > resume.{pdf,json}

[OPTIONS]
  -JSON
        output JSON instead of PDF data
  -URL string
        the API endpoint of resumake.io (default "http://localhost:3001/api/generate/resume")
  -force-single
        force the resulting resume to be a single page
  -template int
        which template to use [values: 1-9] (default 6)
```

## Warnings ⚠

Under active development...

## F.A.Q

> Q: Do I _really_ need to install `Node.js` and `Latex`?

> A: Read the design document `docs/bin/go-resume-design.pdf` and find out for yourself 😉

> Q: What do the resumes look like?

> A: https://resumake.io/generator/templates
