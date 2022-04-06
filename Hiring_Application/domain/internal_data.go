package domain

//internal hiring requirement struct

type RequirementInt struct {
	CandidateId       int      `json:"candidateid"`
	Fullname          string   `json:"fullname"`
	RequiredSkillsets []string `json:"requiredSkillsets"`
	OptionalSkillsets []string `json:"optionalSkillsets"`
}
