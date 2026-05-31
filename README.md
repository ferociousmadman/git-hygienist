The Git-Hygienist CLI

A Go-based command-line tool engineered for validating git commit history against conventional standards to ensure repository hygiene and architectural consistency.

Language | Status | Purpose
:---      |:---    |:---
Go	     | Portfolio / Educational	|  Git Hygiene & Compliance

Project Overview

The Git-Hygienist CLI is an opinionated local development tool designed to enforce "Conventional Commit" standards within a repository's history. Moving beyond the scope of generic utility applications, this tool provides a specialized validation engine that analyzes commit metadata to ensure logs remain human-readable and machine-parsable. The current implementation demonstrates the core validation logic using a simulated commit collection, establishing a robust architectural foundation for live git integration.

This project serves as a cornerstone of a technical portfolio, specifically curated to showcase mastery of the Go programming language's core fundamentals. By prioritizing clean code, type safety, and efficient data handling, it demonstrates a readiness to build professional-grade developer tooling and solve real-world workflow challenges for high-performing engineering teams.

Core Features

* Conventional Prefix Validation: The engine performs rigorous pattern matching to validate standard commit types, including feat:, fix:, and docs:.
* Automated Health Feedback: Implements a "Commit Health" check that evaluates individual entries and provides immediate terminal feedback on compliance status.
* Extensible Logic Engine: Designed with a decoupled validation logic that currently processes a commitObjects collection, allowing for high-performance iteration over repository history.

Technical Implementation & Mastery

Core Go Concepts Applied

Go Feature	Implementation Detail	Developer Proficiency
strings package	Utilizes strings.HasPrefix to evaluate the message field of the Commit struct against conventional standards.	String Analysis & Logic Branching: Demonstrates efficient use of the standard library for pattern-based matching and conditional execution.
structs	Defines a Commit type with specific fields for Hash, Author, and Message.	Type-Safe Data Modeling: Showcases the ability to design domain-specific models that ensure data integrity across the application.
slices	Manages the commitObjects collection, storing and iterating through instances of the Commit struct.	Linear Time Complexity (O(n)): Displays mastery of Go's primary collection type for managing dynamic data sets with efficient traversal patterns.

Future Roadmap

1. Cobra CLI Migration: Transition from basic fmt/println logic to the Cobra CLI library to implement a standard-compliant POSIX interface with support for flags and subcommands.
2. Repo Branch Health Scoring: Implementation of an algorithmic scoring system to provide an objective health metric for repository branches based on commit frequency and hygiene compliance.
3. Structured Output Formatting: Adding support for diverse output formats, including Structured JSON for CI/CD integration and Terminal Tables for enhanced developer readability.
4. Advanced Pattern Identification: Development of logic to detect unreviewed PR patterns and non-standard merge behaviors within the commit history.

Installation

Prerequisites: Go 1.18 or higher.

To evaluate the validation logic locally, follow these steps:

# Clone the public repository
git clone [https://github.com/ferociousmadman/git-hygienist.git]

# Enter the project directory
cd git-hygienist

# Run the validation demonstration
go run main.go


Contact & Repository Info

This project is maintained as a Public Go repository, reflecting a commitment to open-source standards and professional version control practices.

* GitHub Profile: [https://github.com/ferociousmadman]
* LinkedIn: [https://www.linkedin.com/in/skipperdavies/]
