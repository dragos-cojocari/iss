package ui

import (
	"bufio"
	"math/rand/v2"
	"os"
	"strings"
	"sync"
)

// QuoteManager handles loading and cycling through Matrix quotes
type QuoteManager struct {
	quotes       []string
	currentIndex int
	mu           sync.RWMutex
}

var (
	quoteManager     *QuoteManager
	quoteManagerOnce sync.Once
)

// GetQuoteManager returns the singleton quote manager instance
func GetQuoteManager() *QuoteManager {
	quoteManagerOnce.Do(func() {
		quoteManager = &QuoteManager{
			quotes:       []string{},
			currentIndex: 0,
		}
		quoteManager.loadQuotes()
	})
	return quoteManager
}

// loadQuotes loads quotes from the quotes.txt file
func (qm *QuoteManager) loadQuotes() {
	file, err := os.Open("quotes.txt")
	if err != nil {
		// Fallback quotes if file not found
		qm.quotes = []string{
			"There is no spoon.",
			"I know kung fu.",
			"Welcome to the real world.",
			"Free your mind.",
		}
		// Randomize starting position
		qm.currentIndex = rand.IntN(len(qm.quotes))
		return
	}
	defer func() {
		_ = file.Close()
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			qm.quotes = append(qm.quotes, line)
		}
	}

	// Randomize starting position
	if len(qm.quotes) > 0 {
		qm.currentIndex = rand.IntN(len(qm.quotes))
	}
}

// GetQuote returns the current quote
func (qm *QuoteManager) GetQuote() string {
	qm.mu.RLock()
	defer qm.mu.RUnlock()

	if len(qm.quotes) == 0 {
		return "There is no spoon."
	}

	return qm.quotes[qm.currentIndex]
}

// NextQuote cycles to the next quote
func (qm *QuoteManager) NextQuote() string {
	qm.mu.Lock()
	defer qm.mu.Unlock()

	if len(qm.quotes) == 0 {
		return "There is no spoon."
	}

	qm.currentIndex = (qm.currentIndex + 1) % len(qm.quotes)
	return qm.quotes[qm.currentIndex]
}

// RandomQuote returns a random quote
func (qm *QuoteManager) RandomQuote() string {
	qm.mu.Lock()
	defer qm.mu.Unlock()

	if len(qm.quotes) == 0 {
		return "There is no spoon."
	}

	qm.currentIndex = rand.IntN(len(qm.quotes))
	return qm.quotes[qm.currentIndex]
}
