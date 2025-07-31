# AYO League API

A comprehensive REST API for managing a football league system, built with Go, Gin web framework, and GORM ORM.

## Overview

AYO League API provides endpoints for managing teams, players, matches, goals, and generating detailed reports. It features a clean architecture with proper validation, error handling, and database relationships.

## Features

- **Team Management**: Create, read, update, and delete teams
- **Player Management**: Manage players with team associations and jersey numbers
- **Match Management**: Schedule matches between teams and update scores
- **Goal Tracking**: Record goals with detailed information (player, team, match, timing)
- **Match Reports**: Generate comprehensive match reports with statistics
- **Validation**: Comprehensive validation for all entities and relationships
- **Error Handling**: Standardized error responses and validation messages

## Tech Stack

- **Language**: Go 1.23+
- **Web Framework**: Gin 1.10.1
- **ORM**: GORM 1.30.1
- **Database**: MySQL
- **Configuration**: Environment variables with .env support

## Project Structure

```
ayo-api/
├── config/           # Configuration management
├── controllers/      # API endpoint handlers
├── database/         # Database initialization
├── models/          # Database models and relationships
├── routes/          # API route definitions
├── utils/           # Utility functions and helpers
├── main.go          # Application entry point
├── go.mod           # Go module dependencies
└── .env.example     # Environment variables template
```

## Database Models

### Team
- `id`: Auto-generated primary key
- `name`: Team name (unique, required)
- `logo`: Team logo URL (required)
- `year_established`: Year the team was founded (required)
- `address`: Team address (required)
- `city`: Team city (required)
- `players`: One-to-many relationship with Player model

### Player
- `id`: Auto-generated primary key
- `name`: Player name (required)
- `height`: Player height in cm (required)
- `weight`: Player weight in kg (required)
- `position`: Player position (required)
- `team_id`: Foreign key to Team (required)
- `jersey_number`: Unique within team (required)
- `team`: Belongs-to relationship with Team
- `goals`: One-to-many relationship with Goal model

### Match
- `id`: Auto-generated primary key
- `date`: Match date (required)
- `time`: Match time in HH:MM format (required)
- `team1_id`: Foreign key to first team (required)
- `team2_id`: Foreign key to second team (required)
- `team1_score`: Score for team 1 (default: 0)
- `team2_score`: Score for team 2 (default: 0)
- `status`: Match status (scheduled, live, finished)
- `team1`: Belongs-to relationship with Team 1
- `team2`: Belongs-to relationship with Team 2
- `goals`: One-to-many relationship with Goal model

### Goal
- `id`: Auto-generated primary key
- `match_id`: Foreign key to Match (required)
- `player_id`: Foreign key to Player (required)
- `team_id`: Foreign key to Team (required)
- `goal_time`: Minute when goal was scored (required)
- `goal_type`: Type of goal (normal, penalty, own_goal)
- `match`: Belongs-to relationship with Match
- `player`: Belongs-to relationship with Player
- `team`: Belongs-to relationship with Team

## API Endpoints

### Teams
- `GET /api/teams` - Get all teams
- `GET /api/teams/:id` - Get team by ID
- `POST /api/teams` - Create new team
- `PUT /api/teams/:id` - Update team
- `DELETE /api/teams/:id` - Delete team

### Players
- `GET /api/players` - Get all players
- `GET /api/players/:id` - Get player by ID
- `POST /api/players` - Create new player
- `PUT /api/players/:id` - Update player
- `DELETE /api/players/:id` - Delete player

### Matches
- `GET /api/matches` - Get all matches
- `GET /api/matches/:id` - Get match by ID
- `POST /api/matches` - Create new match
- `PUT /api/matches/:id/score` - Update match score
- `GET /api/matches/:id/report` - Get match report

### Goals
- `POST /api/goals` - Record a new goal
- `GET /api/matches/:id/goals` - Get all goals for a match

## Getting Started

### Prerequisites

- Go 1.23 or higher
- MySQL server
- Git

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd ayo-api
```

2. Install dependencies:
```bash
go mod download
```

3. Set up environment variables:
```bash
cp .env.example .env
```

Edit `.env` file with your database credentials:
```
DB_USER=root
DB_PASSWORD=yourpassword
DB_HOST=localhost
DB_PORT=3306
DB_NAME=ayo_league
SERVER_PORT=8080
SERVER_MODE=development
```

4. Create the database:
```sql
CREATE DATABASE ayo_league;
```

5. Run the application:
```bash
go run main.go
```

The server will start on `http://localhost:8080`

## Usage Examples

### Creating a Team

```bash
curl -X POST http://localhost:8080/api/teams \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Manchester United",
    "logo": "https://example.com/manutd.png",
    "year_established": 1878,
    "address": "Old Trafford, Sir Matt Busby Way",
    "city": "Manchester"
  }'
```

### Creating a Player

```bash
curl -X POST http://localhost:8080/api/players \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Cristiano Ronaldo",
    "height": 187,
    "weight": 83,
    "position": "Forward",
    "team_id": 1,
    "jersey_number": 7
  }'
```

### Creating a Match

```bash
curl -X POST http://localhost:8080/api/matches \
  -H "Content-Type: application/json" \
  -d '{
    "date": "2024-01-15",
    "time": "20:00",
    "team1_id": 1,
    "team2_id": 2,
    "status": "scheduled"
  }'
```

### Recording a Goal

```bash
curl -X POST http://localhost:8080/api/goals \
  -H "Content-Type: application/json" \
  -d '{
    "match_id": 1,
    "player_id": 1,
    "team_id": 1,
    "goal_time": 45,
    "goal_type": "normal"
  }'
```

### Getting Match Report

```bash
curl http://localhost:8080/api/matches/1/report
```

## Validation Rules

- **Team**: Name must be unique, all fields required
- **Player**: Jersey number must be unique within team, team must exist
- **Match**: Both teams must exist and be different
- **Goal**: Player must belong to the specified team, team must be playing in the match

## Error Handling

The API returns standardized error responses:

```json
{
  "status": "error",
  "error": "Descriptive error message"
}
```

Common HTTP status codes:
- `200` - Success
- `201` - Created
- `400` - Bad Request (validation errors)
- `404` - Not Found
- `500` - Internal Server Error

## Development

### Running Tests

```bash
go test ./...
```

### Database Migrations

The application automatically runs migrations on startup. Models are defined in the `models/` directory.

### API Documentation

### Interactive Documentation
Visit our interactive API documentation at: [kebimdxst3.apidog.io](https://kebimdxst3.apidog.io)

### Manual Testing
You can also test the API using tools like Postman or curl. All endpoints follow RESTful conventions.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.