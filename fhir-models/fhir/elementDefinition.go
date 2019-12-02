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

// THIS FILE IS GENERATED BY https://github.com/samply/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// ElementDefinition is documented here http://hl7.org/fhir/StructureDefinition/ElementDefinition
type ElementDefinition struct {
	Id                  *string                       `json:"id,omitempty"`
	Extension           []Extension                   `json:"extension,omitempty"`
	ModifierExtension   []Extension                   `json:"modifierExtension,omitempty"`
	Path                string                        `json:"path"`
	Representation      []PropertyRepresentation      `json:"representation,omitempty"`
	SliceName           *string                       `json:"sliceName,omitempty"`
	SliceIsConstraining *bool                         `json:"sliceIsConstraining,omitempty"`
	Label               *string                       `json:"label,omitempty"`
	Code                []Coding                      `json:"code,omitempty"`
	Slicing             *ElementDefinitionSlicing     `json:"slicing,omitempty"`
	Short               *string                       `json:"short,omitempty"`
	Definition          *string                       `json:"definition,omitempty"`
	Comment             *string                       `json:"comment,omitempty"`
	Requirements        *string                       `json:"requirements,omitempty"`
	Alias               []string                      `json:"alias,omitempty"`
	Min                 *int                          `json:"min,omitempty"`
	Max                 *string                       `json:"max,omitempty"`
	Base                *ElementDefinitionBase        `json:"base,omitempty"`
	ContentReference    *string                       `json:"contentReference,omitempty"`
	Type                []ElementDefinitionType       `json:"type,omitempty"`
	MeaningWhenMissing  *string                       `json:"meaningWhenMissing,omitempty"`
	OrderMeaning        *string                       `json:"orderMeaning,omitempty"`
	Example             []ElementDefinitionExample    `json:"example,omitempty"`
	MaxLength           *int                          `json:"maxLength,omitempty"`
	Condition           []string                      `json:"condition,omitempty"`
	Constraint          []ElementDefinitionConstraint `json:"constraint,omitempty"`
	MustSupport         *bool                         `json:"mustSupport,omitempty"`
	IsModifier          *bool                         `json:"isModifier,omitempty"`
	IsModifierReason    *string                       `json:"isModifierReason,omitempty"`
	IsSummary           *bool                         `json:"isSummary,omitempty"`
	Binding             *ElementDefinitionBinding     `json:"binding,omitempty"`
	Mapping             []ElementDefinitionMapping    `json:"mapping,omitempty"`
}
type ElementDefinitionSlicing struct {
	Id            *string                                 `json:"id,omitempty"`
	Extension     []Extension                             `json:"extension,omitempty"`
	Discriminator []ElementDefinitionSlicingDiscriminator `json:"discriminator,omitempty"`
	Description   *string                                 `json:"description,omitempty"`
	Ordered       *bool                                   `json:"ordered,omitempty"`
	Rules         SlicingRules                            `json:"rules"`
}
type ElementDefinitionSlicingDiscriminator struct {
	Id        *string           `json:"id,omitempty"`
	Extension []Extension       `json:"extension,omitempty"`
	Type      DiscriminatorType `json:"type"`
	Path      string            `json:"path"`
}
type ElementDefinitionBase struct {
	Id        *string     `json:"id,omitempty"`
	Extension []Extension `json:"extension,omitempty"`
	Path      string      `json:"path"`
	Min       int         `json:"min"`
	Max       string      `json:"max"`
}
type ElementDefinitionType struct {
	Id            *string                `json:"id,omitempty"`
	Extension     []Extension            `json:"extension,omitempty"`
	Code          string                 `json:"code"`
	Profile       []string               `json:"profile,omitempty"`
	TargetProfile []string               `json:"targetProfile,omitempty"`
	Aggregation   []AggregationMode      `json:"aggregation,omitempty"`
	Versioning    *ReferenceVersionRules `json:"versioning,omitempty"`
}
type ElementDefinitionExample struct {
	Id        *string     `json:"id,omitempty"`
	Extension []Extension `json:"extension,omitempty"`
	Label     string      `json:"label"`
}
type ElementDefinitionConstraint struct {
	Id           *string            `json:"id,omitempty"`
	Extension    []Extension        `json:"extension,omitempty"`
	Key          string             `json:"key"`
	Requirements *string            `json:"requirements,omitempty"`
	Severity     ConstraintSeverity `json:"severity"`
	Human        string             `json:"human"`
	Expression   *string            `json:"expression,omitempty"`
	Xpath        *string            `json:"xpath,omitempty"`
	Source       *string            `json:"source,omitempty"`
}
type ElementDefinitionBinding struct {
	Id          *string         `json:"id,omitempty"`
	Extension   []Extension     `json:"extension,omitempty"`
	Strength    BindingStrength `json:"strength"`
	Description *string         `json:"description,omitempty"`
	ValueSet    *string         `json:"valueSet,omitempty"`
}
type ElementDefinitionMapping struct {
	Id        *string     `json:"id,omitempty"`
	Extension []Extension `json:"extension,omitempty"`
	Identity  string      `json:"identity"`
	Language  *string     `json:"language,omitempty"`
	Map       string      `json:"map"`
	Comment   *string     `json:"comment,omitempty"`
}