package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/QRY91/wherewasi/internal/database"
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wherewasi",
	Short: "AI context generation CLI - Less explaining. More building.",
	Long: `wherewasi generates dense AI context summaries to eliminate 
the 'explain my project again' overhead. Like a ripcord for your 
development workflow - invisible until you need it.`,
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start background tracking",
	Long:  "Begin monitoring your project for context generation",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🚀 wherewasi shadow mode starting...")
		fmt.Println("🥷 Passive tracking enabled")
		fmt.Println("🪂 Ready for ripcord deployment: wherewasi pull")
		// TODO: Implement background daemon
	},
}

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull the ripcord - generate AI context",
	Long:  "Generate dense, AI-ready context for immediate use",
	Run: func(cmd *cobra.Command, args []string) {
		project, _ := cmd.Flags().GetString("project")
		days, _ := cmd.Flags().GetInt("days")
		keyword, _ := cmd.Flags().GetString("keyword")
		clipboard_flag, _ := cmd.Flags().GetBool("clipboard")
		history_flag, _ := cmd.Flags().GetBool("history")
		save_flag, _ := cmd.Flags().GetBool("save")

		fmt.Println("🪂 Pulling ripcord...")

		// Handle history search
		if history_flag {
			if keyword != "" {
				results, err := db.SearchStoredContexts(keyword)
				if err != nil {
					fmt.Printf("⚠️  Could not search history: %v\n", err)
				} else if len(results) > 0 {
					fmt.Println("📚 CONTEXT HISTORY SEARCH:")
					for _, result := range results {
						fmt.Printf("  • [%s] %s | %s\n", result.Project, result.Timestamp.Format("2006-01-02T15:04"), result.SessionInfo)
					}
				} else {
					fmt.Println("📚 No matching contexts found in history")
				}
				return
			} else {
				// Show recent contexts for current project
				currentProject := getProjectName()
				results, err := db.GetRecentContexts(currentProject, 5)
				if err != nil {
					fmt.Printf("⚠️  Could not get history: %v\n", err)
				} else if len(results) > 0 {
					fmt.Printf("📚 RECENT CONTEXTS (%s):\n", currentProject)
					for _, result := range results {
						fmt.Printf("  • 📅 %s | %s\n", result.Timestamp.Format("2006-01-02T15:04"), result.SessionInfo)
					}
				} else {
					fmt.Println("📚 No context history found for this project")
				}
				return
			}
		}

		context := generateEnhancedContext(project, days, keyword)

		// Save context if requested (default true for production usage)
		if save_flag {
			sessionInfo := detectActiveSession()
			if sessionInfo == "" {
				sessionInfo = "Context pull"
			}
			currentProject := getProjectName()
			_, err := db.SaveContext(currentProject, context, sessionInfo, keyword)
			if err != nil {
				fmt.Printf("⚠️  Could not save context: %v\n", err)
			}
		}

		if clipboard_flag {
			err := clipboard.WriteAll(context)
			if err != nil {
				fmt.Printf("⚠️  Could not copy to clipboard: %v\n", err)
				fmt.Println("📋 Context output (copy manually):")
				fmt.Println("\n" + context)
			} else {
				fmt.Println("📋 Context copied to clipboard! Paste and build.")
			}
		} else {
			fmt.Println("\n" + context)
		}
	},
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show tracking status",
	Long:  "Display what's currently being tracked",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🪂 wherewasi ripcord status:")
		fmt.Println("  🥷 Shadow mode: active")
		fmt.Println("  📊 Projects tracked: detecting...")
		fmt.Println("  🧠 Context ready: pull to deploy")
		showTrackedProjects()
	},
}

