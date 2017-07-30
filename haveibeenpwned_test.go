package haveibeenpwned

import "testing"
import "time"

const APIrateLimit = 1500

func TestBreachedAccountWithNoFilters(t *testing.T) {
	time.Sleep(APIrateLimit * time.Millisecond)
	breaches, err := BreachedAccount("test@example.com", "", false, false)
	if err != nil {
		t.Fatalf("response error: %v", err)
	}
	if len(breaches) != 37 {
		t.Errorf("expected 37 results, got %d", len(breaches))
	}
	if breaches[1].Description == "" {
		t.Error("expected a description, got empty")
	}
}

func TestBreachedAccountWithDomainFilter(t *testing.T) {
	time.Sleep(APIrateLimit * time.Millisecond)
	breaches, err := BreachedAccount("test@example.com", "000webhost.com", false, false)
	if err != nil {
		t.Fatalf("response error: %v", err)
	}
	if len(breaches) != 1 {
		t.Errorf("expected 1 results, got %d", len(breaches))
	}
}

func TestBreachedAccountWithDomainFilterAndTruncated(t *testing.T) {
	time.Sleep(APIrateLimit * time.Millisecond)
	breaches, err := BreachedAccount("test@example.com", "000webhost.com", true, false)
	if err != nil {
		t.Fatalf("response error: %v", err)
	}
	if len(breaches) != 1 {
		t.Errorf("expected 1 results, got %d", len(breaches))
	}
	if breaches[0].Description != "" {
		t.Errorf("expected empty description, got %s", breaches[0].Description)
	}
}

func TestUnBreachedAccount(t *testing.T) {
	time.Sleep(APIrateLimit * time.Millisecond)
	breaches, err := BreachedAccount("notanexistingemail@address.com", "", false, false)
	if err != nil {
		t.Fatalf("response error: %v", err)
	}
	if breaches != nil {
		t.Error("expected nil, got breaches")
	}
}

func TestExistingBreach(t *testing.T) {
	time.Sleep(APIrateLimit * time.Millisecond)
	breach, err := Breach("adobe")
	if err != nil {
		t.Fatalf("response error: %v", err)
	}
	if breach.Name == "" {
		t.Error("expected a name, got empty")
	}
}

func TestNonExistingBreach(t *testing.T) {
	time.Sleep(APIrateLimit * time.Millisecond)
	breach, err := Breach("aasdasdasd")
	if err != nil {
		t.Fatalf("response error: %v", err)
	}
	if breach.Name != "" {
		t.Errorf("expected empty, got a name: %s", breach.Name)
	}
}

func TestExistingPasteAccount(t *testing.T) {
	time.Sleep(APIrateLimit * time.Millisecond)
	pastes, err := PasteAccount("test@example.com")
	if err != nil {
		t.Fatalf("response error: %v", err)
	}
	if len(pastes) != 34 {
		t.Errorf("expected 34 results, got %d", len(pastes))
	}
	if pastes[1].ID == "" {
		t.Error("expected a description, got empty")
	}
}

func TestNonExistingPasteAccount(t *testing.T) {
	time.Sleep(APIrateLimit * time.Millisecond)
	pastes, err := PasteAccount("notanexistingemail@address.com")
	if err != nil {
		t.Fatalf("response error: %v", err)
	}
	if pastes != nil {
		t.Errorf("expected no results, got %d", len(pastes))
	}
}

func TestInvalidPasteAccount(t *testing.T) {
	time.Sleep(APIrateLimit * time.Millisecond)
	pastes, err := PasteAccount("test")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if pastes != nil {
		t.Errorf("expected no results, got %d", len(pastes))
	}
}
