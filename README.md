# Raindrop Alfred Workflow

This is a workflow for [Alfred](https://www.alfredapp.com/) that opens a web page from [Raindrop](https://raindrop.io/) bookmark.

<img src="https://user-images.githubusercontent.com/34795067/179398992-3c3c02f6-e6f3-4e34-9501-b77bf1b56b3e.gif" width="480px" />

## Install

### Create Raindrop test token

A test token is required to download bookmarks from Raindrop.

`Settings` -> `Integrations` -> `For Developers` -> `Create new app`

<img src="https://user-images.githubusercontent.com/34795067/179387993-1457f2ea-804d-40f3-b115-bd8a832b711b.png" width="480px" />

`Create new app`

<img src="https://user-images.githubusercontent.com/34795067/179387997-341c5a47-78e2-4145-b73c-7bedfde60a7a.png" width="320px" />

`Create test token`

<img src="https://user-images.githubusercontent.com/34795067/179388038-eef84664-c961-4d5f-8107-342898789348.png" width="320px" />

Save the created test token.

ref. [Obtain access token \- API Documentation](https://developer.raindrop.io/v1/authentication/token)

### Download Raindrop.alfredworkflow

Download the `Raindrop.alfredworkflow` file from the [Releases](https://github.com/yskoht/alfred-raindrop/releases) page.

Then double-click it to install. Enter the test token you just saved.

<img src="https://user-images.githubusercontent.com/34795067/179397877-dbc67e2b-6d33-4f56-9eba-40d7fc95273b.png" width="480px" >

## Usage

### Sync

For fast searching, bookmarks must be downloaded from Raindrop.
A database is created locally by the `raindrop-sync` command.

<img src="https://user-images.githubusercontent.com/34795067/179388125-778481d1-0fad-433d-b6a1-fb26cddaee6e.png" width="480px" />

### Search

Type `r`, and then keywords.

<img src="https://user-images.githubusercontent.com/34795067/179388159-7feddec0-f762-41b9-a3fa-e53c1f1955db.png" width="480px" />

## License

MIT
