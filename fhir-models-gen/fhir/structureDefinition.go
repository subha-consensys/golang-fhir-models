// Copyright 2019 The Samply Development Community
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fhir

import "encoding/json"

// THIS FILE IS GENERATED BY https://github.com/samply/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// StructureDefinition is documented here http://hl7.org/fhir/StructureDefinition/StructureDefinition
type StructureDefinition struct {
	Id                *string                          `json:"id,omitempty"`
	Meta              *Meta                            `json:"meta,omitempty"`
	ImplicitRules     *string                          `json:"implicitRules,omitempty"`
	Language          *string                          `json:"language,omitempty"`
	Text              *Narrative                       `json:"text,omitempty"`
	Extension         []Extension                      `json:"extension,omitempty"`
	ModifierExtension []Extension                      `json:"modifierExtension,omitempty"`
	Url               string                           `json:"url"`
	Identifier        []Identifier                     `json:"identifier,omitempty"`
	Version           *string                          `json:"version,omitempty"`
	Name              string                           `json:"name"`
	Title             *string                          `json:"title,omitempty"`
	Status            PublicationStatus                `json:"status"`
	Experimental      *bool                            `json:"experimental,omitempty"`
	Date              *string                          `json:"date,omitempty"`
	Publisher         *string                          `json:"publisher,omitempty"`
	Contact           []ContactDetail                  `json:"contact,omitempty"`
	Description       *string                          `json:"description,omitempty"`
	UseContext        []UsageContext                   `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept                `json:"jurisdiction,omitempty"`
	Purpose           *string                          `json:"purpose,omitempty"`
	Copyright         *string                          `json:"copyright,omitempty"`
	Keyword           []Coding                         `json:"keyword,omitempty"`
	FhirVersion       *FHIRVersion                     `json:"fhirVersion,omitempty"`
	Mapping           []StructureDefinitionMapping     `json:"mapping,omitempty"`
	Kind              StructureDefinitionKind          `json:"kind"`
	Abstract          bool                             `json:"abstract"`
	Context           []StructureDefinitionContext     `json:"context,omitempty"`
	ContextInvariant  []string                         `json:"contextInvariant,omitempty"`
	Type              string                           `json:"type"`
	BaseDefinition    *string                          `json:"baseDefinition,omitempty"`
	Derivation        *TypeDerivationRule              `json:"derivation,omitempty"`
	Snapshot          *StructureDefinitionSnapshot     `json:"snapshot,omitempty"`
	Differential      *StructureDefinitionDifferential `json:"differential,omitempty"`
}
type StructureDefinitionMapping struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Identity          string      `json:"identity"`
	Uri               *string     `json:"uri,omitempty"`
	Name              *string     `json:"name,omitempty"`
	Comment           *string     `json:"comment,omitempty"`
}
type StructureDefinitionContext struct {
	Id                *string              `json:"id,omitempty"`
	Extension         []Extension          `json:"extension,omitempty"`
	ModifierExtension []Extension          `json:"modifierExtension,omitempty"`
	Type              ExtensionContextType `json:"type"`
	Expression        string               `json:"expression"`
}
type StructureDefinitionSnapshot struct {
	Id                *string             `json:"id,omitempty"`
	Extension         []Extension         `json:"extension,omitempty"`
	ModifierExtension []Extension         `json:"modifierExtension,omitempty"`
	Element           []ElementDefinition `json:"element"`
}
type StructureDefinitionDifferential struct {
	Id                *string             `json:"id,omitempty"`
	Extension         []Extension         `json:"extension,omitempty"`
	ModifierExtension []Extension         `json:"modifierExtension,omitempty"`
	Element           []ElementDefinition `json:"element"`
}
type OtherStructureDefinition StructureDefinition

// MarshalJSON marshals the given StructureDefinition as JSON into a byte slice
func (r StructureDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherStructureDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherStructureDefinition: OtherStructureDefinition(r),
		ResourceType:             "StructureDefinition",
	})
}

// UnmarshalStructureDefinition unmarshals a StructureDefinition.
func UnmarshalStructureDefinition(b []byte) (StructureDefinition, error) {
	var structureDefinition StructureDefinition
	if err := json.Unmarshal(b, &structureDefinition); err != nil {
		return structureDefinition, err
	}
	return structureDefinition, nil
}