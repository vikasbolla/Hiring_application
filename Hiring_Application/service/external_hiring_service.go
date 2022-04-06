package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var externalcandidates []Candidate

// post candidates for external hiring

func ExternalPostRequirement(w http.ResponseWriter, r *http.Request) {
	var candidate Candidate
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

var requirementext RequirementExt

// matching requirement skills with candidate skills

func ExternalMatchRequirement(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&requirementext)
	var matchedskill []Candidate
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
