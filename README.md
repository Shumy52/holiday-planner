# Holiday Planner

A web-based holiday planning and tracking system for organizations that require employees to pre-plan their vacation days and file change requests. The app ensures clear visibility of employee availability across departments and automates paperwork generation, while preserving audit trails and approval workflows.

## Features

- Calendar interface for viewing and scheduling vacations
- Enforced yearly planning: all vacation days must be allocated at the start of the year
- Change request workflow with manager approvals
- Automatic PDF document generation based on approved requests
- Immutable history for audit compliance
- Department-based filtering and user role separation
- Public holiday awareness (automatically imported from national APIs)

## Tech Stack

| Layer        | Technology        |
|--------------|-------------------|
| Frontend     | Angular            |
| Backend      | Go (Gin)           |
| Database     | PostgreSQL         |
| Auth         | Keycloak (OAuth2)  |
| Calendar UI  | FullCalendar.js    | Maybe?
| PDF Export   | Headless LibreOffice + DOCX templates | Maybe?
| DevOps       | Docker Compose     | Maybe?

## Project Structure

```text
vacation-manager/
├── frontend/           # Angular app
├── backend/            # Go API server
├── db/                 # SQL migrations
├── templates/          # DOCX templates for PDF generation
├── scripts/            # Holiday import scripts, etc.
├── docker-compose.yml
└── README.md
