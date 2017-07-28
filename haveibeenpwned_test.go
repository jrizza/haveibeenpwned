package haveibeenpwned

import "testing"
import "time"

func TestBreachedAccountWithNoFilters(t *testing.T) {
	time.Sleep(1500 * time.Millisecond)
	breaches, err := BreachedAccount("example@gmail.com", "", false, false)
	if err != nil {
		t.Fatalf("response error: %v", err)
	}
	if len(breaches) != 39 {
		t.Errorf("expected 39 results, got %d", len(breaches))
	}
	if breaches[0].Description != "" {
		t.Errorf("expected a description, got empty")
	}
}

func TestBreachedAccountWithDomainFilter(t *testing.T) {
	time.Sleep(1500 * time.Millisecond)
	breaches, err := BreachedAccount("example@gmail.com", "000webhost.com", false, false)
	if err != nil {
		t.Fatalf("response error: %v", err)
	}
	if len(breaches) != 1 {
		t.Errorf("expected 1 results, got %d", len(breaches))
	}
}

func TestBreachedAccountWithDomainFilterAndTruncated(t *testing.T) {
	time.Sleep(1500 * time.Millisecond)
	breaches, err := BreachedAccount("example@gmail.com", "000webhost.com", true, false)
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

func TestBreachedAccountWithDomainFilterTruncatedAndUnverified(t *testing.T) {
	time.Sleep(1500 * time.Millisecond)
	breaches, err := BreachedAccount("example@gmail.com", "000webhost.com", true, true)
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
