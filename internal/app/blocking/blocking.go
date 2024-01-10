// internal/app/blocking/blocking.go
package blocking

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/feruzoripov/block-unwanted-hosts/pkg/dns"
)

const hostsFilePath = "etc/hosts"

// FetchAndBlockWebsites fetches and blocks websites.
func FetchAndBlockWebsites() error {
	urls := []string{
		"https://someonewhocares.org/hosts/hosts",
		"https://raw.githubusercontent.com/StevenBlack/hosts/master/hosts",
		"https://adaway.org/hosts.txt",
		"https://pgl.yoyo.org/adservers/serverlist.php?hostformat=hosts;showintro=0&mimetype=plaintext",
	}

	websites, err := fetchWebsitesFromURLs(urls)
	if err != nil {
		return fmt.Errorf("error fetching websites: %s", err)
	}

	log.Println("Blocking websites in the hosts file...")

	if err := blockWebsites(websites); err != nil {
		return fmt.Errorf("error blocking websites: %s", err)
	}

	log.Println("Websites blocked successfully!")

	return nil
}

func fetchWebsitesFromURLs(urls []string) ([]string, error) {
	var websites []string
	visited := make(map[string]struct{})

	for _, url := range urls {
		log.Printf("Fetching websites from %s...", url)

		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		websitesFromResponse := extractWebsitesFromResponse(resp.Body)
		for _, website := range websitesFromResponse {
			// Check for duplicates
			if _, exists := visited[website]; !exists {
				websites = append(websites, website)
				visited[website] = struct{}{}
			}
		}

		log.Printf("Fetched %d unique websites from %s", len(visited), url)
	}

	return websites, nil
}

func extractWebsitesFromResponse(body io.Reader) []string {
	var websites []string
	scanner := bufio.NewScanner(body)

	for scanner.Scan() {
		line := scanner.Text()
		if isHostsEntry(line) {
			websites = append(websites, extractDomainFromHostsEntry(line))
		}
	}

	return websites
}

func isHostsEntry(line string) bool {
	return strings.HasPrefix(line, "127.0.0.1") || strings.HasPrefix(line, "0.0.0.0")
}

func extractDomainFromHostsEntry(line string) string {
	fields := strings.Fields(line)
	if len(fields) > 1 {
		return fields[1]
	}
	return ""
}

func blockWebsites(websites []string) error {
	log.Println("Opening hosts file for appending...")

	file, err := os.OpenFile(hostsFilePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := addCommentLine(file); err != nil {
		return err
	}

	if err := addWebsitesToHostsFile(file, websites); err != nil {
		return err
	}

	if err := flushChangesToFile(file); err != nil {
		return err
	}

	log.Println("DNS cache flushing...")

	if err := dns.ClearCache(); err != nil {
		return err
	}

	log.Println("DNS cache flushed successfully!")

	return nil
}

func addCommentLine(file *os.File) error {
	comment := "# Blocked websites from the specified URLs\n"
	_, err := file.WriteString(comment)
	return err
}

func addWebsitesToHostsFile(file *os.File, websites []string) error {
	for _, website := range websites {
		entry := fmt.Sprintf("127.0.0.1\t%s\n", website)
		if _, err := file.WriteString(entry); err != nil {
			return err
		}
	}
	return nil
}

func flushChangesToFile(file *os.File) error {
	return file.Sync()
}
