package domain

//Requirement for external hiring struct

type RequirementExt struct {
	RequirementId     int      `json:"requirementid"`
	Position          string   `json:"position"`
	RequiredSkillsets []string `json:"requiredSkillsets"`
}
