Git-Hygienist CLI

Git-Hygienist CLI is a terminal-based utility designed to automate the auditing of repository health through commit history analysis. It streamlines developer exerience by validating local commit messages against conventional standards, ensuring maintainable and structured version control across the development lifecycle.

Portfolio Project

This project was intended to demonstrate proficiency in the Go programming language, specifically focusing on the creation of high-performance developer tooling and knowledge of the Go standard library.

Core Features

* Conventional Commit Validation: Implements logic to identify and validate industry-standard prefixes—specifically feat:, fix:, and docs:—as explicitly defined in the core hygiene logic.
* Repository Health Scoring: Derived from the Go Programming Portfolio Roadmap, this feature audits branch health to identify unreviewed patterns and adherence to team standards.
* Standardized Terminal Output: Generates high-fidelity, structured terminal tables for immediate, actionable feedback during the audit process.

Developer Skills & Go Competencies

This project showcases an adherence to idiomatic Go practices and a "standard-library first" philosophy to ensure minimal overhead and maximum portability.

The Go Standard Library

* strings: Utilized for high-performance prefix parsing and conditional validation of commit messages.
* os: Employed for system-level interactions, including handling process output via os.Stdout.
* text/tabwriter: Knowledge of the io.Writer interface is demonstrated through the implementation of tabwriter.NewWriter. By configuring specific parameters—minwidth: 0, tabwidth: 0, padding: 3, padchar: ' ', and the tabwriter.Debug flag—the tool generates aligned terminal tables without external dependencies.

Data Structures & Logic

The application utilizes a Commit struct to model repository data with strict typing:

type Commit struct {
    Hash    string
    Author  string
    Message string
}


The audit engine employs idiomatic slice iteration and conditional filtering to process collections of these structs, evaluating the repository history against compliance rules with O(n) efficiency.

Architected using the Cobra CLI framework, the tool follows the Standard Go Project Layout. By separating command definitions into a cmd/ package (root.go for base configuration and check.go for audit logic), the application achieves a scalable command-dispatch pattern that separates the CLI entry points from core business logic.

The "Zero-Dependency" Pivot

During development, the project transitioned away from external packages such as olekukonko/tablewriter. This was a deliberate architectural decision triggered by a specific API mismatch encountered in the build logs: table.SetHeader undefined (type *tablewriter.Table has no field or method SetHeader).

Rather than troubleshooting external dependency versioning, the project pivoted to a native implementation using text/tabwriter. This minimizes the final binary size, eliminates dependency debt, and showcases an understanding of Go’s built-in formatting capabilities.

Visual Demonstration

When executing the check command, Git-Hygienist produces a structured audit of the current repository state:

+-----------------------+------------------+--------+----------+
| HASH    | AUTHOR           | TYPE   | STATUS   |
+-----------------------+------------------+--------+----------+
| a1b2c3d | ferociousmadman  | feat   | PASS     |
| e5f6g7h | ferociousmadman  | fix    | PASS     |
+-----------------------+------------------+--------+----------+


Installation & Usage

Installation

Build the binary locally using the Go toolchain:

go build -o git-hygienist


For global access, install the binary to your $GOPATH/bin:

go install .


Usage

Execute the health check against your current working directory:

# Running via go run
go run main.go check

# Running the installed binary
git-hygienist check


Repository Structure

* main.go: The application entry point that initializes the command execution.
* cmd/: Package directory containing CLI command logic.
  * root.go: Base command configuration and Cobra initialization.
  * check.go: Implementation of the commit health audit and tabwriter logic.
* go.mod / go.sum: Standard dependency management files.

Contact & Portfolio Context

This project is authored by ferociousmadman. It serves as a practical demonstration of solving real-world developer friction through efficient, idiomatic Go code. It highlights the transition from a beginner's focus on syntax to a higher level focus on architecture, dependency management, and tool performance.