func generateEnhancedContext(project string, days int, keyword string) string {
	var context strings.Builder

	// Header with ecosystem context
	context.WriteString("--- AI CONTEXT DEPLOYMENT ---\n")

	if project != "" {
		context.WriteString(fmt.Sprintf("🎯 FOCUSED ON: %s\n", project))
		if !isValidProject(project) {
			context.WriteString("⚠️  Project not found in current ecosystem\n")
		}
	} else {
		context.WriteString(fmt.Sprintf("🏠 CURRENT PROJECT: %s\n", getProjectName()))
		context.WriteString(fmt.Sprintf("📍 LOCATION: %s\n", getCurrentDir()))
	}

	// Active session detection
	sessionContext := detectActiveSession()
	if sessionContext != "" {
		context.WriteString(fmt.Sprintf("⚡ ACTIVE SESSION: %s\n", sessionContext))
	}

	// Ecosystem awareness
	context.WriteString("\n🧠 QRY ECOSYSTEM CONTEXT:\n")
	context.WriteString("- Building unified local developer AI system\n")
	context.WriteString("- Tools: wherewasi(context), uroboro(content), doggowoof(alerts), qomoboro(time)\n")
	context.WriteString("- Recent breakthrough: Ecosystem intelligence discovery\n")
	context.WriteString("- Current focus: Ripcord implementation for instant AI context\n")

	// Time-filtered context
	if days > 0 {
		context.WriteString(fmt.Sprintf("\n⏰ LAST %d DAYS:\n", days))
		commits := getCommitsSince(days, project)
		if len(commits) > 0 {
			for _, commit := range commits {
				if keyword == "" || strings.Contains(strings.ToLower(commit), strings.ToLower(keyword)) {
					context.WriteString(fmt.Sprintf("  • %s\n", commit))
				}
			}
		} else {
			context.WriteString("  • No commits found in timeframe\n")
		}
	} else {
		// Default recent commits
		commits := getRecentCommits(5)
		if len(commits) > 0 {
			context.WriteString("\n📝 RECENT COMMITS:\n")
			for _, commit := range commits {
				if keyword == "" || strings.Contains(strings.ToLower(commit), strings.ToLower(keyword)) {
					context.WriteString(fmt.Sprintf("  • %s\n", commit))
				}
			}
		}
	}

	// Current state
	if changes := getUncommittedChanges(); len(changes) > 0 {
		context.WriteString("\n🔄 UNCOMMITTED CHANGES:\n")
		for _, change := range changes {
			if keyword == "" || strings.Contains(strings.ToLower(change), strings.ToLower(keyword)) {
				context.WriteString(fmt.Sprintf("  • %s\n", change))
			}
		}
	}

	// Project structure
	context.WriteString("\n📁 KEY FILES:\n")
	keyFiles := getKeyFiles()
	for _, file := range keyFiles {
		context.WriteString(fmt.Sprintf("  • %s\n", file))
	}

	// Recent development insights from chat history
	chatInsights := getRecentChatInsights()
	if len(chatInsights) > 0 {
		context.WriteString("\n💭 RECENT DEVELOPMENT INSIGHTS:\n")
		for _, insight := range chatInsights {
			context.WriteString(fmt.Sprintf("  • %s\n", insight))
		}
	}

	// Enhanced search context if keyword provided
	if keyword != "" {
		context.WriteString(fmt.Sprintf("\n🔍 CROSS-PROJECT SEARCH: '%s'\n", keyword))
		searchResults := searchCrossProject(keyword, project)
		if len(searchResults) > 0 {
			for _, result := range searchResults {
				context.WriteString(fmt.Sprintf("  • %s\n", result))
			}
		} else {
			context.WriteString("  • No matches found across ecosystem\n")
		}
	}

	context.WriteString("\n🎯 READY FOR AI COLLABORATION\n")
	context.WriteString("--- END CONTEXT ---")

	return context.String()
}

func isValidProject(project string) bool {
	parentDir := filepath.Dir(getCurrentDir())
	projectPath := filepath.Join(parentDir, project)
	_, err := os.Stat(projectPath)
	return err == nil
}

