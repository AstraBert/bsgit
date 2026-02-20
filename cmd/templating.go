package cmd

import (
	"bytes"
	"text/template"
	"time"
)

const TemplateReadme string = "# {{.ProjectName}}\n\n{{.Description}}"
const TemplateContributing string = `# Contributing to {{.ProjectName}}

Thank you for your interest in contributing to this project! Please review these guidelines before getting started.

## Issue Reporting

### When to Report an Issue

- You've discovered bugs but lack the knowledge or time to fix them
- You have feature requests but cannot implement them yourself

> ⚠️ **Important:** Always search existing open and closed issues before submitting to avoid duplicates.

### How to Report an Issue

1. Open a new issue
2. Provide a clear, concise title that describes the problem or feature request
3. Include a detailed description of the issue or requested feature

## Code Contributions

### When to Contribute

- You've identified and fixed bugs
- You've optimized or improved existing code
- You've developed new features that would benefit the community

### How to Contribute

1. **Fork the repository and check out a secondary branch**

2. **Make your changes and test**

   Ensure the build succeeds and all tests pass. Add tests for new features.

4. **Verify formatting and linting compliance**
   Ensure your changes pass all linting checks.

5. **Commit your changes**

6. **Submit a pull request**
   Include a comprehensive description of your changes.

---

**Thank you for contributing!**
`

const TemplateLicense string = `The MIT License

Copyright (c) {{.OwnerName}} {{.Year}}

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.`

type Readme struct {
	ProjectName string
	Description string
}

type Contributing struct {
	ProjectName string
}

type License struct {
	Year      int
	OwnerName string
}

func NewLicense(ownerName string) License {
	year := time.Now().Year()
	return License{
		Year:      year,
		OwnerName: ownerName,
	}
}

func GetTemplateBuilder(name, t string) *template.Template {
	return template.Must(template.New(name).Parse(t))
}

func BuildReadme(readme Readme) string {
	t := GetTemplateBuilder("readme", TemplateReadme)
	var buf bytes.Buffer
	t.Execute(&buf, readme)
	return buf.String()
}

func BuildContributing(contributing Contributing) string {
	t := GetTemplateBuilder("contributing", TemplateContributing)
	var buf bytes.Buffer
	t.Execute(&buf, contributing)
	return buf.String()
}

func BuildLicense(license License) string {
	t := GetTemplateBuilder("license", TemplateLicense)
	var buf bytes.Buffer
	t.Execute(&buf, license)
	return buf.String()
}
