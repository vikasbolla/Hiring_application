package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vikasbolla/Hiring_application/Hiring_Application/domain"
)

//external candidate details struct

type ExternalCandidate struct {
	CandidateId int      `json:"candidateid"`
	Fullname    string   `json:"fullname"`
	Skillsets   []string `json:"skillsets"`
}

var externalcandidates []ExternalCandidate

// post candidates for external hiring

func ExternalPostRequirement(w http.ResponseWriter, r *http.Request) {
	var candidate ExternalCandidate
	w.Header().Add("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&candidate)
	externalcandidates = append(externalcandidates, candidate)

	fmt.Fprintf(w, "Success")
}

// function to get details of candidates

func ExternalAllCandidates(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(externalcandidates)
}

var requirementext domain.RequirementExt

// matching requirement skills with candidate skills

func ExternalMatchRequirement(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&requirementext)
	var matchedskill []ExternalCandidate
	for i := 0; i < len(externalcandidates); i++ {
		c := 0
		for j := 0; j < len(externalcandidates[i].Skillsets); j++ {
			for k := 0; k < len(requirementext.RequiredSkillsets); k++ {
				if externalcandidates[i].Skillsets[j] == requirementext.RequiredSkillsets[k] {
					c++
				}
			}
		}
		if c == len(requirementext.RequiredSkillsets) {
			matchedskill = append(matchedskill, externalcandidates[i])
		}
	}
	json.NewEncoder(w).Encode(matchedskill)
}
