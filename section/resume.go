package section

import (
	"Go-Latex-Test/utils"
	"bytes"
)

type Resume struct {
	Contact            Contact              `json:"contact"`
	Skills             []Skill              `json:"skills"`
	Education          []Education          `json:"education"`
	WorkExperience     []WorkExperience     `json:"workExperience"`
	ResearchExperience []ResearchExperience `json:"researchExperience"`
	PersonalProjects   []PersonalProject    `json:"personalProjects"`
	Achievements       []MarkdownSnippet    `json:"achievements"`
	ExtraCurriculars   []MarkdownSnippet    `json:"extraCurriculars"`
	PORs               []POR                `json:"positionsOfResponsibility"`
}

func (resume Resume) Prepare() (*bytes.Buffer, error) {
	var resumeBuffer bytes.Buffer
	preamble, _ := utils.ReadFileToBuffer("section/tex/preamble.text")
	contact, _ := PrepareContactSection(resume.Contact)
	education, _ := PrepareEducationSection(resume.Education)
	workExperience, _ := PrepareWorkExperienceSection(resume.WorkExperience)
	researchExperience, _ := PrepareResearchExperienceSection(resume.ResearchExperience)
	personalProjects, _ := PreparePersonalProjectsSection(resume.PersonalProjects)
	achievements, _ := PrepareAchievementsSection(resume.Achievements)
	pors, _ := PreparePORsSection(resume.PORs)
	closing, _ := utils.ReadFileToBuffer("section/tex/closing.text")

	preamble.WriteTo(&resumeBuffer)
	contact.WriteTo(&resumeBuffer)
	education.WriteTo(&resumeBuffer)
	workExperience.WriteTo(&resumeBuffer)
	researchExperience.WriteTo(&resumeBuffer)
	personalProjects.WriteTo(&resumeBuffer)
	achievements.WriteTo(&resumeBuffer)
	pors.WriteTo(&resumeBuffer)
	closing.WriteTo(&resumeBuffer)

	return &resumeBuffer, nil

}
