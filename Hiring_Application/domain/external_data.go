package domain

type Candidate struct {
	CandidateId int      `json:"candidateid"`
	Fullname    string   `json:"fullname"`
	Skillsets   []string `json:"skillsets"`
}

type RequirementExt struct {
	RequirementId     int      `json:"requirementid"`
	Position          string   `json:"position"`
	RequiredSkillsets []string `json:"requiredSkillsets"`
}