func getCommitsSince(days int, project string) []string {
	since := time.Now().AddDate(0, 0, -days).Format("2006-01-02")

	var cmd *exec.Cmd
	if project != "" && isValidProject(project) {
		parentDir := filepath.Dir(getCurrentDir())
		projectPath := filepath.Join(parentDir, project)
		cmd = exec.Command("git", "-C", projectPath, "log", "--oneline", "--since", since)
	} else {
		cmd = exec.Command("git", "log", "--oneline", "--since", since)
	}

	output, err := cmd.Output()
	if err != nil {
		return []string{"No git history found for timeframe"}
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	var commits []string
	for _, line := range lines {
		if line != "" {
			commits = append(commits, line)
		}
	}
	return commits
}

func searchInProject(keyword string) []string {
	return searchInProjectWithPath(".", keyword)
}

func searchInProjectWithPath(projectPath, keyword string) []string {
	// Enhanced grep search with line numbers and multiple file types including chat histories
	cmd := exec.Command("grep", "-r", "-n", "-i",
		"--include=*.go", "--include=*.md", "--include=*.txt",
		"--include=*.json", "--include=*.yaml", "--include=*.yml",
		keyword, projectPath)
	output, err := cmd.Output()
	if err != nil {
		return []string{}
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	var results []string
	chatHistoryCount := 0

	for i, line := range lines {
		if line != "" && i < 20 { // Increased to 20 results
			// Parse file:line:content format
			parts := strings.SplitN(line, ":", 3)
			if len(parts) >= 3 {
				file := strings.TrimPrefix(parts[0], projectPath+"/")
				lineNum := parts[1]
				content := strings.TrimSpace(parts[2])

				// Special handling for chat history files (cursor_*.md)
				isChatHistory := strings.HasPrefix(file, "cursor_") && strings.HasSuffix(file, ".md")
				if isChatHistory {
					chatHistoryCount++
					if chatHistoryCount <= 3 { // Limit chat history results
						// Enhanced context for chat histories
						lineRange := getChatHistoryContext(filepath.Join(projectPath, file), lineNum, keyword)
						results = append(results, fmt.Sprintf("💬 %s:%s → %s", file, lineRange, content[:60]+"..."))
					}
				} else {
					// Regular file handling
					if len(content) > 80 {
						content = content[:77] + "..."
					}
					results = append(results, fmt.Sprintf("%s:%s → %s", file, lineNum, content))
				}
			}
		}
	}
	return results
}

func getChatHistoryContext(filePath, lineNum, keyword string) string {
	// Try to determine conversation context around the line
	line, err := strconv.Atoi(lineNum)
	if err != nil {
		return lineNum
	}

	// Look for section headers or conversation boundaries
	contextStart := max(1, line-50)
	contextEnd := line + 50

	// Read a chunk around the match to find conversation boundaries
	cmd := exec.Command("sed", "-n", fmt.Sprintf("%d,%dp", contextStart, contextEnd), filePath)
	output, err := cmd.Output()
	if err != nil {
		return lineNum
	}

	lines := strings.Split(string(output), "\n")
	for i, contextLine := range lines {
		// Look for conversation markers or section headers
		if strings.Contains(contextLine, "# ") || strings.Contains(contextLine, "## ") {
			actualLine := contextStart + i
			if actualLine <= line {
				return fmt.Sprintf("%d", actualLine) + "-" + lineNum
			}
		}
	}

	return fmt.Sprintf("%d-%d", contextStart, contextEnd)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func detectActiveSession() string {
	var sessionInfo []string

	// Check for recent file modifications (last 30 minutes)
	recentFiles := getRecentlyModifiedFiles(30)
	if len(recentFiles) > 0 {
		sessionInfo = append(sessionInfo, fmt.Sprintf("Recent edits: %s", strings.Join(recentFiles, ", ")))
	}

	// Check for chat history files indicating active development
	chatFiles := getChatHistoryFiles()
	if len(chatFiles) > 0 {
		// Get the most recent chat file
		mostRecent := getMostRecentFile(chatFiles)
		if mostRecent != "" {
			sessionInfo = append(sessionInfo, fmt.Sprintf("Latest discussion: %s", mostRecent))
		}
	}

	// Check git status for active work
	status := getGitWorkingStatus()
	if status != "" {
		sessionInfo = append(sessionInfo, status)
	}

	if len(sessionInfo) > 0 {
		return strings.Join(sessionInfo, " | ")
	}
	return ""
}

func getRecentlyModifiedFiles(minutes int) []string {
	// Find files modified in last N minutes
	cmd := exec.Command("find", ".", "-name", "*.go", "-o", "-name", "*.md", "-mmin", fmt.Sprintf("-%d", minutes))
	output, err := cmd.Output()
	if err != nil {
		return []string{}
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	var files []string
	for _, line := range lines {
		if line != "" && line != "." {
			file := strings.TrimPrefix(line, "./")
			files = append(files, file)
		}
	}
	return files
}

func getChatHistoryFiles() []string {
	files, err := filepath.Glob("cursor_*.md")
	if err != nil {
		return []string{}
	}
	return files
}

func getMostRecentFile(files []string) string {
	if len(files) == 0 {
		return ""
	}

	var mostRecent string
	var mostRecentTime time.Time

	for _, file := range files {
		info, err := os.Stat(file)
		if err != nil {
			continue
		}
		if info.ModTime().After(mostRecentTime) {
			mostRecentTime = info.ModTime()
			mostRecent = file
		}
	}

	return mostRecent
}

func getGitWorkingStatus() string {
	// Check if there are staged changes (ready to commit)
	cmd := exec.Command("git", "diff", "--cached", "--name-only")
	output, err := cmd.Output()
	if err == nil && strings.TrimSpace(string(output)) != "" {
		return "Staged changes ready to commit"
	}

	// Check if there's active work (modified files)
	cmd = exec.Command("git", "status", "--porcelain")
	output, err = cmd.Output()
	if err == nil && strings.TrimSpace(string(output)) != "" {
		return "Active development in progress"
	}

	return ""
}

func getRecentChatInsights() []string {
	var insights []string

	// Get most recent chat file
	chatFiles := getChatHistoryFiles()
	if len(chatFiles) == 0 {
		return insights
	}

	mostRecent := getMostRecentFile(chatFiles)
	if mostRecent == "" {
		return insights
	}

	// Check if it was modified recently (last hour)
	info, err := os.Stat(mostRecent)
	if err != nil {
		return insights
	}

	if time.Since(info.ModTime()) > time.Hour {
		return insights
	}

	// Extract key insights from recent chat
	insights = extractChatInsights(mostRecent)
	return insights
}

func extractChatInsights(filename string) []string {
	var insights []string

	// Look for key patterns that indicate current work
	patterns := []string{
		"implementing", "enhanced", "added", "fixed", "testing",
		"next steps", "working on", "building", "creating",
	}

	for _, pattern := range patterns {
		cmd := exec.Command("grep", "-i", "-n", pattern, filename)
		output, err := cmd.Output()
		if err != nil {
			continue
		}

		lines := strings.Split(strings.TrimSpace(string(output)), "\n")
		for i, line := range lines {
			if line != "" && i < 3 { // Limit to 3 insights per pattern
				// Extract meaningful context
				parts := strings.SplitN(line, ":", 3)
				if len(parts) >= 3 {
					content := strings.TrimSpace(parts[2])
					if len(content) > 20 && len(content) < 150 {
						// Clean up markdown and formatting
						content = strings.ReplaceAll(content, "**", "")
						content = strings.ReplaceAll(content, "*", "")
						insights = append(insights, content)
					}
				}
			}
		}

		if len(insights) >= 5 { // Limit total insights
			break
		}
	}

	return insights
}

// Database instance (will be initialized in main)
var db *database.DB

func searchCrossProject(keyword string, project string) []string {
	var allResults []string

	// Always search current project first
	currentProject := getProjectName()
	currentResults := searchInProjectWithPath(".", keyword)
	for _, result := range currentResults {
		allResults = append(allResults, fmt.Sprintf("[%s] %s", currentProject, result))
	}

	if project != "" && isValidProject(project) {
		// Search specific project
		parentDir := filepath.Dir(getCurrentDir())
		projectPath := filepath.Join(parentDir, project)
		results := searchInProjectWithPath(projectPath, keyword)
		for _, result := range results {
			allResults = append(allResults, fmt.Sprintf("[%s] %s", project, result))
		}
	} else {
		// Search all ecosystem projects (excluding current)
		parentDir := filepath.Dir(getCurrentDir())
		entries, err := os.ReadDir(parentDir)
		if err != nil {
			return allResults // Return current project results at least
		}

		for _, entry := range entries {
			if entry.IsDir() && !strings.HasPrefix(entry.Name(), ".") &&
				entry.Name() != currentProject && hasGitRepo(filepath.Join(parentDir, entry.Name())) {
				projectPath := filepath.Join(parentDir, entry.Name())
				results := searchInProjectWithPath(projectPath, keyword)
				for _, result := range results {
					allResults = append(allResults, fmt.Sprintf("[%s] %s", entry.Name(), result))
					if len(allResults) >= 20 { // Increased limit for comprehensive search
						return allResults
					}
				}
			}
		}
	}

	return allResults
}

func showTrackedProjects() {
	parentDir := filepath.Dir(getCurrentDir())
	entries, err := os.ReadDir(parentDir)
	if err != nil {
		fmt.Println("  ⚠️  Could not scan project directory")
		return
	}

	fmt.Println("  📂 Ecosystem projects:")
	projectCount := 0
	for _, entry := range entries {
		if entry.IsDir() && !strings.HasPrefix(entry.Name(), ".") {
			projectPath := filepath.Join(parentDir, entry.Name())
			if hasGitRepo(projectPath) {
				fmt.Printf("    • %s\n", entry.Name())
				projectCount++
			}
		}
	}
	fmt.Printf("  🎯 Total: %d projects tracked\n", projectCount)
}

func hasGitRepo(path string) bool {
	gitPath := filepath.Join(path, ".git")
	_, err := os.Stat(gitPath)
	return err == nil
}

func getProjectName() string {
	dir := getCurrentDir()
	return filepath.Base(dir)
}

func getRecentCommits(count int) []string {
	cmd := exec.Command("git", "log", "--oneline", "-n", fmt.Sprintf("%d", count))
	output, err := cmd.Output()
	if err != nil {
		return []string{"No git history found"}
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	var commits []string
	for _, line := range lines {
		if line != "" {
			commits = append(commits, line)
		}
	}
	return commits
}

func getUncommittedChanges() []string {
	cmd := exec.Command("git", "status", "--porcelain")
	output, err := cmd.Output()
	if err != nil {
		return []string{}
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	var changes []string
	for _, line := range lines {
		if line != "" {
			changes = append(changes, line)
		}
	}
	return changes
}

func getKeyFiles() []string {
	var keyFiles []string

	// Common important files
	importantFiles := []string{
		"README.md", "NORTHSTAR.md", "main.go", "go.mod", "package.json",
		"Dockerfile", "requirements.txt", "Cargo.toml",
	}

	for _, file := range importantFiles {
		if _, err := os.Stat(file); err == nil {
			keyFiles = append(keyFiles, file)
		}
	}

	return keyFiles
}

func getCurrentDir() string {
	dir, err := os.Getwd()
	if err != nil {
		return "unknown"
	}
	return dir
}

func init() {
	// Add flags to pull command
	pullCmd.Flags().StringP("project", "p", "", "Target specific project in ecosystem")
	pullCmd.Flags().IntP("days", "d", 0, "Include last N days of history (default: recent commits)")
	pullCmd.Flags().StringP("keyword", "k", "", "Filter context by keyword")
	pullCmd.Flags().BoolP("clipboard", "c", true, "Copy to clipboard (default: true)")
	pullCmd.Flags().Bool("history", false, "Search context history instead of generating new")
	pullCmd.Flags().BoolP("save", "s", true, "Save context to history (default: true)")

	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(pullCmd)
	rootCmd.AddCommand(statusCmd)
}

func main() {
	// Initialize database
	var err error
	db, err = database.NewDB("")
	if err != nil {
		fmt.Printf("⚠️  Failed to initialize database: %v\n", err)
		// Continue without persistence
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
