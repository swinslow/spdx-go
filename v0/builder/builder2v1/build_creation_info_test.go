// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package builder2v1

import (
	"fmt"
	"testing"
)

// ===== CreationInfo section builder tests =====
func TestBuilder2_1CanBuildCreationInfoSection(t *testing.T) {

	namespacePrefix := "https://github.com/swinslow/spdx-docs/spdx-go/testdata-whatever-"
	creatorType := "Organization"
	creator := "Jane Doe LLC"
	testValues := make(map[string]string)
	testValues["Created"] = "2018-10-20T16:48:00Z"
	packageName := "project1"
	verificationCode := "TESTCODE"

	ci, err := BuildCreationInfoSection2_1(packageName, verificationCode, namespacePrefix, creatorType, creator, testValues)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if ci == nil {
		t.Fatalf("expected non-nil CreationInfo, got nil")
	}
	if ci.SPDXVersion != "SPDX-2.1" {
		t.Errorf("expected %s, got %s", "SPDX-2.1", ci.SPDXVersion)
	}
	if ci.DataLicense != "CC0-1.0" {
		t.Errorf("expected %s, got %s", "CC0-1.0", ci.DataLicense)
	}
	if ci.SPDXIdentifier != "SPDXRef-DOCUMENT" {
		t.Errorf("expected %s, got %s", "SPDXRef-DOCUMENT", ci.SPDXIdentifier)
	}
	if ci.DocumentName != "project1" {
		t.Errorf("expected %s, got %s", "project1", ci.DocumentName)
	}
	wantNamespace := fmt.Sprintf("https://github.com/swinslow/spdx-docs/spdx-go/testdata-whatever-project1-%s", verificationCode)
	if ci.DocumentNamespace != wantNamespace {
		t.Errorf("expected %s, got %s", wantNamespace, ci.DocumentNamespace)
	}
	if len(ci.CreatorPersons) != 0 {
		t.Fatalf("expected %d, got %d", 0, len(ci.CreatorPersons))
	}
	if len(ci.CreatorOrganizations) != 1 {
		t.Fatalf("expected %d, got %d", 1, len(ci.CreatorOrganizations))
	}
	if ci.CreatorOrganizations[0] != "Jane Doe LLC" {
		t.Errorf("expected %s, got %s", "Jane Doe LLC", ci.CreatorOrganizations[0])
	}
	if len(ci.CreatorTools) != 1 {
		t.Fatalf("expected %d, got %d", 1, len(ci.CreatorTools))
	}
	if ci.CreatorTools[0] != "github.com/swinslow/spdx-go/v0/builder" {
		t.Errorf("expected %s, got %s", "github.com/swinslow/spdx-go/v0/builder", ci.CreatorTools[0])
	}
	if ci.Created != "2018-10-20T16:48:00Z" {
		t.Errorf("expected %s, got %s", "2018-10-20T16:48:00Z", ci.Created)
	}
}

func TestBuilder2_1CanBuildCreationInfoSectionWithCreatorPerson(t *testing.T) {
	namespacePrefix := "https://github.com/swinslow/spdx-docs/spdx-go/testdata-whatever-"
	creatorType := "Person"
	creator := "John Doe"
	testValues := make(map[string]string)
	testValues["Created"] = "2018-10-20T16:48:00Z"
	packageName := "project1"
	verificationCode := "TESTCODE"

	ci, err := BuildCreationInfoSection2_1(packageName, verificationCode, namespacePrefix, creatorType, creator, testValues)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if ci == nil {
		t.Fatalf("expected non-nil CreationInfo, got nil")
	}
	if len(ci.CreatorPersons) != 1 {
		t.Fatalf("expected %d, got %d", 1, len(ci.CreatorPersons))
	}
	if ci.CreatorPersons[0] != "John Doe" {
		t.Errorf("expected %s, got %s", "John Doe", ci.CreatorPersons[0])
	}
	if len(ci.CreatorOrganizations) != 0 {
		t.Fatalf("expected %d, got %d", 0, len(ci.CreatorOrganizations))
	}
	if len(ci.CreatorTools) != 1 {
		t.Fatalf("expected %d, got %d", 1, len(ci.CreatorTools))
	}
	if ci.CreatorTools[0] != "github.com/swinslow/spdx-go/v0/builder" {
		t.Errorf("expected %s, got %s", "github.com/swinslow/spdx-go/v0/builder", ci.CreatorTools[0])
	}
}

func TestBuilder2_1CanBuildCreationInfoSectionWithCreatorTool(t *testing.T) {
	namespacePrefix := "https://github.com/swinslow/spdx-docs/spdx-go/testdata-whatever-"
	creatorType := "Tool"
	creator := "some-other-tool-2.1"
	testValues := make(map[string]string)
	testValues["Created"] = "2018-10-20T16:48:00Z"
	packageName := "project1"
	verificationCode := "TESTCODE"

	ci, err := BuildCreationInfoSection2_1(packageName, verificationCode, namespacePrefix, creatorType, creator, testValues)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if ci == nil {
		t.Fatalf("expected non-nil CreationInfo, got nil")
	}
	if len(ci.CreatorPersons) != 0 {
		t.Fatalf("expected %d, got %d", 0, len(ci.CreatorPersons))
	}
	if len(ci.CreatorOrganizations) != 0 {
		t.Fatalf("expected %d, got %d", 0, len(ci.CreatorOrganizations))
	}
	if len(ci.CreatorTools) != 2 {
		t.Fatalf("expected %d, got %d", 2, len(ci.CreatorTools))
	}
	if ci.CreatorTools[0] != "github.com/swinslow/spdx-go/v0/builder" {
		t.Errorf("expected %s, got %s", "github.com/swinslow/spdx-go/v0/builder", ci.CreatorTools[0])
	}
	if ci.CreatorTools[1] != "some-other-tool-2.1" {
		t.Errorf("expected %s, got %s", "some-other-tool-2.1", ci.CreatorTools[1])
	}
}

func TestBuilder2_1CanBuildCreationInfoSectionWithInvalidPerson(t *testing.T) {
	namespacePrefix := "https://github.com/swinslow/spdx-docs/spdx-go/testdata-whatever-"
	creatorType := "Whatever"
	creator := "John Doe"
	testValues := make(map[string]string)
	testValues["Created"] = "2018-10-20T16:48:00Z"
	packageName := "project1"
	verificationCode := "TESTCODE"

	ci, err := BuildCreationInfoSection2_1(packageName, verificationCode, namespacePrefix, creatorType, creator, testValues)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if ci == nil {
		t.Fatalf("expected non-nil CreationInfo, got nil")
	}
	if len(ci.CreatorPersons) != 1 {
		t.Fatalf("expected %d, got %d", 1, len(ci.CreatorPersons))
	}
	if ci.CreatorPersons[0] != "John Doe" {
		t.Errorf("expected %s, got %s", "John Doe", ci.CreatorPersons[0])
	}
	if len(ci.CreatorOrganizations) != 0 {
		t.Fatalf("expected %d, got %d", 0, len(ci.CreatorOrganizations))
	}
	if len(ci.CreatorTools) != 1 {
		t.Fatalf("expected %d, got %d", 1, len(ci.CreatorTools))
	}
	if ci.CreatorTools[0] != "github.com/swinslow/spdx-go/v0/builder" {
		t.Errorf("expected %s, got %s", "github.com/swinslow/spdx-go/v0/builder", ci.CreatorTools[0])
	}
}
